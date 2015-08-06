package gotypeformio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type Image struct {
	Id  int    `json:"id"`
	Url string `json:"original_url"`
}

type ImageRequest struct {
	Url  string `json:"url"`
	Type string `json:"type,omitempty"`
}

func CreateImage(authToken string, debug bool, url string) (*Image, error) {
	imageRequest := &ImageRequest{
		Url: url,
	}

	rawSubmission, err := json.Marshal(imageRequest)
	if err != nil {
		return &Image{}, err
	}

	buf := bytes.NewBuffer(rawSubmission)

	req, err := http.NewRequest("POST", TYPEFORM_API+"images", buf)
	if err != nil {
		return &Image{}, err
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
		return &Image{}, err
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
		return &Image{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = NewRequestError(res.StatusCode, res.Request.URL.String(), string(body))
		return &Image{}, err
	}

	var resp Image
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &Image{}, err
	}

	return &resp, nil
}

func GetImage() {
	/*
	   url:
	   required
	   String
	   The url of the string
	   type:

	*/
}

/*

{
    "id": 5,
    "original_url": "http://example.com/original_img.png",
    "type": "choice"
}

*/
