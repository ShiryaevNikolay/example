package user

type User struct {
	ID           string `json:"id" bson:"_id, omitempty"` // bson - для mongodb, _id - сисетмное поле в mongodb, которое само генерирует уникальный id в своей системе
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"` // у JSON "-", т.к. мы никому не будем отдавать пароль
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"` // у JSON "-", т.к. мы никому не будем отдавать пароль
}
