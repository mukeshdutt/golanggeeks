package model

// Menu is model
type Menu struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Category string `json:"category"`
	Status   bool   `json:"status"`
}
