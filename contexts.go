package apiai

import (
        "net/http"
        "encoding/json"
        "bytes"
        "fmt"
        "net/url"
)

type ContextParameter struct {
        Name  string `json:"name"`
        Value string `json:"value"`
}

type Context struct {
        Name     string `json:"name"`
        Lifespan int `json:"lifespan"`
        Params   []ContextParameter `json:"parameters"`
}

func (c *ApiClient) GetContexts() ([]Context, error) {
        req, err := http.NewRequest("GET", c.buildUrl("contexts", map[string]string{
                "sessionId": c.config.sessionId,
        }), nil)
        if err != nil {
                return nil, err
        }
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Content-type", "application/json, charset=utf-8")
        req.Header.Set("Authorization", "Bearer "+c.config.token)

        httpClient := http.DefaultClient
        resp, err := httpClient.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        var contexts []Context
        switch resp.StatusCode {
        case http.StatusOK:
                decoder := json.NewDecoder(resp.Body)
                err = decoder.Decode(&contexts)
                if err != nil {
                        return nil, err
                }
                return contexts, nil
        default:
                return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
        }
}

func (c *ApiClient) GetContext(name string) (*Context, error) {
        req, err := http.NewRequest("GET", c.buildUrl("contexts/"+url.QueryEscape(name), map[string]string{
                "sessionId": c.config.sessionId,
        }), nil)
        if err != nil {
                return nil, err
        }
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Content-type", "application/json, charset=utf-8")
        req.Header.Set("Authorization", "Bearer "+c.config.token)

        httpClient := http.DefaultClient
        resp, err := httpClient.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        var context *Context
        switch resp.StatusCode {
        case http.StatusOK:
                decoder := json.NewDecoder(resp.Body)
                err = decoder.Decode(&context)
                if err != nil {
                        return nil, err
                }
                return context, nil
        default:
                return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
        }
}

func (c *ApiClient) CreateContext(context Context) error {
        body := new(bytes.Buffer)
        err := json.NewEncoder(body).Encode(context)
        if err != nil {
                return err
        }

        req, err := http.NewRequest("POST", c.buildUrl("contexts", map[string]string{
                "sessionId": c.config.sessionId,
        }), body)
        if err != nil {
                return err
        }
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Content-type", "application/json, charset=utf-8")
        req.Header.Set("Authorization", "Bearer "+c.config.token)

        httpClient := http.DefaultClient
        resp, err := httpClient.Do(req)
        if err != nil {
                return err
        }
        defer resp.Body.Close()

        switch resp.StatusCode {
        case http.StatusOK:
                return nil
        default:
                return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
        }
}

func (c *ApiClient) DeleteContexts() error {
        req, err := http.NewRequest("DELETE", c.buildUrl("contexts", map[string]string{
                "sessionId": c.config.sessionId,
        }), nil)
        if err != nil {
                return err
        }
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Content-type", "application/json, charset=utf-8")
        req.Header.Set("Authorization", "Bearer "+c.config.token)

        httpClient := http.DefaultClient
        resp, err := httpClient.Do(req)
        if err != nil {
                return err
        }
        defer resp.Body.Close()

        switch resp.StatusCode {
        case http.StatusOK:
                return nil
        default:
                return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
        }
}

func (c *ApiClient) DeleteContext(name string) error {
        req, err := http.NewRequest("DELETE", c.buildUrl("contexts/"+url.QueryEscape(name), map[string]string{
                "sessionId": c.config.sessionId,
        }), nil)
        if err != nil {
                return err
        }
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Content-type", "application/json, charset=utf-8")
        req.Header.Set("Authorization", "Bearer "+c.config.token)

        httpClient := http.DefaultClient
        resp, err := httpClient.Do(req)
        if err != nil {
                return err
        }
        defer resp.Body.Close()

        switch resp.StatusCode {
        case http.StatusOK:
                return nil
        default:
                return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
        }
}