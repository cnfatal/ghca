<?php

/* 电信天翼飞Young 3.09 相关算号算法 */
namespace ghca;

class Ghca
{
    public static function encrypt($account,$password)
    {
        $account=strtoupper($account);
        $key = 'aI0fC8RslXg6HXaKAUa6kpvcAXszvTcxYP8jmS9sBnVfIqTRdJS1eZNHmBjKN28j';
        $time = time();
        $time_hex = dechex($time);

        //将时间转为4个字符
        $time_byte_one =  chr($time>>24);
        $time_byte_two =  chr($time>>16);
        $time_byte_thr =  chr($time>>8);
        $time_byte_for =  chr($time);

        $password_ascii_sum = 0;
        for ($i=0;$i<=strlen($account);$i++)
        {
            $password_ascii_sum+=ord(substr($account,$i,1));
        }
        $time_mod_pswlen = $time % strlen($password);
        $time_mod = $time_mod_pswlen<1?1:0+$time_mod_pswlen;

        $seed = $time_mod - ($time_mod_pswlen==strlen($password) ? 1 : 0);
        $account_last_four = sprintf('%04x',$password_ascii_sum ^ $seed);
        $split_len = $time_mod -($time_mod==strlen($password)? 1 : 0) + 1;

        $split_password_one = substr($password,0,$split_len);
        $split_password_two = substr($password,$split_len);
        $split_key_one = substr($key,0,60-strlen($split_password_one));
        $split_key_two = substr($key,0,64-strlen($account)-strlen($split_password_two));


        $md5_str = $time_byte_one.$time_byte_two.$time_byte_thr.$time_byte_for.$split_key_one.$split_password_one.$account.$split_key_two.$split_password_two;
        //MD5 原始二进制流 注意 “true” 参数
        $md5_once = md5($md5_str,true);
        $md5_twice = md5($md5_once);
        $md5_final = substr($md5_twice,0,16);

        $encrypt_account = '~ghca'.strtoupper($time_hex.'2023'.$md5_final.$account_last_four.$account);

        $array = array(
            'data'=>array(
                'type'=>'ghca',
                'id'=>$account,
                'attribute'=>array(
                    'account'=>$account,
                    'password'=>$password,
                    'encrypt'=>$encrypt_account
                )
            )
        );
        return $array;
    }
}
