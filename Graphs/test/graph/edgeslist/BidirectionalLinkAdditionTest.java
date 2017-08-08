package graph.edgeslist;

import graph.Edge;
import graph.Graph;
import graph.Node;
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
        g = new EdgesListGraph<>();
        source = new Node<>(1);
        destination = new Node<>(2);
        g.addEdge(1, 2);
        g.addEdge(2, 1);
    }

    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    public void testEdgeExists() {
        assertTrue(g.getEdges().contains(new Edge<>(source, destination)));
    }

    @Test
    public void testNode1To2Connection() {
        assertTrue(g.connectionExists(source, destination));
    }

    @Test
    public void testNode2To1Connection() {
        assertTrue(g.connectionExists(destination, source));
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