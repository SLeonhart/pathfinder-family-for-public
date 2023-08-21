package scheduller

import (
	"context"
	"pathfinder-family/config"

	"pathfinder-family/infrastructure/logger"

	"github.com/jasonlvhit/gocron"
)

type Scheduller struct {
	cfg *config.Config
}

func NewScheduller(cfg *config.Config) *Scheduller {
	return &Scheduller{
		cfg: cfg,
	}
}

func (s *Scheduller) Run() {
	if s.cfg.Sheduller.OperationFrequencySec <= 0 {
		msg := "Can't parse SHEDULLER_OPERATION_FREQUENCY_SEC env variable"
		logger.Error("", msg)
		panic(msg)
	}

	gocron.Every(uint64(s.cfg.Sheduller.OperationFrequencySec)).Seconds().Do(func() {
		ctx := context.Background()
		shedullerFunc(ctx)
	})

	go s.startCron()
}

func (s *Scheduller) startCron() {
	<-gocron.Start()
}

func shedullerFunc(ctx context.Context) error {
	return nil
}
