package controller

import (
	"github.com/artemmarkaryan/wb-parser/internal/sku_getter"
	"time"
)

const poolSize = 10
const coolDown = time.Second / 50

type wildberriesController struct {
	ControllerParent
}

func NewWildberriesController() *wildberriesController {
	return &wildberriesController{
		ControllerParent: *NewControllerParent(
			sku_getter.NewCSVBytesSkuGetter(),
			poolSize,
			coolDown,
		),
	}
}
