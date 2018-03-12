/**
 * 兔子问题的递归算法（斐波那契数列)
 * 有一对兔子，从出生后第三个月起每个月都生一对兔子，小兔子长到第三个月后又生一对兔子，假如兔子都不死，每个月兔子对数为多少
 * 1月1对 2月2对 3月3对  4月5对  5月8对    1 2 3 5 8 13
 * F(N) = f(n-1)+ f(n-2)
 * @param $month
 * @return integer
 * @version 1.0
 * @time 2018.2.27
 * @anthor zwcheng
 */
function counts($month) {
    if ($month == 1 || $month == 2) {
        return $month;
    } else {
        return counts($month-2) + counts($month-1);
    }
}

echo counts(6);
