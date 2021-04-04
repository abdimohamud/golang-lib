package hazardhub

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RiskScore struct {
	Score string `json:"score"`
}

type RisksResponse struct {
	Tornado  RiskScore `json:"enhanced_tornado_params"`
	Tsunami  RiskScore `json:"tsunami"`
	Volcano  RiskScore `json:"volcano"`
	Wildfire RiskScore `json:"wildfire"`
	// Earthquake metrics
	DesignatedFault    RiskScore `json:"designated_fault"`
	Earthquake         RiskScore `json:"earthquake"`
	FaultEarthquake    RiskScore `json:"fault_earthquake"`
	FrackingEarthquake RiskScore `json:"fracking_earthquake"`
	// Flood metrics
	CatFlood  RiskScore `json:"hazardhub_catastrophic_flood"`
	Flood     RiskScore `json:"hazardhub_flood"`
	EFlood    RiskScore `json:"enhanced_hazardhub_flood"`
	FemaFlood RiskScore `json:"fema_all_flood"`
	// Hurricane metrics
	Hurricane  RiskScore `json:"hurricane"`
	EHurricane RiskScore `json:"enhanced_hurricane_params"`
	// WinterStorm metrics
	IceDamage   RiskScore `json:"ice_dam_index"`
	FrozenPipes RiskScore `json:"frozen_pipe_index"`
	SnowLoad    RiskScore `json:"hh_snow_load"`
}

func getRisks(auth, url string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", auth)

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)

	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read hazard hub response: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("hazard hub error: %s", string(b))
	}

	return b, nil
}
