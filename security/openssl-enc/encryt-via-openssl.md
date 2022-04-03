## Encrypt files via openssl

Sometimes you want to store some secrets in your git repository. It is a good practice to have them encrypted if stored externally (publicly).
There are dedicated solutions such as [`git-secret`](https://git-secret.io/) or [`git-crypt`](https://github.com/AGWA/git-crypt). However, you can also achieve similar result with the [`openssl`](https://www.openssl.org/) tool, which is quite often already preinstalled on your host.

## Example

To decrypt [`example.yaml.enc`](./example.yaml.enc) stored in this repository,run:
> You have to provide password `okon`.

```bash
./security/openssl-enc/decrypt.sh
enter aes-256-cbc decryption password:
Files decrypted successfully
```

This script creates [`security/openssl-enc/assets/decrypted/](./assets/decrypted) (git ignored) directory containing file with decoded value:
```yaml
hakuna: matata
```

If you need to update secrets in this repo overwrite plaintext file and execute [`./security/openssl-enc/encrypt.sh`](./encrypt.sh)
