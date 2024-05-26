package models

type Users struct {
	Username string  `json:"username" bson:"username"`
	Password string  `json:"password" bson:"password"`
	IsAdmin  bool    `json:"isAdmin" bson:"isAdmin"`
	DP       string  `json:"dp" bson:"dp"`
	Uploads  []Blogs `json:"blogs" bson:"blogs"`
}

func UserModel() Users {
	return Users{}
}
