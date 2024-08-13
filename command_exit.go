package main

import "os"

func commandExit(ctrl *Controller, input *[]string) error {
	os.Exit(0)
	return nil
}
