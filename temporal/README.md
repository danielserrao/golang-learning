# Introduction

This simple app, creates a workflow on Temporal. Know more about it at https://temporal.io/.

# Steps to execute this app

### Initiate go.mod

- cd </path/to/temporal-directory>
- go mod init golang-learning/temporal

### Install Dependencies

- go get go.temporal.io/sdk/client
- go get go.temporal.io/sdk/workflow


### Make sure that Temporal frontend API port is accessible

Usually the default port is 7233 the host that is running this app should be able to access the port. 

The way to guarantee will depend on your setup, but in this example we will assume that temporal is running in a kubernetes cluster and you are running the script from your local.

For this use port-forward. You can do it using [k9s](https://k9scli.io/) or manually as explained at [Forward a local port to a port on the Pod](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/#forward-a-local-port-to-a-port-on-the-pod).

### Run the app

- go run .