package models

import (
	"golang_microservice_mongodb_kub_jwt_grpc/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Created  time.Time          `bson:"created"`
	Updated  time.Time          `bson:"updated"`
}

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Created:  u.Created.Unix(),
		Updated:  u.Updated.Unix(),
	}
}

func (u *User) FromProtoBuffer(user *pb.User) {
	var err error
	u.Id, err = primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		panic(err)
	}

	u.Name = user.GetName()
	u.Email = user.GetEmail()
	u.Password = user.GetPassword()
	u.Created = time.Unix(user.Created, 0)
	u.Updated = time.Unix(user.Updated, 0)
}
