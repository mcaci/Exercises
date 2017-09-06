package conway;

/**
 * Created by mcaci on 9/6/17.
 */
class DigitSequenceBuilder {

    private StringBuilder sequenceOfSingleDigitsBuilder = new StringBuilder();
    private char currentDigit;

    public DigitSequenceBuilder(String digitSequence) {
        currentDigit = initializeCurrentDigit(digitSequence);
    }

    private char initializeCurrentDigit(String digitSequence) {
        return digitSequence.length() != 0 ? digitSequence.charAt(0) : ' ';
    }

    public String getSequence() {
        return sequenceOfSingleDigitsBuilder.toString();
    }

    public void reset(char sequenceDigit) {
        sequenceOfSingleDigitsBuilder = new StringBuilder();
        currentDigit = sequenceDigit;
    }

    public void append(char sequenceDigit) {
        sequenceOfSingleDigitsBuilder.append(sequenceDigit);
    }

    public char getCurrentDigit() {
        return currentDigit;
    }
}
