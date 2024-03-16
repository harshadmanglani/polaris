---
sidebar_position: 2
title: Get Started
tags: 
    - go
---
# Getting started with Polaris

## Install Go
Make sure you have Go installed. These tutorials were produced using Go 1.21

Check your version of Go with the following command:
```
go version
```

This will return your installed Go version:

```
go version go1.21.5 darwin/arm64
```
## Install Polaris

If you are creating a new project using Polaris, you can start by creating a new directory:

```
mkdir goproject
```

Next, switch to the new directory:
```
cd goproject
```

Then, initialize a Go project in that directory:
```
go mod init
```

Finally, install the Polaris with go get:
```
go get github.com/harshadmanglani/polaris
```