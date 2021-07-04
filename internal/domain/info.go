// Структура Info хранит полученные с сайта данные

package domain

type Info map[string]string

func (i *Info) Map() map[string]string {
	return *i
}

func (i *Info) Pop(key string) {
	delete(*i, key)
}