package godis

type dataProperties struct {
	os      string `json:"$os"`
	browser string `json:"$browser"`
	device  string `json:"$device"`
}

type scopes struct {
	Token      string         `json:"token"`
	Intents    int            `json:"intents"`
	Properties dataProperties `json:"properties"`
}
type identify struct {
	operator int     `json:"op"`
	Data     *scopes `json:"d"`
}
