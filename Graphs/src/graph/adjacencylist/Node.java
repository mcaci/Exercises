package graph.adjacencylist;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/2/17.
 */
public class Node<T> {

    private final T t;
    private final List<Node<T>> adjacentNodes;

    public Node(T t) {
        this.t = t;
        adjacentNodes = new LinkedList<>();
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

    public List<Node<T>> getAdjacentNodes() {
        return adjacentNodes;
    }
}
