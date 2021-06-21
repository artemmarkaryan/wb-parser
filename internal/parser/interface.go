package parser

type Parser interface {
	GetInfo(body []byte) (infos map[string]string, err error)
}
