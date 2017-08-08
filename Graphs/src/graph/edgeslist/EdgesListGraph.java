package graph.edgeslist;

import graph.Edge;
import graph.Node;
import graph.GraphSkeleton;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

/**
 * Created by mcaci on 8/2/17.
 */
public class EdgesListGraph<T> extends GraphSkeleton<T> {

    private List<Edge<T>> edges;

    public EdgesListGraph() {
        edges = new LinkedList<>();
    }

    @Override
    public void addNode(T node) {
        add(new Node<>(node));
    }

    @Override
    protected void addEdgeItself(T source, T destination) {
        getEdges().add(new Edge<>(source, destination));
    }


    @Override
    public List<Edge<T>> getEdges() {
        return edges;
    }

    @Override
    public void removeEdge(Node<T> source, Node<T> destination) {
        Edge<T> edge = new Edge<>(source, destination);
        this.getEdges().removeIf(e -> Objects.equals(e, edge));
    }

    @Override
    public boolean connectionExists(Node<T> source, Node<T> destination) {
        Edge<T> edge = new Edge<>(source, destination);
        return this.getEdges().contains(edge);
    }


}
