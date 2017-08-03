package graph.adjacencylist;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class BidirectionalLinkAdditionTest {

    private Graph<Integer> g;
    private Node<Integer> source;
    private Node<Integer> destination;

    @Before
    public void setUp() throws Exception {
        g = new Graph<>();
        source = new Node<>(1);
        destination = new Node<>(2);
        g.add(source, destination, true);
    }

    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    public void testEdgeExists() {
        assertTrue(source.getAdjacentNodes().contains(destination));
    }

    @Test
    public void testNode1To2Connection() {
        assertTrue(g.connectionExistsBetween(source, destination));
    }

    @Test
    public void testNode2To1Connection() {
        assertTrue(g.connectionExistsBetween(destination, source));
    }

    @Test
    public void testSourceNodeExists() {
        assertTrue(g.getNodes().contains(source));
    }

    @Test
    public void testDestinationNodeExists() {
        assertTrue(g.getNodes().contains(destination));
    }
}