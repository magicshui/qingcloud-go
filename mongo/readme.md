# API List
## Mongo 集群

青云 Mongo 集群默认的 replica set 包含一个 primary 和一个 priority0 节点， 所有备份都会基于 priority0 节点。 青云 Mongo 集群运行于私有网络内，结合青云提供的高性能硬盘和实时副本， 您的数据安全将会得到最大限度的保护。

[x] DescribeMongoNodes
[x] DescribeMongoParameters
[x] ResizeMongos
[x] CreateMongo
[x] StopMongos
[x] StartMongos
[x] DescribeMongos
[x] DeleteMongos
[x] CreateMongoFromSnapshot
[x] ChangeMongoVxnet
[x] AddMongoInstances
[x] RemoveMongoInstances
[x] ModifyMongoAttributes
[x] ModifyMongoInstances
[] GetMongoMonitor