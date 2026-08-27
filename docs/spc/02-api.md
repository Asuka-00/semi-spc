# SPC系统API文档

## 概述

本文档描述SPC系统的HTTP API接口。所有API遵循GVA v2.9.1的约定。

## 认证

### JWT Token
大部分API需要JWT认证：
```http
Authorization: Bearer <token>
```

### API Token
数据采集API (`/spc/collect`) 支持API Token认证，适用于设备自动上传。

## 通用响应格式

### 成功响应
```json
{
  "code": 0,
  "data": {...},
  "msg": "成功"
}
```

### 失败响应
```json
{
  "code": 7,
  "data": {},
  "msg": "错误信息"
}
```

### 分页响应
```json
{
  "code": 0,
  "data": {
    "list": [...],
    "total": 100,
    "page": 1,
    "pageSize": 10
  },
  "msg": "获取成功"
}
```

## API端点

所有SPC API的基础路径为 `/api/spc`

### 厂区管理 (Site)

#### 创建厂区
```http
POST /api/spc/site
Content-Type: application/json
Authorization: Bearer <token>

{
  "code": "FAB1",
  "name": "Demo Fab 1",
  "timezone": "Asia/Shanghai",
  "status": 1,
  "remark": "演示厂区"
}
```

#### 更新厂区
```http
PUT /api/spc/site
Content-Type: application/json
Authorization: Bearer <token>

{
  "ID": 1,
  "code": "FAB1",
  "name": "Demo Fab 1 Updated",
  ...
}
```

#### 删除厂区
```http
DELETE /api/spc/site
Content-Type: application/json
Authorization: Bearer <token>

{
  "ID": 1
}
```

#### 获取厂区详情
```http
GET /api/spc/site?ID=1
Authorization: Bearer <token>
```

#### 获取厂区列表
```http
GET /api/spc/getSiteList?page=1&pageSize=10
Authorization: Bearer <token>
```

### 数据采集 (Collect)

#### 采集数据
这是SPC系统的核心API，用于实时数据采集和OOC/OOS检测。

```http
POST /api/spc/collect
Content-Type: application/json
Authorization: Bearer <token>

{
  "chartCode": "CHART_GATE_CD",
  "lotId": "LOT001",
  "waferId": "W001",
  "equipmentId": 1,
  "chamberId": 1,
  "recipeId": 1,
  "sampleTime": "2026-08-27T10:30:00Z",
  "subgroupNo": 101,
  "values": [45.2, 45.5, 44.8, 45.1, 45.3]
}
```

**请求字段说明:**
- `chartCode` (required) - 控制图代码
- `lotId` (optional) - 批次号
- `waferId` (optional) - 晶圆ID
- `equipmentId` (optional) - 设备ID
- `chamberId` (optional) - 腔室ID
- `recipeId` (optional) - 配方ID
- `sampleTime` (required) - 采样时间
- `subgroupNo` (required) - 子组号
- `values` (required) - 测量值数组

**响应:**
```json
{
  "code": 0,
  "data": {
    "sampleId": 123,
    "oocFlag": false,
    "oosFlag": false,
    "violations": [],
    "alarms": [],
    "message": "数据采集成功"
  },
  "msg": "数据采集成功"
}
```

**异常情况响应:**
```json
{
  "code": 0,
  "data": {
    "sampleId": 124,
    "oocFlag": true,
    "oosFlag": false,
    "violations": [
      {
        "ruleCode": "WE1",
        "message": "点超出控制限",
        "points": [2],
        "severity": "CRIT"
      }
    ],
    "alarms": [45, 46],
    "message": "数据采集成功，检测到异常: 失控"
  },
  "msg": "数据采集成功"
}
```

**注意:** OOC/OOS告警不会导致HTTP错误。API返回200状态码，异常信息在响应体中。

### 告警管理 (Alarm)

#### 获取告警列表
```http
GET /api/spc/getAlarmList?page=1&pageSize=10&status=OPEN&alarmType=OOC
Authorization: Bearer <token>
```

**查询参数:**
- `page` - 页码
- `pageSize` - 每页大小
- `status` (optional) - 告警状态: OPEN/ACK/CLOSED
- `alarmType` (optional) - 告警类型: OOC/OOS

#### 确认告警
```http
POST /api/spc/acknowledgeAlarm?remark=已查看
Content-Type: application/json
Authorization: Bearer <token>

{
  "ID": 45
}
```

#### 关闭告警
```http
POST /api/spc/closeAlarm?remark=已处理完成
Content-Type: application/json
Authorization: Bearer <token>

{
  "ID": 45
}
```

#### 获取告警统计
```http
GET /api/spc/getAlarmStatistics?days=7
Authorization: Bearer <token>
```

**响应:**
```json
{
  "code": 0,
  "data": {
    "byType": [
      {"alarmType": "OOC", "count": 15},
      {"alarmType": "OOS", "count": 8}
    ],
    "byStatus": [
      {"status": "OPEN", "count": 10},
      {"status": "ACK", "count": 8},
      {"status": "CLOSED", "count": 5}
    ],
    "bySeverity": [
      {"severity": "CRIT", "count": 5},
      {"severity": "WARN", "count": 15},
      {"severity": "INFO", "count": 3}
    ]
  },
  "msg": "获取成功"
}
```

### 控制图管理 (Chart)

控制图的CRUD操作与Site类似，遵循相同的API模式：
- `POST /api/spc/createChart` - 创建
- `PUT /api/spc/updateChart` - 更新
- `DELETE /api/spc/deleteChart` - 删除
- `GET /api/spc/findChart` - 查询单个
- `GET /api/spc/getChartList` - 查询列表

### 能力分析 (Capability)

#### 计算过程能力
```http
POST /api/spc/calculateCapability
Content-Type: application/json
Authorization: Bearer <token>

{
  "chartId": 1,
  "windowFrom": "2026-08-20T00:00:00Z",
  "windowTo": "2026-08-27T23:59:59Z"
}
```

**响应:**
```json
{
  "code": 0,
  "data": {
    "ID": 10,
    "chartId": 1,
    "windowFrom": "2026-08-20T00:00:00Z",
    "windowTo": "2026-08-27T23:59:59Z",
    "n": 50,
    "cp": 1.33,
    "cpk": 1.15,
    "pp": 1.28,
    "ppk": 1.10,
    "meanVal": 45.0,
    "stdVal": 1.25
  },
  "msg": "计算成功"
}
```

## 完整的实体API列表

以下实体都有相同的CRUD API模式（create/update/delete/find/getList）：

- Site (厂区)
- Area (区域)
- Equipment (设备)
- Chamber (腔室)
- Technology (技术节点)
- Product (产品)
- ProcessStep (工艺步骤)
- Recipe (配方)
- Lot (批次)
- Wafer (晶圆)
- Parameter (参数)
- Spec (规格)
- Chart (控制图)
- ControlLimit (控制限)
- Rule (规则)
- Sample (样本)
- Alarm (告警)
- Ocap (OCAP方案)
- OcapExecution (OCAP执行)
- Capability (能力分析)

## 错误码

| Code | 说明 |
|------|------|
| 0    | 成功 |
| 7    | 失败/错误 |

更详细的错误信息在 `msg` 字段中。

## 数据采集最佳实践

### 1. 批量采集
对于高频采集场景，建议批量累积后定期上传，而不是每个点都实时调用API。

### 2. 异步处理
数据采集API是同步的，会立即返回OOC/OOS结果。对于不需要实时反馈的场景，可以异步上传。

### 3. 重试机制
网络故障时应实现重试机制，建议指数退避策略。

### 4. API Token轮换
定期轮换API Token以提高安全性。

## Swagger文档

完整的API文档可通过Swagger UI访问：
```
http://localhost:8888/swagger/index.html
```

生成Swagger文档：
```bash
cd server
swag init
```
