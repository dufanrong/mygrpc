# mygRPC

Meta trace的raas top3数据中，root对于下游MS的调用有四种情况：

![image](https://github.com/dufanrong/mygrpc/blob/master/img/whiteboard_exported_image.png)

greeter_client/main.go接受http请求（http://localhost:8082/?case=x），解析case字段并选择相应的server发送请求。

性能测试：

wrk -t4 -c10 -d30s -s scripts/test.lua http://localhost:8082


