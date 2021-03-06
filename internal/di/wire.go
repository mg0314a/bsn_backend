// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"bsn_backend/internal/dao"
	"bsn_backend/internal/server/grpc"
	"bsn_backend/internal/server/http"
	"bsn_backend/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
