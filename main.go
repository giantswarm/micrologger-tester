package main

import (
	"fmt"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

func main() {
	logger, err := micrologger.New(micrologger.Config{})
	if err != nil {
		panic(microerror.Mask(err))
	}

	i := 0
	for range time.Tick(1 * time.Second) {
		logger.Log("level", "info", "message", fmt.Sprintf("Generating entry %v", i))
		i++
	}
}
