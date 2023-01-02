package notification

type NotificationDto map[string]interface{}

func CreateDataPingNotificationDto(ping string) NotificationDto {
	response := map[string]interface{}{
		"type": DATA_PING,
		"ping": ping,
	}

	return response
}
