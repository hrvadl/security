# Security (cryptography) labs

This program can

1. Encrypt/decrypt data using one of the following methods:

   - [Rearrangement](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/rearrangement/rearrangement.go) (lab 1).
   - [Ceasar](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/ceasar/ceasar.go) (lab 2).
   - [Gamma](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/gamma/cipher.go) (lab 3).

2. [Guess the key](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/analysis/decrypt.go) for the Cipher encoded data (lab 4).
3. [Create signature](https://github.com/hrvadl/security/blob/master/internal/sign/domain/sign/signer.go) and verify it using RSA keypair (lab 5).

The app reads inputs and provides outputs to the files. But it can be easily extended to work with the stdin/stdout.

## Folder structure

- [internal/sign](https://github.com/hrvadl/security/tree/master/internal/sign) contains app logic related for the signing/verifying signature of the file.
- [internal/cryptoalgo](https://github.com/hrvadl/security/tree/master/internal/cryptoalgo) contains app logic related to the implementing DIY encryption algorithms.

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

It will look for a specific file in the `static` folder, then sign it using a newly generated private key and then verify the signature with the
pulic key.
