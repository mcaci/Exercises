package graph.adjacencylist;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import static junit.framework.TestCase.assertFalse;
import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 8/2/17.
 */
public class LinkAdditionTest {

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
    public void testEdgeExists() {
        assertTrue(source.getAdjacentNodes().contains(destination));
    }

    @Test
    public void testNode1To2Connection() {
        assertTrue(g.connectionExistsBetween(source, destination));
    }

    @Test
    public void testNode2To1Connection() {
        assertFalse(g.connectionExistsBetween(destination, source));
    }

    @Test
    public void testSourceNodeExists() {
        assertTrue(g.getNodes().contains(source));
    }

    @Test
    public void testDestinationNodeExists() {
        assertTrue(g.getNodes().contains(destination));
    }

    @Test
    public void testNoDuplicateWillBeAdded() {
        g.add(new AdjacencyListNode<>(new Integer(1))); // adding same node as in setup
        assertTrue(g.getNodes().size() == 2);
    }

    @Test
    public void testEdgeIsAddedWhenAddingADuplicateNode() {
        g.add(new AdjacencyListNode<>(1), new AdjacencyListNode<>(5)); // adding same source as in setup
        int numberOfEdgesAdjacentToNode1 = ((AdjacencyListNode<Integer>)g.getNode(1)).getAdjacentNodes().size();
        assertTrue("actual size of the edges is " + numberOfEdgesAdjacentToNode1, numberOfEdgesAdjacentToNode1 == 2);
    }

}