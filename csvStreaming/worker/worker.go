package worker

import "github.com/personhashing/models"

func Worker(p models.Process, input <-chan models.JSON) <-chan models.JSON {
	output := make(chan models.JSON)

	go func() {
		defer close(output)
		for d := range input {
			output <- p(d)
		}
	}()
	return output
}
