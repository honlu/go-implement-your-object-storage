export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672  # 注意冒号哦！
export ES_SERVER=127.0.0.1:9200  # 增加ES_SERVER环境变量的设置
# 设置
sudo ifconfig ens33:1 10.29.1.1/12345
sudo ifconfig ens33:2 10.29.1.2/12345
sudo ifconfig ens33:3 10.29.1.3/12345
sudo ifconfig ens33:4 10.29.1.4/12345
sudo ifconfig ens33:5 10.29.1.5/12345
sudo ifconfig ens33:6 10.29.1.6/12345

sudo ifconfig ens33:7 10.29.2.1/12345
sudo ifconfig ens33:8 10.29.2.2/12345
# 进入dataServer目录，启动服务
cd dataServer
LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 go run dataServer.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 go run dataServer.go &
# 进入apiServer目录，启动服务
cd ../apiServer
echo $RABBITMQ_SERVER  # 如果新创终端，首先查看是否有当前变量
# export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672  # 如果没有，新终端添加变量
LISTEN_ADDRESS=10.29.2.1:12345  go run apiServer.go &
LISTEN_ADDRESS=10.29.2.2:12345  go run apiServer.go &

# 建文件夹，当存储路径
cd ..
for i in `seq 1 6`; do mkdir -p ./tmp/$i/objects; done