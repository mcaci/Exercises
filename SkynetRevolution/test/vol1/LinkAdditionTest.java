package vol1;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkAdditionTest {

    private Graph g;
    private Edge edge;

    @Before
    public void setUp() throws Exception {
        g = new Graph();
        edge = new Edge(1, 2);
        g.add(edge);
    }

    @After
    public void tearDown() throws Exception {
        g = null;
        edge = null;
    }

    @Test
    public void testEdgeExists() {
        assertTrue(g.getEdges().contains(edge));
    }

    @Test
    public void testSourceNodeExists() {
        assertTrue(g.getNodes().contains(edge.one));
    }

    @Test
    public void testDestinationNodeExists() {
        assertTrue(g.getNodes().contains(edge.two));
    }
}