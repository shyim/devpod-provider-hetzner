package hetzner

import (
	"errors"
	"fmt"
)

var (
	ErrMultipleVolumesFound = func(name string) error {
		return fmt.Errorf("multiple volumes with name %s found", name)
	}
	ErrUnknownDiskImage = errors.New("unknown disk image")
	ErrUnknownMachineID = errors.New("unknown machine id")
	ErrUnknownRegion    = errors.New("unknown region")
)