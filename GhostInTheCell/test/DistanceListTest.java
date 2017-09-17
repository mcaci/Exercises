import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 9/8/17.
 */
public class DistanceListTest {

    private static final String[] DISTANCE_INPUT = {"1 2 20", "1 3 15", "2 3 18"};
    private final DistancesMapBuilder distancesMapBuilder = new DistancesMapBuilder();

    @Before
    public void setUp() throws Exception {
        for (String distancePair : DISTANCE_INPUT) {
            distancesMapBuilder.addDistancePair(distancePair);
        }
    }

    @Test
    public void testDistancesListContainsNode1() throws Exception {
        testOriginNodePresence(1);
    }

    @Test
    public void testDistancesListContainsNode2() throws Exception {
        testOriginNodePresence(2);
    }

    @Test
    public void test1315Element() throws Exception {
        testValuePresence(15, 1, 3);
    }

    @Test
    public void test2318Element() throws Exception {
        testValuePresence(18, 2, 3);
    }

    @Test
    public void test1220Element() throws Exception {
        testValuePresence(20, 1, 2);
    }

    private void testOriginNodePresence(int origin) {
        assertTrue(distancesMapBuilder.build().containsKey(origin));
    }

    private void testValuePresence(int expectedDistance, int source, int destination) {
        assertEquals(expectedDistance, (int) distancesMapBuilder.build().get(source).get(destination));
    }

}
