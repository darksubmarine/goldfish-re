# Goldfish-RE
[![Goldfish-RE Build and Tests](https://github.com/darksubmarine/goldfish-re/actions/workflows/ci-cd.yml/badge.svg?branch=main)](https://github.com/darksubmarine/goldfish-re/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/darksubmarine/goldfish-re)](https://goreportcard.com/report/github.com/darksubmarine/goldfish-re)
[![GoDoc](https://pkg.go.dev/badge/github.com/darksubmarine/goldfish-re?status.svg)](https://pkg.go.dev/github.com/darksubmarine/goldfish-re?tab=doc)
[![Release](https://img.shields.io/github/release/darksubmarine/goldfish-re.svg?style=flat-square)](https://github.com/darksubmarine/goldfish-re/releases)


Reactive and Embeddable rules engine library written in pure Go!.

This rules engine has been thought to trigger automatically an event letting you know that
a condition from your ruleset has been satisfied by some updated fact into your context.

The evaluation algorithm is RETE based with focus on evaluation and memory using a Trie struct to improve its performance.

### Supported data types
The rule engines expose different data types to work with:

- **String**: This is a well known `string` type
- **Number**: A number is a representation of an integer value. In this case is a wrapper for a `int64` data type
- **Float**: Floats are also numeric types. They represent the decimal numbers implemented as `float64`
- **Boolean**: The boolean are useful to assert a condition as `true` or `false`
- **Date**: Represents a `time.Time` data type

For more details, please read the [documentation](https://darksubmarine.com/docs/goldfish-re)

## Installation
To install the library into your project please run:

```text
go install github.com/darksubmarine/goldfish-re
```

## Examples
 
 - [Starterkit](https://github.com/darksubmarine/goldfish-re/tree/master/examples/starterkit)
 - [Advance](https://github.com/darksubmarine/goldfish-re/tree/master/examples/advanced)


## License

[Apache License](LICENSE)