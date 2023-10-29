# SVault

Simple key-value secret Vault inspired by Hasicorp Vault. Project tries to explore the inner working of vault engine with golang. SVault uses AES256-GCM encryption for secret store. This is a exploratory project & intend to stay as Experimental.

* Simple KV storage for secrets.
* Industry standard aes256-gcm secret encryption.
* Easy cli interface.

## Usage guide

* Download pre-compiled binaries from [release page](https://github.com/ragul28/svault/releases/latest).

```sh
tar -xvf svault_x.x.x_darwin_amd64.tar.gz
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
svault list
svault get github_token
```

## Basic CLI usage
```
Usage:
  svault [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete secret from vault store
  get         Get secret from vault store
  help        Help about any command
  init        Init vault secret engine
  list        List stored secret from vault store (Flages: -f, --freshInit)
  status      Get the status of the vault engine
  store       Store secret to vault store

Flags:
  -h, --help               help for get
  -m, --masterkey string   Pass masterkey as flag
```

## Build Process

Build using go environment.
```sh
git clone https://github.com/ragul28/svault
cd svault
make build
```