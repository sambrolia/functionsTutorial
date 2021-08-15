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

func prepareTea(resources *teaResources, milkRequired bool, sugarRequired int) ([]string, error) {
	var cmd []string
	cmd = append(cmd, "Add water to kettle and boil")

	if resources.cleanMugs > 0 {
		cmd = append(cmd, "Get a mug")
		resources.cleanMugs = resources.cleanMugs - 1
	} else {
		return nil, fmt.Errorf("There are no clean mugs.")
	}

	if resources.teabagsLeft > 0 {
		cmd = append(cmd, "Add teabag to mug")
		cmd = append(cmd, "Pour boiled water into mug")
		cmd = append(cmd, "Stir")
		cmd = append(cmd, "Remove teabag")
		resources.teabagsLeft = resources.teabagsLeft - 1
	} else {
		return nil, fmt.Errorf("There are no teabags remaining.")
	}

	if sugarRequired > 0 {
		if resources.sugarLeft >= sugarRequired {
			cmd = append(cmd, "Add sugar to mug")
			cmd = append(cmd, "Stir")
			resources.sugarLeft = resources.sugarLeft - sugarRequired
		} else {
			return nil, fmt.Errorf(
				"There are %v spoonfuls of sugar remaining, but %v were requested",
				resources.sugarLeft,
				sugarRequired)
		}
	}

	if milkRequired {
		if resources.isMilkFresh {
			cmd = append(cmd, "Add milk to mug")
			cmd = append(cmd, "Stir")
		} else {
			return nil, fmt.Errorf("Milk required, but has gone off.")
		}
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
