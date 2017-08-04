package graph.edgeslist;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkBreakingTest {

    private EGraph<Integer> g;

    @Before
    public void setUp() throws Exception {
        g = new EGraph<>();
        g.add(1, 2);
    }

    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    // NOTE: will not work if Node does not have equals and hashcode implemented
    public void testBreakExistingLink() {
        g.removeEdge(new Edge<>(1, 2));
        assertTrue(g.getEdges().isEmpty());
    }

    @Test
    public void testBreakNonExistingLink() {
        g.removeEdge(new Edge<>(1, 0));
        assertTrue(g.getEdges().size() == 1);
    }
}