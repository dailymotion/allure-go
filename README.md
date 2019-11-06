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

## Features available

This project is still in progress. Here is the list of available features :
- It is possible to create a test object in the report with the `allure.Test()` method.
- It is possible to create a step object inside a test object of the report with the `allure.Step()` method.
- It is possible to pass parameters to a test object or a step object with the `allure.TestWithParameters()` and `allure.StepWithParameters()` methods.
- It is possible to add attachments to a test object or a step object with the `allure.AddAttachment()` method (depending if it is called inside an `allure.Test()` wrapper or an `allure.Step()` wrapper).

## Improvements needed

Here is the list of improvements that are needed for existing features :
- [ ] Manage failure at step level : the project must be able to mark a Step as failed (red in the report) when an assertion fails. The assertion stacktrace must appear in the report.
- [ ] Manage error at test and step level : the project must be able to mark a Test and Step as broken (orange in the report) when an error happens. The error stacktrace must appear in the report.
- [ ] Add support of Links
- [ ] Add support of Unknown status for Test object
- [ ] Add support of Set up in Execution part of a Test
- [ ] Add support of Tear down in Execution part of a Test
- [ ] Add support of Severity
- [ ] Add support of Flaky tag
- [ ] Add support of Categories
- [ ] Add support of Features