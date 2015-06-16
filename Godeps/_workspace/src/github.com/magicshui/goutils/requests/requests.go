package requests

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func PostRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, err
}
func PostFileRequest(url string, params map[string]string, data []byte, name string) ([]byte, error) {
	b := bytes.NewReader(data)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(name, name)
	if err != nil {
		return nil, err
	}
	defer writer.Close()
	_, err = io.Copy(part, b)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return []byte(""), err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return result, err
}
func PostHttpsRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, err
}

func GetHttpsRequest(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte(""), err
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err
	}
	return result, err
}

func GetRequest(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte(""), err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte(""), err
	}
	return result, err
}

func PutRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", url, b)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, err
}
func DeleteRequest(url string, data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	req, err := http.NewRequest("Delete", url, b)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, err
}
