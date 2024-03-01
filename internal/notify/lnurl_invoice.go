package notify

import (
	"encoding/json"
)

type LnurlPayInvoiceData struct {
	Amount   uint64 `json:"amount"`
	ReplyURL string `json:"reply_url"`
}

type LnurlPayInvoicePayload struct {
	Template string              `json:"template"`
	Data     LnurlPayInvoiceData `json:"data"`
}

func NewLnurlPayInvoicePayload(r NotifyRequestDto) *LnurlPayInvoicePayload {
	if r.Template == "lnurlpay_invoice" {
		bytes, err := json.Marshal(r.Data)
		if err != nil {
			return nil
		}

		var data LnurlPayInvoiceData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil
		}

		return &LnurlPayInvoicePayload{
			Template: r.Template,
			Data:     data,
		}
	}

	return nil
}

func (p *LnurlPayInvoicePayload) ToNotification(token, platform string, appData *string) *Notification {
	return &Notification{
		Template:         p.Template,
		DisplayMessage:   "Invoice requested",
		Type:             platform,
		TargetIdentifier: token,
		AppData:          appData,
		Data: map[string]interface{}{
			"amount":    p.Data.Amount,
			"reply_url": p.Data.ReplyURL,
		},
	}
}
