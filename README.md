<p align="center">
    <img src="https://github.com/KevinFasusi/sigzag/blob/376da187ac04145ff4c3e9828d5ba99d31806c72/.github/images/sigzag-logo.png" width="201" alt="sigzag logo">
</p>

![Static Badge](https://img.shields.io/badge/Under%20Development-494980) [![GitHub license][license-badge]](LICENSE)

SigZag is a small utility for signing digital assets and generating local manifests 
for **data centric workloads**. 

Managing the provenance, lineage and integrity of data ingress, transformation and egress can be challenging.
The difficulty can be exacerbated when using large open data sets. Accidental and in some instance malicious changes to
data can occur when data wrangling and munging.

The SigZag utility is useful for:

- Cryptographically signing content, both the file and data (including database).
- Generating manifests for signed data in JSON format (gRPC TBC).
- Tracking the history of a data assets. 
- Checking, comparing and monitoring the integrity of data in common file formats (csv, xlsx, xlsm, xml, json, jupyter 
notebooks, parquet, txt, zip, etc.).

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

### Download

Download a single file:

```shell
$ sigzag --url https://somefile/url

```

output:

```
27551470 bytes downloaded
sha256: 374ea82b289ec738e968267cac59c7d5ff180f9492250254784b2044e90df5a9
```
Download multiple items using a json file:

```shell
$ sigzag --urls path/to/file.json
```

Json example:
```
[
  { "url": "https://go.dev/dl/go1.22.2.src.tar.gz",
    "sha256": "374ea82b289ec738e968267cac59c7d5ff180f9492250254784b2044e90df5a9",
    "size": "26MB"},
  {"url": "https://go.dev/dl/go1.22.2.darwin-amd64.tar.gz",
    "sha256": "33e7f63077b1c5bce4f1ecadd4d990cf229667c40bfb00686990c950911b7ab7",
    "size": "67MB"},
  {"url": "https://go.dev/dl/go1.22.2.linux-amd64.tar.gz",
    "sha256": "5901c52b7a78002aeff14a21f93e0f064f74ce1360fce51c6ee68cd471216a17",
    "size": "67MB"}
]
```

output:

```
27551470 bytes downloaded
27 MB downloaded
sha256: 374ea82b289ec738e968267cac59c7d5ff180f9492250254784b2044e90df5a9
Match: true
70323302 bytes downloaded
70 MB downloaded
sha256: 33e7f63077b1c5bce4f1ecadd4d990cf229667c40bfb00686990c950911b7ab7
Match: true
68958123 bytes downloaded
68 MB downloaded
sha256: 5901c52b7a78002aeff14a21f93e0f064f74ce1360fce51c6ee68cd471216a17
Match: true
```

## Flags

| Flag               | Description                                                                                |
|:-------------------|:-------------------------------------------------------------------------------------------|
| --url              | Download asset and show sha256 checksum                                                    |
| --urls             | Download assets from a list of urls in a file and generate a manifest containing checksums |
| --root             | Root directory to descend. Defaults to working directory                                   |
| --level            | Directory nesting depth (default==3)                                                       |
| --diff             | Compare two manifests and return the difference if any                                     |
| --asset            | Asset to compile history for from manifests                                                |
| --history          | Manifests to search for an assets history                                                  |
| --tag-file         | Prepends file with string to filename                                                      |
| --compare-manifest | Compare manifest                                                                           |
| --compare-merkle   | Compare merkle tree                                                                        |
| --datasource       | Search url for data sources with common extensions                                         |

<!-- refs -->
[license-badge]: https://img.shields.io/github/license/KevinFasusi/sigzag