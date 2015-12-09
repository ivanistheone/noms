package main

import (
	"os"
	"path/filepath"

	"github.com/attic-labs/noms/d"
	"github.com/attic-labs/noms/tools/fileutil"
	"github.com/attic-labs/noms/tools/runner"
)

func main() {
	// ln -sf ../../js/.babelrc .babelrc hack, because zip files screw up symlinks.
	path, err := filepath.Abs(".babelrc")
	d.Chk.NoError(err)
	fileutil.ForceSymlink("../../js/.babelrc", path)

	runner.ForceRun("./link.sh")
	runner.ForceRun("npm", "install")
	if _, present := os.LookupEnv("NOMS_SERVER"); !present {
		os.Setenv("NOMS_SERVER", "http://localhost:8000")
	}
	if _, present := os.LookupEnv("NOMS_DATASET_ID"); !present {
		os.Setenv("NOMS_DATASET_ID", "crunchbase/index")
	}
	runner.ForceRun("npm", "run", "build")
}