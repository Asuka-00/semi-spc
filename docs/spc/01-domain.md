# SPC系统领域模型

## 概述

本文档定义SPC（Statistical Process Control，统计过程控制）系统的领域模型。该系统为半导体制造行业设计，支持通用的Fab生产环境。

## 上下文关系

```
Site(Fab) → Area → Equipment → Chamber
Material: Lot → Wafer
Process: Technology → Product → ProcessStep → Recipe
SPC: Parameter → Spec → Chart → Sample → Measurement → Alarm → OCAP
```

## 实体定义

### 1. 工厂上下文 (Factory Context)

#### 1.1 SpcSite - 厂区/Fab
生产厂区，顶级组织单元。

字段:
- `code` - 厂区代码 (唯一)
- `name` - 厂区名称
- `timezone` - 时区 (默认: Asia/Shanghai)
- `status` - 状态 (0=禁用 1=启用)

#### 1.2 SpcArea - 区域/Area
厂区内的功能区域，如光刻区、刻蚀区。

字段:
- `site_id` - 所属厂区
- `code` - 区域代码 (唯一)
- `name` - 区域名称

示例: PHOTO(光刻区), ETCH(刻蚀区), CMP(研磨区)

#### 1.3 SpcEquipment - 设备/Equipment
生产设备。

字段:
- `site_id` - 所属厂区
- `area_id` - 所属区域
- `code` - 设备代码 (唯一)
- `name` - 设备名称
- `eqp_type` - 设备类型 (LITHO/ETCH/CVD/PVD/IMP/DIFF/CMP/METROLOGY/OTHER)
- `vendor` - 供应商

#### 1.4 SpcChamber - 腔室/Chamber
设备内的反应腔室。一台设备可包含多个腔室。

字段:
- `equipment_id` - 所属设备
- `code` - 腔室代码 (唯一)
- `name` - 腔室名称

### 2. 物料上下文 (Material Context)

#### 2.1 SpcLot - 批次/Lot
一批一起加工的晶圆。

字段:
- `site_id` - 所属厂区
- `product_id` - 产品
- `lot_id` - 批次号 (唯一)
- `lot_type` - 批次类型 (PROD/ENG/PILOT)
- `qty` - 片数

#### 2.2 SpcWafer - 晶圆/Wafer
批次内的单片晶圆。

字段:
- `lot_id` - 所属批次
- `slot_no` - 槽位号 (1-25)
- `wafer_id` - 晶圆ID (唯一)

### 3. 工艺上下文 (Process Context)

#### 3.1 SpcTechnology - 工艺技术节点/Technology
制造工艺代次，如28nm, 14nm。

字段:
- `code` - 技术代码 (唯一)
- `name` - 技术名称
- `node_nm` - 节点尺寸 (纳米)

#### 3.2 SpcProduct - 产品/Product
基于某技术节点的产品。

字段:
- `technology_id` - 所属技术节点
- `code` - 产品代码 (唯一)
- `name` - 产品名称

#### 3.3 SpcProcessStep - 工艺步骤/Process Step
制造流程中的一个工艺步骤。

字段:
- `code` - 步骤代码 (唯一)
- `name` - 步骤名称
- `step_type` - 步骤类型

示例: GATE_LITHO(栅极光刻), GATE_ETCH(栅极刻蚀)

#### 3.4 SpcRecipe - 配方/Recipe
设备在特定工艺步骤的参数配方。

字段:
- `equipment_id` - 设备
- `process_step_id` - 工艺步骤
- `code` - 配方代码 (唯一)
- `name` - 配方名称
- `version` - 版本号

### 4. SPC控制上下文 (SPC Control Context)

#### 4.1 SpcParameter - 参数/Parameter
需要监控的测量参数。

字段:
- `code` - 参数代码 (唯一)
- `name` - 参数名称
- `data_type` - 数据类型 (VARIABLE/ATTRIBUTE)
- `unit` - 单位 (nm, Å, ℃等)
- `decimal_places` - 小数位数
- `sample_level` - 采样级别 (LOT/WAFER/SITE)

示例: GATE_CD(栅极CD), OXIDE_THK(氧化层厚度)

#### 4.2 SpcSpec - 规格/Specification
产品参数的规格限，支持版本化和生效期。

字段:
- `parameter_id` - 参数
- `product_id` - 产品
- `process_step_id` - 工艺步骤
- `equipment_id` - 设备 (nullable, NULL表示全局规格)
- `version` - 版本号
- `usl` - 规格上限 (Upper Specification Limit)
- `target` - 目标值
- `lsl` - 规格下限 (Lower Specification Limit)
- `effective_from` - 生效开始时间
- `effective_to` - 生效结束时间

#### 4.3 SpcChart - 控制图/Chart
SPC控制图配置。

字段:
- `code` - 控制图代码 (唯一)
- `name` - 控制图名称
- `parameter_id` - 监控参数
- `spec_id` - 关联规格
- `chart_type` - 控制图类型 (I_MR/XBAR_R/XBAR_S/P/NP/C/U/EWMA/CUSUM)
- `subgroup_size` - 子组大小 n
- `ruleset` - 启用规则集 (逗号分隔, 如 "WE1,WE2,NELSON5")
- `limit_method` - 控制限方法 (CALC/MANUAL)
- `ewma_lambda` - EWMA平滑系数 λ (0<λ<1)
- `cusum_k` - CUSUM参考值 K
- `cusum_h` - CUSUM决策区间 H

#### 4.4 SpcControlLimit - 控制限/Control Limit
控制图的控制限，支持历史版本。

字段:
- `chart_id` - 控制图
- `ucl` - 上控制限 (Upper Control Limit)
- `cl` - 中心线 (Center Line)
- `lcl` - 下控制限 (Lower Control Limit)
- `ucl_s` - S图/MR图上控制限
- `cl_s` - S图/MR图中心线
- `lcl_s` - S图/MR图下控制限
- `calc_n` - 计算样本数
- `source` - 来源 (CALC/MANUAL)
- `effective_from` - 生效开始时间
- `effective_to` - 生效结束时间

#### 4.5 SpcRule - 控制规则/Rule
控制图的检测规则配置。

字段:
- `chart_id` - 控制图
- `rule_code` - 规则代码 (WE1/WE2/NELSON1等)
- `enabled` - 是否启用
- `n` - 连续点数
- `k` - σ倍数

#### 4.6 SpcSample - 样本/子组/Sample
一次采样数据，可能包含多个测量点。

字段:
- `chart_id` - 控制图
- `lot_id` - 批次 (nullable)
- `wafer_id` - 晶圆 (nullable)
- `equipment_id` - 设备 (nullable)
- `chamber_id` - 腔室 (nullable)
- `recipe_id` - 配方 (nullable)
- `sample_time` - 采样时间
- `subgroup_no` - 子组号
- `n` - 实际测量点数
- `mean_val` - 均值 X̄
- `range_val` - 极差 R
- `std_val` - 标准差 S
- `ooc_flag` - 失控标志
- `oos_flag` - 超规格标志

#### 4.7 SpcMeasurement - 测量值/Measurement
样本内的单个测量点数据。

字段:
- `sample_id` - 所属样本
- `seq_no` - 序号 (1-n)
- `site_x` - 测量点X坐标
- `site_y` - 测量点Y坐标
- `value` - 测量值 (变量型)
- `defect_count` - 缺陷数 (计数型)

#### 4.8 SpcAlarm - 告警/Alarm
OOC/OOS检测到的异常告警。

字段:
- `sample_id` - 触发样本
- `chart_id` - 控制图
- `alarm_type` - 告警类型 (OOC/OOS)
- `rule_code` - 触发规则代码
- `severity` - 严重度 (INFO/WARN/CRIT)
- `status` - 状态 (OPEN/ACK/CLOSED)
- `hold_lot` - 是否Hold批次

#### 4.9 SpcOcap - OCAP行动方案/Out-of-Control Action Plan
失控时的标准处理流程。

字段:
- `chart_id` - 控制图
- `name` - 方案名称
- `trigger_type` - 触发类型 (OOC/OOS/BOTH)
- `steps_json` - 步骤定义 (JSON)

#### 4.10 SpcOcapExecution - OCAP执行记录/OCAP Execution
OCAP方案的执行跟踪。

字段:
- `alarm_id` - 关联告警
- `ocap_id` - OCAP方案
- `status` - 执行状态 (PENDING/IN_PROGRESS/COMPLETED/CANCELLED)
- `owner` - 负责人
- `started_at` - 开始时间
- `closed_at` - 关闭时间
- `comment` - 处理备注

#### 4.11 SpcCapability - 过程能力分析/Capability Analysis
时间窗口内的过程能力计算结果。

字段:
- `chart_id` - 控制图
- `window_from` - 分析窗口开始
- `window_to` - 分析窗口结束
- `n` - 样本数
- `cp` - 短期能力指数 Cp
- `cpk` - 短期过程能力 Cpk
- `pp` - 长期能力指数 Pp
- `ppk` - 长期过程能力 Ppk
- `mean_val` - 均值 μ
- `std_val` - 标准差 σ

## 关系图

```
                ┌─────────┐
                │  Site   │
                └────┬────┘
                     │
          ┌──────────┴──────────┐
          ▼                     ▼
     ┌────────┐           ┌─────────┐
     │  Area  │           │   Lot   │
     └───┬────┘           └────┬────┘
         │                     │
         ▼                     ▼
  ┌──────────┐           ┌─────────┐
  │Equipment │           │  Wafer  │
  └────┬─────┘           └─────────┘
       │
       ▼
  ┌────────┐
  │Chamber │
  └────────┘

  ┌──────────┐      ┌─────────┐      ┌──────────┐
  │Technology│─────>│ Product │─────>│ProcessStep│
  └──────────┘      └────┬────┘      └──────────┘
                         │
                    ┌────▼────┐
                    │  Spec   │<──────┐
                    └────┬────┘       │
                         │            │
                    ┌────▼────┐   ┌───┴──────┐
                    │  Chart  │<──│Parameter │
                    └────┬────┘   └──────────┘
                         │
                    ┌────▼────┐
                    │ Sample  │
                    └────┬────┘
                         │
              ┌──────────┴──────────┐
              ▼                     ▼
        ┌────────────┐         ┌────────┐
        │Measurement │         │ Alarm  │
        └────────────┘         └───┬────┘
                                   │
                                   ▼
                              ┌─────────┐
                              │  OCAP   │
                              └─────────┘
```

## 数据流

### 采集流程
1. 测量设备采集数据
2. 调用 `/spc/collect` API
3. 系统保存Sample和Measurement
4. 检测OOS（与Spec比较）
5. 检测OOC（与ControlLimit和Rule比较）
6. 生成Alarm（如果异常）
7. 可触发OCAP执行

### 分析流程
1. 选择Chart和时间窗口
2. 提取历史Sample数据
3. 计算过程能力指数 Cp/Cpk/Pp/Ppk
4. 保存CapabilityAnalysis结果
5. 生成能力报告
