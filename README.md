# TF-IDF-Test

- Test application written in Golang.
- Run TF-IDF method

## How to use

In my case, run from OSX 10.10.5(Yosemite)

### Install Golang

Check Golang official!
see: https://golang.org/

```
$ go version
go version go1.5.1 darwin/amd64
```

### Install Mecab

see: http://statsbeginner.hatenablog.com/entry/2015/10/05/221357

```
$ mecab -v
mecab of 0.996
```

### (Optional) Install latest Mecab Dictionary

see: http://diary.overlasting.net/2015-03-13-1.html

### go get

```
go get "github.com/yukihir0/mecab-go"
go get "github.com/djimenez/iconv-go"
```

### Ready?
```
go run app/tf_idf.go
```
