<?php
// swoole process 学习
function runner(\swoole_process $sub_process)
{
    // 执行脚本
//    $sub_process->exec("/usr/local/bin/php", [
//        __DIR__ . '/help.php',
//        'hello1',
//        '>>/Users/zhgxun/Public/html/logs/out.txt',
//        '2>>/Users/zhgxun/Public/html/logs/error.txt'
//    ]);
    $cmd = [
        '/usr/local/bin/php',
        __DIR__ . '/help.php',
        'hello1',
        '>>/Users/zhgxun/Public/html/logs/out.txt',
        '2>>/Users/zhgxun/Public/html/logs/error.txt'
    ];
    $cmd = implode(" ", $cmd);
    passthru($cmd);
}
$process = new \swoole_process("runner", false, 0);
$process->start();
//// 读取输出
//echo $process->read();
//// 传入参数
//$process->write("World\n");
//// 读取输出
//echo $process->read();
\swoole_process::wait(true);
