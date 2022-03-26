package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"go-ddd-blog/internal/domain/idgen"
	idgenImpl "go-ddd-blog/internal/infra/idgen"
	"go-ddd-blog/internal/infra/persist"
	postApp "go-ddd-blog/internal/service/post/api"
	postCmd "go-ddd-blog/internal/service/post/api/command"
	"go-ddd-blog/internal/service/post/domain/event"
	postRepo "go-ddd-blog/internal/service/post/infra/persist"
	userApp "go-ddd-blog/internal/service/user/api"
	userCmd "go-ddd-blog/internal/service/user/api/command"
	eventHandler "go-ddd-blog/internal/service/user/api/event_handler"
	userRepo "go-ddd-blog/internal/service/user/infra/persist"
)

func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(&userRepo.UserPO{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&userRepo.MessagePO{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&postRepo.CommentPO{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&postRepo.PostPO{}); err != nil {
		panic(err)
	}
}

func InitEventHandler(h eventHandler.CommentAddedHandler) {
	event.RegisterPostNewCommentEventHandler(h.Handle)
}

func main() {
	app := fx.New(
		fx.Supply(
			//persist.NewSqliteDB("test.db"),
			persist.NewPostgresDB("host=localhost user=postgres password=123456 dbname=postgres"),
		),
		fx.Provide(
			idgenImpl.NewRandomIDGen,

			postRepo.NewRepo,
			userRepo.NewRepo,

			postRepo.NewQuery,
			userRepo.NewQuery,

			userCmd.NewCommand,
			postCmd.NewCommand,

			userApp.NewApplication,
			postApp.NewApplication,

			eventHandler.NewCommentNotificationHandler,

			Init,
		),

		fx.Invoke(
			InitDB,
			idgen.InitIDGen,
			InitEventHandler,
			func(r *gin.Engine) {
				err := r.Run() // listen and serve on 0.0.0.0:8080
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
