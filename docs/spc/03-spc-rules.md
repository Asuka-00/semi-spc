# SPC规则和计算方法

## SPC控制图类型

### 1. I-MR图 (Individual-Moving Range)
**适用场景:** 单点测量，无法形成子组

**统计量:**
- I: 个体值
- MR: 连续两个值的移动极差

**控制限计算:**
- I图: UCL = X̄ + 2.66×MR̄, LCL = X̄ - 2.66×MR̄
- MR图: UCL = 3.267×MR̄, LCL = 0

### 2. X̄-R图 (Mean-Range)
**适用场景:** 子组大小 n=2~10，最常用

**统计量:**
- X̄: 子组均值
- R: 子组极差 (max - min)

**控制限计算:**
- X̄图: UCL = X̿ + A₂×R̄, LCL = X̿ - A₂×R̄
- R图: UCL = D₄×R̄, LCL = D₃×R̄

**常数表 (n=2~10):**
| n | A₂    | D₃    | D₄    | d₂    |
|---|-------|-------|-------|-------|
| 2 | 1.880 | 0.000 | 3.267 | 1.128 |
| 3 | 1.023 | 0.000 | 2.574 | 1.693 |
| 4 | 0.729 | 0.000 | 2.282 | 2.059 |
| 5 | 0.577 | 0.000 | 2.114 | 2.326 |
| 6 | 0.483 | 0.000 | 2.004 | 2.534 |
| 7 | 0.419 | 0.076 | 1.924 | 2.704 |
| 8 | 0.373 | 0.136 | 1.864 | 2.847 |
| 9 | 0.337 | 0.184 | 1.816 | 2.970 |
| 10| 0.308 | 0.223 | 1.777 | 3.078 |

### 3. X̄-S图 (Mean-Standard Deviation)
**适用场景:** 子组大小 n>10

**统计量:**
- X̄: 子组均值
- S: 子组标准差

**控制限计算:**
- X̄图: UCL = X̿ + A₃×S̄, LCL = X̿ - A₃×S̄
- S图: UCL = B₄×S̄, LCL = B₃×S̄

## OOS检测 (Out Of Spec)

**定义:** 测量值超出规格限

**判定:**
```
if value > USL: OOS (超上限)
if value < LSL: OOS (超下限)
```

**规格单边:**
- 仅上限: 只检查USL
- 仅下限: 只检查LSL
- 无规格: 不检测OOS

## OOC检测 (Out Of Control)

**定义:** 过程失控，违反控制规则

### Western Electric规则

#### WE1: 点超出控制限
**条件:** 任何点超出 UCL 或 LCL

**严重度:** CRIT

**图示:**
```
UCL ─────────────────────
           ×            ← 触发WE1
CL  ─────────────────────
            
LCL ─────────────────────
```

#### WE2: 3点中2点在A区外（同侧）
**条件:** 连续3点中有2点超出 ±2σ（同一侧）

**区域定义:**
- A区上界: CL + 2σ
- A区下界: CL - 2σ

**严重度:** WARN

#### WE3: 5点中4点在C区外（同侧）
**条件:** 连续5点中有4点超出 ±1σ（同一侧）

**区域定义:**
- C区上界: CL + 1σ
- C区下界: CL - 1σ

**严重度:** WARN

#### WE4: 连续8点在中心线同一侧
**条件:** 连续8点全部在CL上方或下方

**严重度:** WARN

**图示:**
```
UCL ─────────────────────
        • • • •         ← 4点在上侧
CL  ─────────────────────
    • • • •             ← 4点在上侧
LCL ─────────────────────
    ← 连续8点同侧，触发WE4
```

### Nelson规则

#### NELSON1: 点超出控制限
同 WE1

#### NELSON2: 连续9点在中心线同一侧
类似WE4但更严格（9点 vs 8点）

#### NELSON3/NELSON5: 连续6点递增或递减
**条件:** 连续6点单调递增或递减

**严重度:** WARN

**用途:** 检测趋势变化

```
      •
    •
  •
•         ← 递增趋势
```

#### NELSON4: 连续14点交替上下
**条件:** 连续14点交替高低

**严重度:** INFO

**用途:** 检测系统性振荡

#### NELSON6: 5点中4点距离中心线>1σ（同侧）
类似WE3

#### NELSON7: 连续15点在中心线±1σ内
**条件:** 15点都很接近CL

**严重度:** INFO

**用途:** 检测过度控制或数据造假

#### NELSON8: 连续8点距离中心线>1σ（两侧）
**条件:** 8点都远离CL（可在不同侧）

**严重度:** WARN

**用途:** 检测混合分布

## 过程能力分析

### 能力指数 Cp/Cpk

**Cp (Process Capability):**
```
Cp = (USL - LSL) / (6σ)
```
- 衡量过程能力（不考虑偏移）
- 假设过程居中

**Cpk (Process Capability Index):**
```
CPU = (USL - μ) / (3σ)
CPL = (μ - LSL) / (3σ)
Cpk = min(CPU, CPL)
```
- 考虑过程偏移
- Cpk ≤ Cp，相等时过程完全居中

**判定标准:**
- Cpk ≥ 1.67: 优秀
- Cpk ≥ 1.33: 良好
- Cpk ≥ 1.00: 尚可
- Cpk < 1.00: 不足

### 性能指数 Pp/Ppk

**区别于 Cp/Cpk:**
- Cp/Cpk: 使用Within标准差（R̄/d₂或S̄），反映短期能力
- Pp/Ppk: 使用Overall标准差（总体σ），反映长期性能

**计算:**
```
Pp = (USL - LSL) / (6σ_overall)
Ppk = min((USL-μ)/(3σ_overall), (μ-LSL)/(3σ_overall))
```

**关系:**
- Pp ≥ Cp (长期变异 ≥ 短期变异)
- Pp/Cp 比值反映过程稳定性

### σ估计方法

**From R (极差法):**
```
σ = R̄ / d₂
```

**From S (标准差法):**
```
σ = S̄ / c₄
```
其中 c₄ ≈ √(2/(n-1)) × Γ(n/2) / Γ((n-1)/2)

**From MR (移动极差法, I-MR图):**
```
σ = MR̄ / 1.128
```

## 规则配置

### 默认规则集
系统默认启用 **WE1**（点超出控制限），这是最基本的失控检测。

### 推荐规则组合

**保守型 (低误报):**
```
WE1
```

**标准型 (平衡):**
```
WE1, WE2, WE4
```

**灵敏型 (高检出):**
```
WE1, WE2, WE3, WE4, NELSON5
```

**全面型 (最大检出):**
```
WE1, WE2, WE3, WE4, NELSON2, NELSON3, NELSON5, NELSON6, NELSON7, NELSON8
```

## 计算引擎实现

### 引擎位置
```
server/service/spc/engine/
├── constants.go      # 控制图常数
├── rules.go          # OOC规则检测
├── capability.go     # 能力指数计算
```

### 单元测试
所有计算逻辑都有完整的单元测试：
```bash
go test -v ./server/service/spc/engine/...
```

### 使用示例

**计算控制限:**
```go
import "github.com/flipped-aurora/gin-vue-admin/server/service/spc/engine"

xBarBar := 100.0  // 总均值
rBar := 5.0       // 平均极差
n := 5            // 子组大小

uclX, clX, lclX, uclR, clR, lclR := engine.XbarRLimits(xBarBar, rBar, n)
```

**OOC检测:**
```go
values := []float64{100, 102, 115, 103, 101}  // 最近的样本值
ucl := 110.0
cl := 100.0
lcl := 90.0
enabledRules := []string{"WE1", "WE2", "WE4"}

violations := engine.CheckOOC(values, ucl, cl, lcl, enabledRules)
for _, v := range violations {
    fmt.Printf("规则 %s: %s\n", v.RuleCode, v.Message)
}
```

**能力分析:**
```go
values := []float64{95, 100, 105, 98, 102, 97, 103}
usl := 110.0
lsl := 90.0

result := engine.CalculateCapability(values, &usl, &lsl, nil, 0)
fmt.Printf("Cp=%.2f, Cpk=%.2f\n", result.Cp, result.Cpk)
```

## 参考文献

1. Montgomery, D.C. (2012). *Statistical Quality Control*. 7th Edition.
2. Western Electric Company (1956). *Statistical Quality Control Handbook*.
3. Nelson, L.S. (1984). *The Shewhart Control Chart—Tests for Special Causes*. Journal of Quality Technology.
