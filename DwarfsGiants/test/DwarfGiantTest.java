import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.Assert.assertEquals;

/**
 * Created by nikiforos on 07/09/17.
 */
public class DwarfGiantTest {

    @Test
    public void testOneLink() throws Exception {
        List<Edge> edges = new ArrayList<>();
        edges.add(new Edge(1, 2));

        DwarfGiant dwarfGiant = new DwarfGiant(edges);
        assertEquals(1, dwarfGiant.compute());
    }

    @Test
    public void testTwoLinks() throws Exception {
        List<Edge> edges = new ArrayList<>();
        edges.add(new Edge(1, 2));
        edges.add(new Edge(2, 3));

        DwarfGiant dwarfGiant = new DwarfGiant(edges);
        assertEquals(2, dwarfGiant.compute());
    }
}
