package main

import "os"

func commandExit(ctrl *Controller) error {
	os.Exit(0)
	return nil
}
