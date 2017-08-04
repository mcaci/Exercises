package graph.edgeslist;

import graph.Node;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static junit.framework.TestCase.assertFalse;
import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkAdditionTest {

    private EGraph<Integer> g;
    private Edge<Integer> edge;

    @Before
    public void setUp() throws Exception {
        g = new EGraph<>();
        edge = new Edge<>(1, 2);
        g.addEdge(1, 2);
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
    public void testNode1To2Connection() {
        Node<Integer> one = edge.getOne();
        Node<Integer> two = edge.getTwo();
        assertTrue(g.connectionExists(one, two));
    }

    @Test
    public void testNode2To1Connection() {
        Node<Integer> one = edge.getOne();
        Node<Integer> two = edge.getTwo();
        assertFalse(g.connectionExists(two, one));
    }

    @Test
    public void testSourceNodeExists() {
        assertTrue(g.getNodes().contains(edge.getOne()));
    }

    @Test
    public void testDestinationNodeExists() {
        assertTrue(g.getNodes().contains(edge.getTwo()));
    }
}