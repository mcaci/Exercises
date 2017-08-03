package graph;

/**
 * Created by mcaci on 8/2/17.
 */
public class Edge {

    private final Node one;
    private final Node two;

    public Edge(Node one, Node two) {
        this.one = one;
        this.two = two;
    }

    public Edge(int source, int destination) {
        this(new Node(source), new Node(destination));
    }

    public Edge reverse() {
        return new Edge(getTwo(), getOne());
    }

    public Node getOne() {
        return one;
    }

    public Node getTwo() {
        return two;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Edge edge = (Edge) o;

        if (one != null ? !one.equals(edge.one) : edge.one != null) return false;
        return two != null ? two.equals(edge.two) : edge.two == null;
    }

    @Override
    public int hashCode() {
        int result = one != null ? one.hashCode() : 0;
        result = 31 * result + (two != null ? two.hashCode() : 0);
        return result;
    }
}
