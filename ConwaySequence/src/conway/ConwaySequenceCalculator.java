package conway;

/**
 * Created by mcaci on 9/5/17.
 */
public class ConwaySequenceCalculator {

    public String conway(String sequence, int level) {
        for (int i = 1; i < level; i++) {
            sequence = conwayIteration(sequence);
        }
        return sequence;
    }

    private String conwayIteration(String sequence) {
        StringBuilder conwaySequenceBuilder = new StringBuilder();
        DigitSequenceBuilder digitSequenceBuilder = new DigitSequenceBuilder(sequence);
        for (char sequenceDigit : sequence.toCharArray()) {
            if(sequenceDigit != digitSequenceBuilder.getCurrentDigit()) {
                computeAndAppend(conwaySequenceBuilder, digitSequenceBuilder.getSequence());
                digitSequenceBuilder.reset(sequenceDigit);
            }
            digitSequenceBuilder.append(sequenceDigit);
        }
        computeAndAppend(conwaySequenceBuilder, digitSequenceBuilder.getSequence());
        return conwaySequenceBuilder.toString();
    }

    private void computeAndAppend(StringBuilder conwaySequenceBuilder, String sequenceOfSingleDigits) {
        String conwaySequencePart = computeConwayOnSequencesOfOneSingleDigit(sequenceOfSingleDigits);
        conwaySequenceBuilder.append(conwaySequencePart);
    }

    private String computeConwayOnSequencesOfOneSingleDigit(String sequence) {
        StringBuilder conwaySequenceBuilder = new StringBuilder();
        conwaySequenceBuilder.append(sequence.length());
        conwaySequenceBuilder.append(sequence.charAt(0));
        return conwaySequenceBuilder.toString();
    }
}
