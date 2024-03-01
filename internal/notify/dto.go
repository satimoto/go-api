package notify

type NotifyRequestDto struct {
	Template string      `json:"template"`
	Data     interface{} `json:"data"`
}

type Notification struct {
	Template         string
	DisplayMessage   string
	Type             string
	TargetIdentifier string
	AppData          *string
	Data             map[string]interface{}
}
