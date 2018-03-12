<?php
/**
 * 一群猴子排成一圈，按1，2，...，n依次编号。
 * 然后从第1只开始数，数到第m只,把它踢出圈，
 * 从它后面再开始数，再数到第m只，在把它踢出去...，
 * 如此不停的进行下去，直到最后只剩下一只猴子为止，那只猴子就叫做大王。
 * @param $m
 * @param $n
 * @return integer
 * @version 1.0
 * @date 2018.1.27
 * @anthor zwcheng
 */
/*function monkey($m,$n){
    $arr[$m] = $m;
    $i = 0;//设置数组指针
    while (count($arr) > 1) {
        //遍历数组，判断当前猴子是否为出局序号，
        //如果是则出局，否则放到数组最后
        if (($i + 1) % $n == 0) {
            unset($arr[$i]);
        } else {
            array_push($arr, $arr[$i]);
            //本轮非出局猴子放数组尾部
            unset($arr[$i]);
            //删除
        }
        $i++;
    }
    return $arr;
}

print_r(monkey(5,3));
