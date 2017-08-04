package graph.adjacencylist;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkBreakingTest {

    private Graph<Integer> g;
    private AdjacencyListNode<Integer> source;
    private AdjacencyListNode<Integer> destination;

    @Before
    public void setUp() throws Exception {
        g = new Graph<>();
        source = new AdjacencyListNode<>(1);
        destination = new AdjacencyListNode<>(2);
        g.add(source, destination);
    }
    @After
    public void tearDown() throws Exception {
        g = null;
    }

    @Test
    // NOTE: will not work if Node does not have equals and hashcode implemented
    public void testBreakExistingLink() {
        g.removeEdge(source, destination);
        assertTrue(source.getAdjacentNodes().isEmpty());
    }

    @Test
    public void testBreakNonExistingLink() {
        g.removeEdge(source, new AdjacencyListNode<>(0));
        assertTrue(source.getAdjacentNodes().size() == 1);
    }
}