# Allure Golang Integrations

allure-go is a Go package that provides support for Allure reports in Golang : https://github.com/allure-framework/allure2

## Installing

To start using allure-go, install Go and run `go get`:

```sh
$ go get -u github.com/dailymotion/allure-go
```

This will retrieve the library.

## Using

First of all, you need to set an environment variable to define the `allure-results` folder location:
```
export ALLURE_RESULTS_PATH=/some/path
```
This should be the path where the `allure-results` folder exists or should be created, not the path of the folder itself.

The `allure-results` folder is automatically created if it doesn't exist with `drwxr-xr-x` file system permission.

To implement this library in your tests, follow the [examples](example/example_test.go).

Execute tests with the usual go test command :
```
go test ./example
```

This will automatically generate an `allure-results` folder in the path defined by ALLURE_RESULTS_PATH.

To see the report in html, generate it with the allure command line :
```
allure serve $ALLURE_RESULTS_PATH/allure-results
```
This will automatically generate and open the HTML reports in a browser.

The results file are compatible with the [Jenkins plugin](https://wiki.jenkins.io/display/JENKINS/Allure+Plugin).

## Optional environment variable

Allure-go will retrieve the absolute path of your testing files (for example, /Users/myuser/Dev/myProject/tests/myfile_test.go) and will display this path in the reports.

To make it cleaner, you can trim prefix the path of your project by defining the `ALLURE_WORKSPACE_PATH` with the value of your project root path :
```
export ALLURE_WORKSPACE_PATH=/another/path
```

You will now get the relative path of your test files from your project root level.

## Goals

This project is an open source repository with multiple goals to achieve :
- [x] Provide a first level of integration able to build a json container in an `allure-results` folder for a Test file and be able to generate an Allure report based on this json container.
- [ ] Integrate Steps in the project. A Step() method will create the json object describing a step in the json container. This method can be called inside a method used by a test or directly in a test. It needs to find which container is related to the current test.
- [ ] Integrate Attachments in the project. An attachment is a file in the `allure-results` folder that can be referred in a container or a step.

The end goal is to provide the same features than https://docs.qameta.io/allure/#_java

There are possible issues and open questions that we need to address. For example, how does the history work (an Allure report can display the result of previous executions).