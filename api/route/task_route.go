package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khoahase151217/go-clean-api-architecture/api/controller"
	"github.com/khoahase151217/go-clean-api-architecture/bootstrap"
	"github.com/khoahase151217/go-clean-api-architecture/domain"
	"github.com/khoahase151217/go-clean-api-architecture/mongo"
	"github.com/khoahase151217/go-clean-api-architecture/repository"
	"github.com/khoahase151217/go-clean-api-architecture/usecase"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
