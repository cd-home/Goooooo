[TOC]

# Gooooooo

Gooooooo

### 启动项目

~~~makefile
make run   app=admin mode=dev
make build app=admin
make upx   app=admin mode=dev level=2
~~~

### DEBUG
~~~
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "args": ["server", "--app=admin", "--mode=dev"],
            "program": "${fileDirname}"
        }
        
    ]
}
~~~