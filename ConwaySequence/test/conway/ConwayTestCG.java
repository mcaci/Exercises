package conway;

import org.junit.Before;
import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

/**
 * Created by mcaci on 9/5/17.
 */
public class ConwayTestCG {

    ConwaySequenceCalculatorCG conwaySequenceCalculator;

    @Before
    public void setUp() {
        this.conwaySequenceCalculator = new ConwaySequenceCalculatorCG();
    }

    private String computeSequence(String sequence, int level) {
        return conwaySequenceCalculator.conway(sequence, level);
    }

    @Test
    public void testConwayWithBase1Level1() throws Exception {
        assertEquals("1", computeSequence("1", 1));
    }

    @Test
    public void testConwayWithBase1Level2() throws Exception {
        assertEquals("1 1", computeSequence("1", 2));
    }

    @Test
    public void testConwayWithBase111Level2() throws Exception {
        assertEquals("3 1", computeSequence("111", 2));
    }

    @Test
    public void testConwayWithBase12Level2() throws Exception {
        assertEquals("1 1 1 2", computeSequence("12", 2));
    }

    @Test
    public void testConwayWithBase1231Level2() throws Exception {
        assertEquals("1 1 1 2 1 3 1 1", computeSequence("1231", 2));
    }

    @Test
    public void testConwayWithBase12245551Level2() throws Exception {
        assertEquals("1 1 2 2 1 4 3 5 1 1", computeSequence("12245551", 2));
    }

    @Test
    public void testConwayWithBase1Level3() throws Exception {
        assertEquals("2 1", computeSequence("1", 3));
    }

    @Test
    public void testConwayWithBase2Level3() throws Exception {
        assertEquals("1 1 1 2", computeSequence("2", 3));
    }

    @Test
    public void testConwayWithBase12245551Level3() throws Exception {
        assertEquals("2 1 2 2 1 1 1 4 1 3 1 5 2 1", computeSequence("12245551", 3));
    }

    @Test
    public void testConwayWithBase12245551Level6() throws Exception {
        assertEquals("3 1 2 2 1 1 2 2 1 1 1 3 1 2 2 1 1 4 1 3 2 1 2 3 2 1 1 5 3 1 2 2 1 1", computeSequence("12245551", 6));
    }

    @Test
    public void testConwayWithBase25Level10() throws Exception {
        assertEquals("3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 2 1 1 5", computeSequence("25", 10));
    }

    @Test
    public void testConwayWithBase33Level25() throws Exception {
        assertEquals("3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 1 2 3 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 2 2 1 1 2 1 3 2 1 1 3 1 2 1 1 1 3 2 2 2 1 1 2 3 1 1 3 1 1 2 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 1 2 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 1 1 1 2 1 3 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 2 1 2 2 1 1 1 3 1 2 2 1 1 2 1 3 2 1 1 3 1 2 1 1 1 3 2 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 2 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 1 2 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 1 1 1 2 1 3 3 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 3 3 1 1 2 1 3 2 1 1 3 2 2 1 3 2 1 1 2 3 1 1 3 2 1 3 2 2 1 1 2 1 1 1 3 1 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 2 1 2 2 2 1 1 2 1 1 2 3 2 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 1 1 1 2 1 3 2 1 3 2 1 1 2 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 1 1 1 2 1 3 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 1 2 3 1 1 2 1 1 2 3 2 2 2 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 1 2 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 1 2 3 1 1 2 1 1 2 3 2 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 1 1 1 2 1 3 1 1 1 2 1 3 2 1 1 2 3 1 1 3 2 1 3 2 2 1 1 2 1 1 1 3 1 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 2 1 2 2 2 1 1 2 1 1 2 3 2 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 3 3 1 1 2 1 3 2 1 2 3 1 2 3 1 1 2 1 1 1 3 1 1 2 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 3 2 2 1 3 2 1 1 3 2 1 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 3 1 1 2 1 3 2 1 2 3 2 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 2 2 1 3 2 1 1 2 3 1 1 3 2 1 3 2 2 1 1 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 3 2 2 1 1 3 3 2 1 1 3 2 2 1 1 2 2 1 1 2 1 3 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 1 2 3 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 2 2 1 1 2 1 3 2 1 1 3 3 1 1 2 1 3 2 1 1 3 2 1 2 2 2 1 2 2 1 1 1 3 1 2 2 1 1 3 1 2 1 1 1 3 2 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 1 1 1 2 1 3 3 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 3 3 1 1 2 1 3 2 1 1 3 2 2 1 3 2 1 1 2 3 1 1 3 2 1 3 2 2 1 1 2 1 1 1 3 1 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 2 1 2 2 2 1 1 2 1 1 2 3 2 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 1 3 1 1 1 2 3 1 1 3 3 2 1 1 1 2 1 3 1 2 2 1 1 2 3 1 1 3 1 1 1 2 3 1 1 2 1 1 1 3 3 1 1 2 1 1 1 3 1 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 3 2 2 1 1 1 3 1 2 2 1 1 2 1 3 1 1 1 2 1 3 1 2 2 1 1 2 1 3 2 1 1 3 2 1 3 2 2 1 1 2 3 1 1 3 1 1 2 2 2 1 2 3 2 1 1 2 1 1 1 3 1 2 2 1 1 3 2 2 1 1 1 3 1 2 2 1 1 2 1 3 2 1 1 3 1 1 1 2 3 1 1 3 2 2 3 1 1 2 1 1 1 3 2 1 3 2 2 1 2 3 1 2 2 1 1 3 2 2 2 1 2 2 2 1 1 2 1 1 2 3 2 2 2 1 1 3 3", computeSequence("33", 25));
    }


}
