package hazardhub

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetGeoRisks(auth, url string, lat, lng float64) (map[int]*RiskProfile, error) {
	key := fmt.Sprintf("%f,%f", lat, lng)

	p, ok := getParsed(key)
	if ok {
		return p, nil
	}

	r, ok := getRaw(key)
	if ok {
		p := parse(r)
		uploadParsed(key, p)
		return p, nil
	}

	params := map[string]string{
		"lat": strconv.FormatFloat(lat, 'f', 6, 64),
		"lng": strconv.FormatFloat(lng, 'f', 6, 64),
	}

	b, err := getRisks(auth, url, params)
	if err != nil {
		return nil, fmt.Errorf("unable to get risks(%s): %s", key, err)
	}
	uploadRaw(key, b)

	var resp RisksResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		tmplt := "unable to parse hazard hub response(%s): %s"
		return nil, fmt.Errorf(tmplt, string(b), err)
	}

	p = parse(&resp)
	uploadParsed(key, p)
	return p, nil
}
