<?php
/**
 * 几种常见的排序方法
 */
class Sorting
{
	// 1. 冒泡法排序
	// 在要排序的一组数中, 对当前还未排好的序列, 从前往后对相邻的两个数依次进行比较和调整, 让较大的数往下沉, 较小的往上冒
	// 即每当两相邻的数比较后发现它们的排序与排序要求相反时, 就将它们互换
	public function bubble($arr = [])
	{
		$length = count($arr);
		// 外层控制循环的次数, 比数据少循环一次即可
		for ($i = 1; $i < $length; $i++) {
			// 内层比较数据, 每次均要比较所有剩下的数, 交换方式决定升序或降序
			for ($j = 0; $j < $length - $i; $j++) {
				if ($arr[$j] > $arr[$j+1]) {
					$temp = $arr[$j+1];
					$arr[$j+1] = $arr[$j];
					$arr[$j] = $temp;
				}
			}
		}
		return $arr;
	}
	// 2. 快速排序 
	// 选择一个基准元素, 通常选择第一个元素或者最后一个元素
	// 通过一趟扫描, 将待排序列分成两部分, 一部分比基准元素小, 一部分大于等于基准元素
	// 此时基准元素在其排好序后的正确位置, 然后再用同样的方法递归地排序划分的两部分
	public function quick($arr = [])
	{
		$length = count($arr);
		if ($length <= 1) {
			return $arr;
		}
		$base = $arr[0];
		$left = $right = [];
		for ($i = 1; $i < $length; $i++) {
			$arr[$i] >= $base ? $right[] = $arr[$i] : $left[] = $arr[$i];
		}
		$l = $this->quick($left);
		$r = $this->quick($right);
		return array_merge($l, [$base], $r);
	}
	// 3. 插入排序
	// 在要排序的一组数中, 假设前面的数已经是排好顺序的, 现在要把第n个数插到前面的有序数中, 使得这n个数也是排好顺序的
	// 如此反复循环, 直到全部排好顺序
	public function insert($arr = [])
	{
		$len = count($arr); 
	    for ($i = 1; $i < $len; $i++) {
	        $tmp = $arr[$i];
	        // 内层循环控制，比较并插入
	        for ($j = $i - 1; $j >= 0; $j--) {
	            if ($tmp < $arr[$j]) {
	                // 发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
	                $arr[$j+1] = $arr[$j];
	                $arr[$j] = $tmp;
	            } else {
	                // 如果碰到不需要移动的元素，由于是已经排序好是数组，则前面的就不需要再次比较了。
	                break;
	            }
	        }
	    }
	    return $arr;
	}
}
$sort = new Sorting();
$arr = [
	10, 2, 6, 4, 3
];
print_r($sort->bubble($arr));
print_r($sort->quick($arr));
print_r($sort->insert($arr));
