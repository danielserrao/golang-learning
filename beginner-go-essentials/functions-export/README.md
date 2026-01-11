# Introduction

This small module is an example of how to make functions private or public to be reused by other modules.

In this file, we explain on a high level how to make functions private or public and how to test this module. For more detailed explanations, see the code and related comments.


# Concepts

In Go, a function is exported (i.e., visible outside its package) if its name starts with an uppercase letter.

Conversely, if a function's name starts with a lowercase letter, it is unexported (i.e., only visible within its package).


# Steps to test this module
 
- cd </path/to/functions-export>
- go mod init golang-learning/functions-export
- go run .