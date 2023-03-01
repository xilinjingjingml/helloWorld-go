# 新手引导

创建第一个mcbeam服务

### 增加mcbeam集群入口
```
mcbeam env add mcbeam proxy.mcbeam.cc
```
### 向管理员申请帐号,切换到对应的namespace, 登录mcbeam集群
```
mcbeam user namespace set myproject
mcbeam login
```
### 使用模板创建helloworld服务
```
mcbeam new helloworld
```
### 生成proto协议
```
cd helloworld
make proto
```
### 部署helloworld服务到mcbeam集群
```
mcbeam run .
```
### 查看服务运行状态
```
mcbeam status
```
### 查看服务输出日志
```
mcbeam logs helloworld
```
### 通过命令行访问helloworld服务
```
mcbeam helloworld --name=Mars
```
### 通过浏览器访问helloworld服务
```
curl "https://myproject.mcbeam.cc/helloworld?name=Mars"
```
### 查看详细使用说明
```
mcbeam --help
```