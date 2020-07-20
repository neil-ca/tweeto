package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User modelo*/
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Surname     string             `bson:"surname" json:"surname,omitempty"`
	DateOfBirth time.Time          `bson:"dateofbirth" json:"dateofbirth,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempt"`
	Banner      string             `bson:"banner" json:"banner,omitempt"`
	Biography   string             `bson:"biography" json:"biography,omitempt"`
	Location    string             `bson:"location" json:"location,omitempt"`
	SiteWeb     string             `bson:"siteweb" json:"siteweb,omitempt"`

}
