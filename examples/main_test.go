package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/vpofe/go-http-client/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests for package examples")

	gohttp.StartMockServer()

	os.Exit(m.Run())
}
