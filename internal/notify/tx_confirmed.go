package notify

import (
	"encoding/json"
)

type TxConfirmedData struct {
	TxID string `json:"tx_id"`
}

type TxConfirmedPayload struct {
	Template string          `json:"template"`
	Data     TxConfirmedData `json:"data"`
}

func NewTxConfirmedPayload(r NotifyRequestDto) *TxConfirmedPayload {
	if r.Template == "tx_confirmed" {
		bytes, err := json.Marshal(r.Data)
		if err != nil {
			return nil
		}

		var data TxConfirmedData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil
		}

		return &TxConfirmedPayload{
			Template: r.Template,
			Data:     data,
		}
	}

	return nil
}

func (p *TxConfirmedPayload) ToNotification(token, platform string, appData *string) *Notification {
	return &Notification{
		Template:         p.Template,
		DisplayMessage:   "Transaction confirmed",
		Type:             platform,
		TargetIdentifier: token,
		AppData:          appData,
		Data: map[string]interface{}{
			"tx_id": p.Data.TxID,
		},
	}
}
