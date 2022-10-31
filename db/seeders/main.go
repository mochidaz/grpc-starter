package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grpc-starter/common/config"
	gdb "grpc-starter/common/gorm"
	"grpc-starter/common/postgres"
	emailEntity "grpc-starter/modules/notification/v1/entity"
	"grpc-starter/modules/user/v1/entity"
	"math/rand"
)

func main() {
	cfg, err := config.NewConfig(".env")

	if err != nil {
		panic(err)
	}

	pgInstance, err := postgres.NewPool(&cfg.Postgres)

	if err != nil {
		panic(err)
	}

	db, err := gdb.NewPostgresGormDB(pgInstance)

	if err != nil {
		panic(err)
	}

	createSampleUser(db)

	createDummyUsers(db)

	users := []*entity.User{}

	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Find(&users).
		Error; err != nil {
		panic(err)
	}

	createDummyEmails(db, users)

}

func createSampleUser(db *gorm.DB) {
	userId := uuid.New()

	if err := db.WithContext(context.Background()).
		Model(&entity.User{}).
		Create(entity.NewUser(
			userId,
			"rahman",
			"rahmanhakim2435@gmail.com",
			"rahman",
			"080000",
			"system",
		)).
		Error; err != nil {
		panic(err)
	}

}

func createDummyUsers(db *gorm.DB) {
	for i := 0; i < 99; i++ {
		userId := uuid.New()

		if err := db.WithContext(context.Background()).
			Model(&entity.User{}).
			Create(entity.NewUser(
				userId,
				fmt.Sprintf("dummy-%d", i),
				fmt.Sprintf("dummy-email-%d@dummy.com", i),
				"dummy",
				fmt.Sprintf("0800000%d", i),
				"system",
			)).
			Error; err != nil {
			panic(err)
		}
	}
}

func createDummyEmails(db *gorm.DB, users []*entity.User) {
	for i := 0; i < 500; i++ {

		random := rand.Intn(len(users))

		if err := db.WithContext(context.Background()).
			Model(&emailEntity.EmailSent{}).
			Create(emailEntity.NewEmailSent(
				users[random].ID.String(),
				"mailer@mail.com",
				users[random].Email,
				"dummy",
				"dummy",
				"sent",
				"dummy category",
				"system",
			)).
			Error; err != nil {
			panic(err)
		}
	}
}
