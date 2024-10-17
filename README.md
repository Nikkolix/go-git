# go-git

Package go-git is a git cmd tool for go modules.
It's used to increase the version tag by one (v0.0.1 -> v0.0.2).

## Installation
```sh
go get github.com/Nikkolix/go-git
```

## Usage

```sh
./go-git.exe -msg <commit message>
```

## Description

```sh
git add .
git commit -m <commit message>
git tag v0.0.<n+1> master
git push -u origin v0.0.<n+1>
git push
```

where n is the last version tag used
