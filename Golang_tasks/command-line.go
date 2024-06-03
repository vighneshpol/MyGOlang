package main

import (
    "fmt"
    "io"
    "log"
    "os/exec"
)

func main() {

    cmd := exec.Command("cat")
    stdin, err := cmd.StdinPipe()
}