package ozon

import (
	"github.com/artemmarkaryan/wb-parser/internal/controller/base_controller"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"time"
)

const poolSize = 10
const coolDown = time.Second / 100

type ozonController struct {
	base_controller.BaseController
}


func NewOzonController() *base_controller.BaseController {
	return base_controller.NewBaseController(
		poolSize,
		poolSize,
		ozonController{},
	)
}

func (r ozonController) parse(data *[]byte) (infoArr []domain.Info) {
	return
}