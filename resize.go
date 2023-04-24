package goterm

import (
  "fmt"
  "os/exec"
)


func Resize(w, h int) (err error) {
  cmd := exec.Command("resize", "-s", fmt.Sprintf("%d", h), fmt.Sprintf("%d", w))
  return cmd.Run()
}
