//Package recaptchago provides ways to validate Google recaptcha tokens.
//Currently only simple func is provided. In future I might make a
//http.Handler middleware...
package recaptchago

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//ValidateToken takes token returned by user and secret key provided by recapcha,
//and validates it using recaptha's API.
func ValidateToken(token, secret string) (bool, error) {
	values := make(url.Values)
	values.Set("secret", secret)
	values.Set("response", token)
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", values)
	if err != nil {
		return false, err
	}
	//log.Println(resp)
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	r := make(map[string]interface{})
	err = json.Unmarshal(data, &r)
	if err != nil {
		return false, err
	}
	//log.Println(string(data))
	s := r["success"].(bool)
	return s, nil
}
