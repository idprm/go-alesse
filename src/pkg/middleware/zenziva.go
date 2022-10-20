package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/idprm/go-alesse/src/pkg/util/localconfig"
)

type NotifDoctorRequest struct {
	Userkey  string `json:"userkey"`
	Passkey  string `json:"passkey"`
	Instance string `json:"instance"`
	Nohp     string `json:"nohp"`
	Pesan    string `json:"pesan"`
}

func ZenzivaSendSMS(cfg *localconfig.Secret, msisdn string, message string) (string, error) {
	url := cfg.ZV.Url + "/api/WAsendMsg/"

	request := NotifDoctorRequest{
		Userkey:  cfg.ZV.UserName,
		Passkey:  cfg.ZV.Password,
		Instance: cfg.ZV.Instance,
		Nohp:     msisdn,
		Pesan:    message,
	}

	payload, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}
