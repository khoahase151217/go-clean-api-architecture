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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
