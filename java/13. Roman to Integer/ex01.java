package ex;

import java.util.HashMap;

public class ex01 {
    public static void main(String[] args) {

        int output = romanToInt(args[0]);
        System.out.println(output);
    }

    public static int romanToInt(String s) {
        s = s.toUpperCase();
        int actualNumber;
        int lastNumber = 0;
        int total = 0;

        HashMap<Character, Integer> map = new HashMap<>();
        char[] roman = {'I', 'V', 'X', 'L', 'C', 'D', 'M'};
        int[] decimal = {1, 5, 10, 50, 100, 500, 1000};

        for (int i = 0; i < decimal.length ; i++) {
            map.put(roman[i], decimal[i]);
        }
        for (int i = s.length() - 1; i >= 0; i--) {
            actualNumber = map.get(s.charAt(i));

            if (lastNumber > actualNumber) {
                total -= actualNumber;
            } else {
                total += actualNumber;
            }
            lastNumber = actualNumber;
        }
        return total;
    }
}
