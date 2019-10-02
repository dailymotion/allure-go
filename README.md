# Allure Golang Integrations

allure-go is a Go package that provides support for Allure reports in Golang : https://github.com/allure-framework/allure2

## Installing

To start using allure-go, install Go and run `go get`:

```sh
$ go get -u github.com/dailymotion/allure-go
```

This will retrieve the library.

## Using

To implement this library in your tests, follow the examples in the [examples](example/example_test.go).
To run test, use the usual go test command :
```go test ./example```
This will automatically generate an `allure-results` folder at the root level of your project.
To see the report in html, generate it with the allure command line :
```allure serve allure-results```

## Goals

This project is an open source repository with multiple goals to achieve :
- Provide a first level of integration able to build a json container in an `allure-reports` folder for a Test file and be able to generate an Allure report based on this json container. => DONE
- Integrate Steps in the project. A Step() method will create the json object describing a step in the json container. This method can be called inside a method used by a test or directly in a test. It needs to find which container is related to the current test.
- Integrate Attachments in the project. An attachment is a file in the `allure-results` folder that can be referred in a container or a step.

The end goal is to provide the same features than https://docs.qameta.io/allure/#_java

There are possible issues and open questions that we need to address. For example, how does the history work (an Allure report can display the result of previous executions).