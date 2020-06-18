package organization

import "time"

type createdResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
