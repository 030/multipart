# go-multipart

[![CI](https://github.com/030/multipart/workflows/Go/badge.svg?event=push)](https://github.com/030/multipart/actions?query=workflow%3AGo)
[![GoDoc Widget]][godoc]
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/030/multipart?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/multipart)](https://goreportcard.com/report/github.com/030/multipart)
[![StackOverflow SE Questions](https://img.shields.io/stackexchange/stackoverflow/t/multipart.svg?logo=stackoverflow)](https://stackoverflow.com/tags/multipart)
[![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/multipart.svg?logo=stackexchange)](https://devops.stackexchange.com/tags/multipart)
[![ServerFault SE Questions](https://img.shields.io/stackexchange/serverfault/t/multipart.svg?logo=serverfault)](https://serverfault.com/tags/multipart)
[![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/multipart?logo=docker&logoColor=white)](https://hub.docker.com/r/utrecht/multipart)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/utrecht/multipart?logo=docker&logoColor=white&sort=semver)
![Issues](https://img.shields.io/github/issues-raw/030/multipart.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/multipart.svg)
![Total downloads](https://img.shields.io/github/downloads/030/multipart/total.svg)
![GitHub forks](https://img.shields.io/github/forks/030/multipart?label=fork&style=plastic)
![GitHub watchers](https://img.shields.io/github/watchers/030/multipart?style=plastic)
![GitHub stars](https://img.shields.io/github/stars/030/multipart?style=plastic)
![License](https://img.shields.io/github/license/030/multipart.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/multipart.svg)
![Contributors](https://img.shields.io/github/contributors/030/multipart.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/multipart.svg)
![Last commit](https://img.shields.io/github/last-commit/030/multipart.svg)
![Release date](https://img.shields.io/github/release-date/030/multipart.svg)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/030/multipart?logo=github&sort=semver)](https://github.com/030/multipart/releases/latest)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=bugs)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=coverage)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_multipart)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_multipart&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_multipart)
[![codecov](https://codecov.io/gh/030/multipart/branch/main/graph/badge.svg)](https://codecov.io/gh/030/multipart)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/multipart)
[![Chocolatey](https://img.shields.io/chocolatey/dt/multipart)](https://chocolatey.org/packages/multipart)
[![multipart](https://snapcraft.io//multipart/badge.svg)](https://snapcraft.io/multipart)
[![codebeat badge](https://codebeat.co/badges/f4aa5086-a4d5-41cd-893a-5da816ee9107)](https://codebeat.co/projects/github-com-030-multipart-main)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

[godoc]: https://godoc.org/github.com/030/multipart
[godoc widget]: https://godoc.org/github.com/030/multipart?status.svg

## Quickstart

Start a Nexus3 server:

```bash
docker run --rm -d -p 9000:8081 sonatype/nexus3:3.16.0
```

Upload artifacts once it has been started:

```bash
./multipart -url \
    http://localhost:9000/service/rest/v1/components?repository=maven-releases \
    -user admin -pass admin123 -F \
    "maven2.asset1=@test/testdata/file1.pom,\
    maven2.asset1.extension=pom,\
    maven2.asset2=@test/testdata/file1.jar,\
    maven2.asset2.extension=jar,\
    maven2.asset3=@test/testdata/file1.war,\
    maven2.asset3.extension=war,\
    maven2.asset4=@test/testdata/file1.ear,\
    maven2.asset4.extension=ear,\
    maven2.asset5=@test/testdata/file1-sources.jar,\
    maven2.asset5.classifier=sources,\
    maven2.asset5.extension=jar"
```

[godoc]: https://godoc.org/github.com/030/go-multipart
[godoc widget]: https://godoc.org/github.com/030/go-multipart?status.svg
