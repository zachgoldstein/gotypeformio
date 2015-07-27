package typeform

type FormResults struct {
	Answers []struct {
		Data struct {
			Type  string `json:"type"`
			Value struct {
				Label string      `json:"label"`
				Other interface{} `json:"other"`
			} `json:"value"`
		} `json:"data"`
		FieldID float64 `json:"field_id"`
	} `json:"answers"`
	ID    string `json:"id"`
	Token string `json:"token"`
}
