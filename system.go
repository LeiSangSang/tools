package tools

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	logfile *os.File
	SignalChan chan os.Signal
)

func Dup(path string) {
	var err error
	logfile, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}

	err = syscall.Dup2(int(logfile.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		panic(err)
	}
}

func MonitorSignal() {
	SignalChan = make(chan os.Signal, 2)
	signal.Notify(SignalChan,
		syscall.SIGILL,  //非法指令(程序错误、试图执行数据段、栈溢出等)
		syscall.SIGFPE,  //算术运行错误(浮点运算错误、除数为零等)
		syscall.SIGSEGV, //无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
		syscall.SIGPIPE, //消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
		syscall.SIGTERM, //结束程序
		syscall.SIGTSTP, //停止进程
		syscall.SIGBUS,  //非法地址(内存地址对齐错误)
		syscall.SIGSYS,  //无效的系统调用(SVr4)
		syscall.SIGXCPU, //超过CPU时间资源限制
		syscall.SIGXFSZ, //超过文件大小限制
	)
}
