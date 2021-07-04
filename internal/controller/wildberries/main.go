package wildberries

import (
	"github.com/artemmarkaryan/wb-parser/internal/controller/base_controller"
	"time"
)

const poolSize = 10
const coolDown = time.Second / 50

type wildberriesController struct {
	base_controller.BaseController
}

func NewWildberriesController() *wildberriesController {
	c := wildberriesController{}
	c.SetCoolDown(coolDown)
	c.SetConnectionPoolSize(poolSize)
	return &c
}
