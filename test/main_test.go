package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wavesplatform/gowaves/pkg/ride"
	"io"
	"os"
	"testing"
)

func TestCompile(t *testing.T) {
	f, err := os.Open("../ride/main.ride")
	if !assert.NoError(t, err, "os.Open") {
		t.FailNow()
	}

	body, err := io.ReadAll(f)
	if !assert.NoError(t, err, "io.ReadAll") {
		t.FailNow()
	}

	//client.NewUtils().ScriptCompile()

	tree, err := ride.Parse(body)
	if !assert.NoError(t, err, "ride.Parse") {
		t.FailNow()
	}

	_, err = ride.Compile(tree)
	if !assert.NoError(t, err, "ride.Compile") {
		t.FailNow()
	}
}
