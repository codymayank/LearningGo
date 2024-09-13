### Table of Contents

* [File Structure](#file-structure)
* [Installing Go dependencies](#installing-go-dependencies)
* [What is go mod tidy?](#what-is-go-mod-tidy)
* [Go basic Programming Concepts](#go-basic-programming-concepts)
  * [How to run a Go program](#how-to-run-a-go-program)
  * [Loops and DataStructures](#loops-and-datastructures)
    * [variables](#variables)
    * [Functions](#functions)
  * [Error Handling](#error-handling)
  * [Try Catch in Go](#try-catch-in-go)
  * [Logging in Go](#logging-in-go)
  * [How to export a function in Go](#how-to-export-a-function-in-go)
* [Resource links](#resource-links)

##### About GO

##### What is Go?

- Go is designed to be simple and easy to use, with a clean and concise syntax that makes it easy to read and write
  code.
- Go is designed to be fast and efficient, with a focus on performance and scalability.
- Go is designed to be reliable and robust, with built-in support for error handling and testing.
- Go is designed to be secure, with built-in support for concurrency and parallelism.
- Go is designed to be portable, with support for multiple platforms and architectures.
- Go is designed to be open source, with a large and active community of developers contributing to its development.

##### When to use Go?

- Go is a great choice for building web applications, APIs, and microservices.
- Go is a great choice for building command-line tools, utilities, system software and network services.
- Go is a great choice for building high-performance and scalable applications.

#### File Structure

`go.mod` = like a dependency module where you mention all the dependencies of your project like package.json or
requirements.txt

- For the actual development, the module value is the name of the github repository where the project is hosted
  `go.sum` = like a lock file where you mention the exact version of the dependencies
  `main.go` = the main file of your project

**NOTE:** In go, code is organized into packages and these packages are part of modules. A module is a collection of
related
packages that are versioned together.

#### Installing Go dependencies

To install the go dependencies, look up for the pkg.go.dev website and search for the package you want to install. For
example, if you want to install the `gorilla/mux` package, you can search for it on the website and you will find the
import path. You can then use the `go get` command to install the package. For example: `go get github.com/gorilla/mux`

#### What is go mod tidy?

`go mod tidy` is a command that removes any dependencies that are not used in the project. It also adds any dependencies

#### Go basic Programming Concepts

##### How to run a Go program

```aiignore
go run main.go
```

##### Loops and DataStructures

###### variables

- Go is a statically typed language, which means that the type of a variable is known at compile time.

```aiignore
var x int = 5
```

If you want to declare and initialize a variable at the same time, you can use the `:=` operator.

```aiignore
y := 10 which is similar to var y int = 10
```

###### Functions

- Functions in Go are defined using the `func` keyword. The return type of the function is specified after the parameter
  list.

```aiignore
func add(x int, y int) int {
    return x + y
}
```

##### Error Handling

- In Go, errors are values that can be returned from functions. You can use the `errors.New` function to create a new
  error value.

```aiignore
err := errors.New("Something went wrong")
```

##### Try Catch in Go

In go, equivalent of null is `nil` and equivalent of `try..catch` is `defer..recover`

##### Logging in Go

```aiignore
import "log"
log.Println("Hello, World!") - This will log the message to the console
log.Fatalf("An error occurred: %v", err) - This will log the error and exit the program
log.Panicf("An error occurred: %v", err) - This will log the error and panic ( Consider Panic more similar to typescript `throw` )
```

Here `%v` is a placeholder for the value of the `err` variable while `%s` is a placeholder for the string value of the
`err` variable.

##### How to export a function in Go

There is no special keyword to export a function in Go. If you want a function to be accessible from other packages, you
just need to `start the function name with an uppercase letter`. For example

```shell
# Here Add is exported
func Add(x int, y int) int { 
    return x + y
}
```

```aiignore

#### How to publish a Go module
 
To publish a go module, follow the following procedure:
```aiignore
go mod init github.com/username/repo
```

#### Resource links

- Tutorial Doc: https://go.dev/doc/tutorial/