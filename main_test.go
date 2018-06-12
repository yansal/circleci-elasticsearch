package main

import (
	"context"
	"testing"

	"github.com/olivere/elastic"
)

func TestMain(t *testing.T) {
	client, err := elastic.NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(exists)
}
