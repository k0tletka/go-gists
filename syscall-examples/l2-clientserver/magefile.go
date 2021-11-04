//+build mage

package main

import (
    "errors"
    "os"

    "github.com/magefile/mage/sh"
)

const (
    outputPath = "./artifacts"
)

var (
    // Errors
    ErrMustRunFromRoot = errors.New("build must run with root privileges")

    // Build config
    binaryList = map[string]string{
        "./client.go": outputPath + "/client",
        "./server.go": outputPath + "/server",
    }
)

func Build() error {
    // Check root privileges
    if os.Geteuid() != 0 {
        return ErrMustRunFromRoot
    }

    for mfp, op := range binaryList {
        if err := buildBinary(mfp, op); err != nil {
            return err
        }
    }

    return nil
}

// Build binaries
func buildBinary(mainFilePath, outputPath string) error {
    var err error

    if err = sh.RunV("go", "build", "-o", outputPath, mainFilePath); err != nil {
        return err
    }

    if err = sh.RunV("setcap", "cap_net_raw+ep", outputPath); err != nil {
        return err
    }

    return nil
}
