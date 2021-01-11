package write

type Widget struct {
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Value       int    `json:"value"`
}
