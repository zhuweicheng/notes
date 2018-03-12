<?php

function flow()
{
    // 假如当前登陆用户的COOKIE为该值
    $_COOKIE['id'] = 'vW8ySjnxAeLF5cV5-bo4NeL94Mu1SzDMm3myEIdQktiHtW0G';
    if (isset($_COOKIE['id'])) {
        $length = strlen($_COOKIE['id']);
        $num = 0;
        for ($i = 0; $i < $length; $i++) {
            $num += ord($_COOKIE['id'][$i]);
        }
        return $num % 100;
    }
    return rand(1, 100);
}
echo "当前权重\n";
echo flow();
echo "\n";
