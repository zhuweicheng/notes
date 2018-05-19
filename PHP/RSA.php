<?php

/**
 * RSA加密PHP版本
 */
class RSA
{
    /**
     * 对字符串进行RSA加密
     *
     * @param string $channelPass 待加密字符串
     * @return bool|string 加密后十六进制字符串
     */
    public static function getSign($channelPass = '')
    {
        // 不同平台之间转化过的对等公钥文件
        $pubKeyFile = "file://" . dirname(__FILE__) . "/publickey.pem";
        // 读取公钥内容
        $pubKey = file_get_contents($pubKeyFile);
        /**
         * 从证书中提取公钥并准备使用
         * @link http://php.net/manual/zh/function.openssl-pkey-get-public.php
         */
        $res = openssl_pkey_get_public($pubKey);
        if ($res) {
            /**
             * 用公钥加密数据
             * @link http://php.net/manual/zh/function.openssl-public-encrypt.php
             */
            $opt = openssl_public_encrypt($channelPass, $result, $res);
            if ($opt) {
                /**
                 * 函数把包含数据的二进制字符串转换为十六进制值
                 * @link http://php.net/manual/zh/function.bin2hex.php
                 */
                return bin2hex($result);
            }
        }

        return false;
    }
}

echo RSA::getSign('Test password');
echo "\n\n";
