Terraform Provider Encode
=======================

Terraform provider that implements text encodings other than base64.

Currently this is done via data sources. 
This is a stop-gap solution until Terraform supports plugins that provide functions:

- [Allow plugins to export custom functions #2771](https://github.com/hashicorp/terraform/issues/2771)
- [Extending terraform with custom functions #27696](https://github.com/hashicorp/terraform/issues/27696)

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) >= v1.2.x (Coould work with older versions too, untested.)
- [Go](https://golang.org/doc/install) >= 1.16

Building The Provider
---------------------

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command: 
```sh
$ make install
```

Adding Dependencies
---------------------

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.


Using the provider
----------------------

Fill this in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `make install`. This will build the provider and put the provider binary in the `~/.terraform.d/plugins/` directory.

In order to run the full suite of Acceptance tests, run `make testacc`.

```sh
$ make testacc
```
