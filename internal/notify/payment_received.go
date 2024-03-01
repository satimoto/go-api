package notify

import (
	"encoding/json"
)

type PaymentReceivedData struct {
	PaymentHash string `json:"payment_hash"`
}

type PaymentReceivedPayload struct {
	Template string              `json:"template"`
	Data     PaymentReceivedData `json:"data"`
}

func NewPaymentReceivedPayload(r NotifyRequestDto) *PaymentReceivedPayload {
	if r.Template == "payment_received" {
		bytes, err := json.Marshal(r.Data)
		if err != nil {
			return nil
		}

		var data PaymentReceivedData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil
		}

		return &PaymentReceivedPayload{
			Template: r.Template,
			Data:     data,
		}
	}

	return nil
}

func (p *PaymentReceivedPayload) ToNotification(token, platform string, appData *string) *Notification {
	return &Notification{
		Template:         p.Template,
		DisplayMessage:   "Incoming payment",
		Type:             platform,
		TargetIdentifier: token,
		AppData:          appData,
		Data: map[string]interface{}{
			"payment_hash": p.Data.PaymentHash,
		},
	}
}
