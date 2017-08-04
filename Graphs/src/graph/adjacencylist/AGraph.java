package graph.adjacencylist;

import graph.Graph;
import graph.Node;
import graph.edgeslist.Edge;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

/**
 * Created by mcaci on 8/2/17.
 */
public class AGraph<T> implements Graph<T> {

    private List<Node<T>> nodes;

    public AGraph() {
        nodes = new LinkedList<>();
    }

    @Override
    public void addNode(T node) {
        add(new AdjacencyListNode<>(node));
    }

    @Override
    public void addEdge(T source, T destination) {
        addNode(source);
        addNode(destination);
        Node<T> sourceNode = getNode(source);
        Node<T> destinationNode = getNode(destination);
        ((AdjacencyListNode<T>) sourceNode).getAdjacentNodes().add(destinationNode);
    }

    @Override
    public void removeNode(Node<T> node) {
        nodes.removeIf(n -> n.equals(node));
    }

    @Override
    public void removeEdge(Node<T> source, Node<T> destination) {
        ((AdjacencyListNode<T>) source).getAdjacentNodes().removeIf(n -> Objects.equals(n, destination));
    }

    @Override
    public boolean connectionExists(Node<T> source, Node<T> destination) {
        return ((AdjacencyListNode<T>) source).getAdjacentNodes().contains(destination);
    }

    @Override
    public List<Edge<T>> getEdges() {
        List<Edge<T>> edges = new LinkedList<>();
        for (Node<T> node : nodes) {
            for (Node<T> adjacentNode : ((AdjacencyListNode<T>) node).getAdjacentNodes()) {
                edges.add(new Edge<>(node, adjacentNode));
            }
        }
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
