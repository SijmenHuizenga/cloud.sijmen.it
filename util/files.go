package util

import (
	"github.com/GeertJohan/go.rice"
	"io/ioutil"
	"os"
)

const zzf = 0644

func CopyResourceToDisk(filename string) error {
	resources, e := rice.FindBox("res")
	if e != nil { return e }

	fidisklayout, e := resources.String(filename)
	if e != nil { return e }

	return ioutil.WriteFile(filename, []byte(fidisklayout), zzf)
}

func DeleteFileOnDisk(filename string) error {
	return os.Remove(filename)
}