package model

type Profile struct {
	Id        int    `db:"id" json:"id,omitempty"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type Response struct {
	HttpCode int    `db:"http_code"`
	Link     string `db:"link"`
}
