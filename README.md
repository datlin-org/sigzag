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
    "asset": "sigzag/.gitignore",
    "sha256": "3f65909ea8ecb4d4ad55863aec833a10b7d6592a405bf1f2fa2078b8017be353",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/LICENSE",
    "sha256": "6b789c0126d3800da6b229a29b7b47fbd02e3b432dc63e50a144e71c2dc24fe6",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/README.md",
    "sha256": "a6c99b38459ac583106c3de4c245d93d2e9a343b443bb91ebeedb06aa91a8ca8",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/pdf/testdata2.pdf",
    "sha256": "66d443c39adea589b5bcd508d5b2d85356c7921b51d3636be40f28757b384d21",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/pdf/testdata5.pdf",
    "sha256": "f7d0391e2d3dfb284aef716743690c035d43d5b72e1f5b569a08c08b095b9dba",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/testdata1.md",
    "sha256": "e6dbed6368694ce31739583c9f93d7fa97aaa7e0d0549ff00810ee193e10e392",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/testdata2.md",
    "sha256": "c8821204fd0c26e338d7895303087bb6be33b197bae47a8523d4c1805d5c96bd",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/pdf/testdata1.pdf",
    "sha256": "47da23b69a3f5329d41886bccef8eb3e112d477e8a5ef3da4994ed2e78376b0f",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/testdata3.md",
    "sha256": "9b0bfdf44481eb979dd3ed67cede9b4c0e265d2769a8858f35877f705743e3bd",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  },
  {
    "asset": "sigzag/testdata/pdf/testdata3.pdf",
    "sha256": "976a0c9710bdc2e745c8d5a5e466d3181bd5e899b5d46170efbf379dd46bd320",
    "timestamp": "Wed Mar 27 10:27:37 GMT 2024"
  }
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

Prepend output file with alternative string:

```shell
$ ./sigzag  --output-file tango
```

Output:
```
tango-manifest-71cad6089ac2a094f068c302d31e11949c0125543eb432032268ccc658ffc3de.json
tango-merkletree-af78ae7ee5e3151d8d2d3d80da098b5f24fa53302b6db09f6e3f869e072a9e0c.json
```

## Flags

| Flag               | Description                                              |
|:-------------------|:---------------------------------------------------------|
| --root             | Root directory to descend. Defaults to working directory |
| --level            | Directory nesting depth (default==3)                     |
| --output-file      | Prepends file with string                                |
| --compare-manifest | Compare manifest                                         |
| --compare-merkle   | Compare merkle tree                                      |