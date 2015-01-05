package core

import (
  "strings"
)

const (
  Major = "0"
  Minor = "0"
  Revision = "0"
)

var Version = func() string {
  return strings.Join([]string{Major, Minor, Revision}, ".")
}
