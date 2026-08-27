# SPC 生产数据摄取契约

本文档定义半导体fab生产线与SPC系统之间的数据摄取接口契约。

## 摄取端点

### POST /api/spc/collect

单次数据采集，采集一个子组的测量数据。

**认证方式**：
- JWT Token (Header: `x-token`)
- API Token (Header: `X-API-Token`，适用于MES/SECS主机集成)

**幂等性**：
- 通过Header `X-Idempotency-Key` 或 body `idempotencyKey` 提供幂等性键
- 相同的key + chartCode返回原始结果，不会创建重复样本
- 建议格式：`{host}_{timestamp}_{chartCode}_{subgroupNo}` 或UUID

**请求Headers**：
```
Content-Type: application/json
x-token: <JWT Token>
X-Idempotency-Key: <optional, 幂等性键>
```

**请求Body字段**：

| 字段 | 类型 | 必填 | 说明 | 示例 |
|------|------|------|------|------|
| chartCode | string | 是 | 控制图代码 | `CHART_CD_001` |
| lotId | string | 否 | 批次ID | `LOT2024001` |
| waferId | string | 否 | 晶圆ID | `W001` |
| equipmentId | uint | 否 | 设备ID | `1` |
| chamberId | uint | 否 | 腔室ID | `1` |
| recipeId | uint | 否 | 配方ID | `1` |
| sampleTime | string (ISO8601) | 是 | 采样时间 | `2024-01-15T10:30:00Z` |
| subgroupNo | int | 否 | 子组号（可选，用于幂等性） | `123` |
| values | array of float64 | 是 | 测量值数组，长度必须等于控制图的subgroupSize | `[1.23, 1.25, 1.24, 1.22, 1.26]` |
| idempotencyKey | string | 否 | 幂等性键（可在body或header提供） | `mes-host_1705311000_CHART_CD_001_123` |

**成功响应（HTTP 200）**：
```json
{
  "code": 0,
  "data": {
    "sampleId": 101,
    "oocFlag": false,
    "oosFlag": false,
    "violations": [],
    "alarms": [],
    "message": "数据采集成功"
  },
  "msg": "采集成功"
}
```

**OOC/OOS异常响应（HTTP 200，业务层失控/超规格）**：
```json
{
  "code": 0,
  "data": {
    "sampleId": 102,
    "oocFlag": true,
    "oosFlag": false,
    "violations": [
      {
        "ruleCode": "WE1",
        "severity": "CRIT",
        "message": "点超出3σ控制限",
        "position": 5
      }
    ],
    "alarms": [201, 202],
    "message": "数据采集成功，检测到异常: 失控"
  },
  "msg": "采集成功"
}
```

**错误响应示例**：

| 错误信息 | HTTP状态 | Code | 原因 |
|----------|----------|------|------|
| `控制图不存在: CHART_XXX` | 200 | 7 | chartCode不存在 |
| `控制图未启用` | 200 | 7 | chart.status=0 |
| `测量值数量必须等于子组大小` | 200 | 7 | len(values) != chart.subgroupSize |
| `亚组号已存在` | 200 | 7 | 相同chartId + subgroupNo重复（非幂等键场景） |
| `规格配置不存在` | 200 | 7 | chart.specId无效 |
| `未授权` | 400 | - | JWT/API Token无效或缺失 |

### POST /api/spc/collectCsv

批量CSV数据采集。

**认证方式**：同上

**请求Headers**：
```
Content-Type: multipart/form-data
x-token: <JWT Token>
```

**请求Body**：
- `file`: CSV文件（multipart/form-data）

**CSV格式**：

CSV列（第一行header）：
```
chartCode,lotId,waferId,equipmentId,chamberId,recipeId,sampleTime,subgroupNo,value1,value2,value3,...
```

示例：
```csv
chartCode,lotId,waferId,sampleTime,subgroupNo,value1,value2,value3,value4,value5
CHART_CD_001,LOT2024001,W001,2024-01-15T10:30:00Z,1,1.23,1.25,1.24,1.22,1.26
CHART_CD_001,LOT2024001,W002,2024-01-15T10:31:00Z,2,1.24,1.26,1.25,1.23,1.27
```

**注意**：
- `value1`, `value2`, ... 列数应匹配控制图的`subgroupSize`
- CSV不支持幂等性键（如需幂等，请使用单次API + idempotency key）

**成功响应（HTTP 200）**：
```json
{
  "code": 0,
  "data": {
    "totalRows": 50,
    "successCount": 48,
    "failCount": 2,
    "errors": [
      {"row": 5, "error": "控制图未启用"},
      {"row": 12, "error": "测量值数量必须等于子组大小"}
    ]
  },
  "msg": "CSV处理完成"
}
```

## cURL示例

### 单次采集（幂等）
```bash
curl -X POST "https://fab.example.com/api/spc/collect" \
  -H "Content-Type: application/json" \
  -H "x-token: YOUR_JWT_TOKEN" \
  -H "X-Idempotency-Key: mes-host_1705311000_CHART_CD_001_123" \
  -d '{
    "chartCode": "CHART_CD_001",
    "lotId": "LOT2024001",
    "waferId": "W001",
    "equipmentId": 1,
    "chamberId": 1,
    "sampleTime": "2024-01-15T10:30:00Z",
    "subgroupNo": 123,
    "values": [1.23, 1.25, 1.24, 1.22, 1.26]
  }'
```

### CSV批量采集
```bash
curl -X POST "https://fab.example.com/api/spc/collectCsv" \
  -H "x-token: YOUR_JWT_TOKEN" \
  -F "file=@samples.csv"
```

## 集成检查清单

在生产部署前，请确认：

- [ ] **认证**：已配置API Token或JWT signing-key（不使用默认值）
- [ ] **幂等性**：MES/SECS主机已实现X-Idempotency-Key生成逻辑
- [ ] **错误处理**：客户端正确处理HTTP 200 + code=7的业务错误
- [ ] **OOC/OOS告警**：客户端能解析violations和alarms字段
- [ ] **Lot Hold**：如果控制图配置了`holdLot=true`，CRIT告警会自动Hold批次，客户端需同步lot状态
- [ ] **网络重试**：实现指数退避重试（带相同idempotency key）
- [ ] **时区**：sampleTime使用UTC或明确时区（ISO8601格式）
- [ ] **CSV编码**：UTF-8编码，BOM可选

## 生产环境安全配置

**必须修改的默认值**（在`server/config.yaml`）：

```yaml
# JWT签名密钥（默认值：qmPlus）
jwt:
  signing-key: "YOUR_PRODUCTION_SECRET_KEY_HERE"  # 最少32字符

# MySQL root密码
mysql:
  username: spc_user
  password: "YOUR_MYSQL_PASSWORD"  # 不使用root或默认密码

# Redis密码
redis:
  password: "YOUR_REDIS_PASSWORD"
```

**默认GVA管理员账号**：
- 用户名：`admin`
- 密码：`123456`
- **用途**：仅用于初始化`/init`，之后立即修改密码或禁用

**生产部署后**：
1. 修改admin密码
2. 创建专用服务账号（通过GVA用户管理）
3. 配置Casbin权限（仅授予SPC相关API权限）
4. 为MES集成配置API Token（如使用）

## 错误码速查

| code | 含义 |
|------|------|
| 0    | 成功 |
| 7    | 业务错误（msg字段包含详细原因） |
| 400  | 认证失败或请求格式错误 |

## 联系与支持

如遇集成问题，请提供：
- cURL命令或HTTP请求日志
- 完整的响应body（包括code和msg）
- 控制图配置截图（SPC系统中的chart配置）

---
**版本**: v1.0  
**最后更新**: 2024-01-15  
**维护**: SPC系统团队
