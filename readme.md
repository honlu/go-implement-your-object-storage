# 分布式对象存储-原理、架构及Go语言实现
原始仓库：https://github.com/stuarthu/go-implement-your-object-storage

## 2022-10-21
创建新分支并切换，此项目也将从main切换到honlu分支。原始项目使用较久的包管理方式。本仓库致力于使用go module管理包，更新配置和部分代码。
```
git checkout -b honlu
git add .
git commit -m "update branch"
git push --set-upstream origin honlu
```
环境相关：
虚拟机：Ubuntu 20.04
Go版本：go 1.13
注意：由于代码是在ubuntu上测试通过的，但此仓库是windows上再修改和上传的，所以有些格式会发生转化。比如LF will be replaced by CRLF。

## To do List
- [x] chapter1 对象存储简介 [2022-10-20 完成]
- [x] chapter2 可扩展的分布式系统 [2022-10-21 完成]
- [ ] chapter3 元数据服务
- [ ] chapter4 数据校验和去重
- [ ] chapter5 数据冗余和即时修复
- [ ] chapter6 断电续传
- [ ] chapter7 数据压缩
- [ ] chapter8 数据维护

### chapter1
只需要添加一个go.mod，并修改server.go中引用包设置。

### chapter2
将apiServer和dataServer分别作为一个项目，进行go module管理，并修改其中的一些包引用修改。  
将apiServer和dataServer引用src/lib中进行require设置。并对src/lib中objectstream和rabbitmq进行项目管理。