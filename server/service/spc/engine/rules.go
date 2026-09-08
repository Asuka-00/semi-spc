package engine

import (
	"fmt"
	"math"
)

// RuleViolation 规则违反记录
type RuleViolation struct {
	RuleCode   string         `json:"ruleCode"`
	Message    string         `json:"message"`
	MessageKey string         `json:"messageKey"`
	Params     map[string]any `json:"params,omitempty"`
	Points     []int          `json:"points"`
	Severity   string         `json:"severity"`
}

func violation(code, key, fallback, severity string, points []int, params map[string]any) RuleViolation {
	if params == nil {
		params = map[string]any{}
	}
	return RuleViolation{
		RuleCode:   code,
		Message:    fallback,
		MessageKey: key,
		Params:     params,
		Points:     points,
		Severity:   severity,
	}
}

func sigmaOf(ucl, cl float64) float64 {
	if ucl == cl {
		return 0
	}
	return (ucl - cl) / 3.0
}

// CheckOOS 检查是否超规格 (Out Of Spec)
func CheckOOS(value float64, usl, lsl *float64) (bool, string) {
	if usl != nil && value > *usl {
		return true, "超上规格限 USL"
	}
	if lsl != nil && value < *lsl {
		return true, "超下规格限 LSL"
	}
	return false, ""
}

// CheckOOC 使用规则代码列表（默认参数）检测 OOC
func CheckOOC(values []float64, ucl, cl, lcl float64, enabledRules []string) []RuleViolation {
	configs := make([]RuleConfig, 0, len(enabledRules))
	for _, code := range enabledRules {
		configs = append(configs, DefaultRuleConfig(code))
	}
	return CheckOOCWithConfig(values, ucl, cl, lcl, configs)
}

// CheckOOCWithConfig 使用可配置规则参数检测 OOC
func CheckOOCWithConfig(values []float64, ucl, cl, lcl float64, rules []RuleConfig) []RuleViolation {
	violations := []RuleViolation{}
	if len(values) == 0 {
		return violations
	}
	for _, raw := range rules {
		cfg := raw.withDefaults()
		switch cfg.kind() {
		case KindBeyondLimit:
			if v := checkBeyondLimit(values, ucl, cl, lcl, cfg); v != nil {
				violations = append(violations, *v)
			}
		case KindKOfNZone:
			violations = append(violations, checkKOfNZone(values, ucl, cl, cfg)...)
		case KindConsecutiveSide:
			if v := checkConsecutiveSide(values, cl, cfg); v != nil {
				violations = append(violations, *v)
			}
		case KindConsecutiveTrend:
			violations = append(violations, checkConsecutiveTrend(values, cfg)...)
		case KindConsecutiveAlt:
			if v := checkConsecutiveAlternate(values, cfg); v != nil {
				violations = append(violations, *v)
			}
		case KindConsecutiveWithin:
			if v := checkConsecutiveWithin(values, ucl, cl, cfg); v != nil {
				violations = append(violations, *v)
			}
		case KindConsecutiveBeyond:
			if v := checkConsecutiveBeyondBoth(values, ucl, cl, cfg); v != nil {
				violations = append(violations, *v)
			}
		}
	}
	return violations
}

func checkBeyondLimit(values []float64, ucl, cl, lcl float64, cfg RuleConfig) *RuleViolation {
	upper, lower := ucl, lcl
	if cfg.K > 0 {
		sig := sigmaOf(ucl, cl)
		upper = cl + cfg.K*sig
		lower = cl - cfg.K*sig
	}
	for i, v := range values {
		if v > upper || v < lower {
			msg := fmt.Sprintf("点超出控制限(K=%.2fσ)", cfg.K)
			item := violation(cfg.RuleCode, "spc.rule.msg.beyond", msg, cfg.Severity, []int{i}, map[string]any{"k": cfg.K})
			return &item
		}
	}
	return nil
}

func checkKOfNZone(values []float64, ucl, cl float64, cfg RuleConfig) []RuleViolation {
	out := []RuleViolation{}
	n, hits := cfg.N, cfg.Hits
	if n < 2 || hits < 1 || len(values) < n {
		return out
	}
	sig := sigmaOf(ucl, cl)
	upper := cl + cfg.K*sig
	lower := cl - cfg.K*sig
	for i := n - 1; i < len(values); i++ {
		countUpper, countLower := 0, 0
		pointsUpper, pointsLower := []int{}, []int{}
		for j := 0; j < n; j++ {
			idx := i - n + 1 + j
			v := values[idx]
			if v > upper {
				countUpper++
				pointsUpper = append(pointsUpper, idx)
			}
			if v < lower {
				countLower++
				pointsLower = append(pointsLower, idx)
			}
		}
		if countUpper >= hits {
			msg := fmt.Sprintf("%d点中有%d点超出+%.2fσ", n, hits, cfg.K)
			out = append(out, violation(cfg.RuleCode, "spc.rule.msg.zoneUpper", msg, cfg.Severity, pointsUpper, map[string]any{"n": n, "hits": hits, "k": cfg.K}))
		}
		if countLower >= hits {
			msg := fmt.Sprintf("%d点中有%d点低于-%.2fσ", n, hits, cfg.K)
			out = append(out, violation(cfg.RuleCode, "spc.rule.msg.zoneLower", msg, cfg.Severity, pointsLower, map[string]any{"n": n, "hits": hits, "k": cfg.K}))
		}
	}
	return out
}

func checkConsecutiveSide(values []float64, cl float64, cfg RuleConfig) *RuleViolation {
	n := cfg.N
	if n < 2 || len(values) < n {
		return nil
	}
	for i := n - 1; i < len(values); i++ {
		allAbove, allBelow := true, true
		for j := 0; j < n; j++ {
			v := values[i-n+1+j]
			if v <= cl {
				allAbove = false
			}
			if v >= cl {
				allBelow = false
			}
		}
		if allAbove || allBelow {
			points := make([]int, n)
			for j := 0; j < n; j++ {
				points[j] = i - n + 1 + j
			}
			side := "above"
			sideZh := "上方"
			if allBelow {
				side = "below"
				sideZh = "下方"
			}
			msg := fmt.Sprintf("连续%d点在中心线%s", n, sideZh)
			item := violation(cfg.RuleCode, "spc.rule.msg.side", msg, cfg.Severity, points, map[string]any{"n": n, "side": side})
			return &item
		}
	}
	return nil
}

func checkConsecutiveTrend(values []float64, cfg RuleConfig) []RuleViolation {
	out := []RuleViolation{}
	n := cfg.N
	if n < 3 || len(values) < n {
		return out
	}
	for i := n - 1; i < len(values); i++ {
		increasing, decreasing := true, true
		for j := 1; j < n; j++ {
			prev := values[i-n+j]
			cur := values[i-n+1+j]
			if cur <= prev {
				increasing = false
			}
			if cur >= prev {
				decreasing = false
			}
		}
		if increasing || decreasing {
			points := make([]int, n)
			for j := 0; j < n; j++ {
				points[j] = i - n + 1 + j
			}
			trend := "up"
			trendZh := "递增"
			if decreasing {
				trend = "down"
				trendZh = "递减"
			}
			msg := fmt.Sprintf("连续%d点%s", n, trendZh)
			out = append(out, violation(cfg.RuleCode, "spc.rule.msg.trend", msg, cfg.Severity, points, map[string]any{"n": n, "trend": trend}))
		}
	}
	return out
}

func checkConsecutiveAlternate(values []float64, cfg RuleConfig) *RuleViolation {
	n := cfg.N
	if n < 3 || len(values) < n {
		return nil
	}
	for i := n - 1; i < len(values); i++ {
		alternating := true
		for j := 2; j < n; j++ {
			a := values[i-n+1+j]
			b := values[i-n+j]
			c := values[i-n-1+j]
			if (a-b)*(b-c) >= 0 {
				alternating = false
				break
			}
		}
		if alternating {
			points := make([]int, n)
			for j := 0; j < n; j++ {
				points[j] = i - n + 1 + j
			}
			msg := fmt.Sprintf("连续%d点交替上下", n)
			item := violation(cfg.RuleCode, "spc.rule.msg.alternate", msg, cfg.Severity, points, map[string]any{"n": n})
			return &item
		}
	}
	return nil
}

func checkConsecutiveWithin(values []float64, ucl, cl float64, cfg RuleConfig) *RuleViolation {
	n := cfg.N
	if n < 2 || len(values) < n {
		return nil
	}
	sig := sigmaOf(ucl, cl)
	upper := cl + cfg.K*sig
	lower := cl - cfg.K*sig
	for i := n - 1; i < len(values); i++ {
		allWithin := true
		for j := 0; j < n; j++ {
			v := values[i-n+1+j]
			if v > upper || v < lower {
				allWithin = false
				break
			}
		}
		if allWithin {
			points := make([]int, n)
			for j := 0; j < n; j++ {
				points[j] = i - n + 1 + j
			}
			msg := fmt.Sprintf("连续%d点在中心线±%.2fσ内", n, cfg.K)
			item := violation(cfg.RuleCode, "spc.rule.msg.within", msg, cfg.Severity, points, map[string]any{"n": n, "k": cfg.K})
			return &item
		}
	}
	return nil
}

func checkConsecutiveBeyondBoth(values []float64, ucl, cl float64, cfg RuleConfig) *RuleViolation {
	n := cfg.N
	if n < 2 || len(values) < n {
		return nil
	}
	sig := sigmaOf(ucl, cl)
	for i := n - 1; i < len(values); i++ {
		allOutside := true
		for j := 0; j < n; j++ {
			if math.Abs(values[i-n+1+j]-cl) <= cfg.K*sig {
				allOutside = false
				break
			}
		}
		if allOutside {
			points := make([]int, n)
			for j := 0; j < n; j++ {
				points[j] = i - n + 1 + j
			}
			msg := fmt.Sprintf("连续%d点距离中心线>%.2fσ(两侧)", n, cfg.K)
			item := violation(cfg.RuleCode, "spc.rule.msg.beyondBoth", msg, cfg.Severity, points, map[string]any{"n": n, "k": cfg.K})
			return &item
		}
	}
	return nil
}

// 兼容旧单测的包装函数
func checkWE1(values []float64, ucl, lcl float64) *RuleViolation {
	cl := (ucl + lcl) / 2
	return checkBeyondLimit(values, ucl, cl, lcl, DefaultRuleConfig("WE1"))
}

func checkWE2(values []float64, ucl, cl, lcl float64) []RuleViolation {
	return checkKOfNZone(values, ucl, cl, DefaultRuleConfig("WE2"))
}

func checkWE3(values []float64, ucl, cl, lcl float64) []RuleViolation {
	return checkKOfNZone(values, ucl, cl, DefaultRuleConfig("WE3"))
}

func checkWE4(values []float64, cl float64) *RuleViolation {
	return checkConsecutiveSide(values, cl, DefaultRuleConfig("WE4"))
}

func checkNELSON2(values []float64, cl float64) *RuleViolation {
	return checkConsecutiveSide(values, cl, DefaultRuleConfig("NELSON2"))
}

func checkNELSON3(values []float64) []RuleViolation {
	return checkConsecutiveTrend(values, DefaultRuleConfig("NELSON3"))
}

func checkNELSON5(values []float64) []RuleViolation {
	// 旧实现将 NELSON5 视为趋势规则，单测仍覆盖该包装
	return checkConsecutiveTrend(values, DefaultRuleConfig("NELSON3"))
}

func checkNELSON4(values []float64) *RuleViolation {
	return checkConsecutiveAlternate(values, DefaultRuleConfig("NELSON4"))
}

func checkNELSON6(values []float64, ucl, cl, lcl float64) []RuleViolation {
	return checkKOfNZone(values, ucl, cl, DefaultRuleConfig("NELSON6"))
}

func checkNELSON7(values []float64, ucl, cl, lcl float64) *RuleViolation {
	return checkConsecutiveWithin(values, ucl, cl, DefaultRuleConfig("NELSON7"))
}

func checkNELSON8(values []float64, ucl, cl, lcl float64) *RuleViolation {
	return checkConsecutiveBeyondBoth(values, ucl, cl, DefaultRuleConfig("NELSON8"))
}
