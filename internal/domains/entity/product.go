package entity

type Product struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}
