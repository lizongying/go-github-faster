# go github faster

get github faster ip

通过修改Hosts文件，可以缓解部分网络github访问很慢的问题。

[go github faster](https://github.com/lizongying/go-github-faster)

## Usage

可执行文件可以在Releases中找到 [releases](https://github.com/lizongying/go-crawler-example/releases)

![image](./screenshots/screenshot.png)

1. 测试最快的IP地址

   如:
   ```shell
   # web
   ./github_faster_darwin_arm64 -q
   
   # git
   ./github_faster_darwin_arm64 -p ssh -m git -q
   ```
    * -p tcp/ssh, 使用的协议, 一般web使用tcp(https)协议, git使用ssh协议. 默认tcp
    * -m web/api/git, 通常需要绑定web(网页访问)和git(git操作). 默认web
    * -q 安静模式，不输出连接时间, 注意, 有虽然连接时间很快但是连接失败的情况. 默认会输出连接时间

2. 修改host文件

   例: 把如下内容添加到host文件
   ```
   # web
   20.27.177.113 github.com
   20.205.243.166 github.com
   20.207.73.82 github.com
   20.200.245.247 github.com
   20.248.137.48 github.com
   20.233.83.145 github.com
   20.29.134.23 github.com
   185.199.108.1 github.com
   20.175.192.147 github.com
   20.201.28.151 github.com
   
   # api
   20.27.177.113 github.com
   20.27.177.118 github.com
   20.205.243.166 github.com
   20.205.243.160 github.com
   20.207.73.82 github.com
   20.207.73.83 github.com
   20.200.245.248 github.com
   20.200.245.247 github.com
   20.233.83.149 github.com
   20.248.137.50 github.com
   20.29.134.19 github.com
   20.29.134.23 github.com
   20.233.83.145 github.com
   20.248.137.48 github.com
   20.175.192.147 github.com
   20.175.192.146 github.com
   20.87.245.4 github.com
   20.201.28.151 github.com
   20.201.28.152 github.com
    ```

    * unix `/etc/hosts`
    * windows `C:\Windows\System32\drivers\etc\hosts`

## Build

```shell
make 
```

## 赞赏

![image](./screenshots/appreciate.jpeg)