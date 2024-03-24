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
Json file Output:
```
manifest-0b60f8c9b4fa19bcb391276a3cdd3363d9efa0532201262cdc6c6e0881928dfa.json
```

Json contents:

```
[
  {
    "asset": "config",
    "sha256": "13620931f1e68e7196450da8af6fe88f0a15edc3da9da7fd641f892d4ef24557"
  },
  {
    "asset": "description",
    "sha256": "85ab6c163d43a17ea9cf7788308bca1466f1b0a8d1cc92e26e9bf63da4062aee"
  },
  {
    "asset": "fsmonitor-watchman.sample",
    "sha256": "e0549964e93897b519bd8e333c037e51fff0f88ba13e086a331592bf801fa1d0"
  },
  {
    "asset": "applypatch-msg.sample",
    "sha256": "0223497a0b8b033aa58a3a521b8629869386cf7ab0e2f101963d328aa62193f7"
  }...
  ]
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

