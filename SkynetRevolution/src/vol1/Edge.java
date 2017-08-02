package vol1;

/**
 * Created by mcaci on 8/2/17.
 */
public class Edge extends Pair<Node> {

    public Edge(int source, int destination) {
        super(new Node(source), new Node(destination));
    }
}
