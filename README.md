# Security (cryptography) labs

This program can encrypt/decrypt data using one of the following methods:

- Rearrangement
- Ceasar
- Gamma

This program can also guess key for the Cipher encoded data.

## How to run?

Make sure you have [taskfile](https://taskfile.dev/) and [Go](https://go.dev/) installed. Then from the root of the repo run:

```sh
task run-cryptoalgo
```

This will encrypt/decrypt data using DIY algorithm. The program will ask you to give it file
with the key (if needed) as well as file with the input data. Then, you'll be asked to decide,
whether you want to encrypt or decrypt it. The result will be stored in the given output file.

To run a program, which will create and verify signature for a given file you can run:

```sh
task run-sign
```
