# 终端在chapter1目录下，创建文件夹
mkdir tmp
mkdir ./tmp/objects
# 启动服务
LISTEN_ADDRESS=:12345 STORAGE_ROOT=./tmp go run server.go


# 再起一个终端，测试服务
curl -v 127.0.0.1:12345/objects/test   # 查看是否可以get一个名为test的对象，没有会404
curl -v 127.0.0.1:12345/objects/test -XPUT -d"this is a test object"  # PUT一个test对象.返回200
# 然后测试再次是否有test对象。返回200
curl -v 127.0.0.1:12345/objects/test 