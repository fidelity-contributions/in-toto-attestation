# in-toto attestation implementation tests

We provide a set of basic tests for the different language
bindings for the in-toto attestation layers.

## Testing the Go bindings

The go packages `go/v1` and `go/predicates` provide a number of tests
for the statement and predicate layers.

### Running the Go tests

To run all tests:

```shell
make go_test
```

### Writing new Go tests

Please use the standard [Go testing package] to write tests
for new predicates. For example tests, take a look at the `*_test.go`
files in the `go/` directory tree.

At a minimum, we suggest testing JSON marshalling and unmarshalling
of the Go language bindings.

## Testing the Python bindings

The `tests/python` directory contains a number of tests for the statement
and predicate layers.

### Running the Python tests

To run all tests:

```shell
make py_test
```

### Writing new Python tests

Please use the standard [Python unittest package] to write tests
for new predicates. For example tests, take a look at the `test_*.py`
modules in the `tests/python/` directory tree.

At a minimum, we suggest testing JSON marshalling and unmarshalling
of the Python language bindings.

## Testing the Rust bindings

The Rust package `rust/` provides a few tests for the statement and predicate
layers.

Note: Support for Rust is still very early, so consider these bindings to not
be stable or well-tested.

### Pre-requisites

Install a Rust toolchain (1.85 or later)
(see https://www.rust-lang.org/tools/install)

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Then, install required libraries

```bash
sudo apt install libssl-dev protobuf-compiler
```

### Running the Rust tests

To run all tests (including doctests):

```bash
cd rust
cargo test
```

[Go testing package]: https://pkg.go.dev/testing
[Python unittest package]: https://docs.python.org/3/library/unittest.html
