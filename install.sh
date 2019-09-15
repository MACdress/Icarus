#!/usr/bin/env bash

#之后的代码，如果返回一个非0值，则整个脚本立即退出
set -e

if [ ! -f install.sh ]; then
	echo 'install must be run within its container folder' 1>&2
	exit 1
fi

#使用GoMod
export GO111MODULE=on

#创建日志文件夹
if [ ! -d log ]; then
	mkdir log
fi