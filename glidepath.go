package main

import "os/exec"

func main() {
  app := "export"
  arg0 := "GOPATH=$(glide gopath)"

  cmd := exec.Command(app, arg0)
  out, err := cmd.Output()

  if err != nil {
    println(err.Error())
    return
  }

  print(string(out))
}
