package cli_test

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/marco-souza/omg/cmd/cli"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	c := cli.NewCli()

	t.Run("should parse flags and args", func(t *testing.T) {
		os.Args = []string{"cmd", "-h", "-o", "output.txt", "arg1", "arg2"}

		c.Parse()

		assert.True(t, *c.ShowHelp)
		assert.Equal(t, "output.txt", *c.OutputPath)
		assert.Equal(t, []string{"arg1", "arg2"}, c.Args)
	})

	t.Run("should show usage with flags", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		flag.CommandLine.SetOutput(w)
		c.Usage()
		w.Close()

		output, err := io.ReadAll(r)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, string(output), "Usage of cmd:")
		assert.Contains(t, string(output), "-output")
		assert.Contains(t, string(output), "-help")
	})
}
