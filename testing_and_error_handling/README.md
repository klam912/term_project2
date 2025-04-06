# Term Project

This folder contains a Go program that prints greeting messages for a list of names and demonstrates Go's testing and error handling capabilities.

## Requirements

- Go 1.24 or higher

## Installation

### Step 1: Install Go

First, ensure Go is installed on your system. If not, you can follow the instructions on the official Go installation page: [Install Go](https://golang.org/doc/install).

### Step 2: Clone the Repository

Clone the repository to your local machine:
```bash
git clone https://github.com/klam912/term_project2.git
cd term_project2/testing_and_error_handling/
```

### Step 3: Build the program
To build the Go program, run:
```bash
go build
```

### Step 4: Install the program
Once you have built the Go program, you must install the program into executable binary.
Before you install, you must add the Go install directory to your system's shell path so that you don't need to specify where the executable is.

You can find your Go install path, where the go command will install the current package.
```bash
go list -f '{{.Target}}'
```
Then, add the path to $PATH.
```bash
export PATH=$PATH:/path/to/your/install/directory
```

Once you have updated the shell path, install the package.
```bash
go install
```

### Step 5: Run the program 
You can now run your program by typing its name.
```bash
main
```