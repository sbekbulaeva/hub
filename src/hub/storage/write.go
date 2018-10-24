package storage

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"hub/aws"
	"hub/config"
	"hub/util"
)

func Write(data []byte, files *Files) []error {
	errs := make([]error, 0)

	for _, file := range files.Files {
		failedToWrite := false
		switch file.Kind {
		case "s3":
			err := aws.WriteS3(file.Path, bytes.NewReader(data))
			if err != nil {
				msg := fmt.Sprintf("Unable to write `%s` %s file: %v", file.Path, files.Kind, err)
				if aws.IsSlowDown(err) && (len(files.Files) > 1 || config.Force) {
					util.Warn("%s", msg)
					failedToWrite = true
				} else {
					errs = append(errs, errors.New(msg))
				}
			}

		case "fs":
			out, err := os.Create(file.Path)
			if err != nil {
				log.Fatalf("Unable to open `%s` %s file for write: %v", file.Path, files.Kind, err)
			}
			wrote, err := out.Write(data)
			err2 := out.Close()
			if err != nil || wrote != len(data) || err2 != nil {
				if err == nil && err2 != nil {
					err = err2
				}
				err = fmt.Errorf("Unable to write `%s` %s file (wrote %d out of %d bytes): %s",
					file.Path, files.Kind, wrote, len(data), util.Errors2(err))
				errs = append(errs, err)
			}
		}

		if config.Verbose && !failedToWrite {
			log.Printf("Wrote %s `%s`", files.Kind, file.Path)
		}
	}

	if len(errs) == 0 {
		errs = nil
	}

	return errs
}
