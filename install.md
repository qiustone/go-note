# Install Golang for Windows
1. 去官网下载golang的最新安装包 https://golang.google.cn/dl/
2. 双击运行，根据提示安装程序，我的安装路径为 C:\go, 工程目录为 D:\code\go, D:\code\go\bin, D:\code\go\pkg, D:\code\go\src
3. 打开环境变量，在系统变量中添加 GOROOT(C:\Go)，在用户变量中GOPATH(D:\code\go), 
4. 在系统变量中的Path里添加%GOROOT%\bin, 在用户变量中的Path里添加%GOPATH%\bin
5. 打开CMD, 分别输入go version 和 go env 如果能正常输出相关信息，则表示安装成功。

||系统变量|用户变量|
|:---|:---|:---|
|添加| GOROOT (C:\Go)  |GOPATH (D:\code\go)  |
|在Path中添加| %GOROOT%\bin  |%GOPATH%\bin  |

