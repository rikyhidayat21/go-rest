//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService() *SimpleService {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil
}
