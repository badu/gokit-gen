package main

import (
	"bytes"
	"log"
	"testing"
)

func Test(t *testing.T) {
	log.SetFlags(0)
	log.SetPrefix("gokit-gen: ")

	args := []string{"", "/home/badu/appetize/protorepo/src/bitbucket.org/appetize_backend/protorepo/stock/stock_transaction.proto", "./../../templates", "bitbucket.org/appetize_backend/stock-transactions/pkg/stock", "debug"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
}
