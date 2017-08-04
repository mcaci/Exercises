package graph.edgeslist;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static junit.framework.TestCase.assertFalse;
import static org.junit.Assert.assertTrue;

import graph.Node;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkAdditionTest {

    private Graph<Integer> g;
    private Edge<Integer> edge;

    @Before
    public void setUp() throws Exception {
        g = new Graph<>();
        edge = new Edge<>(1, 2);
        g.add(1, 2);
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
        assertTrue(g.connectionExistsBetween(one, two));
    }

    @Test
    public void testNode2To1Connection() {
        Node<Integer> one = edge.getOne();
        Node<Integer> two = edge.getTwo();
        assertFalse(g.connectionExistsBetween(two, one));
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