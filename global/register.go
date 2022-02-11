package global

import (
	"errors"
	"fmt"
	"os"
)

var (
// model *GlobalRegister
// once  sync.Once
)

type IRegister interface {
	Init() error
}

type GlobalRegister struct {
	servers []IRegister
}

// func instance() *GlobalRegister {
// 	once.Do(func() {
// 		model = &GlobalRegister{}
// 	})
// 	return model
// }

func Register(models ...IRegister) *GlobalRegister {
	if len(models) <= 0 {
		// panic(errors.New("No services have been loaded yet."))
		fmt.Errorf("%v", errors.New("No services have been loaded yet."))
		os.Exit(1)
	}
	model := &GlobalRegister{}
	model.servers = make([]IRegister, len(models))
	for i, m := range models {
		model.servers[i] = m
	}
	return model
}

func (g *GlobalRegister) Init() {
	for _, svc := range g.servers {
		must(svc.Init())
	}
}

func must(err error) {
	if err != nil {
		fmt.Errorf("%v", errors.New("No services have been loaded yet."))
		os.Exit(1)
	}
}
