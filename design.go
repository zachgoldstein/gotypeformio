package typeform

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type Design struct {
	Colors            Colors `json:"colors"`
	Font              string `json:"font"`
	BackgroundImageId int    `json:"background_image_id"`
	Id                int    `json:"id,omitempty"`
}

func NewDesign(colors *Colors, font string, backgroundImageId int) *Design {
	return &Design{
		Colors:            *colors,
		Font:              font,
		BackgroundImageId: backgroundImageId,
	}
}

type Colors struct {
	Answer     string `json:"answer,omitempty"`
	Background string `json:"background,omitempty"`
	Button     string `json:"button,omitempty"`
	Question   string `json:"question,omitempty"`
}

func NewColors(answer string, background string, button string, question string) *Colors {
	return &Colors{
		Answer:     answer,
		Background: background,
		Button:     button,
		Question:   question,
	}
}

func CreateDesign(authToken string, debug bool, designReq *Design) (*Design, error) {

	rawSubmission, err := json.Marshal(designReq)
	if err != nil {
		return &Design{}, err
	}

	buf := bytes.NewBuffer(rawSubmission)

	req, err := http.NewRequest("POST", TYPEFORM_API+"designs", buf)
	if err != nil {
		return &Design{}, err
	}

	req.Header.Add("x-api-token", authToken)

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump request: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full request: \n %v \n", string(dump))
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &Design{}, err
	}

	if debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump response: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full response: \n %v \n", string(dump))
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &Design{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = NewRequestError(res.StatusCode, res.Request.URL.String(), string(body))
		return &Design{}, err
	}

	var resp Design
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &Design{}, err
	}

	return &resp, nil
}
