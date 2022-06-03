package utility

import (
  "os"
)

func ReadFile(path string) (string, error) {
  contents, err := os.ReadFile(path)
  return string(contents), err
}
