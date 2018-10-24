package storage

import (
	"log"
	"strings"

	"hub/util"
)

func printFiles(files []File, kind string) {
	log.Printf("%s %s:", strings.Title(kind), util.Plural(len(files), "file"))
	for _, file := range files {
		locked := ""
		if file.Locked {
			locked = " [locked]"
		}
		if file.Exist {
			log.Printf("\t%s%s; size = %d; mod time = %v", file.Path, locked, file.Size, file.ModTime)
		} else {
			log.Printf("\t%s (not found)%s", file.Path, locked)
		}
	}
}
