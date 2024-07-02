# Cron Expression Parser in Go

This is a simple Go program that parses a cron expression and expands its fields into a list of values.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Notes](#notes)

## Installation

1. **Install Go**:
   Ensure you have Go installed on your machine. You can download and install it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**:
   Clone this repository or copy the `cron_parser.go` file to your local machine.

## Usage

### Running the Program

1. **Navigate to the Directory**:
   Open a terminal and change to the directory containing the `cron_parser.go` file.

   ```sh
   cd path/to/your/directory

2. **Run with go run**:
    Run the program with a cron expression as an argument.
    `go run cron_parser.go "*/15 0 1,15 * 1-5 /usr/bin/find"`

3. **Or Compile and Run**:
    Alternatively, you can compile the program into an executable and then run it.
    `go build cron_parser.go`
    
    `./cron_parser "*/15 0 1,15 * 1-5 /usr/bin/find"`


## Example

Given the example cron expression `*/15 0 1,15 * 1-5 /usr/bin/find`


The output might look like this:

    minute         0 15 30 45
    hour           0
    day of month   1 15
    month          1 2 3 4 5 6 7 8 9 10 11 12
    day of week    1 2 3 4 5

## Notes
- The program parses the time fields of a cron expression: minute, hour, day of month, month, and day of week.
- It does not process the command part of the cron expression.
- The cron expression should have exactly 6 fields.
- The program exits with an error message if the number of fields in the cron expression is incorrect.