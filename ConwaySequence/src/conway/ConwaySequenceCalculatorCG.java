package conway;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

/**
 * Created by mcaci on 9/5/17.
 */
public class ConwaySequenceCalculatorCG {

    public String conway(String sequence, int level) {
        int R = Integer.parseInt(sequence);
        int L = level;

        List<Integer> sequenceLine = getSequenceLine(R, L);
        return printSequenceLine(sequenceLine);
    }


    private static List<Integer> getSequenceLine(int first, int lineNumber) {

        List<Integer> currentLine = new ArrayList<>();
        currentLine.add(Integer.valueOf(first));

        for (int i = 2; i <= lineNumber; i++) {
            currentLine = getNextSequenceLine(currentLine);
        }

        return currentLine;
    }

    private static List<Integer> getNextSequenceLine(List<Integer> currentLine) {

        List<Integer> nextLine = new ArrayList<>();

        int startIndex = 0, currentIndex = 1;
        while (currentIndex < currentLine.size()) {
            if (currentLine.get(startIndex) != currentLine.get(currentIndex)) {
                nextLine.add(currentIndex - startIndex);
                nextLine.add(currentLine.get(startIndex));
                startIndex = currentIndex;
            }
            currentIndex++;
        }

        nextLine.add(currentIndex - startIndex);
        nextLine.add(currentLine.get(startIndex));

        return nextLine;
    }

    private static String printSequenceLine(List<Integer> sequenceLine) {
        return sequenceLine.stream().map(i -> i.toString()).collect(Collectors.joining(" "));
    }
}
