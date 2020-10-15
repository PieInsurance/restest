# Restest

A test REST API.

## Setup

This code uses [Go](https://golang.org/) and you are expected to [download](https://golang.org/dl/) and run it locally.

To run the application, run this command:

```
go run main.go
```

this will stand up a service on port `8000`.

## Assignment

We want to test the functionality of our [temperature](https://en.wikipedia.org/wiki/Temperature) API. We _think_ it
works well, but need more system level testing.

## Examples:

These examples are from a [well known table](https://en.wikipedia.org/wiki/Temperature#Examples) and work.

```http://localhost:8000/1811```

```
[
    "1.811000E+00 kK",
    "1.537850e+03 °C",
    "1.600221E+03 nm"
]
```

```http://localhost:8000/373.1339```

```
[
    "3.731339E-01 kK",
    "9.998390e+01 °C",
    "7.766649E+03 nm"
]
```

```http://localhost:8000/16E6```

```
[
    "1.600000E+01 MK",
    "1.599973e+07 °C",
    "1.811250E-01 nm"
]
```
