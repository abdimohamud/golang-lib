package hazardhub

import (
	"fmt"
	"math"
	"os"
)

var env = os.Getenv("ENVIRONMENT")

var scoreToInt = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"F": 4,
}

func parse(r *RisksResponse) map[int]*RiskProfile {
	riskProfiles := [12]*RiskProfile{
		pandemic(),
		roadSafety(),
		wildFire(r),
		heatWave(),
		homeSafety(),
		flood(r),
		tsunami(r),
		hurricane(r),
		winterStorm(r),
		volcano(r),
		tornado(r),
		earthquake(r),
	}

	doc := map[int]*RiskProfile{}
	for _, rP := range riskProfiles {
		doc[rP.ID] = rP
	}

	return doc
}

func composeScore(scores ...RiskScore) int {
	sum := 0
	num := 0
	for _, r := range scores {
		if r.Score == "" {
			continue
		}
		score, ok := scoreToInt[r.Score]
		if !ok {
			fmt.Printf("no mapping found for: %s", r.Score)
			continue
		}
		sum += score
		num += 1
	}

	if num == 0 {
		return 1
	}

	avg := float64(sum) / float64(num)
	// round up to err on caution
	avg += 0.5

	return int(math.Max(avg, 1))
}

func maxScore(scores ...RiskScore) int {
	var max int

	for _, r := range scores {
		if r.Score == "" {
			continue
		}
		score, ok := scoreToInt[r.Score]
		if !ok {
			fmt.Printf("no mapping found for: %s", r.Score)
			continue
		}
		if score > max {
			max = score
		}
	}

	return max
}
