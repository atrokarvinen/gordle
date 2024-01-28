package models

type User struct {
	Id    int    `json:"id"`
	Games []Game `json:"games"`
}
