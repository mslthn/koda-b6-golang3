# Search Name Program

### Description

Simple command-line application written in Go to search for a name inside a list of names.The program asks the user to input a keyword and then searches for matching names from a predefined slice.

### Features

This project demonstrates:
- Passing slices as function parameters
- Creating custom packages
- Basic CLI interaction
- Searching/filtering data
- Modular project structure

## Package Used

### Main Package

Entry point of the application

Utility:
- Create list of names
- Accept user input from terminal
- Call the search function from lib package

### fmt Package

Utility:
- Print output to terminal (fmt.Print)
- Read user input (fmt.Scan)

### lib Package

Utility:
- Receive slice of names
- Receive search keyword
- Compare names with keyword
- Return matching results
- Return empty slice if no match found

### Screenshot
![screenshot](./img/Screenshot%20from%202026-02-11%2018-51-25.png)