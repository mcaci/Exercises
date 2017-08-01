/**
 * Created by mcaci on 8/1/17.
 */
public class HappyNumberChecker {

    boolean isHappy(int number) {
        while (number > 4) {
            number = sumDigits(number);
        }
        return number == 1;
    }

    private int sumDigits(int number) {
        int sum = 0;
        while (number > 0) {
            int digit = number % 10;
            sum += digit * digit;
            number = number / 10;
        }
        return sum;
    }
}