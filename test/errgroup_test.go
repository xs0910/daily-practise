package test

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
)

func TestErrGroup(t *testing.T) {
	var g errgroup.Group
	var urls = []string{
		"https://www.golang.org/",
		"https://www.google.com/",
	}

	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
