package notify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/appleboy/go-fcm"
	"github.com/satimoto/go-datastore/pkg/util"
)

type Notifier interface {
	ToNotification(token, platform string, appData *string) *Notification
}

func (r *NotifyResolver) PostNotify(rw http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")
	platform := request.URL.Query().Get("platform")
	appData := util.NilString(request.URL.Query().Get("app_data"))

	notifyRequest, err := r.getNotifyRequest(request)
	if err != nil {
		log.Printf("API075: Error getting notify request: %v err=%v", token, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	notification := r.getNotification(token, platform, appData, notifyRequest)
	message := r.getMessage(notification)
	if message == nil {
		log.Printf("API076: Error getting notification: %v, type=%v, data=%#v", token, notifyRequest.Template, notifyRequest.Data)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := r.NotificationService.SendNotification(message); err != nil {
		log.Printf("API077: Error sending notification: %v, type=%v, data=%#v", token, notifyRequest.Template, notifyRequest.Data)
		log.Printf("API077: err=%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (r *NotifyResolver) getNotifyRequest(request *http.Request) (NotifyRequestDto, error) {
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return NotifyRequestDto{}, err
	}

	body := NotifyRequestDto{}
	if err = json.Unmarshal(bodyBytes, &body); err != nil {
		return NotifyRequestDto{}, err
	}

	return body, nil
}

func (r *NotifyResolver) getNotification(token, platform string, appData *string, request NotifyRequestDto) *Notification {
	var payloadHandler Notifier

	switch request.Template {
	case "address_txs_confirmed":
		payloadHandler = NewAddressTxsConfirmedPayload(request)
	case "lnurlpay_info":
		payloadHandler = NewLnurlPayInfoPayload(request)
	case "lnurlpay_invoice":
		payloadHandler = NewLnurlPayInvoicePayload(request)
	case "payment_received":
		payloadHandler = NewPaymentReceivedPayload(request)
	case "tx_confirmed":
		payloadHandler = NewTxConfirmedPayload(request)
	}

	if payloadHandler != nil {
		return payloadHandler.ToNotification(token, platform, appData)
	}
	return nil
}

func (r *NotifyResolver) getMessage(notification *Notification) *fcm.Message {
	if notification != nil {
		payload, err := json.Marshal(notification.Data)
		if err != nil {
			return nil
		}

		data := make(map[string]interface{})
		data["notification_type"] = notification.Template
		data["notification_payload"] = string(payload)
		if notification.AppData != nil {
			data["app_data"] = *notification.AppData
		}

		message := fcm.Message{
			To:               notification.TargetIdentifier,
			ContentAvailable: false,
			MutableContent:   true,
			Priority:         "high",
			Data:             data,
		}

		if strings.EqualFold(notification.Type, "ios") {
			message.Notification = &fcm.Notification{
				Title: notification.DisplayMessage,
			}
		}

		return &message 
	}
	return nil
}
