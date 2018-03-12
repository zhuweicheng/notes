<?php
class Test
{
	protected $num = 3;
	protected $works = [];
	protected $ppid = null;
	public function __construct()
	{
		// 初始化一些环境参数
		$this->ppid = getmypid();
		echo "该脚本启动的进程ID -- {$this->ppid}\n";
	}
	
 	// 启动进程
	public function start()
	{
		for ($i = 0; $i < $this->num; $i++) {
			// 启动进程后执行回调函数
			$process = new \swoole_process(function(\swoole_process $work) use ($i) {
				echo "{$i} - 所有进程都启动完毕之后才会执行到该处...\n";
				// 在这里运行一个php命令
			});
			// 成功启动子进程后返回子进程id
			$pid = $process->start();
			// 将成功启动的工作存入工作表
			$this->works[$pid] = $process;
		}
		// 需要注意的是，Process的signal方法是一个异步方法，其底层会开启事件循环，
        // 因此使用了signal方法的进程在主代码执行完后并不会主动退出，需要调用exit、发送信号等方式关闭。
		$this->signal();
		return true;
	}

	// 等待子进程执行完毕的信号监听函数
	public function signal()
	{
		// SIGTERM是杀或的killall命令发送到进程默认的信号
		\swoole_process::signal(SIGTERM, function ($signo) {
            echo "进程接收到信号 SIGTERM 退出\n";
        });
        // SIGUSR1
        \swoole_process::signal(SIGUSR1, function ($signo) {
            echo "进程接收到信号 SIGUSR1 退出\n";
        });

        // 子进程正常执行完毕的信号 SIGCHLD
        \swoole_process::signal(SIGCHLD, function ($signo) {
        	while ( ($ret = \swoole_process::wait()) !== false) {
        		echo "子进程执行完毕\n";
        		print_r($ret);
        		unset($this->works[$ret['pid']]);
        	}
        	if (!count($this->works)) {
        		// 此处要结束脚本进程比较麻烦，直接exit()不知道有什么副作用
        		exit("当前脚本已经执行完毕，直接结束\n");
        	}
        });
	}
}
print_r((new Test())->start());
