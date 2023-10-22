func Exit(code int) // 让当前程序以给出的状态码（code）退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。

func Getuid() int // 获取调用者的用户id

func Geteuid() int // 获取调用者的有效用户id

func Getgid() int // 获取调用者的组id

func Getegid() int // 获取调用者的有效组id

func Getgroups() ([]int, error) // 获取调用者所在的所有组的组id

func Getpid() int // 获取调用者所在进程的进程id

func Getppid() int // 获取调用者所在进程的父进程的进程id

func Hostname() (name string, err error) // 获取主机名

func Getenv(key string) string // 获取某个环境变量

func Setenv(key, value string) error // 设置一个环境变量,失败返回错误，经测试当前设置的环境变量只在 当前进程有效（当前进程衍生的所以的go程都可以拿到，子go程与父go程的环境变量可以互相获取）；进程退出消失

func Clearenv() // 删除当前程序已有的所有环境变量。不会影响当前电脑系统的环境变量，这些环境变量都是对当前go程序而言的