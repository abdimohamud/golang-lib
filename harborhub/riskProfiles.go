package hazardhub

import (
	"fmt"
)

type RiskTreatment struct {
	Text  string
	Color string
}

var riskTreatments = map[int]*RiskTreatment{
	1: &RiskTreatment{Text: "Don't worry", Color: "42B9FF"},
	2: &RiskTreatment{Text: "Be careful", Color: "21C95A"},
	3: &RiskTreatment{Text: "Use caution", Color: "F9B506"},
	4: &RiskTreatment{Text: "Be ready", Color: "F4240F"},
	5: &RiskTreatment{Text: "Be ready", Color: "F4240F"},
}

type RiskProfile struct {
	ID         int    `json:"id"`
	Level      int    `json:"risk_level"`
	LevelColor string `json:"risk_level_color"`
	LevelText  string `json:"risk_level_text"`
}

func NewRiskProfile(id, level int) *RiskProfile {
	rP := RiskProfile{ID: id, Level: level}
	treatment, ok := riskTreatments[level]
	if !ok {
		fmt.Printf("no treatment found for ID: %d, Level: %d\n", id, level)
		return &rP
	}
	rP.LevelColor = treatment.Color
	rP.LevelText = treatment.Text
	return &rP
}

func pandemic() *RiskProfile {
	return NewRiskProfile(1, 3)
}

func roadSafety() *RiskProfile {
	return NewRiskProfile(2, 4)
}

func wildFire(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(3, composeScore(r.Wildfire))
}

func heatWave() *RiskProfile {
	return NewRiskProfile(4, 3)
}

func homeSafety() *RiskProfile {
	return NewRiskProfile(5, 4)
}

func flood(r *RisksResponse) *RiskProfile {
	l := maxScore(r.CatFlood, r.FemaFlood, r.Flood, r.EFlood)
	return NewRiskProfile(6, l)
}

func tsunami(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(7, composeScore(r.Tsunami))
}

func hurricane(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(8, composeScore(r.Hurricane, r.EHurricane))
}

func winterStorm(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(9, maxScore(r.IceDamage, r.FrozenPipes, r.SnowLoad))
}

func volcano(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(10, composeScore(r.Volcano))
}

func tornado(r *RisksResponse) *RiskProfile {
	return NewRiskProfile(11, composeScore(r.Tornado))
}

func earthquake(r *RisksResponse) *RiskProfile {
	l := maxScore(r.DesignatedFault, r.Earthquake, r.FaultEarthquake, r.FrackingEarthquake)
	return NewRiskProfile(12, l)
}
