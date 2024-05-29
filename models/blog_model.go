package models

type Blogs struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	Content   string `json:"content" bson:"content"`
	HeroImage string `json:"heroImage bson:"heroImage"`
}

func BlogModel() Blogs {
	return Blogs{}
}
