# Sigzag

Sigzag (*work in progress*) is a small utility for signing digital assets and generating manifests.

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
go build main.go -o sigzag
```

### Linux

Move the binary from the build step into `/opt/utils` directory and add to your distribution's `PATH` variable:
```
export PATH="$PATH:/opt/utils"
```

## Example

Generate a manifest for all the files in a directory:

```shell
$ sigzag --root  path/to/directory/
```
Json file Output:
```
manifest-2024328-8853-402cd35c615b384c807cb1adb71923ff1ac0f8da06aadd0eb20a568b6a1f3609.json
```

The manifest file name is composed as `manifest-date-time-SHA256.json`

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

### Compare Manifests

Compare the manifests:

```shell
$ sigzag --compare-manifest manifest-2024328-8853-402cd35c615b384c807cb1adb71923ff1ac0f8da06aadd0eb20a568b6a1f3609.json \
manifest-2024328-8857-18fb3f4bcad83115ca08d2de456b63582acc7bf97e233f83f985ce63b4a9c50d.json
```
Output:
```
Equal:true
```

### Diff Manifests

```shell
$ sizg --diff manifest-2024328-8853-402cd35c615b384c807cb1adb71923ff1ac0f8da06aadd0eb20a568b6a1f3609.json \
manifest-2024328-8857-18fb3f4bcad83115ca08d2de456b63582acc7bf97e233f83f985ce63b4a9c50d.json
```
writes to `diff-date-time.json` file:

```
diff-2024328-81357.json
```

### Compile History for an Asset

The following will search all manifest (using the wildcard) jsons for entries equal to the asset specified:

```shell
$ sigzag --asset testdata/testdata1.md --history manifest-*.json
```

writes output to a `history-date-time.json` file:

```
history-2024329-135354.json
```

file contents:
```
[
  {
    "Asset": "testdata/testdata1.md",
    "History": [
      {
        "asset": "testdata/testdata1.md",
        "sha256": "84ef1496749b2316580228132c2b5b4f084f9ba7e3eaa7227e4c7d3e72a956ec",
        "timestamp": "Fri Mar 29 13:48:49 GMT 2024"
      },
      {
        "asset": "testdata/testdata1.md",
        "sha256": "e6dbed6368694ce31739583c9f93d7fa97aaa7e0d0549ff00810ee193e10e392",
        "timestamp": "Fri Mar 29 13:53:11 GMT 2024"
      },
      {
        "asset": "testdata/testdata1.md",
        "sha256": "977f147f643fd9d1fe78c54afcdcb078d46b8bbc6acbba3d39206fb63482ecfc",
        "timestamp": "Fri Mar 29 13:53:17 GMT 2024"
      },
      {
        "asset": "testdata/testdata1.md",
        "sha256": "b0fc3d0414b5c7b91a4760915efaffdd447b0a416143ce7ef57656033394f7eb",
        "timestamp": "Fri Mar 29 13:53:20 GMT 2024"
      },
      {
        "asset": "testdata/testdata1.md",
        "sha256": "b8bd10ad80ec1c041a2c12591f0aa148a003a42591a30bb544fa55e27e51820a",
        "timestamp": "Fri Mar 29 13:53:21 GMT 2024"
      }
    ]
  }
]
```

### Rename Json
Prepend output file with alternative string:

```shell
$ sigzag  --tag-file tango
```

Output:
```
tango-manifest-2024328-8853-402cd35c615b384c807cb1adb71923ff1ac0f8da06aadd0eb20a568b6a1f3609.json
tango-merkletree-2024328-8853-2a04c17860212ddce9ea2c1f921da29d834f762e700b609f281478a72ff63192.json
```

## Flags

| Flag               | Description                                              |
|:-------------------|:---------------------------------------------------------|
| --root             | Root directory to descend. Defaults to working directory |
| --level            | Directory nesting depth (default==3)                     |
| --diff             | Compare two manifests and return the difference if any   |
| --asset            | Asset to compile history for from manifests              |
| --history          | Manifests to search for an assets history                |
| --tag-file         | Prepends file with string to filename                    |
| --compare-manifest | Compare manifest                                         |
| --compare-merkle   | Compare merkle tree                                      |