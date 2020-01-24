# Allure Golang Integration

`allure-go` is a Go package that provides all required tools to transform your tests into [Allure](https://github.com/allure-framework/allure2) reports.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
    1. [Specifying `allure-results` folder location](#specifying-allure-results-folder-location)
    2. [Optional environment variable](#optional-environment-variable) 
3. [Features](#features)
    1. [Test Wrapper](#test-wrapper)
    2. [Nested Steps](#nested-steps)
    3. [Option Slices](#option-slices)
        1. [Description](#description)
        2. [Action](#action)
        3. [Parameter](#parameter)
        4. [Parameters](#parameters)
    4. [Test-specific options](#test-specific-options)
        1. [Lead](#lead)
        2. [Owner](#owner)
        3. [Epic](#epic)
        4. [Severity](#severity)
        5. [Story](#story)
        6. [Feature](#feature)
    4. [Parameters](#parameters-as-a-feature)
    5. [Attachments](#attachments)
    6. [Setup/Teardown](#setupteardown)
    7. [Environment files](#environment-files)
4. [Feature Roadmap](#feature-roadmap) 


## Installation
In order to retrieve `allure-go` simply run the following command (assuming `Go` is installed):
```sh
$ go get -u github.com/dailymotion/allure-go
```

## Usage

### Specifying allure-results folder location
Golang's approach to evaluating test scripts does not allow to establish a single location for test results 
programmatically, the following environment variable is required for `allure-go` to work as expected. 
In order to specify the location of a report run the following:
```shell script
export ALLURE_RESULTS_PATH=</some/path>
```

`allure-results` folder will be created in that location. 

### Optional environment variable
Allure-go will retrieve the absolute path of your testing files (for example, /Users/myuser/Dev/myProject/tests/myfile_test.go) and will display this path in the reports.

To make it cleaner, you can trim prefix the path of your project by defining the `ALLURE_WORKSPACE_PATH` with the value of your project root path :
```shell script
export ALLURE_WORKSPACE_PATH=/another/path
```

You will now get the relative path of your test files from your project root level.

## Features

### Test Wrapper
`allure-go` test reporting relies on wrapping test scripts inside a call to `allure.Test()` function:
```go
allure.Test(t *testing.T, testOptions ...Option)
```
This function allows adding required modules representing a test along with additional non-essential modules like labels.
For basic usage refer to the following example:
```go
package test

import (
        "fmt"
        "github.com/dailymotion/allure-go"
        "testing"
)
 
func TestStep(t *testing.T) {
    allure.Test(t, allure.Action(func() {
                    fmt.Println("This block of code is a test")
                }))
}
```  

### Nested Steps
`allure-go` takes advantage of Allure Reports feature of nested steps and allows marking portions of the code as steps 
even if said portions are found outside test scripts in functions that are eventually called by a test script.
This is a great feature for enabling levels of detail for test reporting.
```go
package test

import (
        "fmt"
        "github.com/dailymotion/allure-go"
        "testing"
)
 
func TestStep(t *testing.T) {
    allure.Test(t, allure.Action(func() {
                    allure.Step(allure.Description("Action"),
                                    allure.Action(func() {
                                            fmt.Println("This block of code is a step")
                                    }))
                    PerformAction()
                }))
}

func PerformAction() {
    allure.Step(allure.Description("Action"),
                allure.Action(func() {
                        fmt.Println("This block of code is also a step")
                }))
}
```

### Option Slices
Option slices allow providing as much or as little detail for the test scripts as needed.

#### Description
```go
allure.Description("Description of a test or a step")
```
Provides a name for the step and a description for the test. This option is required
 
#### Action
```go
allure.Action(func() {})
```
This option is a wrapper for the portion of code that is either your test or your step. This option is required.

#### Parameter
```go
allure.Parameter(name string, value interface{})
```
This option specifies parameters for the test or step accordingly.
This particular option can be called multiple times to add multiple parameters.

#### Parameters
```go
allure.Parameters(parameters map[string]interface{})
```
This option allows specifying multiple parameters at once. 
This option can be called multiple times as well and will not interfere with previous parameter assignments.

### Test-specific options

#### Lead
```go
allure.Lead(lead string)
```
This option specifies a lead of the feature.

#### Owner
```go
allure.Owner(owner string)
```
This option specifies an owner of the feature.

#### Epic
```go
allure.Epic(epic string)
```
This option specifies an epic this test belongs to.

#### Severity
```go
allure.Severity(severity severity.Severity)
```
This option specifies a severity level of the test script to allow better prioritization.

#### Story
```go
allure.Story(story string)
```
This option specifies a story this test script belongs to. 
This particular option can be called multiple times to connect this test script to multiple stories. 

#### Feature
```go
allure.Feature(feature string)
```
This option specifies a feature this test script belongs to.
This particular option can be called multiple times to connect this test script to multiple features.

### Parameters as a feature
`allure-go` allows specifying parameters on both test and step level to further improvement informativeness of your 
test scripts and bring it one step closer to being documentation. 
In order to specify a parameter refer to the following example:
```go
package test

import (
           "github.com/dailymotion/allure-go"
           "testing"
        )     

func TestParameter(t *testing.T) {
    allure.Test(t, 
                allure.Description("Test that has parameters"),
                allure.Parameter("testParameter", "test"),
                allure.Action(func() {
                    allure.Step(allure.Description("Step with parameters"),
                                allure.Action(func() {}),
                                allure.Parameter("stepParameter", "step"))
                }))
}
```
Allure parameters integrate neatly with Golang's parameterized tests too:
```go
package test

import (
           "github.com/dailymotion/allure-go"
           "testing"
        )

func TestParameterized(t *testing.T) {
    for i := 0; i < 5; i++ {
    		t.Run("", func(t *testing.T) {
    			allure.Test(t,
    				allure.Description("Test with parameters"),
    				allure.Parameter("counter", i),
    				allure.Action(func() {}))
    		})
    	}
}
```

### Attachments
`allure-go` providing an ability to attach content of various MIME types to test scripts and steps.
The most common MIME types are available as constants in `allure-go` 
(`image/png`, `application/json`, `text/plain` and `video/mpeg`).
In order to attach content to a test or step refer to the following example:
```go
package test

import (
            "errors"
            "github.com/dailymotion/allure-go"
            "log"
            "testing"
   )

func TestImageAttachmentToStep(t *testing.T) {
	allure.Test(t, 
        allure.Description("testing image attachment"), 
        allure.Action(func() {
            allure.Step(allure.Description("adding an image attachment"), 
                allure.Action(func() {
                    err := allure.AddAttachment("image", allure.ImagePng, []byte("<byte array representing an image>"))
                    if err != nil {
                        log.Println(err)
                    }
        }))
	}))
}
```

### Setup/Teardown
Golang does not directly follow the setup/teardown approach of other languages like Java, C# and Python.
This does not prevent us from logically separating said phases in test scripts and taking 
advantage of separating these phases in the test reports too.
`allure.BeforeTest` and `allure.AfterTest` functions have to be called in sequence with the test wrapper.

`allure.BeforeTest` goes first, then `allure.Test` and finally `allure.AfterTest`. 
This is done to allow various logical conclusions, such as upon failure of `allure.BeforeTest` actual test body 
will be skipped, etc.

Please refer to the following example of setup/teardown usage:
```go
package test

import (
        "github.com/dailymotion/allure-go"
        "testing"
)
func TestAllureSetupTeardown(t *testing.T) {
	allure.BeforeTest(t,
		allure.Description("setup"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Setup step 1"),
				allure.Action(func() {}))
		}))

	allure.Test(t,
		allure.Description("actual test"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Test step 1"),
				allure.Action(func() {}))
		}))

	allure.AfterTest(t,
		allure.Description("teardown"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Teardown step 1"),
				allure.Action(func() {}))
		}))
}
```  
In case of utilizing setup/teardown in a parameterized test, all of the test phases have to be called 
inside `t.Run(...)` call.

### Environment Files
`allure-go` allows taking advantage of the environment file feature of Allure reports. You can specify report-specific 
variables that you want to appear in the report such as browser kind and version, OS, environment name, etc.
In order to do that create an `environment.xml` or `environment.properties` file as instructed [here](https://docs.qameta.io/allure/#_environment) and define an
environment variable for `allure-go` to incorporate in the results:
```shell script
export ALLURE_ENVIRONMENT_FILE_PATH=<path to environment file>
```

## Feature Roadmap
`allure-go` is still in active development and is not yet stabilized or finalized.

Here is a list of improvements that are needed for existing features :
- [X] Manage failure at step level : the project must be able to mark a Step as failed (red in the report) when an assertion fails. The assertion stacktrace must appear in the report.
- [X] Manage error at test and step level : the project must be able to mark a Test and Step as broken (orange in the report) when an error happens. The error stacktrace must appear in the report.
- [X] Add support of Links
- [ ] Add support of Unknown status for Test object
- [X] Add support of Set up in Execution part of a Test
- [X] Add support of Tear down in Execution part of a Test
- [X] Add support of Severity
- [ ] Add support of Flaky tag
- [X] Add support of Categories
- [X] Add support of Features
- [X] Add support for environment files
- [ ] Add support for history