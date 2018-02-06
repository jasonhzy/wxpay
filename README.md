### 配置环境变量$GOPATH

为了更加方便的操作，请将 $GOPATH/bin 加入到你的 $PATH 变量中

    # 如果您还没添加 $GOPATH 变量
    $ echo 'export GOPATH="/Users/jason/Sites/project/golang"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件

    # 如果您已经添加了 $GOPATH 变量
    $ echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件

### 安装或者升级 Beego 和 Bee 的开发工具

    $ go get -u github.com/astaxie/beego
    $ go get -u github.com/beego/bee

### 将项目wxpay复制到$GOPATH/src下，运行bee run wxpay
    ______
    | ___ \
    | |_/ /  ___   ___
    | ___ \ / _ \ / _ \
    | |_/ /|  __/|  __/
    \____/  \___| \___| v1.9.1
    2018/02/06 17:09:32 INFO     ▶ 0001 Using 'wxpay' as 'appname'
    2018/02/06 17:09:32 INFO     ▶ 0002 Initializing watcher...
    2018/02/06 17:09:33 SUCCESS  ▶ 0003 Built Successfully!
    2018/02/06 17:09:33 INFO     ▶ 0004 Restarting 'wxpay'...
    2018/02/06 17:09:33 SUCCESS  ▶ 0005 './wxpay' is running...
    2018/02/06 17:09:33.725 [I] [asm_amd64.s:2086] http server Running on http://:8080

即表明成功运行: http://localhost:8080