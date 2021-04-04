package hazardhub

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetZipRisks(auth, url, state, zip string) (map[int]*RiskProfile, error) {
	if _, err := strconv.Atoi(zip); err != nil {
		return nil, fmt.Errorf("invalid zipcode: %s", zip)
	}

	p, ok := getParsed(zip)
	if ok {
		return p, nil
	}

	r, ok := getRaw(zip)
	if ok {
		p := parse(r)
		uploadParsed(zip, p)
		return p, nil
	}

	params := map[string]string{"state": state, "zip": zip}
	b, err := getRisks(auth, url, params)
	if err != nil {
		return nil, fmt.Errorf("unable to get risks(%s): %s", zip, err)
	}
	uploadRaw(zip, b)

	var resp RisksResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		tmplt := "unable to parse hazard hub response(%s): %s"
		return nil, fmt.Errorf(tmplt, string(b), err)
	}

	p = parse(&resp)
	uploadParsed(zip, p)
	return p, nil
}
