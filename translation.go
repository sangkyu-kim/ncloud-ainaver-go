package ainaver

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const TranslationUrl string = DefaultRestUrl + "/nmt/v1/translation"

type PostTranslation struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	Text        string `json:"text"`
	Honorific   bool   `json:"honorific"`
	ReplaceInfo string `json:"replaceInfo"`
}

type TranslationResponse struct {
	Message ResultMessage `json:"message"`
}

type ResultMessage struct {
	Type    string            `json:"@type"`
	Service string            `json:"@service"`
	Version string            `json:"@version"`
	Result  TranslationResult `json:"result"`
}

type TranslationResult struct {
	TranslatedText string `json:"translatedText"`
}

func (c *Client) PostTranslation(ctx context.Context, reqBody *PostTranslation) (*TranslationResponse, error) {
	rb, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", TranslationUrl, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var result *TranslationResponse

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
