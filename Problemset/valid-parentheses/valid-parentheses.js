
// @Title: 有效的括号 (Valid Parentheses)
// @Author: 846188037@qq.com
// @Date: 2020-07-11 21:05:00
// @Runtime: 68 ms
// @Memory: 34.6 MB

var isValid = function(s) {
    try{
    var s=Array.from(s);
    var s1=new Array();
    while (s.length!=0){
        if (s[0] == "(" || s[0] == "[" || s[0] == "{"){
            s1.push(s[0])
            s.shift()
        }
        switch(s[0]){
            case ")":
                if(s1.pop()=="("){
                    s.shift()
                    break
                }
                else{
                    return false
                }
            case "]":
                if(s1.pop()=="["){
                    s.shift()
                    break
                }
                else{
                    return false
                }
            case "}":
                if(s1.pop()=="{"){
                    s.shift()
                    break
                }
                else{
                    return false
                }
        }

    }
    if(s1.length != 0){
        return false
    }
    else{
        return true
    }}
    catch(error){return false}


};
