package graph.edgeslist;

/**
 * Created by mcaci on 8/2/17.
 */
public class Edge<T> {

    private final Node<T> one;
    private final Node<T> two;

    public Edge(Node<T> one, Node<T> two) {
        this.one = one;
        this.two = two;
    }

    public Edge(T one, T two) {
        this(new Node<>(one), new Node<>(two));
    }

    public Edge<T> reverse() {
        return new Edge<>(getTwo(), getOne());
    }

    public Node<T> getOne() {
        return one;
    }

    public Node<T> getTwo() {
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
