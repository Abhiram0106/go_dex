package main

import (
	"errors"
	"fmt"
)

func commandMapf(ctrl *Controller, input *[]string) error {
	if ctrl.nextURL == nil && ctrl.previousURL != nil {
		return errors.New("You are on the first page")
	}

	res, err := ctrl.httpClient.GetLocations(ctrl.nextURL)
	if err != nil {
		return err
	}
	ctrl.nextURL = res.Next
	ctrl.previousURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(ctrl *Controller, input *[]string) error {
	if ctrl.previousURL == nil {
		return errors.New("You are on the first page")
	}

	res, err := ctrl.httpClient.GetLocations(ctrl.previousURL)
	if err != nil {
		return err
	}
	ctrl.nextURL = res.Next
	ctrl.previousURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
