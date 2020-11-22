# WeChatToolerFrameWork
一个基于WeChat的自动收发框架

# 组成
## WeChatSetup2.6.8.52 版本
## Tool.dll 负责底层HOOK
## Patch.dll 负责提供HTTP接口和Tool.dll通信


# 目前支持的功能
# 发送文本消息
# 发送图片文件

# 接收所有类型的消息(文本、视频、文件、图片、群聊消息、系统提示、公众号消息等)
# 接收文本类型消息(当消息类型=1时)
# 接收图片消息(自动保存到微信主目录下的\Img)   附带发送者的WXID
# 接收语音消息(自动保存到微信主目录下的\Voice) 附带发送者的WXID


# 使用方法
  ##1 下载对应版本的微信安装程序链接：https://pan.baidu.com/s/1-8gR9NCxE4u4ArWU6DvPPg  提取码：yhec 
  ##2 下载修改过后的微信主程序WeChat2.exe https://github.com/blacknight2018/WeChatControlPC/blob/master/WeChat2.exe
  ##3 将Patch.dll和Tool.dll以及2中下载的修改过后的微信主程序WeChat2.exe覆盖到你的微信主目录
  ##4 运行WeChat2.exe,如果没问题你会看到如下消息
 
