package sarif

type SarifReport2 struct {
	Schema  string `json:"$schema"`
	Version string `json:"version"`
	Runs    []struct {
		Tool struct {
			Driver struct {
				Name    string `json:"name"`
				Version string `json:"version"`
				Rules   []struct {
					ID              string `json:"id"`
					Name            string `json:"name"`
					FullDescription struct {
						Text string `json:"text"`
					} `json:"fullDescription"`
					HelpURI string `json:"helpUri"`
				} `json:"rules"`
			} `json:"driver"`
		} `json:"tool"`
		Results []struct {
			RuleID  string `json:"ruleId"`
			Level   string `json:"level"`
			Message struct {
				Text string `json:"text"`
			} `json:"message"`
			Locations []struct {
				PhysicalLocation struct {
					Address struct {
						AbsoluteAddress string `json:"absoluteAddress"`
					} `json:"address"`
				} `json:"physicalLocation"`
			} `json:"locations"`
		} `json:"results"`
	} `json:"runs"`
}
