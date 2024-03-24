# SigZag

Sigzag is a small utility for signing digital assets and generating a manifests.

- Cryptographically sign content
- Generate a manifest
- Compare manifests
- Diff manifests
- Manifests in Json and gRPC format

## Install

Clone repository:

```shell
git clone https://github.com/KevinFasusi/sigzag.git
```

Build:

```shell
go build cmd/main.go

```

## Example

Generate a manifest for all the files in a directory:

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

