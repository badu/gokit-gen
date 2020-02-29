package main

import (
	"bytes"
	"log"
	"testing"
)

/**

args := []string{"", "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/pkg/proto/testdata/test.proto", "./../../templates", "bitbucket.org/appetize_backend/protorepo/stock/bitbucket.org/appetize_backend/stock_transaction", "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/pkg/proto/testdata/"}
*/
func Test(t *testing.T) {
	log.SetFlags(0)
	log.SetPrefix("gokit-gen: ")

	args := []string{"", "/home/badu/appetize/protorepo/src/bitbucket.org/appetize_backend/protorepo/stock/stock_transaction.proto", "./../../templates", "bitbucket.org/appetize_backend/protorepo/stock/bitbucket.org/appetize_backend/stock_transaction", "/home/badu/appetize/protorepo/src/bitbucket.org/appetize_backend/protorepo/stock/bitbucket.org/appetize_backend/stock_transaction/"}

	var stdout bytes.Buffer

	err := run(args, &stdout)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
}
