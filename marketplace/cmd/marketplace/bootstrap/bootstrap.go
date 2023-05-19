package bootstrap

import (
	. "barbz.dev/marketplace/internal/infrastructure/repository/ad"
	"barbz.dev/marketplace/internal/infrastructure/server"
	. "barbz.dev/marketplace/internal/pkg/application/ad"
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type config struct {
	// API configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8000"`
	ShutdownTimeout time.Duration `default:"10s"`
}

var cfg config

func Run() error {
	// Load config
	if err := envconfig.Process("marketplace", &cfg); err != nil {
		return err
	}

	dependencies := loadServices()
	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, dependencies)
	return srv.Run(ctx)
}

func loadServices() (mapDependencies map[string]interface{}) {
	saveAd, findAllAds, findAdById := initAdsDependencies()
	mapDependencies["ad.SaveAd"] = saveAd
	mapDependencies["ad.FindAllAds"] = findAllAds
	mapDependencies["ad.FindAdById"] = findAdById

	return
}

func initAdsDependencies() (save SaveAd, findAll FindAllAds, findById FindAdById) {
	ads := NewInMemoryRepository()
	save = NewSaveAd(ads)
	findAll = NewFindAllAds(ads)
	findById = NewFindAdById(ads)
	return
}
