package test

import (
	"context"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/ride"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCompile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	f, err := os.Open("../ride/main.ride")
	if !assert.NoError(t, err, "os.Open") {
		t.FailNow()
	}

	body, err := io.ReadAll(f)
	if !assert.NoError(t, err, "io.ReadAll") {
		t.FailNow()
	}

	sc, _, err := client.NewUtils(client.Options{
		BaseUrl: "https://nodes.wavesnodes.com",
		Client:  &http.Client{Timeout: 3 * time.Second},
	}).ScriptCompile(ctx, string(body))
	if !assert.NoError(t, err, "client.NewUtils.ScriptCompile") {
		t.FailNow()
	}

	src, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(sc.Script, "base64:"))
	if !assert.NoError(t, err, "base64.StdEncoding.DecodeString") {
		t.FailNow()
	}

	tree, err := ride.Parse(src)
	if !assert.NoError(t, err, "ride.Parse") {
		t.FailNow()
	}

	_, err = ride.Compile(tree)
	if !assert.NoError(t, err, "ride.Compile") {
		t.FailNow()
	}
}
