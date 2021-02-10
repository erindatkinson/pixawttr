package convert

import (
	"os/exec"

	"github.com/pkg/errors"
)

//Merge makes a composite image
func Merge(bgImage string, weatherImage string, outFile string) error {
	cmd := exec.Command("composite", "-geometry", "+50+50", weatherImage, bgImage, outFile)
	combined, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(combined))
	}

	return nil
}
