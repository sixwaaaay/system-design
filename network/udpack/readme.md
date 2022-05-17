# 使用UDP实现可靠的数据传输

1. ack 代表期望收到的数据包序号
2. 如果接受ack超时，则认为丢包，重传
3. 如果是类似TCP的流水线，则可以通过快速重传提升一定的效率
4. 只传输中间丢失的数据包，不传输全部数据包是可行的，但是不能完全不阻塞，因为缓冲区是有限的