# GoShell
用GoSocket写的shell，练练手。

## 怎么用

### 服务端

```bash
cd src/server
go build
# 在目标服务器上运行
server.exe
# 默认监听20000端口，密码是MagicWorld，修改了重新build即可。
```

### 客户端

```bash
cd src/client
go build
# 在本机上运行
client.exe
# 默认连接127.0.0.1的20000端口，修改IP即可。
# 密码是MagicWorld由服务端定，修改即可。
```



## 效果图

![image-20200823231939322](https://github.com/Kai5174/GoShell/README.assets/image-20200823231939322.png)



## 问题

中文乱码