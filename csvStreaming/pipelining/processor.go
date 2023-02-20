package pipelining

import (
	"github.com/personhashing/models"
	"github.com/personhashing/worker"
)

type JSONReader interface {
	Read() (models.JSON, error)
}

type JSONWriter interface {
	Write(models.JSON) error
}

func PipelineProcessor(r JSONReader, w JSONWriter, p []models.Process) error {
	InCh := make(chan models.JSON)
	var (
		tempIn <-chan models.JSON
		outCh  <-chan models.JSON
	)

	go func() {
		for d, err := r.Read(); err == nil; d, err = r.Read() {
			InCh <- d
		}
		close(InCh)
	}()
	tempIn = InCh

	for i := range p {
		outCh = worker.Worker(p[i], tempIn)
		tempIn = outCh
	}

	for d := range outCh {
		if err := w.Write(d); err != nil {
			return err
		}
	}
	return nil

}
