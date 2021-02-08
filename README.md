# :boom: Command Utilities with Golang :)

my command utils for dev :)

---  

## Features

- github
    - commit: read filtered github commit messages

---  

## Getting Started

> ### 1. Go get & install

```bash
$ go get -u github.com/zacscoding/gotils/...
$ gotils -h
```  

> ### 2. Go install (with git clone)

```bash
$ mkdir -p $GOPATH/src/github.com/zacscoding
$ cd $GOPATH/src/github.com/zacscoding
$ git clone https://github.com/zacscoding/gotils.git
$ cd gotils/cmd/gotils
$ go install

// check bin in $GOBIN
$ ls -la $GOBIN/gotils 
-rwxrwxr-x 1 app app 9417592 Nov  7 01:29 /home/app/go/bin/gotils 
```  

---  

## Commands  

- [Github](#Github)  

### Github  

**Commit Command**  

`Commit Command` is useful when u don't known proper commit messages well like update bla.. or change bla..

> example  

```bash
$ gotils github commit --user zacscoding --repo go-workspace --keywords gorm,tests --maxcommits 3
+----------------------------------------------------------------------------------------------------------------------------------------------------+
| Matched Commits from /zacscoding/go-workspace. Keywords:gorm,tests                                                                                 |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
| # | Message                                           | URL                                                                                        |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
| 1 | ratelimit: tests rate limit with http client      | https://github.com/zacscoding/go-workspace/commit/dfc640940e70c54eafa9b484926dd9e659ce9330 |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
| 2 | sarama: tests partition assignment strategies     | https://github.com/zacscoding/go-workspace/commit/48bd90378bf1283d2c7e05c84e638d2729367d5a |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
| 3 | gormV1: tests select for update                   | https://github.com/zacscoding/go-workspace/commit/0d5d901405a1f0ab380451c91f8488069cd8073c |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
| 4 | gormV2: tests DBResolver for splitting read/write | https://github.com/zacscoding/go-workspace/commit/5a41a54d8d54e32c0bf86f1216a2adce70739b96 |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
|   |                                                   | Total Read Commits: #6                                                                     |
+---+---------------------------------------------------+--------------------------------------------------------------------------------------------+
```  

---  
