package vol1.graph;

/**
 * Created by mcaci on 8/2/17.
 */
public class Edge extends Pair<Node> {

    public Edge(int source, int destination) {
        super(new Node(source), new Node(destination));
    }

    public Edge(Node one, Node two) {
        super(one, two);
    }

    public Edge reverse() {
        return new Edge(getTwo(), getOne());
    }
}
