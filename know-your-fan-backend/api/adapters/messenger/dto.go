package messenger

type ClientCreatedEvent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Document []byte `json:"document"`
	Valid    *bool  `json:"valid"`
}
