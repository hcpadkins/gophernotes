// this file was generated by gomacro command: import _b "hash/adler32"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"hash/adler32"
)

// reflection: allow interpreted code to import "hash/adler32"
func init() {
	Packages["hash/adler32"] = Package{
	Binds: map[string]Value{
		"Checksum":	ValueOf(adler32.Checksum),
		"New":	ValueOf(adler32.New),
		"Size":	ValueOf(adler32.Size),
	}, Untypeds: map[string]string{
		"Size":	"int:4",
	}, 
	}
}