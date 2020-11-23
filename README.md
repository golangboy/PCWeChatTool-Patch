# PCWeChatTooler
一个基于WeChat的自动收发框架

# 组成
#### WeChatSetup2.6.8.52 版本(https://pan.baidu.com/s/1-8gR9NCxE4u4ArWU6DvPPg  提取码：yhec )
#### Tool.dll 负责底层HOOK
#### Patch.dll 负责提供HTTP接口和Tool.dll通信

## Patch.dll
#### 你可以通过git本仓库代码运行build.cmd编译出Patch.dll 或者 下载我已经编译后的out\Patch.dll文件

## Tool.dll
#### 你可以通过使用Visual Studio 编译其源码 https://github.com/blacknight2018/WeToolelDll 得到

# 目前支持的功能
#### 发送文本消息
#### 发送图片文件
#### 接收所有类型的消息(文本、视频、文件、图片、群聊消息、系统提示、公众号消息等)
#### 接收文本类型消息(当消息类型=1时)
#### 接收图片消息(自动保存到微信主目录下的\Img)   附带发送者的WXID
#### 接收语音消息(自动保存到微信主目录下的\Voice) 附带发送者的WXID


# 使用方法
#### 1.下载WeChatSetup2.6.8.52 版本的微信安装程序
#### 2.下载修改过后的微信主程序WeChat2.exe https://github.com/blacknight2018/WeChatControlPC/blob/master/WeChat2.exe
#### 3.将Patch.dll和Tool.dll以及2中下载的修改过后的微信主程序WeChat2.exe覆盖到你的微信主目录
#### 4.运行WeChat2.exe,如果没问题你会看到如下消息
![avatar](https://github.com/blacknight2018/WeChatToolerFrameWork/blob/master/Images/Img1.jpg)

# 你需要提供以下HTTP接口接收消息

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
    "type": "1",              当type = 1时 文本消息  type = 10000时 系统消息
}
```
###### 结果示例
> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img2.jpg)  

> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img3.jpg)  

> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img4.jpg)  

> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img5.jpg)  

  

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
###### 结果示例
> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img6.jpg)
 


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
###### 结果示例
> ![avatar](https://github.com/blacknight2018/PCWeChatTool-Patch/blob/master/Images/Img7.jpg)
 
 

# 你可以通过以下接口发送消息

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
> 发送文件(目前暂时只支持图片格式)

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
