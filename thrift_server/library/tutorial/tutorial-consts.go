// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tutorial

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"reardrive/thrift_server/shared"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = shared.GoUnusedProtection__
const INT32CONSTANT = 9853
var MAPCONSTANT map[string]string

func init() {
MAPCONSTANT = map[string]string{
  "goodnight": "moon",
  "hello": "world",
}

}

