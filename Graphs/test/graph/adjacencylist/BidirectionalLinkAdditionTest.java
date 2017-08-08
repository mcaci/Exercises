package graph.adjacencylist;

import graph.Graph;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class BidirectionalLinkAdditionTest {

    private Graph<Integer> g;
    private AdjacencyListNode<Integer> source;
    private AdjacencyListNode<Integer> destination;

    @Before
    public void setUp() throws Exception {
        g = new AdjacencyListGraph<>();
        g.addEdge(1, 2);
        g.addEdge(2, 1);
        source = (AdjacencyListNode<Integer>) g.getNode(1);
        destination = (AdjacencyListNode<Integer>) g.getNode(2);
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