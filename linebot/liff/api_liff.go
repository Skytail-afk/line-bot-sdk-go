/**
 * LIFF server API
 * LIFF Server API.
 *
 * The version of the OpenAPI document: 1.0.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py

package liff

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type LiffAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// LiffAPIOption type
type LiffAPIOption func(*LiffAPI) error

// New returns a new bot client instance.
func NewLiffAPI(channelToken string, options ...LiffAPIOption) (*LiffAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &LiffAPI{
		channelToken: channelToken,
		httpClient:   http.DefaultClient,
	}

	u, err := url.ParseRequestURI("https://api.line.me")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithContext method
func (call *LiffAPI) WithContext(ctx context.Context) *LiffAPI {
	call.ctx = ctx
	return call
}

func (client *LiffAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *LiffAPI) Url(endpointPath string) string {
	newPath := path.Join(client.endpoint.Path, endpointPath)
	u := *client.endpoint
	u.Path = newPath
	return u.String()
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) LiffAPIOption {
	return func(client *LiffAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpoint function
func WithEndpoint(endpoint string) LiffAPIOption {
	return func(client *LiffAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// AddLIFFApp
// Create LIFF app
// Adding the LIFF app to a channel
// Parameters:
//        addLiffAppRequest

// https://developers.line.biz/en/reference/liff-server/#add-liff-app
func (client *LiffAPI) AddLIFFApp(

	addLiffAppRequest *AddLiffAppRequest,

) (*AddLiffAppResponse, error) {
	_, body, error := client.AddLIFFAppWithHttpInfo(

		addLiffAppRequest,
	)
	return body, error
}

// AddLIFFApp
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
// Create LIFF app
// Adding the LIFF app to a channel
// Parameters:
//        addLiffAppRequest

// https://developers.line.biz/en/reference/liff-server/#add-liff-app
func (client *LiffAPI) AddLIFFAppWithHttpInfo(

	addLiffAppRequest *AddLiffAppRequest,

) (*http.Response, *AddLiffAppResponse, error) {
	path := "/liff/v1/apps"

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(addLiffAppRequest); err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(http.MethodPost, client.Url(path), &buf)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := AddLiffAppResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// DeleteLIFFApp
// Delete LIFF app from a channel
// Deletes a LIFF app from a channel.
// Parameters:
//        liffId             ID of the LIFF app to be updated

// https://developers.line.biz/en/reference/liff-server/#delete-liff-app
func (client *LiffAPI) DeleteLIFFApp(

	liffId string,

) (struct{}, error) {
	_, body, error := client.DeleteLIFFAppWithHttpInfo(

		liffId,
	)
	return body, error
}

// DeleteLIFFApp
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
// Delete LIFF app from a channel
// Deletes a LIFF app from a channel.
// Parameters:
//        liffId             ID of the LIFF app to be updated

// https://developers.line.biz/en/reference/liff-server/#delete-liff-app
func (client *LiffAPI) DeleteLIFFAppWithHttpInfo(

	liffId string,

) (*http.Response, struct{}, error) {
	path := "/liff/v1/apps/{liffId}"

	path = strings.Replace(path, "{liffId}", liffId, -1)

	req, err := http.NewRequest(http.MethodDelete, client.Url(path), nil)
	if err != nil {
		return nil, struct{}{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}

// GetAllLIFFApps
// Get all LIFF apps
// Gets information on all the LIFF apps added to the channel.
// Parameters:

// https://developers.line.biz/en/reference/liff-server/#get-all-liff-apps
func (client *LiffAPI) GetAllLIFFApps() (*GetAllLiffAppsResponse, error) {
	_, body, error := client.GetAllLIFFAppsWithHttpInfo()
	return body, error
}

// GetAllLIFFApps
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
// Get all LIFF apps
// Gets information on all the LIFF apps added to the channel.
// Parameters:

// https://developers.line.biz/en/reference/liff-server/#get-all-liff-apps
func (client *LiffAPI) GetAllLIFFAppsWithHttpInfo() (*http.Response, *GetAllLiffAppsResponse, error) {
	path := "/liff/v1/apps"

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetAllLiffAppsResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// UpdateLIFFApp
// Update LIFF app from a channel
// Update LIFF app settings
// Parameters:
//        liffId             ID of the LIFF app to be updated
//        updateLiffAppRequest

// https://developers.line.biz/en/reference/liff-server/#update-liff-app
func (client *LiffAPI) UpdateLIFFApp(

	liffId string,

	updateLiffAppRequest *UpdateLiffAppRequest,

) (struct{}, error) {
	_, body, error := client.UpdateLIFFAppWithHttpInfo(

		liffId,

		updateLiffAppRequest,
	)
	return body, error
}

// UpdateLIFFApp
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
// Update LIFF app from a channel
// Update LIFF app settings
// Parameters:
//        liffId             ID of the LIFF app to be updated
//        updateLiffAppRequest

// https://developers.line.biz/en/reference/liff-server/#update-liff-app
func (client *LiffAPI) UpdateLIFFAppWithHttpInfo(

	liffId string,

	updateLiffAppRequest *UpdateLiffAppRequest,

) (*http.Response, struct{}, error) {
	path := "/liff/v1/apps/{liffId}"

	path = strings.Replace(path, "{liffId}", liffId, -1)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(updateLiffAppRequest); err != nil {
		return nil, struct{}{}, err
	}
	req, err := http.NewRequest(http.MethodPut, client.Url(path), &buf)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}
