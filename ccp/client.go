//
// This client library provides Create, Read, Update, and Delete operations for Cisco Container Platform.
//

package ccp

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"reflect"
)

//import "encoding/json"

type Client struct {
	Username string
	Password string
	BaseURL  string
}

var jar, err = cookiejar.New(nil)

func NewClient(username, password, baseURL string) *Client {

	return &Client{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {

	var client *http.Client

	req.Header.Add("Content-Type", "application/json")
	//req.SetBasicAuth(s.Username, s.Password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr, Jar: jar}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode && 201 != resp.StatusCode && 202 != resp.StatusCode && 204 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Bool(value bool) *bool {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int(value int) *int {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int64(value int64) *int64 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func String(value string) *string {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float32(value float32) *float32 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float64(value float64) *float64 {
	return &value
}

//modified from unexported nonzero function in the validtor package
//https://github.com/go-validator/validator/blob/v2/builtins.go
func nonzero(v interface{}) bool {
	st := reflect.ValueOf(v)
	nonZeroValue := false
	switch st.Kind() {
	case reflect.Ptr, reflect.Interface:
		nonZeroValue = st.IsNil()
	case reflect.Invalid:
		nonZeroValue = true // always invalid
	case reflect.Struct:
		nonZeroValue = false // always valid since only nil pointers are empty
	default:
		return true
	}

	if nonZeroValue {
		return true
	}
	return false
}
