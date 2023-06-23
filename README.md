
# Loggo

[![GitHub version](https://badge.fury.io/gh/haywardrl%2Floggo.svg)](https://badge.fury.io/gh/haywardrl%2Floggo)

Loggo is a simple and flexible logging library for Go, designed with simplicity and efficiency in mind. With multiple logging levels, format string support, and customizable output, Loggo provides the basic essentials without being overly complex.

## Table of Contents

-   [Getting Started](#getting-started)
-   [Installation](#installation)
-   [Usage](#usage)
-   [Features](#features)

## Getting Started

To install Loggo, ensure you have installed Go (version 1.14 or later is recommended), then follow these instructions:

## Installation

Use the `go get` command to download and install Loggo:


``` shell
go get github.com/haywardrl/loggo
```

## Usage

Import Loggo into your Go program. When you create a new logger, specify the logging level. Then call the appropriate method (`Debugf`, `Infof`, `Warnf`, `Errorf`, `Fatalf`) to log messages:

``` go
package main

import (
    "github.com/your-username/loggo"
)

func main() {
    logger := loggo.New(loggo.LevelInfo)

    logger.Infof("This is an info message")
    logger.Warnf("This is a warning message")
    logger.Errorf("This is an error message")
}
```

In this example, `Debug` messages are ignored because the logger's level is set to `LevelInfo`.

## Features

**Flexible Logging Levels**: Loggo supports multiple logging levels, including Debug, Info, Warn, Error, and Fatal. This lets you control the verbosity of your logging output.

**Formatting Support**: Loggo's logging methods support string formatting, similar to `fmt.Printf`. This allows you to easily include variable content in your log messages.

**Customizable Output**: By default, Loggo writes to `os.Stdout`, but you can customize this to write to any `io.Writer`.

