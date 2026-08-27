package engine

import "math"

// RuleViolation 规则违反记录
type RuleViolation struct {
	RuleCode string  // 规则代码 WE1/NELSON1等
	Message  string  // 违反描述
	Points   []int   // 违反点的索引
	Severity string  // 严重度 INFO/WARN/CRIT
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

// CheckOOC 检查是否失控 (Out Of Control)
// 使用Western Electric和Nelson规则
func CheckOOC(values []float64, ucl, cl, lcl float64, enabledRules []string) []RuleViolation {
	violations := []RuleViolation{}

	if len(values) == 0 {
		return violations
	}

	ruleMap := make(map[string]bool)
	for _, r := range enabledRules {
		ruleMap[r] = true
	}

	// WE1 / NELSON1: 点超出控制限
	if ruleMap["WE1"] || ruleMap["NELSON1"] {
		v := checkWE1(values, ucl, lcl)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	// WE2: 3点中有2点在同侧A区外(2σ-3σ)
	if ruleMap["WE2"] {
		v := checkWE2(values, ucl, cl, lcl)
		violations = append(violations, v...)
	}

	// WE3: 5点中有4点在同侧C区外(1σ-3σ)
	if ruleMap["WE3"] {
		v := checkWE3(values, ucl, cl, lcl)
		violations = append(violations, v...)
	}

	// WE4: 连续8点在中心线同一侧
	if ruleMap["WE4"] {
		v := checkWE4(values, cl)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	// NELSON5: 连续6点递增或递减
	if ruleMap["NELSON5"] {
		v := checkNELSON5(values)
		violations = append(violations, v...)
	}

	// NELSON2: 连续9点在中心线同一侧(与WE4类似但更严格)
	if ruleMap["NELSON2"] {
		v := checkNELSON2(values, cl)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	// NELSON3: 连续6点递增或递减(与NELSON5相同)
	if ruleMap["NELSON3"] {
		v := checkNELSON3(values)
		violations = append(violations, v...)
	}

	// NELSON4: 连续14点交替上下
	if ruleMap["NELSON4"] {
		v := checkNELSON4(values)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	// NELSON6: 5点中有4点距离中心线>1σ同侧
	if ruleMap["NELSON6"] {
		v := checkNELSON6(values, ucl, cl, lcl)
		violations = append(violations, v...)
	}

	// NELSON7: 连续15点在中心线±1σ内
	if ruleMap["NELSON7"] {
		v := checkNELSON7(values, ucl, cl, lcl)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	// NELSON8: 连续8点距离中心线>1σ且在两侧
	if ruleMap["NELSON8"] {
		v := checkNELSON8(values, ucl, cl, lcl)
		if v != nil {
			violations = append(violations, *v)
		}
	}

	return violations
}

// WE1 / NELSON1: 点超出控制限 (Beyond limits)
func checkWE1(values []float64, ucl, lcl float64) *RuleViolation {
	for i, v := range values {
		if v > ucl || v < lcl {
			return &RuleViolation{
				RuleCode: "WE1",
				Message:  "点超出控制限",
				Points:   []int{i},
				Severity: "CRIT",
			}
		}
	}
	return nil
}

// WE2: 3点中有2点在同侧A区外(2σ-3σ)
func checkWE2(values []float64, ucl, cl, lcl float64) []RuleViolation {
	violations := []RuleViolation{}
	sigma := (ucl - cl) / 3.0

	zoneAUpperLower := cl + 2*sigma
	zoneALowerUpper := cl - 2*sigma

	for i := 2; i < len(values); i++ {
		window := values[i-2 : i+1]
		countUpper := 0
		countLower := 0
		pointsUpper := []int{}
		pointsLower := []int{}

		for j, v := range window {
			idx := i - 2 + j
			if v > zoneAUpperLower {
				countUpper++
				pointsUpper = append(pointsUpper, idx)
			}
			if v < zoneALowerUpper {
				countLower++
				pointsLower = append(pointsLower, idx)
			}
		}

		if countUpper >= 2 {
			violations = append(violations, RuleViolation{
				RuleCode: "WE2",
				Message:  "3点中有2点超出+2σ(A区上界)",
				Points:   pointsUpper,
				Severity: "WARN",
			})
		}
		if countLower >= 2 {
			violations = append(violations, RuleViolation{
				RuleCode: "WE2",
				Message:  "3点中有2点低于-2σ(A区下界)",
				Points:   pointsLower,
				Severity: "WARN",
			})
		}
	}
	return violations
}

// WE3: 5点中有4点在同侧C区外(1σ-3σ)
func checkWE3(values []float64, ucl, cl, lcl float64) []RuleViolation {
	violations := []RuleViolation{}
	sigma := (ucl - cl) / 3.0

	zoneCUpperLower := cl + sigma
	zoneCLowerUpper := cl - sigma

	for i := 4; i < len(values); i++ {
		window := values[i-4 : i+1]
		countUpper := 0
		countLower := 0
		pointsUpper := []int{}
		pointsLower := []int{}

		for j, v := range window {
			idx := i - 4 + j
			if v > zoneCUpperLower {
				countUpper++
				pointsUpper = append(pointsUpper, idx)
			}
			if v < zoneCLowerUpper {
				countLower++
				pointsLower = append(pointsLower, idx)
			}
		}

		if countUpper >= 4 {
			violations = append(violations, RuleViolation{
				RuleCode: "WE3",
				Message:  "5点中有4点超出+1σ(C区上界)",
				Points:   pointsUpper,
				Severity: "WARN",
			})
		}
		if countLower >= 4 {
			violations = append(violations, RuleViolation{
				RuleCode: "WE3",
				Message:  "5点中有4点低于-1σ(C区下界)",
				Points:   pointsLower,
				Severity: "WARN",
			})
		}
	}
	return violations
}

// WE4: 连续8点在中心线同一侧
func checkWE4(values []float64, cl float64) *RuleViolation {
	if len(values) < 8 {
		return nil
	}

	for i := 7; i < len(values); i++ {
		window := values[i-7 : i+1]
		allAbove := true
		allBelow := true

		for _, v := range window {
			if v <= cl {
				allAbove = false
			}
			if v >= cl {
				allBelow = false
			}
		}

		if allAbove || allBelow {
			points := make([]int, 8)
			for j := 0; j < 8; j++ {
				points[j] = i - 7 + j
			}
			side := "上方"
			if allBelow {
				side = "下方"
			}
			return &RuleViolation{
				RuleCode: "WE4",
				Message:  "连续8点在中心线" + side,
				Points:   points,
				Severity: "WARN",
			}
		}
	}
	return nil
}

// NELSON2: 连续9点在中心线同一侧
func checkNELSON2(values []float64, cl float64) *RuleViolation {
	if len(values) < 9 {
		return nil
	}

	for i := 8; i < len(values); i++ {
		window := values[i-8 : i+1]
		allAbove := true
		allBelow := true

		for _, v := range window {
			if v <= cl {
				allAbove = false
			}
			if v >= cl {
				allBelow = false
			}
		}

		if allAbove || allBelow {
			points := make([]int, 9)
			for j := 0; j < 9; j++ {
				points[j] = i - 8 + j
			}
			side := "上方"
			if allBelow {
				side = "下方"
			}
			return &RuleViolation{
				RuleCode: "NELSON2",
				Message:  "连续9点在中心线" + side,
				Points:   points,
				Severity: "WARN",
			}
		}
	}
	return nil
}

// NELSON3 / NELSON5: 连续6点递增或递减
func checkNELSON3(values []float64) []RuleViolation {
	return checkNELSON5(values)
}

func checkNELSON5(values []float64) []RuleViolation {
	violations := []RuleViolation{}
	if len(values) < 6 {
		return violations
	}

	for i := 5; i < len(values); i++ {
		window := values[i-5 : i+1]
		increasing := true
		decreasing := true

		for j := 1; j < len(window); j++ {
			if window[j] <= window[j-1] {
				increasing = false
			}
			if window[j] >= window[j-1] {
				decreasing = false
			}
		}

		if increasing || decreasing {
			points := make([]int, 6)
			for j := 0; j < 6; j++ {
				points[j] = i - 5 + j
			}
			trend := "递增"
			if decreasing {
				trend = "递减"
			}
			violations = append(violations, RuleViolation{
				RuleCode: "NELSON5",
				Message:  "连续6点" + trend,
				Points:   points,
				Severity: "WARN",
			})
		}
	}
	return violations
}

// NELSON4: 连续14点交替上下
func checkNELSON4(values []float64) *RuleViolation {
	if len(values) < 14 {
		return nil
	}

	for i := 13; i < len(values); i++ {
		window := values[i-13 : i+1]
		alternating := true

		for j := 2; j < len(window); j++ {
			if (window[j]-window[j-1])*(window[j-1]-window[j-2]) >= 0 {
				alternating = false
				break
			}
		}

		if alternating {
			points := make([]int, 14)
			for j := 0; j < 14; j++ {
				points[j] = i - 13 + j
			}
			return &RuleViolation{
				RuleCode: "NELSON4",
				Message:  "连续14点交替上下",
				Points:   points,
				Severity: "INFO",
			}
		}
	}
	return nil
}

// NELSON6: 5点中有4点距离中心线>1σ同侧
func checkNELSON6(values []float64, ucl, cl, lcl float64) []RuleViolation {
	violations := []RuleViolation{}
	sigma := (ucl - cl) / 3.0

	for i := 4; i < len(values); i++ {
		window := values[i-4 : i+1]
		countUpper := 0
		countLower := 0
		pointsUpper := []int{}
		pointsLower := []int{}

		for j, v := range window {
			idx := i - 4 + j
			if v > cl+sigma {
				countUpper++
				pointsUpper = append(pointsUpper, idx)
			}
			if v < cl-sigma {
				countLower++
				pointsLower = append(pointsLower, idx)
			}
		}

		if countUpper >= 4 {
			violations = append(violations, RuleViolation{
				RuleCode: "NELSON6",
				Message:  "5点中有4点距离中心线>1σ(上方)",
				Points:   pointsUpper,
				Severity: "WARN",
			})
		}
		if countLower >= 4 {
			violations = append(violations, RuleViolation{
				RuleCode: "NELSON6",
				Message:  "5点中有4点距离中心线>1σ(下方)",
				Points:   pointsLower,
				Severity: "WARN",
			})
		}
	}
	return violations
}

// NELSON7: 连续15点在中心线±1σ内
func checkNELSON7(values []float64, ucl, cl, lcl float64) *RuleViolation {
	if len(values) < 15 {
		return nil
	}

	sigma := (ucl - cl) / 3.0
	upperBound := cl + sigma
	lowerBound := cl - sigma

	for i := 14; i < len(values); i++ {
		window := values[i-14 : i+1]
		allWithin := true

		for _, v := range window {
			if v > upperBound || v < lowerBound {
				allWithin = false
				break
			}
		}

		if allWithin {
			points := make([]int, 15)
			for j := 0; j < 15; j++ {
				points[j] = i - 14 + j
			}
			return &RuleViolation{
				RuleCode: "NELSON7",
				Message:  "连续15点在中心线±1σ内",
				Points:   points,
				Severity: "INFO",
			}
		}
	}
	return nil
}

// NELSON8: 连续8点距离中心线>1σ且在两侧
func checkNELSON8(values []float64, ucl, cl, lcl float64) *RuleViolation {
	if len(values) < 8 {
		return nil
	}

	sigma := (ucl - cl) / 3.0

	for i := 7; i < len(values); i++ {
		window := values[i-7 : i+1]
		allOutside := true

		for _, v := range window {
			if math.Abs(v-cl) <= sigma {
				allOutside = false
				break
			}
		}

		if allOutside {
			points := make([]int, 8)
			for j := 0; j < 8; j++ {
				points[j] = i - 7 + j
			}
			return &RuleViolation{
				RuleCode: "NELSON8",
				Message:  "连续8点距离中心线>1σ(两侧)",
				Points:   points,
				Severity: "WARN",
			}
		}
	}
	return nil
}
