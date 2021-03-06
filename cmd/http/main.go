/**
 * @Author: Hongker
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:08
 */

package main

import (
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/component/mysql"
	"github.com/ebar-go/gateway/cmd/http/route"
	handlerImpl "github.com/ebar-go/gateway/cmd/http/handler/impl"
	"github.com/ebar-go/gateway/internal/domain/entity"
	repositoryImpl "github.com/ebar-go/gateway/internal/domain/repository/impl"
	serviceImpl "github.com/ebar-go/gateway/internal/domain/service/impl"
	"log"
)

func main() {
	app := ego.App()

	if err := app.LoadConfig("conf/app.yaml"); err != nil {
		log.Fatalf("load config failed: %v\n", err)
	}
	repositoryImpl.Inject(app.Container())
	serviceImpl.Inject(app.Container())
	handlerImpl.Inject(app.Container())

	if err := app.Container().Invoke(AutoMigrate); err != nil {
		log.Fatalf("auto migrate failed: %v\n", err)
	}

	if err := app.LoadRouter(route.Loader); err != nil {
		log.Fatalf("load router failed: %v\n", err)
	}

	app.ServeHTTP()

	app.Run()
}

func AutoMigrate(db mysql.Database) error {
	return db.GetInstance().Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&entity.UpstreamEntity{}, &entity.EndpointEntity{})
}
