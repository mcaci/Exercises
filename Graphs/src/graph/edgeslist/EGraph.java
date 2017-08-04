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

    public List<Edge<T>> getEdges() {
        return edges;
    }

    public List<Node<T>> getNodes() {
        return nodes;
    }

    public void removeEdge(Edge<T> edge) {
        edges.removeIf(e -> e.equals(edge));
    }

    private void add(T nodeContent) {
        add(new Node<>(nodeContent));
    }

    private void add(Node<T> node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    public void add(Node<T> source, Node<T> destination) {
        add(source,destination,false);
    }

    public void add(Node<T> source, Node<T> destination, boolean isBidirectional) {
        add(source);
        add(destination);
        edges.add(new Edge<>(source, destination));
        if(isBidirectional){
            edges.add(new Edge<>(destination, source));
        }
    }

    public void add(T sourceNodeContent, T destinationNodeContent) {
        add(sourceNodeContent,destinationNodeContent,false);
    }

    public void add(T sourceNodeContent, T destinationNodeContent, boolean isBidirectional) {
        Node<T> source = new Node<>(sourceNodeContent);
        Node<T> destination = new Node<>(destinationNodeContent);
        add(source, destination, isBidirectional);
    }



    public boolean connectionExistsBetween(Node<T> one, Node<T> two) {
        return edges.contains(new Edge<>(one, two));
    }
}
