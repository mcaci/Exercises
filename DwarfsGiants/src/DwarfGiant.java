import java.util.Collection;

/**
 * Created by nikiforos on 07/09/17.
 */
public class DwarfGiant {


    private final Collection<Edge> edges;

    public DwarfGiant(Collection<Edge> edges) {
        this.edges = edges;
    }

    public int compute() {
        return edges.size();
    }
}
