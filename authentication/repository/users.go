package repository

import (
	"golang_microservice_mongodb_kub_jwt_grpc/db"

	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

type UsersRepository interface{}

type usersRepository struct {
	c mongo.Collection
}

func NewUserRepository(conn db.Connection) UsersRepository {
	return &usersRepository{c: *conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save() {

}
