package graph.adjacencylist;

import graph.Node;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkBreakingTest {

    private AdjacencyListGraph<Integer> g;

    @Before
    public void setUp() throws Exception {
        g = new AdjacencyListGraph<>();
        g.addEdge(1,2);
    }
    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    // NOTE: will not work if Node does not have equals and hashcode implemented
    public void testBreakExistingLink() {
        AdjacencyListNode<Integer> source = (AdjacencyListNode<Integer>) g.getNode(1);
        Node<Integer> destination = g.getNode(2);
        g.removeEdge(source, destination);
        assertTrue(source.getAdjacentNodes().isEmpty());
    }

    @Test
    public void testBreakNonExistingLink() {
        AdjacencyListNode<Integer> source = (AdjacencyListNode<Integer>) g.getNode(1);
        g.removeEdge(source, new AdjacencyListNode<>(0));
        assertTrue(source.getAdjacentNodes().size() == 1);
    }
}