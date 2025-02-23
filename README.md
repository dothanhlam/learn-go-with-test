# learn-go-with-test

This repository follows the origin website https://quii.gitbook.io/learn-go-with-tests, which adding some modification to inspect the GOlang features. I added "codility" package with takes some tests that I did during the interviews

If you are new beginner with golang, I recommend you setup the Go properly on your working computer before starting the test. You need to install gvm from this repo https://github.com/moovweb/gvm

Then use the gvm to install go 1.21.8. This is the version I used to test against all the testcases. Keep in mind that the repo is for my personal purpose, so do not blame for any reason you use and take the code with your own risk as well

## create new test suite
To create new test suite, make new folder with desired name. For example, "banana"
```bash
$ mkdir banana
```
Then, change into the folder and init the go module
```bash
$ cd banana
$ go mod init banana
```
## run the test
Change into the folder you want to take a test, for example "banana", run a command
```bash
$ cd banana
$ go test -v
```
