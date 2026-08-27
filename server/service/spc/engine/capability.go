package engine

import "math"

// CapabilityResult 过程能力分析结果
type CapabilityResult struct {
	N       int     // 样本数
	Mean    float64 // 均值 μ
	Sigma   float64 // 标准差 σ
	Cp      float64 // 短期能力指数
	Cpk     float64 // 短期过程能力
	Pp      float64 // 长期能力指数
	Ppk     float64 // 长期过程能力
	CPU     float64 // 上单边能力指数
	CPL     float64 // 下单边能力指数
	PPU     float64 // 上单边长期能力指数
	PPL     float64 // 下单边长期能力指数
	USL     *float64
	LSL     *float64
	Target  *float64
	SigmaWT float64 // Within标准差(从R或S估计)
	SigmaOV float64 // Overall标准差(直接计算)
}

// CalculateCapability 计算过程能力指数
// values: 所有个体测量值或子组均值
// usl, lsl, target: 规格限
// withinSigma: Within标准差(从R-bar/d2或S-bar估计), 如果为0则从values计算
func CalculateCapability(values []float64, usl, lsl, target *float64, withinSigma float64) *CapabilityResult {
	if len(values) < 2 {
		return nil
	}

	result := &CapabilityResult{
		N:      len(values),
		Mean:   CalculateMean(values),
		USL:    usl,
		LSL:    lsl,
		Target: target,
	}

	// Overall标准差(总体变异)
	result.SigmaOV = CalculateStdDev(values)

	// Within标准差(组内变异)
	if withinSigma > 0 {
		result.SigmaWT = withinSigma
	} else {
		result.SigmaWT = result.SigmaOV
	}

	result.Sigma = result.SigmaWT

	// 计算Cp和Cpk (使用Within标准差)
	if usl != nil && lsl != nil {
		result.Cp = (*usl - *lsl) / (6.0 * result.SigmaWT)

		cpu := (*usl - result.Mean) / (3.0 * result.SigmaWT)
		cpl := (result.Mean - *lsl) / (3.0 * result.SigmaWT)
		result.CPU = cpu
		result.CPL = cpl
		result.Cpk = math.Min(cpu, cpl)
	} else if usl != nil {
		result.CPU = (*usl - result.Mean) / (3.0 * result.SigmaWT)
		result.Cpk = result.CPU
	} else if lsl != nil {
		result.CPL = (result.Mean - *lsl) / (3.0 * result.SigmaWT)
		result.Cpk = result.CPL
	}

	// 计算Pp和Ppk (使用Overall标准差)
	if usl != nil && lsl != nil {
		result.Pp = (*usl - *lsl) / (6.0 * result.SigmaOV)

		ppu := (*usl - result.Mean) / (3.0 * result.SigmaOV)
		ppl := (result.Mean - *lsl) / (3.0 * result.SigmaOV)
		result.PPU = ppu
		result.PPL = ppl
		result.Ppk = math.Min(ppu, ppl)
	} else if usl != nil {
		result.PPU = (*usl - result.Mean) / (3.0 * result.SigmaOV)
		result.Ppk = result.PPU
	} else if lsl != nil {
		result.PPL = (result.Mean - *lsl) / (3.0 * result.SigmaOV)
		result.Ppk = result.PPL
	}

	return result
}

// CalculateCapabilityFromSubgroups 从子组数据计算过程能力
// subgroupMeans: 子组均值
// subgroupRanges: 子组极差 (用于估计σ)
// subgroupStds: 子组标准差 (用于估计σ, 与subgroupRanges二选一)
// n: 子组大小
func CalculateCapabilityFromSubgroups(
	subgroupMeans []float64,
	subgroupRanges []float64,
	subgroupStds []float64,
	n int,
	usl, lsl, target *float64,
) *CapabilityResult {
	if len(subgroupMeans) < 2 {
		return nil
	}

	// Overall标准差
	sigmaOV := CalculateStdDev(subgroupMeans)

	// Within标准差
	var sigmaWT float64
	if len(subgroupRanges) > 0 {
		rBar := CalculateMean(subgroupRanges)
		sigmaWT = EstimateSigmaFromR(rBar, n)
	} else if len(subgroupStds) > 0 {
		sBar := CalculateMean(subgroupStds)
		sigmaWT = EstimateSigmaFromS(sBar, n)
	} else {
		sigmaWT = sigmaOV
	}

	return CalculateCapability(subgroupMeans, usl, lsl, target, sigmaWT)
}
