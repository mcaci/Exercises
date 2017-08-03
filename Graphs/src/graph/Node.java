package graph;

/**
 * Created by mcaci on 8/2/17.
 */
public class Node<T> {

    private final T t;

    public Node(T t) {
        this.t = t;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Node<?> node = (Node<?>) o;

        return t != null ? t.equals(node.t) : node.t == null;
    }

    @Override
    public int hashCode() {
        return t != null ? t.hashCode() : 0;
    }
}
