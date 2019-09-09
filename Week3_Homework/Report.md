# 安装go语言开发环境实验报告
### 开发环境
Ubuntu18.04.3
### 1.安装 VSCode 编辑器
首先下载安装包：[VSCode下载链接](https://code.visualstudio.com/docs/?dv=linux64_deb)

然后拷贝到虚拟机中，在终端执行命令(x对应下载的版本)

```
sudo dpkg -i code_x.xx.x-xxxxxxxxxx_amd64.deb 
```

安装完成后重启，就可以在Application中找到VSCode了。

### 2.安装 golang
#### 2.1安装
在终端执行命令

``` 
sudo apt-get install golang
``` 

![](/IGM/01.jpg)

安装好之后在终端输入```go version```来查看go语言的版本

![](/IGM/02.jpg)

#### 2.2设置环境变量
- **创建工作空间**
   
    在终端中输入命令```mkdir $HOME/gowork```
- **配置的环境变量**
   
    在终端中输入命令```sudo gedit ~/.bashrc```,在打开的文件中最末尾添加两行代码
    ```
	export GOPATH=$HOME/gowork
	export PATH=$PATH:$GOPATH/bin
    ```
