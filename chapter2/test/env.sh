# 个人笔记：
# 最先执行，配置服务节点
# 因为Ubuntu的内核支持接口别名，所以可以在服务器的网络接口通过使用别名绑定多个地址。
# 终端通过ifconfig可以查看当前服务器的网络接口
# enp0s3是本机的网络接口，要根据实际情况填写

# 在同一台服务器上绑定多个地址的指令
sudo ifconfig enp0s3:1 10.29.1.1/12345
sudo ifconfig enp0s3:2 10.29.1.2/12345
sudo ifconfig enp0s3:3 10.29.1.3/12345
sudo ifconfig enp0s3:4 10.29.1.4/12345
sudo ifconfig enp0s3:5 10.29.1.5/12345
sudo ifconfig enp0s3:6 10.29.1.6/12345

sudo ifconfig enp0s3:7 10.29.2.1/12345
sudo ifconfig enp0s3:8 10.29.2.2/12345

# 安装RabbitMQ
sudo apt-get install rabbitmq-server
# 下载rabbitmqadmin管理工具
sudo rabbitmq-plugins enable rabbitmq_management
wget localhost:15672/cli/rabbitmqadmin
# 创建apiServers和dataServers这两个exchanges
python3 rabbitmqadmin declare exchange name=apiServers type=fanout
python3 rabbitmqadmin declare exchange name=dataServers type=fanout
# 添加用户test,密码test
sudo rabbitmqctl add_user test test
# 给test用户添加访问所有exchange的权限：
sudo rabbitmqctl set_permissions -p / test ".*" ".*" ".*"


# 回到chapter2目录下，准备启动8个服务程序
# 启动前创建相应的$STORAGE_ROOT目录以及子目录objects
for i in `seq 1 6`; do mkdir -p ./tmp/$i/objects; done

export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672
# 进入dataServer目录，启动服务。注意下面STORAGE_ROOT 的值！
cd dataServer
LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=../tmp/1 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=../tmp/2 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=../tmp/3 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=../tmp/4 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=../tmp/5 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=../tmp/6 go run dataServer.go &

# 进入apiServer目录，启动服务
cd ../apiServer
echo $RABBITMQ_SERVER  # 如果新创终端，首先查看是否有当前变量
# export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672
LISTEN_ADDRESS=10.29.2.1:12345  go run apiServer.go &
LISTEN_ADDRESS=10.29.2.2:12345  go run apiServer.go &

# 用curl命令作为客户端来访问服务节点10.29.2.2:12345,put一个test2对象
curl -v 10.29.2.2:12345/objects/test2 -XPUT -d"this is object test2"

# 上传成功后，用locate命令看看test2对象被保存在哪个数据服务节点上
curl 10.29.2.2:12345/locate/test2
# 现在换一个接口服务节点get这个对象
curl 10.29.2.1:12345/objects/test2