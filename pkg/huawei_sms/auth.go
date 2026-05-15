package huaweisms

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"regexp"
)

func (h *HuaweiClient) fetchToken() error {
	body, _, err := h.get(h.BaseURL, nil)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`name="csrf_token"\s+content="(\S+)"`)
	if match := re.FindStringSubmatch(string(body)); len(match) > 1 {
		h.Token = match[1]
	}
	return nil
}

func (h *HuaweiClient) loginState() (*stateResponse, error) {
	headers := make(map[string]string)
	if h.Token != "" {
		headers["__RequestVerificationToken"] = h.Token
	}

	body, _, err := h.get(h.BaseURL+"/api/user/state-login", headers)
	if err != nil {
		return nil, err
	}

	var resp stateResponse
	return &resp, xml.Unmarshal(body, &resp)
}

func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func (h *HuaweiClient) encodePassword(passwordType int) string {
	if passwordType == 4 {
		pwd := base64.StdEncoding.EncodeToString([]byte(sha256Hex(h.Password)))
		hash := sha256Hex(h.Username + pwd + h.Token)
		return base64.StdEncoding.EncodeToString([]byte(hash))
	}
	return base64.StdEncoding.EncodeToString([]byte(h.Password))
}

func (h *HuaweiClient) login(passwordType int) error {
	req := loginRequest{
		Username:     h.Username,
		Password:     h.encodePassword(passwordType),
		PasswordType: passwordType,
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	if h.Token != "" {
		headers["__RequestVerificationToken"] = h.Token
	}

	_, resp, err := h.post(h.BaseURL+"/api/user/login", xmlData, headers)
	if err != nil {
		return err
	}

	// 更新 Token
	for _, key := range []string{
		"__RequestVerificationTokenone",
		"__RequestVerificationTokentwo",
		"__RequestVerificationToken",
	} {
		if token := resp.Header.Get(key); token != "" {
			h.Token = token
			break
		}
	}
	return nil
}
