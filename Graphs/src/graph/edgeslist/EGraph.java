package graph.edgeslist;

import graph.Graph;
import graph.Node;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

/**
 * Created by mcaci on 8/2/17.
 */
public class EGraph<T> implements Graph<T> {

    private List<Edge<T>> edges;
    private List<Node<T>> nodes;

    public EGraph() {
        edges = new LinkedList<>();
        nodes = new LinkedList<>();
    }

    @Override
    public void addNode(T node) {
        add(new Node<>(node));
    }

    @Override
    public void addEdge(T source, T destination) {
        addNode(source);
        addNode(destination);
        edges.add(new Edge<>(source, destination));
    }

    @Override
    public void removeNode(Node<T> node) {
        nodes.removeIf(n -> n.equals(node));
    }

    @Override
    public void removeEdge(Node<T> source, Node<T> destination) {
        Edge<T> edge = new Edge<>(source, destination);
        edges.removeIf(e -> Objects.equals(e, edge));
    }

    @Override
    public boolean connectionExists(Node<T> source, Node<T> destination) {
        Edge<T> edge = new Edge<>(source, destination);
        return edges.contains(edge);
    }

    @Override
    public List<Edge<T>> getEdges() {
        return edges;
    }

    @Override
    public List<Node<T>> getNodes() {
        return nodes;
    }

    private void add(Node<T> node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

}
