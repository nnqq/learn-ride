package test

import (
	"context"
	"encoding/base64"
	"fmt"
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

func getTree() (*ride.Tree, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	f, err := os.Open("../ride/main.ride")
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	body, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	sc, _, err := client.NewUtils(client.Options{
		BaseUrl: "https://nodes.wavesnodes.com",
		Client:  &http.Client{Timeout: 3 * time.Second},
	}).ScriptCompile(ctx, string(body))
	if err != nil {
		return nil, fmt.Errorf("client.NewUtils.ScriptCompile: %w", err)
	}

	src, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(sc.Script, "base64:"))
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString: %w", err)
	}

	tree, err := ride.Parse(src)
	if err != nil {
		return nil, fmt.Errorf("ride.Parse: %w", err)
	}

	return tree, nil
}

func TestCompile(t *testing.T) {
	tree, err := getTree()
	if !assert.NoError(t, err, "getTree") {
		t.FailNow()
	}

	_, err = ride.Compile(tree)
	if !assert.NoError(t, err, "ride.Compile") {
		t.FailNow()
	}
}
