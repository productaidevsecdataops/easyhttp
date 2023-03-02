package easyhttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	timeoutInSecs = 29
)

/*
Used for simple JSON Gets
*/
func Get(url string) ([]byte, *http.Response, error) {
	timeout := time.Duration(timeoutInSecs * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil

}

// GetWithBasicAuth returns byte array of the response, http response and any errors
func GetWithBasicAuth(url string, username string, password string) ([]byte, *http.Response, error) {
	timeout := time.Duration(timeoutInSecs * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(username, password)
	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil

}

// @ Needs a url to call and a bearertoken as a string
// RETURNS a byte array response, http.response and error
// Makes an http get request with a bearer token
func GetWithBearer(url string, token string) ([]byte, *http.Response, error) {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + token

	timeout := time.Duration(timeoutInSecs * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Add("Authorization", bearer)

	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil

}

// @ Needs a url to call and a bearertoken as a string
// RETURNS a byte array response, http.response and error
// Makes an http get request with a bearer token
func GetWithBearerGraph(url string, token string) ([]byte, *http.Response, error) {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + token

	timeout := time.Duration(timeoutInSecs * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("ConsistencyLevel", "eventual")
	request.Header.Add("Authorization", bearer)

	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil

}

/*
A simple POST function
*/
func Post(url string, requestBody []byte) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

func PostWithBasicAuth(url string, requestBody []byte, u string, p string) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth(u, p)
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple POST function
*/
func PostWithBearer(url string, requestBody []byte, token string) ([]byte, *http.Response, error) {

	var bearer = "Bearer " + token

	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", bearer)

	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple PUT function
*/
func Put(url string, requestBody []byte) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple PUT function
*/
func PutWithBasicAuth(url string, requestBody []byte, u string, p string) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth(u, p)
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple PUT function
*/
func PutWithBearer(url string, requestBody []byte, token string) ([]byte, *http.Response, error) {

	bearer := "Bearer " + token

	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", bearer)
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

// PATCHING //////////////////////////////////////////////////////////////////
/*
A simple PUT function
*/
func Patch(url string, requestBody []byte) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple PUT function
*/
func PatchWithBasicAuth(url string, requestBody []byte, u string, p string) ([]byte, *http.Response, error) {
	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth(u, p)
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}

/*
A simple PUT function
*/
func PatchWithBearer(url string, requestBody []byte, token string) ([]byte, *http.Response, error) {

	bearer := "Bearer " + token

	// We will set our default timeout
	timeout := time.Duration(timeoutInSecs * time.Second)

	// We create a client with the timeout attached
	client := http.Client{
		Timeout: timeout,
	}

	// We are going to create a new request
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBody))

	// We set the header to application/json
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", bearer)
	if err != nil {
		return nil, nil, err
	}
	// We execute our request
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	// At this point we know we are successful so we can defer the close
	// https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-iii-36a1ab3d6ef1
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			log.Println("Error Deferring resp.Body.Close (io.Closer)")
		}
	}(resp.Body)

	// We parse our response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// We return a byte array
	return body, resp, nil
}
