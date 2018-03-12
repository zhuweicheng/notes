<?php
/**
 * 一根绳子，奇数长度时从一端减去1米，偶数长度时从中间剪断，将这根绳子全部剪成1米长的小段，剪了几刀？
 * @param $length
 * @return $count
 * @version 1.0
 * @time 2018.1.20
 * @anthor zwcheng
 */
function cut($length)
{
    $count = 0;
    // 按递减的绳子长度遍历
    while ($length >= 1) {
        echo sprintf("当前绳子长度 %d 米\n", $length);
        // 偶数长度时从中间剪断
        if ($length % 2 == 0) {
            // 每次处理绳子长度的一半
            echo sprintf("每次处理绳子长度的一半, 绳子长度 %d 米, 一半长度 %d 米\n", $length, $length / 2);
            // 该递归处理绳子长度的一半, 直到计算出该半截绳子所使用的刀数为止
            $count += cut($length / 2);
            // 接下来该一半绳子继续循环上一次的操作路线, 完成后半截绳子所使用的刀数
            $length /= 2;
        // 奇数长度时从一端减去1米长, 记录1刀
        } else {
            // 绳子长度为1米时记录刀数1刀
            echo sprintf("绳子长度: %d 米, 从一端减去1米，剩余: %d 米\n", $length, $length - 1);
            $length--;
            $count++;
        }
    }
    return $count;
}
$num = $argv[1] ?? 1;
echo sprintf("当前绳子长度为: %d 米, 裁剪开始...\n\n", $num);
echo sprintf("\n裁剪结束, 该绳子一共剪了 %d 刀\n", cut($num) - 1);
