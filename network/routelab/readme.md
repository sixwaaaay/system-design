# 配置linux 的路由

三个机器，ip分别为:
- A: 10.1.11.2
- B: 10.1.11.3
- C: 10.1.11.4

B 作为路由器
A 是client
C 是server

配置 A 到 C 的下一跳为 B
配置 B 到 C 直接跳转

A 配置如下

```
ip route static 10.1.11.4 255.255.255.0 10.1.11.3
```

B 配置如下

```
ip route static 10.1.11.4 255.255.255.0 10.1.11.4
```

