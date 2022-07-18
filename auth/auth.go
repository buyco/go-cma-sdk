package auth

import (
	"context"
	"golang.org/x/xerrors"
	"time"
)

const (
	authPath  = "/as/token.oauth2"
	grantType = "client_credentials"
)

type Result struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Client struct {
	httpClient   HTTPClient
	url          string
	clientId     string
	clientSecret string
	scope        string
}

func NewClient(httpClient HTTPClient, url, clientId, clientSecret, scope string) *Client {
	return &Client{
		httpClient:   httpClient,
		url:          url,
		clientId:     clientId,
		clientSecret: clientSecret,
		scope:        scope,
	}
}

func (a Client) Token(ctx context.Context) (string, time.Duration, error) {
	result := Result{}
	errRes := ErrorResponse{}
	r, err := a.httpClient.NewRequest().
		SetResult(&result).
		SetError(&errRes).
		SetHeaders(a.buildHeaders()).
		SetContext(ctx).
		SetFormData(a.buildParams()).
		Post(a.url + authPath)
	if err != nil {
		return "", 0, err
	}
	if r.IsError() {
		return "", 0, xerrors.New(errRes.Error)
	} else if r.IsSuccess() {
		return result.AccessToken, time.Duration(result.ExpiresIn) * time.Second, nil
	}
	return "", 0, xerrors.Errorf("unknown error. status code %d; body: %s", r.GetStatusCode(), r.String())
}

func (a Client) buildParams() map[string]string {
	return map[string]string{
		"grant_type":    grantType,
		"client_id":     a.clientId,
		"client_secret": a.clientSecret,
		"scope":         a.scope,
	}
}

func (a Client) buildHeaders() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
}
