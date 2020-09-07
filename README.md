# SVault

Simple key-value secret Vault inspired by Hasicorp Vault. Project try to explore the inner working of vault engine with golang. SVault uses AES256-GCM encryption for secret store.This a exploitative project intend to stay as Experimental.

* Secret KV storage
* Secret encryption AES256-GCM
* Easy cli interface  

## Build Process

Build using go environment.
```sh
git clone https://github.com/ragul28/svault
cd svault
make build
```

## CheckList

- [ ] Make cipher as common package.
- [ ] Add persistance storage using boltdb.
- [ ] Implement cli like interface for init, store & get
- [ ] Explore the stream interface from crypto/cipher lib for interface chaining. 
- [ ] Add rest/jrpc interface.