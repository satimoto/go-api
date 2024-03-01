package notify

import (
	"encoding/json"
)

type AddressTxsConfirmedData struct {
	Address string `json:"address"`
}

type AddressTxsConfirmedPayload struct {
	Template string                  `json:"template"`
	Data     AddressTxsConfirmedData `json:"data"`
}

func NewAddressTxsConfirmedPayload(r NotifyRequestDto) *AddressTxsConfirmedPayload {
	if r.Template == "address_txs_confirmed" {
		bytes, err := json.Marshal(r.Data)
		if err != nil {
			return nil
		}
		
		var data AddressTxsConfirmedData
		if err := json.Unmarshal(bytes, &data); err != nil {
			return nil
		}

		return &AddressTxsConfirmedPayload{
			Template: r.Template,
			Data:     data,
		}
	}

	return nil
}

func (p *AddressTxsConfirmedPayload) ToNotification(token, platform string, appData *string) *Notification {
	return &Notification{
		Template:         p.Template,
		DisplayMessage:   "Address transactions confirmed",
		Type:             platform,
		TargetIdentifier: token,
		AppData:          appData,
		Data: map[string]interface{}{
			"address": p.Data.Address,
		},
	}
}
