package engine

import (
	"math"
	"testing"
)

func TestCalculateCapability(t *testing.T) {
	values := []float64{95.0, 100.0, 105.0, 98.0, 102.0, 97.0, 103.0, 99.0, 101.0, 96.0}
	usl := 110.0
	lsl := 90.0

	result := CalculateCapability(values, &usl, &lsl, nil, 0)
	if result == nil {
		t.Fatal("CalculateCapability returned nil")
	}

	if result.N != 10 {
		t.Errorf("Expected N=10, got %d", result.N)
	}

	expectedMean := 99.6
	if math.Abs(result.Mean-expectedMean) > 0.1 {
		t.Errorf("Expected mean≈%f, got %f", expectedMean, result.Mean)
	}

	if result.Cp <= 0 {
		t.Errorf("Expected Cp > 0, got %f", result.Cp)
	}

	if result.Cpk <= 0 {
		t.Errorf("Expected Cpk > 0, got %f", result.Cpk)
	}

	if result.Pp <= 0 {
		t.Errorf("Expected Pp > 0, got %f", result.Pp)
	}

	if result.Ppk <= 0 {
		t.Errorf("Expected Ppk > 0, got %f", result.Ppk)
	}

	if result.Cpk > result.Cp {
		t.Error("Cpk should not exceed Cp")
	}

	t.Logf("Mean=%f, Sigma=%f, Cp=%f, Cpk=%f, Pp=%f, Ppk=%f",
		result.Mean, result.Sigma, result.Cp, result.Cpk, result.Pp, result.Ppk)
}

func TestCalculateCapabilityOneSided(t *testing.T) {
	values := []float64{95.0, 100.0, 105.0, 98.0, 102.0}
	usl := 110.0

	result := CalculateCapability(values, &usl, nil, nil, 0)
	if result == nil {
		t.Fatal("CalculateCapability returned nil for one-sided spec")
	}

	if result.CPU <= 0 {
		t.Errorf("Expected CPU > 0, got %f", result.CPU)
	}

	if result.Cpk != result.CPU {
		t.Errorf("Expected Cpk=CPU for upper-only spec, got Cpk=%f, CPU=%f", result.Cpk, result.CPU)
	}

	t.Logf("Mean=%f, Sigma=%f, CPU=%f, Cpk=%f", result.Mean, result.Sigma, result.CPU, result.Cpk)
}

func TestCalculateCapabilityFromSubgroups(t *testing.T) {
	subgroupMeans := []float64{100.0, 102.0, 98.0, 101.0, 99.0}
	subgroupRanges := []float64{5.0, 6.0, 4.0, 5.5, 4.5}
	n := 5
	usl := 110.0
	lsl := 90.0

	result := CalculateCapabilityFromSubgroups(subgroupMeans, subgroupRanges, nil, n, &usl, &lsl, nil)
	if result == nil {
		t.Fatal("CalculateCapabilityFromSubgroups returned nil")
	}

	if result.N != 5 {
		t.Errorf("Expected N=5, got %d", result.N)
	}

	if result.SigmaWT <= 0 {
		t.Errorf("Expected SigmaWT > 0, got %f", result.SigmaWT)
	}

	if result.SigmaOV <= 0 {
		t.Errorf("Expected SigmaOV > 0, got %f", result.SigmaOV)
	}

	if result.Cp <= 0 {
		t.Errorf("Expected Cp > 0, got %f", result.Cp)
	}

	t.Logf("Mean=%f, SigmaWT=%f, SigmaOV=%f, Cp=%f, Cpk=%f, Pp=%f, Ppk=%f",
		result.Mean, result.SigmaWT, result.SigmaOV, result.Cp, result.Cpk, result.Pp, result.Ppk)
}

func TestCalculateCapabilityPerfectCentering(t *testing.T) {
	values := []float64{100.0, 100.0, 100.0, 100.0, 100.0}
	usl := 110.0
	lsl := 90.0

	result := CalculateCapability(values, &usl, &lsl, nil, 1.0)
	if result == nil {
		t.Fatal("CalculateCapability returned nil")
	}

	if math.Abs(result.Mean-100.0) > 0.001 {
		t.Errorf("Expected mean=100.0, got %f", result.Mean)
	}

	expectedCp := (110.0 - 90.0) / (6.0 * 1.0)
	if math.Abs(result.Cp-expectedCp) > 0.01 {
		t.Errorf("Expected Cp≈%f, got %f", expectedCp, result.Cp)
	}

	if math.Abs(result.Cpk-result.Cp) > 0.01 {
		t.Error("For perfectly centered process, Cpk should equal Cp")
	}
}

func TestCalculateCapabilityOffCenter(t *testing.T) {
	values := []float64{95.0, 95.0, 95.0, 95.0, 95.0}
	usl := 110.0
	lsl := 90.0

	result := CalculateCapability(values, &usl, &lsl, nil, 2.0)
	if result == nil {
		t.Fatal("CalculateCapability returned nil")
	}

	if math.Abs(result.Mean-95.0) > 0.001 {
		t.Errorf("Expected mean=95.0, got %f", result.Mean)
	}

	if result.Cpk >= result.Cp {
		t.Error("For off-center process, Cpk should be less than Cp")
	}

	if result.CPL >= result.CPU {
		t.Error("For process closer to LSL, CPL should be less than CPU")
	}

	t.Logf("Mean=%f, Cp=%f, Cpk=%f, CPU=%f, CPL=%f",
		result.Mean, result.Cp, result.Cpk, result.CPU, result.CPL)
}
