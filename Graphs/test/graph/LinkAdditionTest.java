package graph;

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
    public void testNode1To2Connection() {
        Node one = edge.getOne();
        Node two = edge.getTwo();
        assertTrue(g.connectionExistsBetween(one, two));
    }

    @Test
    public void testNode2To1Connection() {
        Node one = edge.getOne();
        Node two = edge.getTwo();
        assertTrue(g.connectionExistsBetween(two, one));
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