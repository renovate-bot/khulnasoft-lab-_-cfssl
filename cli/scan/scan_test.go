package scan

import (
	"testing"

	"github.com/khulnasoft-lab/cfssl/cli"
)

var hosts = []string{"www.khulnasoft.com", "google.com"}

func TestScanMain(t *testing.T) {
	err := scanMain(hosts, cli.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = scanMain(nil, cli.Config{Hostname: "www.khulnasoft.com, google.com", List: true})
	if err != nil {
		t.Fatal(err)
	}
}
