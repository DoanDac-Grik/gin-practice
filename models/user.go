package models

type User struct {
	Id       string `json:"id" bson:"id"`
	Username string `json:"username" bscon:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Level    string `json:"level" bson:"level" binding:"required"`
}
