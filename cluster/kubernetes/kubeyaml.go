package kubernetes

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// KubeYAML is a placeholder value for calling the helper executable
// `kubeyaml`.
type KubeYAML struct {
}

// Image calls the kubeyaml subcommand `image` with the arguments given.
func (k KubeYAML) Image(in []byte, ns, kind, name, container, image string) ([]byte, error) {
	args := []string{"image", "--namespace", ns, "--kind", kind, "--name", name}
	args = append(args, "--container", container, "--image", image)
	return execKubeyaml(in, args)
}

// Annotate calls the kubeyaml subcommand `annotate` with the arguments as given.
func (k KubeYAML) Annotate(in []byte, ns, kind, name string, policies ...string) ([]byte, error) {
	args := []string{"annotate", "--namespace", ns, "--kind", kind, "--name", name}
	args = append(args, policies...)
	return execKubeyaml(in, args)
}

func execKubeyaml(in []byte, args []string) ([]byte, error) {
	kubeyaml, err := exec.LookPath("kubeyaml")
	if err != nil {
		kubeyaml = os.ExpandEnv("${PLUGINS_PATH}/kubeyaml/kubeyaml")
	}
	cmd := exec.Command(kubeyaml, args...)
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	cmd.Stdin = bytes.NewBuffer(in)
	cmd.Stdout = out
	cmd.Stderr = errOut

	err = cmd.Run()
	if err != nil {
		return nil, errors.New(strings.TrimSpace(errOut.String()))
	}
	return out.Bytes(), nil
}
