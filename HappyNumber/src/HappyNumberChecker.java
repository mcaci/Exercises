/**
 * Created by mcaci on 8/1/17.
 */
public class HappyNumberChecker {

    public boolean isHappy(int number) {
        return checkHappinessCondition(number);
    }

    private boolean checkHappinessCondition(int number) {
        while (number > 9) {
            number = sumDigits(number);
        }
        return number == 1 || number == 7;
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