# EncryptedBase64 - a lightweight Vigenère-inspired cipher

This is a lightweight, simple implementation of a **Vigenère cipher-based algorithm** with an additional step of making an output a **valid Base64 string**.

The difference from a _traditional_ Vigenère implementation is we are using an _entire_ space of byte ordinals (not _just letters_), thus making the implementation compatible with strings containing numbers, punctuation and control characters. One can even encrypt secrets for storage i.e. in configuration files, making them harder (but not impossible!) to crack in case of a config leak*.

The encryption format is compatible and interchangeable with the one used in [this PHP snippet](https://gist.github.com/kvasilov48/29b7e0d3a0f2d8c0b996).

_* -> Of course if the encryption key is not included in the same config, otherwise it's trivial!_

## Practical advice

* Keys longer than the actual data being processed make no practical sense, as only the number of key bytes (from its beginning) equal to the number of bytes in the input string is actually used for encryption (see [the definition of a Vigenère cipher and how it works](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher))
* Never store the key in the same place as input or output data, if you're going to use the algorithm for secrets protection.
* Never use this library to encrypt stuff you're going to put into a VCS (GitHub, GitLab, SVN, etc.) - secrets are never safe there, regardless of their form!
* The exact same key used to encrypt the input data is also needed to decrypt it. If you try to use a different key, you will get either total gibberish or at least a distorted message (depending on how much the used key differs from the needed one).

## Functions

### `Encrypt(inputString string, encryptionKey string) string`

Encrypts the given `string` using a given key (also `string`). Returns `string` containing encrypted data, Base64-encoded.

### `Decrypt(inputString string, encryptionKey string) (string, error)`

Tries to decode the given Base64 `string`, than decrypt its contents using a given key (also `string`). Returns `string` containing decrypted data. If a failure occured during Base64 decoding, an `error` will be passed and the returned string will be empty.

## Contributions

You are welcome to post your contributions in pull requests, as well as post feature requests and bug reports using GitHub Issues. We'll try to review all those and make something good from them :)

## License

The code is licensed under the BSD license. You may use and rework it at your discretion, provided the authorship is retained. Please refer to the `LICENSE.md` file for full terms.
