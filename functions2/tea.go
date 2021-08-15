package main

import (
	"fmt"
)

type teaResources struct {
	cleanMugs int
	teabagsLeft int
	isMilkFresh bool
	sugarLeft int // amount in teaspoons
}

func getMug(resources *teaResources) ([]string, error) {
	var cmd []string
	if resources.cleanMugs > 0 {
		resources.cleanMugs = resources.cleanMugs - 1
		cmd = append(cmd, "Get a mug")
		return cmd, nil
	} else {
		return nil, fmt.Errorf("There are no clean mugs.")
	}
}

func addTeabag(resources *teaResources) ([]string, error) {
	var cmd []string
	if resources.teabagsLeft > 0 {
		resources.teabagsLeft = resources.teabagsLeft - 1
		cmd = append(cmd, "Add teabag to mug")
		cmd = append(cmd, "Pour boiled water into mug")
		cmd = append(cmd, "Stir")
		cmd = append(cmd, "Remove teabag")
		return cmd, nil
	} else {
		return nil, fmt.Errorf("There are no teabags remaining.")
	}
}

func addSugar(resources *teaResources, sugarRequired int) ([]string, error) {
	var cmd []string
	if resources.sugarLeft >= sugarRequired {
		cmd = append(cmd, "Add sugar to mug")
		cmd = append(cmd, "Stir")
		resources.sugarLeft = resources.sugarLeft - sugarRequired
		return cmd, nil
	} else {
		return nil, fmt.Errorf(
			"There are %v spoonfuls of sugar remaining, but %v were requested",
			resources.sugarLeft,
			sugarRequired)
	}
}

func addMilk(resources *teaResources) ([]string, error) {
	var cmd []string
	if resources.isMilkFresh {
			cmd = append(cmd, "Add milk to mug")
			cmd = append(cmd, "Stir")
			return cmd, nil
		} else {
			return nil, fmt.Errorf("Milk required, but has gone off.")
		}
}

func prepareTea(resources *teaResources, milkRequired bool, numSugarRequired int) ([]string, error) {
	var cmd []string
	var err error
	var stepCmd []string
	cmd = append(cmd, "Add water to kettle and boil")

	stepCmd, err = getMug(resources)
	if err != nil {
		return nil, err
	}
	cmd = append(cmd, stepCmd...)
	stepCmd = nil

	stepCmd, err = addTeabag(resources)
	if err != nil {
		return nil, err
	}
	cmd = append(cmd, stepCmd...)
	stepCmd = nil

	if (numSugarRequired > 0) {
		stepCmd, err = addSugar(resources, numSugarRequired)
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, stepCmd...)
		stepCmd = nil
	}

	if (milkRequired) {
		stepCmd, err = addMilk(resources)
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, stepCmd...)
		stepCmd = nil
	}

	return cmd, nil
}


func makeCupOfTea(resources *teaResources, milkRequired bool, sugarRequired int) {
	fmt.Printf(
		"\nRequest recieved to prepare tea, milk required: %v, sugars required: %v \n",
		milkRequired,
		sugarRequired)

	var commands, err = prepareTea(resources, milkRequired, sugarRequired)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for index, command := range commands {
		fmt.Printf("%v) %v.\n", index+1, command)
	}
}

func main() {
	var resources = teaResources{cleanMugs: 4, teabagsLeft: 4, isMilkFresh: true, sugarLeft: 10}
	makeCupOfTea(&resources, false, 0)
	makeCupOfTea(&resources, false, 1)
	makeCupOfTea(&resources, true, 0)
	makeCupOfTea(&resources, true, 6)
}
