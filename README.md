# go-multipart

[![GoDoc Widget]][GoDoc]
[![Build Status](https://travis-ci.org/030/go-multipart.svg?branch=master)](https://travis-ci.org/030/go-multipart)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/go-multipart)](https://goreportcard.com/report/github.com/030/go-multipart)
![Issues](https://img.shields.io/github/issues-raw/030/go-multipart.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/go-multipart.svg)
![Total downloads](https://img.shields.io/github/downloads/030/go-multipart/total.svg)
![License](https://img.shields.io/github/license/030/go-multipart.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/go-multipart.svg)
![Contributors](https://img.shields.io/github/contributors/030/go-multipart.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/go-multipart.svg)
![Last commit](https://img.shields.io/github/last-commit/030/go-multipart.svg)
![Release date](https://img.shields.io/github/release-date/030/go-multipart.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/go-multipart.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=bugs)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=coverage)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_go-multipart&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_go-multipart)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2845/badge)](https://bestpractices.coreinfrastructure.org/projects/2845)
[![codecov](https://codecov.io/gh/030/go-multipart/branch/master/graph/badge.svg)](https://codecov.io/gh/030/go-multipart)
[![BCH compliance](https://bettercodehub.com/edge/badge/030/go-multipart?branch=master)](https://bettercodehub.com/results/030/go-multipart)

## Usage

If one would like to upload a file, then the '@' character should be used like
in curl (see [Testing paragraph](#testing)):

```
./go-multipart -url \
    http://localhost:9999/service/rest/v1/components?repository=maven-releases \
    -user admin -pass admin123 -F \
    "maven2.asset1=@utils/test-files-multipart/file1.pom,\
    maven2.asset1.extension=pom,\
    maven2.asset2=@utils/test-files-multipart/file1.jar,\
    maven2.asset2.extension=jar,\
    maven2.asset3=@utils/test-files-multipart/file1-sources.jar,\
    maven2.asset3.extension=sources.jar"
```

## Testing

```
./integration-tests.sh
```

[GoDoc]: https://godoc.org/github.com/030/go-multipart
[GoDoc Widget]: https://godoc.org/github.com/030/go-multipart?status.svg
