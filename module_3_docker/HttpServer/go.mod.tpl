module main

go 1.20

require (
	github.com/golang/glog v1.1.1 // indirect
)

// Go代码必须放在工作空间内。它其实就是一个目录，其中包含三个子目录：
//
// src 目录包含Go的源文件，它们被组织成包（每个目录都对应一个包），
// pkg 目录包含包对象，
// bin 目录包含可执行命令。

// Go源文件中的第一个语句必须是
//
//package 名称
//这里的 名称 即为导入该包时使用的默认名称。 （一个包中的所有文件都必须使用相同的 名称。）
//
//Go的约定是包名为导入路径的最后一个元素：作为 “crypto/rot13” 导入的包应命名为 rot13。
//
//可执行命令必须使用 package main。
//
//链接成单个二进制文件的所有包，其包名无需是唯一的，只有导入路径（它们的完整文件名） 才是唯一的。
//
//共多关于Go的命名约定见 实效Go编程。
