# SPC 生产数据摄取契约

本文档定义半导体fab生产线与SPC系统之间的数据摄取接口契约。

## 认证方式

所有数据摄取接口均需要认证，支持两种方式：

### 方式1：JWT Token（用户登录）
- Header: `x-token: <JWT_TOKEN>`
- 获取方式：用户通过GVA前端登录后获得
- 适用场景：手动测试、前端页面调用

### 方式2：API Token（服务集成）
- Header: `X-API-Token: <API_TOKEN>` 或 `x-token: <API_TOKEN>`
- 获取方式：在GVA管理后台【系统管理 / API Token】中签发
  1. 登录GVA管理后台（http://your-server:port/）
  2. 进入【系统管理】→【API Token】
  3. 点击【签发Token】
  4. 选择用户、角色、有效期（天数，-1为永久）
  5. 输入备注（如："MES服务集成"）
  6. 复制生成的Token
- 适用场景：MES/SECS主机、自动化脚本集成
- **重要**：API Token本质是长期有效的JWT，妥善保管

**认证失败错误**：
- 未提供token：`未提供认证凭据，请在header中提供x-token（JWT）或X-API-Token（API Token）`
- token无效：`API Token无效或已失效`
- token过期：`API Token已过期`

## 摄取端点

### POST /api/spc/collect

单次数据采集，采集一个子组的测量数据。

**幂等性**：
- 通过Header `X-Idempotency-Key` 或 body `idempotencyKey` 提供幂等性键
- 相同的key + chartCode返回原始结果，不会创建重复样本
- 建议格式：`{host}_{timestamp}_{chartCode}_{subgroupNo}` 或UUID

**请求Headers**：
```
Content-Type: application/json
X-API-Token: <your_api_token>  (或 x-token: <your_jwt>)
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
| `未提供认证凭据...` | 200 | 7 | 缺少x-token或X-API-Token header |
| `API Token无效或已失效` | 200 | 7 | token不存在或status=false |
| `API Token已过期` | 200 | 7 | token.expiresAt < now |
| `控制图不存在: CHART_XXX` | 200 | 7 | chartCode不存在 |
| `控制图未启用` | 200 | 7 | chart.status=0 |
| `测量值数量必须等于子组大小` | 200 | 7 | len(values) != chart.subgroupSize |
| `亚组号已存在` | 200 | 7 | 相同chartId + subgroupNo重复（非幂等键场景） |
| `规格配置不存在` | 200 | 7 | chart.specId无效 |

### POST /api/spc/collectCsv

批量CSV数据采集。

**认证方式**：同上（X-API-Token或x-token header）

**请求Headers**：
```
Content-Type: multipart/form-data
X-API-Token: <your_api_token>  (或 x-token: <your_jwt>)
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

### 签发API Token（管理员操作）
```bash
# 登录后使用JWT签发API Token
curl -X POST "https://fab.example.com/sysApiToken/createApiToken" \
  -H "Content-Type: application/json" \
  -H "x-token: <ADMIN_JWT_TOKEN>" \
  -d '{
    "userId": 2,
    "authorityId": 888,
    "days": 365,
    "remark": "MES服务集成"
  }'

# 响应示例：
# {
#   "code": 0,
#   "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
#   "msg": "签发成功"
# }
```

### 单次采集（使用API Token，幂等）
```bash
curl -X POST "https://fab.example.com/api/spc/collect" \
  -H "Content-Type: application/json" \
  -H "X-API-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
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

### CSV批量采集（使用API Token）
```bash
curl -X POST "https://fab.example.com/api/spc/collectCsv" \
  -H "X-API-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -F "file=@samples.csv"
```

## 集成检查清单

在生产部署前，请确认：

- [ ] **签发API Token**：在GVA后台【系统管理/API Token】签发长期token
- [ ] **配置认证头**：MES/SECS主机在HTTP请求中添加`X-API-Token`或`x-token`
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
  signing-key: "YOUR_PRODUCTION_SECRET_KEY_HERE"  # 最少32字符，API Token也使用此密钥

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
4. 在【系统管理/API Token】为MES集成签发长期token（建议365天或永久）
5. 妥善保管API Token（如丢失可在后台作废重新签发）

## API Token生命周期管理

### 签发Token
- 管理员登录GVA后台
- 【系统管理】→【API Token】→【签发Token】
- 选择用户、角色、有效期（-1为永久）
- 记录token和备注（如"MES主机A"）

### 查询Token
- 【系统管理】→【API Token】→【Token列表】
- 可按用户筛选，查看状态和过期时间

### 作废Token
- 在Token列表点击【作废】
- Token立即失效（加入黑名单），无法再使用
- 如需恢复，重新签发新token

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
- API Token签发记录截图（脱敏token值）

---
**版本**: v1.1  
**最后更新**: 2024-01-15（新增API Token认证）  
**维护**: SPC系统团队
