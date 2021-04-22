
// @Title: 整数反转 (Reverse Integer)
// @Author: 846188037@qq.com
// @Date: 2020-06-08 21:01:23
// @Runtime: 92 ms
// @Memory: 34.7 MB

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    x=String(x);
    var s="";
    if(x.charAt(0)=="-")
    {  
        for(var i=x.length-1;i>0;i--)
        {
            s+=x.charAt(i)
        }
        s="-"+s;

    }
    else{
        for(var i=x.length-1;i>-1;i--)
        {
            s+=x.charAt(i)
        }

    }
    s=Number(s);
    if(s<Math.pow(-2,31)||s>=Math.pow(2,31))
    {
        s=0;
    }

   
    return s

};
