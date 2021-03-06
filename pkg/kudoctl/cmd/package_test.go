package cmd

import (
	"os"
	"testing"

	"github.com/kudobuilder/kudo/pkg/kudoctl/files"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var packageCmdArgs = []struct {
	name         string
	arg          []string
	errorMessage string
}{
	{"expect exactly one argument", []string{}, "expecting exactly one argument - directory of the operator to package"}, // 1
	{"empty string argument", []string{""}, "invalid operator in path: "},                                                // 2
	{"invalid operator", []string{"foo"}, "invalid operator in path: foo"},                                               // 3
	{"valid operator", []string{"/opt/zk"}, ""},                                                                          // 4
}

func TestTableNewBundleCmd(t *testing.T) {
	fs := afero.NewMemMapFs()
	files.CopyOperatorToFs(fs, "../bundle/testdata/zk", "/opt")
	for _, test := range packageCmdArgs {
		newCmdBundle := newPackageCmd(fs, os.Stdout)
		err := newCmdBundle.RunE(newCmdBundle, test.arg)
		if err != nil {
			assert.Equal(t, test.errorMessage, err.Error(), test.name)
		}
	}
}
