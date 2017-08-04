package graph.adjacencylist;

import graph.Node;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/4/17.
 */
public class AdjacencyListNode<T> extends Node<T> {

    private final List<Node<T>> adjacentNodes;

    public AdjacencyListNode(T t) {
        super(t);
        adjacentNodes = new LinkedList<>();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        return super.equals(o);
    }

    @Override
    public int hashCode() {
        int result = super.hashCode();
        result = 31 * result + (adjacentNodes != null ? adjacentNodes.hashCode() : 0);
        return result;
    }

    public List<Node<T>> getAdjacentNodes() {
        return adjacentNodes;
    }
}
