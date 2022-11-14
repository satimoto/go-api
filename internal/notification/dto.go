package notification

type NotificationDto map[string]interface{}

func CreateDataPingNotificationDto() NotificationDto {
	response := map[string]interface{}{
		"type": DATA_PING,
	}

	return response
}
