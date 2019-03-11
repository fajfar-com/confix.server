package configStrorage

import "github.com/fajfar-com/confix.server/internal"

type ConfigStorage interface {
	Get(app string)
}

type fooStorage struct {
	applications []internal.Application
}

