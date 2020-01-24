package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building 2...")

	//##RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

	cmd := exec.Command("CGO_ENABLED=0", "GOOS=linux", "GOARCH=amd64", "go", "build", "-o", buildPath(), root())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
