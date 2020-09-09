# SVault

Simple key-value secret Vault inspired by Hasicorp Vault. Project tries to explore the inner working of vault engine with golang. SVault uses AES256-GCM encryption for secret store. This is a exploratory project & intend to stay as Experimental.

* Simple KV storage for secrets.
* Industry standard aes256-gcm secret encryption.
* Easy cli interface.

## Basic Usage

* Download pre-compiled binaries from [release page](https://github.com/ragul28/svault/releases).

```sh
tar -xvf svault_linux_arm64.tar.gz
chmod +x svault
sudo mv svault /usr/local/bin
```

* Init svault secret engine which generates svault masterkey.
```sh
svault init
```

* Store & get secret from key vault.
```sh
export MASTER_KEY=<master_key>

svault store github_token xxxxxxxx
svault get github_token
```

## Build Process

Build using go environment.
```sh
git clone https://github.com/ragul28/svault
cd svault
make build
```

## CheckList

- [x] Implement cli like interface for init, store & get
- [ ] Add persistance storage using boltdb.
- [ ] Explore the stream interface from crypto/cipher lib for interface chaining. 
- [ ] Add rest/jrpc interface.