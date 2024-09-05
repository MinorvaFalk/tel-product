package main

import (
	"tel/product/config"
	apihandler "tel/product/internal/api/handler"
	apirepository "tel/product/internal/api/repository"
	apiusecase "tel/product/internal/api/usecase"
	"tel/product/internal/router"
	"tel/product/pkg/datasource"
	"tel/product/pkg/logger"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	db := datasource.NewGorm()
	repo := apirepository.New(db)
	uc := apiusecase.New(repo)
	h := apihandler.New(uc)

	router := router.New(h)
	router.MapApiHandler()

	router.Run()
}
