# WeChatToolerFrameWork
一个基于WeChat的自动收发框架

# 组成
#### WeChatSetup2.6.8.52 版本
#### Tool.dll 负责底层HOOK
#### Patch.dll 负责提供HTTP接口和Tool.dll通信


# 目前支持的功能
#### 发送文本消息
#### 发送图片文件
#### 接收所有类型的消息(文本、视频、文件、图片、群聊消息、系统提示、公众号消息等)
#### 接收文本类型消息(当消息类型=1时)
#### 接收图片消息(自动保存到微信主目录下的\Img)   附带发送者的WXID
#### 接收语音消息(自动保存到微信主目录下的\Voice) 附带发送者的WXID


# 使用方法
#### 下载对应版本的微信安装程序链接：https://pan.baidu.com/s/1-8gR9NCxE4u4ArWU6DvPPg  提取码：yhec 
#### 下载修改过后的微信主程序WeChat2.exe https://github.com/blacknight2018/WeChatControlPC/blob/master/WeChat2.exe
#### 将Patch.dll和Tool.dll以及2中下载的修改过后的微信主程序WeChat2.exe覆盖到你的微信主目录
#### 运行WeChat2.exe,如果没问题你会看到如下消息
![avatar](https://github.com/blacknight2018/WeChatToolerFrameWork/blob/master/Images/Img1.jpg)

# 你需要提供以下HTTP接口

### http://localhost:19730/msg
> 接收消息,包含所有类型

###### 参数格式
> JSON

###### HTTP接口方式
> POST

###### 接收参数
``` 
{
    "wx_id": "wxid_xxxxxxx",  发送者,可以是公众号、群、私人
    "msg": "yyyyyy",           
    "type": "1",              当type = 1时 一定是文本消息,否则是一个XML结构。
}
```


### http://localhost:19730/img
> 接收图片消息

###### 参数格式
> JSON

###### HTTP接口方式
> POST

###### 接收参数
``` 
{
    "wx_id": "wxid_xxxxxxx",           发送者,可以是群、私人
    "img_local_abs_path": "c:/xxxxx",  图片的本地绝对路径          
}
```



### http://localhost:19730/voice
> 接收语音消息

###### 参数格式
> JSON

###### HTTP接口方式
> POST

###### 接收参数
``` 
{
    "wx_id": "wxid_xxxxxxx",           发送者,可以是群、私人
    "voice_data": "yyyyyyyyyyyyyyyy",  SKLI格式的音频数据,以BASE64编码返回,关于SKLI的解码和编码GitHub上有开源代码     
}
```


# 你通过以下接口发送消息

### http://localhost:8888/text
> 发送文本消息

###### 参数格式
> JSON

###### HTTP接口方式
> POST

###### 接收参数
``` 
{
    "wx_id": "wxid_xxxxxxx",           发送者,可以是群(@chatrom)
    "text": "yyyyyyyyyyyyyyyy",        文本消息
}
```

### http://localhost:8888/file
> 发送文件(目前暂时支支持图片格式)

###### 参数格式
> JSON

###### HTTP接口方式
> POST

###### 接收参数
``` 
{
    "wx_id": "wxid_xxxxxxx",           发送者,可以是群(@chatrom)
    "path": "yyyyyyyyyyyyyyyy",        文件所在的本地绝对路径
}
```
