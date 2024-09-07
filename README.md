# Security (cryptography) labs

This program can

1. Encrypt/decrypt data using one of the following methods:

   - [Rearrangement](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/rearrangement/rearrangement.go) (lab 1).
   - [Caesar](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/ceasar/ceasar.go) (lab 2).
   - [Gamma](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/cipher/gamma/cipher.go) (lab 3).

2. [Guess the key](https://github.com/hrvadl/security/blob/master/internal/cryptoalgo/domain/analysis/decrypt.go) for the Cipher encoded data (lab 4).
3. [Create signature](https://github.com/hrvadl/security/blob/master/internal/sign/domain/sign/signer.go) and verify it using RSA keypair (lab 5).

The app reads inputs and provides outputs to the files. But it can be easily extended to work with the stdin/stdout.

## Folder structure

- [internal/sign](https://github.com/hrvadl/security/tree/master/internal/sign) contains app logic related for the signing/verifying signature of the file.
- [internal/cryptoalgo](https://github.com/hrvadl/security/tree/master/internal/cryptoalgo) contains app logic related to the implementing DIY encryption algorithms.

## How to run ciphers?

**NOTE**: Program works correcly only with latin alphabet letters, both uppercase and lowercase. It's not recommended to use
cyryllic letters or any other symbols.

Make sure you have [taskfile](https://taskfile.dev/) and [Go](https://go.dev/) installed. Then to encrypt/decrypt data using
Caesar cipher from the root of the repo run:

```sh
task encrypt-caesar
```

It will look for `in.txt` and `key.txt` files in the `static/cipher/rearrangement` folder, then it will encrypt `in.txt` using `key.txt`
and create `encrypted.cipher.txt` and `decrypted.cipher.txt`. The `key.txt` should contain only one integer digit which will be used as a key(shift).
In example: `echo "4" >> key.txt`

To run a program, which will encrypt/decrypt data using Gamma cipher you can run:

```sh
task encrypt-gamma
```

It will look for `in.txt` files in the `static/cipher/rearrangement` folder, then it will encrypt `in.txt`
and create `encrypted.cipher.txt` and `decrypted.cipher.txt`.

To run a program, which will encrypt/decrypt data using rearrangement cipher you can run:

```sh
task encrypt-rearrangement
```

It will look for `in.txt` and `key.txt` files in the `static/cipher/rearrangement` folder, then it will encrypt `in.txt` using `key.txt`
and create `encrypted.cipher.txt` and `decrypted.cipher.txt`. The `key.txt` should contain comma separated integer list with indexes of the
word. For example, to encrypt word "Kyiv" (**four** letter) we should provide list with **four indexes** `echo "4,3,2,1" >> key.txt`

## How to create signature?

To run a program, which will create and verify signature for a given file you can run:

```sh
task sign
```

It will look for a file named `in.txt` in the `static/sign` folder, then sign it using a newly generated private key and then verify the signature with the
pulic key.
