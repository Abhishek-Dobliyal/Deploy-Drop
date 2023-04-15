package main

import (
  "github.com/Abhishek-Dobliyal/deploy-drop/cmd"
  "github.com/common-nighthawk/go-figure"
)

func main() {
  figure.NewColorFigure("Deploy-Drop", "doom", "blue", true).Print()
  cmd.Execute()
}