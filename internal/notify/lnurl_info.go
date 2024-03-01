package notify

import (
	"encoding/json"
)

type LnurlPayInfoData struct {
	CallbackURL string `json:"callback_url"`
	ReplyURL    string `json:"reply_url"`
}

type LnurlPayInfoPayload struct {
	Template string           `json:"template"`
	Data     LnurlPayInfoData `json:"data"`
}

func NewLnurlPayInfoPayload(r NotifyRequestDto) *LnurlPayInfoPayload {
	if r.Template == "lnurlpay_info" {
		bytes, err := json.Marshal(r.Data)
		if err != nil {
			return nil
		}
		
		var data LnurlPayInfoData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil
		}

		return &LnurlPayInfoPayload{
			Template: r.Template,
			Data:     data,
		}
	}

	return nil
}

func (p *LnurlPayInfoPayload) ToNotification(token, platform string, appData *string) *Notification {
	return &Notification{
		Template:         p.Template,
		DisplayMessage:   "Receiving payment",
		Type:             platform,
		TargetIdentifier: token,
		AppData:          appData,
		Data: map[string]interface{}{
			"callback_url": p.Data.CallbackURL,
			"reply_url":    p.Data.ReplyURL,
		},
	}
}
