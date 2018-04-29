package util

import (
	"github.com/GeertJohan/go.rice"
	"io/ioutil"
	"os"
)

const zzf = 0644

func CopyResourceToDisk(resourcename string, targetname string) error {
	cmdPrinter.Println("[copying resource "+resourcename+" to disk on "+targetname+"]")
	resources, e := rice.FindBox("res")
	if e != nil { return e }

	str, e := resources.String(resourcename)
	if e != nil { return e }

	return ioutil.WriteFile(targetname, []byte(str), zzf)
}

func DeleteFileOnDisk(filename string) error {
	return os.Remove(filename)
}