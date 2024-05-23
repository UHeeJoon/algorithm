import java.util.Stack;
class Solution {
    boolean solution(String s) {
        Stack<Character> stk = new Stack<>();
        for(char c : s.toCharArray()) {
            if(c=='(') {
                stk.push(c);
            } else {
                if(stk.isEmpty()){
                    return false;
                }
                stk.pop();
            }
        }
        return stk.isEmpty();
    }
}