package graph;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/2/17.
 */
public abstract  class GraphSkeleton<T> implements Graph<T> {

    private final List<Node<T>> nodes;

    public GraphSkeleton() {
        this.nodes = new LinkedList<>();
    }

    @Override
    public void addEdge(T source, T destination) {
        addNode(source);
        addNode(destination);
        addEdgeItself(source, destination);
    }

    protected abstract void addEdgeItself(T source, T destination);

    protected void add(Node<T> node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    @Override
    public List<Node<T>> getNodes() {
        return nodes;
    }

    @Override
    public void removeNode(Node<T> node) {
        nodes.removeIf(n -> n.equals(node));
    }
}
