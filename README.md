# 毒鸡汤 - GoLang

## 资料

golang 版本：11

mysql 版本：5.5

## 数据库

1. 导入数据：记事本打开 model/mysql.sql ，在 mysql 中导入
2. 配置数据库：记事本打开 config/mysql.toml ，对应参数进行编辑

## 编译与运行（Golang）

> Env
```
export CGO_ENABLED=0
export GOOS=linux 
export GOARCH=mipsle

export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64

export CGO_ENABLED=0
export GOOS=darwin
export GOARCH=amd64
```

> Build
```
go run main.go

go build

go build -ldflags="-w -s" -trimpath
```

## 说明

### 这是我学习 go 语言的项目开发的，轻喷

## 原作者

[毒鸡汤（PHP）](https://github.com/egotong/nows)
