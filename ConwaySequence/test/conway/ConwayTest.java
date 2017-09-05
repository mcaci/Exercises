package conway;

import org.junit.Before;
import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

/**
 * Created by mcaci on 9/5/17.
 */
public class ConwayTest {

    ConwaySequenceCalculator conwaySequenceCalculator;

    @Before
    public void setUp() {
        this.conwaySequenceCalculator = new ConwaySequenceCalculator();
    }

    private String computeSequence(String sequence, int level) {
        return conwaySequenceCalculator.compute(sequence, level);
    }

    @Test
    public void testConwayWithBase1Level1() throws Exception {
        assertEquals("11", computeSequence("1", 1));
    }

    @Test
    public void testConwayWithBase111Level1() throws Exception {
        assertEquals("31", computeSequence("111", 1));
    }

    @Test
    public void testConwayWithBase12Level1() throws Exception {
        assertEquals("1112", computeSequence("12", 1));
    }

    @Ignore
    @Test
    public void testConwayWithBase1Level2() throws Exception {
        assertEquals("21", computeSequence("1", 2));
    }

    @Ignore
    @Test
    public void testConwayWithBase2Level2() throws Exception {
        assertEquals("1112", computeSequence("2", 2));
    }
}
