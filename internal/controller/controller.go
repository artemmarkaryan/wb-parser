package controller

import (
	"bytes"
)


type Info map[string]string

func (i *Info) Map() map[string]string {
	return *i
}

type Controller interface {
	ProcessBytes(data *[]byte) (buff *bytes.Buffer, err error)
}

