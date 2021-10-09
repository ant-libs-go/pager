# Pager

pager是一个便捷的翻页解析库

[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/ant-libs-go/pager?status.png)](http://godoc.org/github.com/ant-libs-go/pager)
[![Go Report Card](https://goreportcard.com/badge/github.com/ant-libs-go/pager)](https://goreportcard.com/report/github.com/ant-libs-go/pager)

## 特性

* 支持常规页码类翻页的因子计算及渲染输出
* 支持瀑布流类翻页的因子计算及渲染输出 [TODO]

## 安装

	go get github.com/ant-libs-go/pager

## 快速开始

```golang
p := pager.New(99, 1, 20)
p.ItemCount()   // 获取数据总条数
p.PageCount()   // 获取总页数
p.pageSize()    // 获取每页数据条数
p.CurrentPage() // 获取当前页码
p.Offset()      // 获取查询偏移量
p.Render()      // 将API输出的信息对外统一暴露
```
