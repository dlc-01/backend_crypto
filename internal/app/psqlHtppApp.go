package app

import (
	"database/sql"
	"github.com/dlc-01/BackendCrypto/internal/conf"
	"github.com/dlc-01/BackendCrypto/internal/controller"
	cryptoService "github.com/dlc-01/BackendCrypto/internal/service/crypto"
	"github.com/dlc-01/BackendCrypto/pkg/coincap"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql/repositories"
	"github.com/dlc-01/BackendCrypto/pkg/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Log    *logger.Logger
	DB     *sql.DB
}

func NewApp() (*App, error) {
	cfg, err := conf.InitConf()
	if err != nil {
		return nil, err
	}

	log, err := logger.Initialize(cfg.Logger)
	if err != nil {
		return nil, err
	}

	db, err := psql.NewSQLClient(*cfg)
	if err != nil {
		log.Panicf("error while starting psql client: %v", err)
		return nil, err
	}

	coinCapClient := coincap.NewCoinCapClient()

	cryptoRepo := repositories.NewCryptoRepo(db)
	cryptoService := cryptoService.NewCryptoService(cryptoRepo, *coinCapClient)
	cryptoController := controller.NewCryptoController(cryptoService)

	router := gin.Default()
	router.GET("/cryptos", cryptoController.GetAllCryptos)
	router.GET("/cryptos/:symbol", cryptoController.GetCryptoBySymbol)
	router.POST("/crypto", cryptoController.CreateCrypto)
	router.PUT("/cryptos", cryptoController.UpdateCrypto)
	router.DELETE("/cryptos/:symbol", cryptoController.DeleteCrypto)

	return &App{
		Router: router,
		Log:    log,
		DB:     db,
	}, nil
}

func (app *App) Run(addr string) {
	go func() {
		if err := app.Router.Run(addr); err != nil {
			app.Log.Fatalf("failed to run server: %v", err)
		}
	}()
}

func (app *App) Shutdown() {
	app.Log.Info("Shutting down application...")
	if err := app.DB.Close(); err != nil {
		app.Log.Errorf("error while closing database connection: %v", err)
	}
	app.Log.Info("Application stopped gracefully.")
}
