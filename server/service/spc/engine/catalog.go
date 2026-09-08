package engine

// RuleKind 规则判定类型，决定使用哪些可配置参数
const (
	KindBeyondLimit        = "beyond_limit"         // 单点超出 Kσ
	KindKOfNZone           = "k_of_n_zone"          // N 点中 Hits 点超出 Kσ 同侧
	KindConsecutiveSide    = "consecutive_side"     // 连续 N 点在中心线同侧
	KindConsecutiveTrend   = "consecutive_trend"    // 连续 N 点递增或递减
	KindConsecutiveAlt     = "consecutive_alternate" // 连续 N 点交替上下
	KindConsecutiveWithin  = "consecutive_within"   // 连续 N 点落在 ±Kσ 内
	KindConsecutiveBeyond  = "consecutive_beyond_both" // 连续 N 点均在 ±Kσ 外（两侧）
)

// RuleConfig 单条可配置规则
type RuleConfig struct {
	RuleCode string  `json:"ruleCode"`
	N        int     `json:"n"`
	Hits     int     `json:"hits"`
	K        float64 `json:"k"`
	Severity string  `json:"severity"`
}

// RuleDefinition 规则目录（默认值 + 可配置项说明）
type RuleDefinition struct {
	Code     string   `json:"code"`
	Family   string   `json:"family"`
	Kind     string   `json:"kind"`
	N        int      `json:"n"`
	Hits     int      `json:"hits"`
	K        float64  `json:"k"`
	Severity string   `json:"severity"`
	Params   []string `json:"params"`
	NameKey  string   `json:"nameKey"`
	DescKey  string   `json:"descKey"`
}

var ruleCatalog = []RuleDefinition{
	{Code: "WE1", Family: "WE", Kind: KindBeyondLimit, N: 1, Hits: 1, K: 3, Severity: "CRIT", Params: []string{"k", "severity"}, NameKey: "spc.rule.WE1.name", DescKey: "spc.rule.WE1.desc"},
	{Code: "WE2", Family: "WE", Kind: KindKOfNZone, N: 3, Hits: 2, K: 2, Severity: "WARN", Params: []string{"n", "hits", "k", "severity"}, NameKey: "spc.rule.WE2.name", DescKey: "spc.rule.WE2.desc"},
	{Code: "WE3", Family: "WE", Kind: KindKOfNZone, N: 5, Hits: 4, K: 1, Severity: "WARN", Params: []string{"n", "hits", "k", "severity"}, NameKey: "spc.rule.WE3.name", DescKey: "spc.rule.WE3.desc"},
	{Code: "WE4", Family: "WE", Kind: KindConsecutiveSide, N: 8, Hits: 8, K: 0, Severity: "WARN", Params: []string{"n", "severity"}, NameKey: "spc.rule.WE4.name", DescKey: "spc.rule.WE4.desc"},
	{Code: "NELSON1", Family: "NELSON", Kind: KindBeyondLimit, N: 1, Hits: 1, K: 3, Severity: "CRIT", Params: []string{"k", "severity"}, NameKey: "spc.rule.NELSON1.name", DescKey: "spc.rule.NELSON1.desc"},
	{Code: "NELSON2", Family: "NELSON", Kind: KindConsecutiveSide, N: 9, Hits: 9, K: 0, Severity: "WARN", Params: []string{"n", "severity"}, NameKey: "spc.rule.NELSON2.name", DescKey: "spc.rule.NELSON2.desc"},
	{Code: "NELSON3", Family: "NELSON", Kind: KindConsecutiveTrend, N: 6, Hits: 6, K: 0, Severity: "WARN", Params: []string{"n", "severity"}, NameKey: "spc.rule.NELSON3.name", DescKey: "spc.rule.NELSON3.desc"},
	{Code: "NELSON4", Family: "NELSON", Kind: KindConsecutiveAlt, N: 14, Hits: 14, K: 0, Severity: "INFO", Params: []string{"n", "severity"}, NameKey: "spc.rule.NELSON4.name", DescKey: "spc.rule.NELSON4.desc"},
	{Code: "NELSON5", Family: "NELSON", Kind: KindKOfNZone, N: 3, Hits: 2, K: 2, Severity: "WARN", Params: []string{"n", "hits", "k", "severity"}, NameKey: "spc.rule.NELSON5.name", DescKey: "spc.rule.NELSON5.desc"},
	{Code: "NELSON6", Family: "NELSON", Kind: KindKOfNZone, N: 5, Hits: 4, K: 1, Severity: "WARN", Params: []string{"n", "hits", "k", "severity"}, NameKey: "spc.rule.NELSON6.name", DescKey: "spc.rule.NELSON6.desc"},
	{Code: "NELSON7", Family: "NELSON", Kind: KindConsecutiveWithin, N: 15, Hits: 15, K: 1, Severity: "INFO", Params: []string{"n", "k", "severity"}, NameKey: "spc.rule.NELSON7.name", DescKey: "spc.rule.NELSON7.desc"},
	{Code: "NELSON8", Family: "NELSON", Kind: KindConsecutiveBeyond, N: 8, Hits: 8, K: 1, Severity: "WARN", Params: []string{"n", "k", "severity"}, NameKey: "spc.rule.NELSON8.name", DescKey: "spc.rule.NELSON8.desc"},
}

var catalogIndex map[string]RuleDefinition

func init() {
	catalogIndex = make(map[string]RuleDefinition, len(ruleCatalog))
	for _, item := range ruleCatalog {
		catalogIndex[item.Code] = item
	}
}

// RuleCatalog 返回全部规则定义
func RuleCatalog() []RuleDefinition {
	out := make([]RuleDefinition, len(ruleCatalog))
	copy(out, ruleCatalog)
	return out
}

// DefaultRuleConfig 按规则代码返回默认配置
func DefaultRuleConfig(code string) RuleConfig {
	if def, ok := catalogIndex[code]; ok {
		return RuleConfig{RuleCode: code, N: def.N, Hits: def.Hits, K: def.K, Severity: def.Severity}
	}
	return RuleConfig{RuleCode: code, N: 1, Hits: 1, K: 3, Severity: "WARN"}
}

func (c RuleConfig) withDefaults() RuleConfig {
	d := DefaultRuleConfig(c.RuleCode)
	if c.RuleCode == "" {
		c.RuleCode = d.RuleCode
	}
	if c.N <= 0 {
		c.N = d.N
	}
	if c.Hits <= 0 {
		c.Hits = d.Hits
	}
	if c.K == 0 && d.K != 0 {
		c.K = d.K
	}
	if c.Severity == "" {
		c.Severity = d.Severity
	}
	if c.Hits > c.N && c.N > 0 {
		c.Hits = c.N
	}
	return c
}

func (c RuleConfig) kind() string {
	if def, ok := catalogIndex[c.RuleCode]; ok {
		return def.Kind
	}
	return KindBeyondLimit
}
