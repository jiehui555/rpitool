package huaweisms

import "encoding/xml"

// GetSMSList 获取收件箱短信列表（BoxType=1）
func (h *HuaweiClient) GetSMSList() ([]Message, error) {
	if err := h.fetchToken(); err != nil {
		return nil, err
	}

	state, err := h.loginState()
	if err != nil {
		return nil, err
	}

	// 未登录则自动登录
	if state.State != "0" {
		if err := h.login(state.PasswordType); err != nil {
			return nil, err
		}
	}

	req := smsRequest{
		PageIndex:       1,
		ReadCount:       20,
		BoxType:         1,
		SortType:        0,
		Ascending:       0,
		UnreadPreferred: 0,
	}

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	if h.Token != "" {
		headers["__RequestVerificationToken"] = h.Token
	}

	body, _, err := h.post(h.BaseURL+"/api/sms/sms-list", xmlData, headers)
	if err != nil {
		return nil, err
	}

	var resp smsResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp.Messages.Message, nil
}
