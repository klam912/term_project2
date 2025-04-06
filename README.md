# Term project part 2
This repo contains folders each demonstrating a unique/interesting feature of Go.

## Getting started
Clone the git repo
```bash
git clone https://github.com/klam912/term_project2.git
cd term_project2
```

## Run any of the `.go` files in each folder
Select a folder to examine a feature of Go:
```bash
cd folder-name
```

If there's already a `go.mod` file, do not run `go mod init` again. 
If not, run:
```bash
go mod init term_project2/folder-name
```

Once you have the `go.mod` file, run the file:
```bash
go run file-name.go
```

### Running Fuzz Test
Fuzz test is a unique feature in Go to automate the testing process (more efficient than unit test).
To run `fuzz_test.go` file, run:
```bash
go test -fuzz=FuzzTestName
```
For the `-fuzz` flag, you must have Fuzz as the first part of the word and then insert whatever test name you want to specifically run. Otherwise, `go test` would run all test.

In the function MSE, there's a commented line that can show you what happens when FuzzTest catches an error from its generated inputs.
FuzzTest will keep finding more edge cases and won't stop running until it encounters one.
To stop the test, exit the program using CTRL-C.