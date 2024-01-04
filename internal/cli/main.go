package cli

import (
	"github.com/Distributed-Lab-Testing/example-svc/internal/config"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service"
	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

var (
	app = kingpin.New("example-svc", "Example service for demonstration purposes. Manages notes in the database")

	runCmd     = app.Command("run", "run command")
	serviceCmd = runCmd.Command("service", "run service's API")

	migrateCmd     = app.Command("migrate", "migrate command")
	migrateUpCmd   = migrateCmd.Command("up", "migrate database up")
	migrateDownCmd = migrateCmd.Command("down", "migrate database down")
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		service.Run(cfg)
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}
	return true
}
