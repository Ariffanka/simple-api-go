package model

type Bio struct {
	ID       int    `json:id`
	Username string `json:username`
	Email    string `json:email`
}
