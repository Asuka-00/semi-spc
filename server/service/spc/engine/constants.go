package engine

import "math"

// SPC控制图常数表
// 参考: Montgomery, D.C. (2012). Statistical Quality Control. 7th Edition.

// ControlChartConstants SPC控制图常数
type ControlChartConstants struct {
	N  int     // 子组大小
	A2 float64 // X̄图控制限因子(用R)
	D3 float64 // R图下控制限因子
	D4 float64 // R图上控制限因子
	A3 float64 // X̄图控制限因子(用S)
	B3 float64 // S图下控制限因子
	B4 float64 // S图上控制限因子
	d2 float64 // 极差均值因子
	E2 float64 // I-MR图MR控制限因子
}

// GetConstants 获取控制图常数 (n=2到10)
func GetConstants(n int) *ControlChartConstants {
	constants := map[int]*ControlChartConstants{
		2:  {N: 2, A2: 1.880, D3: 0.000, D4: 3.267, A3: 2.659, B3: 0.000, B4: 3.267, d2: 1.128, E2: 2.660},
		3:  {N: 3, A2: 1.023, D3: 0.000, D4: 2.574, A3: 1.954, B3: 0.000, B4: 2.568, d2: 1.693, E2: 0.000},
		4:  {N: 4, A2: 0.729, D3: 0.000, D4: 2.282, A3: 1.628, B3: 0.000, B4: 2.266, d2: 2.059, E2: 0.000},
		5:  {N: 5, A2: 0.577, D3: 0.000, D4: 2.114, A3: 1.427, B3: 0.000, B4: 2.089, d2: 2.326, E2: 0.000},
		6:  {N: 6, A2: 0.483, D3: 0.000, D4: 2.004, A3: 1.287, B3: 0.030, B4: 1.970, d2: 2.534, E2: 0.000},
		7:  {N: 7, A2: 0.419, D3: 0.076, D4: 1.924, A3: 1.182, B3: 0.118, B4: 1.882, d2: 2.704, E2: 0.000},
		8:  {N: 8, A2: 0.373, D3: 0.136, D4: 1.864, A3: 1.099, B3: 0.185, B4: 1.815, d2: 2.847, E2: 0.000},
		9:  {N: 9, A2: 0.337, D3: 0.184, D4: 1.816, A3: 1.032, B3: 0.239, B4: 1.761, d2: 2.970, E2: 0.000},
		10: {N: 10, A2: 0.308, D3: 0.223, D4: 1.777, A3: 0.975, B3: 0.284, B4: 1.716, d2: 3.078, E2: 0.000},
	}

	if c, ok := constants[n]; ok {
		return c
	}

	if n > 10 {
		return &ControlChartConstants{
			N:  n,
			A2: 3.0 / (math.Sqrt(float64(n)) * math.Sqrt(float64(n))),
			D3: math.Max(0, 1.0-3.0/(math.Sqrt(float64(n))*math.Sqrt(2.0))),
			D4: 1.0 + 3.0/(math.Sqrt(float64(n))*math.Sqrt(2.0)),
			A3: 3.0 / (math.Sqrt(float64(n)) * math.Sqrt(2.0)),
			B3: math.Max(0, 1.0-3.0/math.Sqrt(float64(2*n-2))),
			B4: 1.0 + 3.0/math.Sqrt(float64(2*n-2)),
			d2: 1.0 + 0.8862*math.Log(float64(n)),
			E2: 0.000,
		}
	}

	return nil
}

// XbarRLimits 计算X̄-R图控制限
func XbarRLimits(xBarBar, rBar float64, n int) (uclX, clX, lclX, uclR, clR, lclR float64) {
	c := GetConstants(n)
	if c == nil {
		return 0, 0, 0, 0, 0, 0
	}

	clX = xBarBar
	uclX = xBarBar + c.A2*rBar
	lclX = xBarBar - c.A2*rBar

	clR = rBar
	uclR = c.D4 * rBar
	lclR = c.D3 * rBar

	return
}

// XbarSLimits 计算X̄-S图控制限
func XbarSLimits(xBarBar, sBar float64, n int) (uclX, clX, lclX, uclS, clS, lclS float64) {
	c := GetConstants(n)
	if c == nil {
		return 0, 0, 0, 0, 0, 0
	}

	clX = xBarBar
	uclX = xBarBar + c.A3*sBar
	lclX = xBarBar - c.A3*sBar

	clS = sBar
	uclS = c.B4 * sBar
	lclS = c.B3 * sBar

	return
}

// IMRLimits 计算I-MR图控制限
func IMRLimits(xBar, mrBar float64) (uclI, clI, lclI, uclMR, clMR, lclMR float64) {
	c := GetConstants(2)
	if c == nil {
		return 0, 0, 0, 0, 0, 0
	}

	clI = xBar
	uclI = xBar + c.E2*mrBar
	lclI = xBar - c.E2*mrBar

	clMR = mrBar
	uclMR = c.D4 * mrBar
	lclMR = c.D3 * mrBar

	return
}

// CalculateMean 计算均值
func CalculateMean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// CalculateRange 计算极差
func CalculateRange(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	min, max := values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

// CalculateStdDev 计算标准差
func CalculateStdDev(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	mean := CalculateMean(values)
	sumSq := 0.0
	for _, v := range values {
		diff := v - mean
		sumSq += diff * diff
	}
	return math.Sqrt(sumSq / float64(len(values)-1))
}

// CalculateMovingRange 计算移动极差
func CalculateMovingRange(values []float64) []float64 {
	if len(values) < 2 {
		return []float64{}
	}
	mrs := make([]float64, len(values)-1)
	for i := 1; i < len(values); i++ {
		mrs[i-1] = math.Abs(values[i] - values[i-1])
	}
	return mrs
}

// EstimateSigmaFromR 从R估计σ
func EstimateSigmaFromR(rBar float64, n int) float64 {
	c := GetConstants(n)
	if c == nil {
		return 0
	}
	return rBar / c.d2
}

// EstimateSigmaFromS 从S估计σ
func EstimateSigmaFromS(sBar float64, n int) float64 {
	c := GetConstants(n)
	if c == nil {
		return 0
	}
	c4 := math.Sqrt(2.0/float64(n-1)) * math.Gamma(float64(n)/2.0) / math.Gamma(float64(n-1)/2.0)
	return sBar / c4
}

// EstimateSigmaFromMR 从移动极差估计σ (I-MR图)
func EstimateSigmaFromMR(mrBar float64) float64 {
	return mrBar / 1.128
}
