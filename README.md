# Github-Asset-Mirror

![Golang](https://img.shields.io/github/actions/workflow/status/starudream/github-asset-mirror/golang.yml?label=golang&style=for-the-badge)
![Docker](https://img.shields.io/github/actions/workflow/status/starudream/github-asset-mirror/docker.yml?label=docker&style=for-the-badge)
![Release](https://img.shields.io/github/v/release/starudream/github-asset-mirror?include_prereleases&sort=semver&style=for-the-badge)
![License](https://img.shields.io/github/license/starudream/github-asset-mirror?style=for-the-badge)

## Usage

```text
Usage of ./bin/app:
  -addr string
    	http server address (default "0.0.0.0:80")
  -name string
    	github repo (default "{{.repo}}-{{.os}}-{{.arch}}-{{.platform}}-{{.ver}}.zip")
  -owner string
    	github owner
  -proxy string
    	available proxy: ghproxy fastgit
  -repo string
    	github repo
  -storage string
    	file storage path (default "/storage")
```

## Docker

![Version](https://img.shields.io/docker/v/starudream/github-asset-mirror?sort=semver&style=for-the-badge)
![Size](https://img.shields.io/docker/image-size/starudream/github-asset-mirror?sort=semver&style=for-the-badge)
![Pull](https://img.shields.io/docker/pulls/starudream/github-asset-mirror?style=for-the-badge)

```bash
docker pull starudream/github-asset-mirror
```

```bash
docker run -d --name github-asset-mirror --restart always -p 80:80 -v $(pwd)/storage:/storage starudream/github-asset-mirror
```

## Example

```shell
./app -owner starudream -repo secret-tunnel -proxy ghproxy
```

## License

[Apache License 2.0](./LICENSE)
