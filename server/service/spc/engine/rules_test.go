package engine

import (
	"testing"
)

func TestCheckOOS(t *testing.T) {
	usl := 110.0
	lsl := 90.0

	oos, msg := CheckOOS(115.0, &usl, &lsl)
	if !oos {
		t.Error("Expected OOS for value above USL")
	}
	if msg != "超上规格限 USL" {
		t.Errorf("Expected USL message, got: %s", msg)
	}

	oos, msg = CheckOOS(85.0, &usl, &lsl)
	if !oos {
		t.Error("Expected OOS for value below LSL")
	}
	if msg != "超下规格限 LSL" {
		t.Errorf("Expected LSL message, got: %s", msg)
	}

	oos, _ = CheckOOS(100.0, &usl, &lsl)
	if oos {
		t.Error("Expected no OOS for value within spec")
	}
}

func TestCheckWE1(t *testing.T) {
	ucl := 110.0
	lcl := 90.0

	values := []float64{100.0, 105.0, 115.0, 102.0}
	v := checkWE1(values, ucl, lcl)
	if v == nil {
		t.Error("Expected WE1 violation")
	}
	if v.Points[0] != 2 {
		t.Errorf("Expected point index 2, got %d", v.Points[0])
	}

	values = []float64{100.0, 105.0, 108.0, 102.0}
	v = checkWE1(values, ucl, lcl)
	if v != nil {
		t.Error("Expected no WE1 violation")
	}
}

func TestCheckWE4(t *testing.T) {
	cl := 100.0

	values := []float64{101.0, 102.0, 103.0, 104.0, 105.0, 106.0, 107.0, 108.0}
	v := checkWE4(values, cl)
	if v == nil {
		t.Error("Expected WE4 violation for 8 points above CL")
	}
	if v.RuleCode != "WE4" {
		t.Errorf("Expected rule code WE4, got %s", v.RuleCode)
	}

	values = []float64{101.0, 102.0, 99.0, 104.0, 105.0, 106.0, 107.0, 108.0}
	v = checkWE4(values, cl)
	if v != nil {
		t.Error("Expected no WE4 violation")
	}
}

func TestCheckNELSON3(t *testing.T) {
	values := []float64{100.0, 101.0, 102.0, 103.0, 104.0, 105.0}
	violations := checkNELSON3(values)
	if len(violations) == 0 {
		t.Error("Expected NELSON3 violation for 6 increasing points")
	}
	if violations[0].RuleCode != "NELSON3" {
		t.Errorf("Expected rule code NELSON3, got %s", violations[0].RuleCode)
	}

	values = []float64{105.0, 104.0, 103.0, 102.0, 101.0, 100.0}
	violations = checkNELSON3(values)
	if len(violations) == 0 {
		t.Error("Expected NELSON3 violation for 6 decreasing points")
	}

	values = []float64{100.0, 101.0, 102.0, 101.0, 104.0, 105.0}
	violations = checkNELSON3(values)
	if len(violations) > 0 {
		t.Error("Expected no NELSON3 violation")
	}
}

func TestConfigurableRuleN(t *testing.T) {
	cl := 100.0
	values := []float64{101, 102, 103, 104}
	cfg := RuleConfig{RuleCode: "WE4", N: 4, Severity: "WARN"}
	v := checkConsecutiveSide(values, cl, cfg.withDefaults())
	if v == nil {
		t.Fatal("Expected WE4 with N=4 to fire")
	}
	if len(v.Points) != 4 {
		t.Errorf("Expected 4 points, got %d", len(v.Points))
	}

	cfg.N = 8
	v = checkConsecutiveSide(values, cl, cfg.withDefaults())
	if v != nil {
		t.Error("Expected no violation when N=8 but only 4 points")
	}
}

func TestConfigurableKOfN(t *testing.T) {
	ucl, cl := 110.0, 100.0
	values := []float64{107, 108, 101}
	cfg := RuleConfig{RuleCode: "WE2", N: 3, Hits: 2, K: 2, Severity: "WARN"}
	vs := checkKOfNZone(values, ucl, cl, cfg.withDefaults())
	if len(vs) == 0 {
		t.Fatal("Expected WE2 with custom 2-of-3 beyond 2σ")
	}
}

func TestCheckOOC(t *testing.T) {
	ucl := 110.0
	cl := 100.0
	lcl := 90.0

	values := []float64{100.0, 105.0, 115.0}
	enabledRules := []string{"WE1"}
	violations := CheckOOC(values, ucl, cl, lcl, enabledRules)
	if len(violations) == 0 {
		t.Error("Expected OOC violations")
	}
	if violations[0].RuleCode != "WE1" {
		t.Errorf("Expected WE1 violation, got %s", violations[0].RuleCode)
	}

	values = []float64{101.0, 102.0, 103.0, 104.0, 105.0, 106.0, 107.0, 108.0}
	enabledRules = []string{"WE4"}
	violations = CheckOOC(values, ucl, cl, lcl, enabledRules)
	if len(violations) == 0 {
		t.Error("Expected WE4 violation")
	}
}

func TestCheckNELSON2(t *testing.T) {
	cl := 100.0
	values := []float64{101.0, 102.0, 103.0, 104.0, 105.0, 106.0, 107.0, 108.0, 109.0}
	v := checkNELSON2(values, cl)
	if v == nil {
		t.Error("Expected NELSON2 violation for 9 points above CL")
	}

	values = []float64{101.0, 102.0, 99.0, 104.0, 105.0, 106.0, 107.0, 108.0, 109.0}
	v = checkNELSON2(values, cl)
	if v != nil {
		t.Error("Expected no NELSON2 violation")
	}
}

func TestCheckNELSON4(t *testing.T) {
	values := []float64{1.0, 2.0, 1.5, 2.5, 2.0, 3.0, 2.5, 3.5, 3.0, 4.0, 3.5, 4.5, 4.0, 5.0}
	v := checkNELSON4(values)
	if v == nil {
		t.Error("Expected NELSON4 violation for 14 alternating points")
	}

	values = []float64{1.0, 2.0, 3.0, 2.0, 3.0, 4.0, 3.0, 4.0, 5.0, 4.0, 5.0, 6.0, 5.0, 6.0}
	v = checkNELSON4(values)
	if v != nil {
		t.Error("Expected no NELSON4 violation")
	}
}
