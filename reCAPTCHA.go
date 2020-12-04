package tools

import (
	"encoding/json"
)

const (
	recaptchaUrl    = `https://recaptcha.net/recaptcha/api/siteverify`
)

type ReCAPTCHAResult struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func CheckReCAPTCHA(ip, response,recaptchaSecret string) (bool,ReCAPTCHAResult,error) {
	params := `secret=` + recaptchaSecret + `&response=` + response + `&remoteip=` + ip
	result, err := CurlPost(recaptchaUrl, params, 2)
	if err != nil {
		return false, ReCAPTCHAResult{}, err
	}
	var res ReCAPTCHAResult
	err = json.Unmarshal(result, &res)
	if err!= nil {
		return false, ReCAPTCHAResult{}, err
	}
	if !res.Success {
		return false, res, nil
	}
	return true,res,nil
}
