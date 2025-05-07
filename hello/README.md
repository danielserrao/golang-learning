# Introduction

Just a very simple golang app called hello to understand the most basic.


# Concepts

### go.mod

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. 
That module is defined by a go.mod file that tracks the modules that provide those packages. 
That go.mod file stays with your code, including in your source code repository.


# Steps to execute this app
 
- cd </path/to/hello-directory>
- go mod init golang-learning/hello
- go run .


# References

- https://go.dev/doc/tutorial/getting-started