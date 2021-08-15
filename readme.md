# Example project to show correct usage of functions:

https://github.com/sambrolia/functionsTutorial

## Context:
The file tea.go is a basic example of a function which puts together instructions to make a cup of tea, based on inputs such as if milk or sugar is required.  
A structure is used to store resources such as how much sugar is remaining, or whether the milk is gone off.  
A simple test file is added which just tests a single usecase (no sugar remaining) for illustration purposes.  

## Changes:
The main changes between functions1 and functions2 is that the code is broken down into smaller functions to achieve specific goals.  

## Result: 
Improvements in:  
- Readability
- Testability
- Maintainability
- Reuse for other usecases (recycling addMilk or addSugar functions for prepareCoffee functionality)
