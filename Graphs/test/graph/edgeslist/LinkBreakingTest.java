package graph.edgeslist;

import graph.Node;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkBreakingTest {

    private EdgesListGraph<Integer> g;

    @Before
    public void setUp() throws Exception {
        g = new EdgesListGraph<>();
        g.addEdge(1, 2);
    }

    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    // NOTE: will not work if Node does not have equals and hashcode implemented
    public void testBreakExistingLink() {
        Node<Integer> source = g.getNode(1);
        Node<Integer> destination = g.getNode(2);
        g.removeEdge(source, destination);
        assertTrue(g.getEdges().isEmpty());
    }

    @Test
    public void testBreakNonExistingLink() {
        Node<Integer> source = g.getNode(1);
        Node<Integer> destination = g.getNode(0);
        g.removeEdge(source, destination);
        assertTrue(g.getEdges().size() == 1);
    }
}