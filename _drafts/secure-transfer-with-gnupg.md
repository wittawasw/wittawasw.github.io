---
layout: post
title: Secure Transfer with GnuPG
---

```shell
# install
brew install gnupg

# generate
# RSA 3072
gpg --full-generate-key

# list keys
gpg -k
```

dd if=/dev/zero of=largefile bs=1M count=1000

From the list of GPG keys, copy the long form of the GPG key ID you'd like to use. In this example, the GPG key ID is 3AA5C34371567BD2:
Shell


$ gpg --list-secret-keys --keyid-format=long
/Users/hubot/.gnupg/secring.gpg
------------------------------------
sec   4096R/3AA5C34371567BD2 2016-03-10 [expires: 2017-03-10]
uid                          Hubot <hubot@example.com>
ssb   4096R/4BB6D45482678BE3 2016-03-10

Paste the text below, substituting in the GPG key ID you'd like to use. In this example, the GPG key ID is 3AA5C34371567BD2:

gpg --armor --export 3AA5C34371567BD2
# Prints the GPG key ID, in ASCII armor format

Copy your GPG key, beginning with -----BEGIN PGP PUBLIC KEY BLOCK----- and ending with -----END PGP PUBLIC KEY BLOCK-----.
