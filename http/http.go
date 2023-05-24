package http

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func Get(link string, timeout time.Duration, headers map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: timeout * time.Millisecond}
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		if uerr, ok := err.(*url.Error); ok {
			if nerr, ok := uerr.Err.(net.Error); ok {
				if nerr.Timeout() {
					// 记下这种写法
				}
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("GET " + link + " readall " + err.Error())
	}
	return body, nil
}
func Post(link string, pl []byte, timeout time.Duration, headers map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: timeout * time.Millisecond}
	buffer := bytes.NewBuffer(pl)
	request, err := http.NewRequest("POST", link, buffer)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	} else { // default
		request.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("POST " + link + " readall " + err.Error())
	}
	return body, nil
}
func Delete(link string, timeout time.Duration, headers map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: timeout * time.Millisecond}
	request, err := http.NewRequest("DELETE", link, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("DELETE " + link + " " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("DELETE " + link + " readall " + err.Error())
	}
	return body, nil
}
func Put(link string, timeout time.Duration, headers map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: timeout * time.Millisecond}
	request, err := http.NewRequest("PUT", link, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("PUT " + link + " " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("PUT " + link + " readall " + err.Error())
	}
	return body, nil
}
func ClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	if net.ParseIP(ip) != nil {
		return ip
	}
	return ""
}
