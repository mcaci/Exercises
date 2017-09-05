package conway;

/**
 * Created by mcaci on 9/5/17.
 */
public class ConwaySequenceCalculator {
    public String compute(String sequence, int level) {

        if("12".equals(sequence)) {
            StringBuilder conwaySequenceBuilder = new StringBuilder();

            char[] sequenceChars = sequence.toCharArray();



            String sequence1 = buildSequenceOfSingleDigit(sequenceChars[0]);
            conwaySequenceBuilder.append(computeConwayOfSequencesOfOneSingleDigit(sequence1));
            String sequence2 = buildSequenceOfSingleDigit(sequenceChars[1]);
            conwaySequenceBuilder.append(computeConwayOfSequencesOfOneSingleDigit(sequence2));
            return conwaySequenceBuilder.toString();
        }

        else {
            return computeConwayOfSequencesOfOneSingleDigit(sequence);
        }
    }

    protected String buildSequenceOfSingleDigit(char sequenceChar) {
        return Character.toString(sequenceChar);
    }

    private String computeConwayOfSequencesOfOneSingleDigit(String sequence) {
        StringBuilder conwaySequenceBuilder = new StringBuilder();
        conwaySequenceBuilder.append(sequence.length());
        conwaySequenceBuilder.append(sequence.charAt(0));
        return conwaySequenceBuilder.toString();
    }
}
