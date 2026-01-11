# Introduction

Just a very simple golang app called hello to understand the most basic.

In this file, we explain the high level conecots and how to test this module. For more detailed explanations, see the code and related comments.


# Concepts

### go.mod

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. 
That module is defined by a go.mod file that tracks the modules that provide those packages. 
That go.mod file stays with your code, including in your source code repository.


# Steps to execute this module
 
- cd </path/to/hello>
- go mod init golang-learning/hello
- go run .


# other useful commands

### go help 

See commands that can be executed by go

### go mod tidy

Ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

When to run it:

✅ After adding a new import that is outside the standard library
✅ After removing an import
✅ After copying code from another project that uses extra modules
✅ After updating Go version or module paths
✅ Before committing your code to make sure go.mod and go.sum are clean

# References

- https://go.dev/doc/tutorial/getting-started