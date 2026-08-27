# SPC系统开发约定

## 命名规范

### 数据库表

- **前缀**: 所有SPC相关表使用 `spc_` 前缀
- **命名风格**: `snake_case`，全小写，下划线分隔
- **单复数**: 使用单数形式（与GVA `sys_user` 保持一致）
- **注释**: 每个表和字段必须有中文COMMENT

示例:
```sql
CREATE TABLE `spc_site` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `code` varchar(64) NOT NULL COMMENT '厂区代码',
  `name` varchar(200) NOT NULL COMMENT '厂区名称',
  ...
) COMMENT='厂区/Fab';
```

### 字段约定

#### 必须字段
每个表都继承 `GVA_MODEL`，包含：
- `id` - 主键
- `created_at` - 创建时间
- `updated_at` - 更新时间
- `deleted_at` - 软删除时间

#### 业务键
- 字段名: `code`
- 类型: `varchar(64) NOT NULL`
- 约束: 唯一索引 `uniqueIndex:idx_spc_xxx_code`
- 说明: 非删除行间唯一

#### 外键
- 命名: `xxx_id`（如 `site_id`, `equipment_id`）
- 类型: `uint` 或 `*uint`（可空）
- 索引: 普通索引 `index:idx_spc_xxx_yyy`
- 约束: 仅逻辑外键，不创建物理外键约束

#### 状态字段
- 字段名: `status`
- 类型: `tinyint`
- 默认值: `1`
- 含义: `0=禁用 1=启用`（除非实体有独立状态机）

#### 备注字段
- 字段名: `remark`
- 类型: `varchar(500)`
- 可空

## 代码组织

### 目录结构
```
server/
├── model/spc/           # 数据模型
├── service/spc/         # 业务逻辑
│   └── engine/         # SPC计算引擎
├── api/v1/spc/         # HTTP处理器
└── router/spc/         # 路由注册

web/
├── src/
│   ├── api/spc/        # 前端API调用
│   └── view/spc/       # 前端页面
```

### Go代码规范

#### 导出类型注释
```go
// SpcSite 厂区/Fab
type SpcSite struct {
    ...
}
```

#### 文件命名
- `spc_site.go` - snake_case
- 一个实体一个文件

#### Model vs Request/Response
- `server/model/spc/` - 数据库模型
- `server/model/spc/request/` - 请求参数
- `server/model/spc/response/` - 响应结构

#### 业务逻辑分离
- Service层: 纯业务逻辑，不依赖gin.Context
- API层: HTTP处理，参数绑定，调用Service
- 无业务逻辑在Handler中

### API约定

#### 路由风格
遵循GVA v2.9.1风格，使用动词前缀：
- `/spc/createXxx` - 创建
- `/spc/updateXxx` - 更新
- `/spc/deleteXxx` - 删除
- `/spc/findXxx` - 查询单个
- `/spc/getXxxList` - 查询列表

不使用 RESTful 风格（POST `/spc/xxx`, PUT `/spc/xxx/:id`）

#### 响应格式
统一信封格式：
```json
{
  "code": 0,
  "data": {...},
  "msg": "成功"
}
```

#### 分页
请求:
```json
{
  "page": 1,
  "pageSize": 10
}
```

响应:
```json
{
  "list": [...],
  "total": 100,
  "page": 1,
  "pageSize": 10
}
```

#### 权限
- 写操作: JWT + Casbin鉴权
- 数据采集API: `/spc/collect` 支持JWT或API Token

#### Swagger文档
每个Handler必须包含完整的swaggo注释：
```go
// CreateSpcSite
// @Tags      SpcSite
// @Summary   创建厂区
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      spc.SpcSite  true  "厂区信息"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /spc/site [post]
```

## Vue前端规范

### 文件头注释
每个SFC文件顶部标注模块和用途：
```vue
<template>
  <!-- 厂区管理 - 列表页面 -->
  ...
</template>
```

### API调用
使用 `web/src/api/spc/` 下的封装函数，不直接调用axios

### 组件命名
- 文件: PascalCase `SiteManage.vue`
- 组件名: PascalCase `SiteManage`

## 枚举和常量

### 设备类型 (EqpType)
```
LITHO     - 光刻
ETCH      - 刻蚀
CVD       - 化学气相沉积
PVD       - 物理气相沉积
IMP       - 离子注入
DIFF      - 扩散
CMP       - 化学机械研磨
METROLOGY - 测量
OTHER     - 其他
```

### 控制图类型 (ChartType)
```
I_MR      - 个体-移动极差图
XBAR_R    - 均值-极差图
XBAR_S    - 均值-标准差图
P         - 不合格品率图
NP        - 不合格品数图
C         - 缺陷数图
U         - 单位缺陷数图
EWMA      - 指数加权移动平均图
CUSUM     - 累积和图
```

### SPC规则代码 (RuleCode)
```
WE1-WE4       - Western Electric规则
NELSON1-NELSON8 - Nelson规则
```

### 告警类型 (AlarmType)
```
OOC - 失控 (Out Of Control)
OOS - 超规格 (Out Of Spec)
```

### 严重度 (Severity)
```
INFO - 信息
WARN - 警告
CRIT - 严重
```

## 数据类型

### 参数数据类型 (DataType)
```
VARIABLE  - 变量型（连续数据）
ATTRIBUTE - 计数型（离散数据）
```

### 批次类型 (LotType)
```
PROD  - 量产批
ENG   - 工程批
PILOT - 试产批
```

### 采样级别 (SampleLevel)
```
LOT   - 批次级
WAFER - 晶圆级
SITE  - 测量点级
```

## Git提交规范

遵循 Conventional Commits:
```
feat(spc): 添加新功能
fix(spc): 修复Bug
docs(spc): 文档变更
style(spc): 代码格式
refactor(spc): 重构
test(spc): 测试
chore(spc): 构建/工具变更
```

## 行业术语缩写

在代码和注释中可使用以下半导体行业标准缩写：
- **EQP** - Equipment (设备)
- **CD** - Critical Dimension (关键尺寸)
- **USL** - Upper Specification Limit (规格上限)
- **LSL** - Lower Specification Limit (规格下限)
- **UCL** - Upper Control Limit (控制上限)
- **LCL** - Lower Control Limit (控制下限)
- **CL** - Center Line (中心线)
- **OOC** - Out Of Control (失控)
- **OOS** - Out Of Spec (超规格)
- **OCAP** - Out-of-Control Action Plan (失控行动方案)
- **Cp/Cpk** - Process Capability Index (过程能力指数)
- **Pp/Ppk** - Process Performance Index (过程性能指数)
- **Fab** - Fabrication (晶圆厂)

但COMMENT中必须给出中文全称。
