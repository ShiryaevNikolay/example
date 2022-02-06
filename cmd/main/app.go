package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/ShiryaevNikolay/example/internal/config"
	"github.com/ShiryaevNikolay/example/internal/user"
	"github.com/ShiryaevNikolay/example/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	// cfgMongo := cfg.MongoDB
	// mongoDbClient, err := mongodb.NewClient(
	// 	context.Background(),
	// 	cfgMongo.Host,
	// 	cfgMongo.Port,
	// 	cfgMongo.Username,
	// 	cfgMongo.Password,
	// 	cfgMongo.Database,
	// 	cfgMongo.AuthDB,
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// storage := db.NewStorage(mongoDbClient, cfgMongo.Collection, logger)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" { // если сокет
		/*
			/path/to/binary
			Dir() -- /path/to
		*/
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // получаем абсолютный путь к папке, где находится бинарник
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock") // склеиваем путь до папки с сокетом (app.sock - сами придумали)

		logger.Info("lesten unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else { // в остальных случаях порт
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening post %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		panic(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
