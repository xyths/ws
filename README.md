# ws
a websocket command line tool

之所以有这个小工具，是因为`wscat`没法用`nohup`运行在后台。

而我需要一个简单的客户端测试我的服务，它只干一件非常简单的事情：建立连接以后保持活跃，打印服务器推送过来的消息。
