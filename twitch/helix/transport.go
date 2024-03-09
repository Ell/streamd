package helix

type Transport struct {
	Method    string `json:"method"`
	SessionId string `json:"session_id,omitempty"`
	Callback  string `json:"callback,omitempty"`
}
