# SigZag

Sigzag is a library for fingerprinting digital asset, generating manifests and managing 
digital inventory using traditional hashing functions and [Hometree](https://github.com/KevinFasusi/hometree) a homomorphic merkle tree.

- Cryptographically sign content
- Generate a manifest
- Compare manifests
- Diff manifests
- SPDX compliant formats (TBC)
- Manifests in Json and gRPC format

## Install


## Example

Generate a manifest and merkletree for all the files in a directory:

```shell
$ ./sigzag --root  path/to/directory/
```
Output:
```
manifest-0b60f8c9b4fa19bcb391276a3cdd3363d9efa0532201262cdc6c6e0881928dfa.json
```

Compare the manifests:

```shell
$ ./sigzag --compare-manifest manifest-0b60f8c9b4fa19bcb391276a3cdd3363d9efa0532201262cdc6c6e0881928dfa.json \
    manifest-0b60f8c9b4fa19bcb391276a3cdd3363d9efa0532201262cdc6c6e0881928dfa.json 
```
Output:
```
Equal:true
```

