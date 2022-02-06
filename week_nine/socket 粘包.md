### socket 粘包
TCP 以流的形式传输，没有明确的边界

### 解决方案
#### fix length：
通信双方以固定的长度进行请求的发送，接收；需要进行数据填充并且且长度的设置不够灵活

> send: a空空  
> recv: a   
#### delimiter based：
通信的双方以固定的分割符进行请求的划分；限制请求内不能出现约定的分割符
> send: a\nb\nc\n  
recv: a  
recv: b  
recv: c  
#### length field based frame decoder：
在数据头中存储数据正文的大小，当读取的数据小于数据头中的大小时，继续读取数据，直到读取的数据长度等于数据头中的长度时才停止。
案例参考 imdecode
参考链接：https://www.cnblogs.com/vipstone/p/14239160.html