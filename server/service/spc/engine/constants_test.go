package engine

import (
	"math"
	"testing"
)

func TestGetConstants(t *testing.T) {
	c := GetConstants(5)
	if c == nil {
		t.Fatal("GetConstants(5) returned nil")
	}
	if c.A2 != 0.577 {
		t.Errorf("Expected A2=0.577, got %f", c.A2)
	}
	if c.D4 != 2.114 {
		t.Errorf("Expected D4=2.114, got %f", c.D4)
	}
	if c.d2 != 2.326 {
		t.Errorf("Expected d2=2.326, got %f", c.d2)
	}
}

func TestXbarRLimits(t *testing.T) {
	xBarBar := 100.0
	rBar := 5.0
	n := 5

	uclX, clX, lclX, uclR, clR, _ := XbarRLimits(xBarBar, rBar, n)

	if clX != 100.0 {
		t.Errorf("Expected CL_X=100.0, got %f", clX)
	}

	expectedUCLX := 100.0 + 0.577*5.0
	if math.Abs(uclX-expectedUCLX) > 0.001 {
		t.Errorf("Expected UCL_X=%f, got %f", expectedUCLX, uclX)
	}

	expectedLCLX := 100.0 - 0.577*5.0
	if math.Abs(lclX-expectedLCLX) > 0.001 {
		t.Errorf("Expected LCL_X=%f, got %f", expectedLCLX, lclX)
	}

	if clR != 5.0 {
		t.Errorf("Expected CL_R=5.0, got %f", clR)
	}

	expectedUCLR := 2.114 * 5.0
	if math.Abs(uclR-expectedUCLR) > 0.001 {
		t.Errorf("Expected UCL_R=%f, got %f", expectedUCLR, uclR)
	}
}

func TestCalculateMean(t *testing.T) {
	values := []float64{10.0, 20.0, 30.0}
	mean := CalculateMean(values)
	if math.Abs(mean-20.0) > 0.001 {
		t.Errorf("Expected mean=20.0, got %f", mean)
	}
}

func TestCalculateRange(t *testing.T) {
	values := []float64{10.0, 15.0, 25.0, 20.0}
	r := CalculateRange(values)
	if math.Abs(r-15.0) > 0.001 {
		t.Errorf("Expected range=15.0, got %f", r)
	}
}

func TestCalculateStdDev(t *testing.T) {
	values := []float64{2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0}
	std := CalculateStdDev(values)
	expectedStd := 2.138
	if math.Abs(std-expectedStd) > 0.01 {
		t.Errorf("Expected std≈%f, got %f", expectedStd, std)
	}
}

func TestCalculateMovingRange(t *testing.T) {
	values := []float64{10.0, 12.0, 15.0, 13.0}
	mrs := CalculateMovingRange(values)
	if len(mrs) != 3 {
		t.Errorf("Expected 3 moving ranges, got %d", len(mrs))
	}
	if math.Abs(mrs[0]-2.0) > 0.001 {
		t.Errorf("Expected MR[0]=2.0, got %f", mrs[0])
	}
	if math.Abs(mrs[1]-3.0) > 0.001 {
		t.Errorf("Expected MR[1]=3.0, got %f", mrs[1])
	}
	if math.Abs(mrs[2]-2.0) > 0.001 {
		t.Errorf("Expected MR[2]=2.0, got %f", mrs[2])
	}
}

func TestEstimateSigmaFromR(t *testing.T) {
	rBar := 5.0
	n := 5
	sigma := EstimateSigmaFromR(rBar, n)
	c := GetConstants(n)
	expectedSigma := rBar / c.d2
	if math.Abs(sigma-expectedSigma) > 0.001 {
		t.Errorf("Expected sigma=%f, got %f", expectedSigma, sigma)
	}
}

func TestEstimateSigmaFromMR(t *testing.T) {
	mrBar := 2.26
	sigma := EstimateSigmaFromMR(mrBar)
	expectedSigma := 2.26 / 1.128
	if math.Abs(sigma-expectedSigma) > 0.01 {
		t.Errorf("Expected sigma≈%f, got %f", expectedSigma, sigma)
	}
}
