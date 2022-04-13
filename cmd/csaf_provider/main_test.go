package main

import (
	"os"
	"testing"
)

// call realMain() with Args that skip over params used by "go test"
//
// use like
//   go test -c -vet=off -covermode=atomic -o csaf_provider.debug
//   cp csaf_provider.debug /usr/lib/cgi-bin/
//
//   pushd /usr/lib/cgi-bin
//   mv csaf_provider.go csaf_provider2.go
//   echo '#!/bin/bash
//   exec /usr/lib/cgi-bin/csaf_provider.debug -test.coverprofile=/tmp/csaf_provider-itest-${EPOCHREALTIME}.cov -- "$@"
//   ' >csaf_provider.go
//   chmod a+x csaf_provider.go
//
// then do a cgi-bin action on the provider like using the uploader
//
// If you want to merge several runs, try
//   go install github.com/wadey/gocovmerge@b5bfa59ec0adc420475f97f89b58045c721d761c
//   ~/go/bin/gocovmerge /tmp/csaf_provider-ites*.cov >csaf_provider-itest-merged.cov
//   cd ~/csaf_distribution
//   go tool cover -func=~/csaf_provider-itest-merged.cov

func TestMain(t *testing.T) {
	var endOfTestParams int
	for i, a := range os.Args[1:] {
		if a == "--" {
			endOfTestParams = i + 1
		}
	}

	if endOfTestParams == 0 {
		t.Skip("skipping integration test, no `--` parameter found")
	}
	realMain(os.Args[endOfTestParams+1:])
}
