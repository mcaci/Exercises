package graph.adjacencylist;

import graph.GraphSkeleton;
import graph.Node;
import graph.Edge;

import java.util.LinkedList;
import java.util.List;
import java.util.Objects;

/**
 * Created by mcaci on 8/2/17.
 */
public class AdjacencyListGraph<T> extends GraphSkeleton<T> {

    @Override
    public void addNode(T node) {
        add(new AdjacencyListNode<>(node));
    }

    @Override
    protected void addEdgeItself(T source, T destination) {
        Node<T> sourceNode = getNode(source);
        Node<T> destinationNode = getNode(destination);
        ((AdjacencyListNode<T>) sourceNode).getAdjacentNodes().add(destinationNode);
    }

    @Override
    public List<Edge<T>> getEdges() {
        List<Edge<T>> edges = new LinkedList<>();
        for (Node<T> node : getNodes()) {
            for (Node<T> adjacentNode : ((AdjacencyListNode<T>) node).getAdjacentNodes()) {
                edges.add(new Edge<>(node, adjacentNode));
            }
        }
        return edges;
    }

    @Override
    public void removeEdge(Node<T> source, Node<T> destination) {
        ((AdjacencyListNode<T>) source).getAdjacentNodes().removeIf(n -> Objects.equals(n, destination));
    }

    @Override
    public boolean connectionExists(Node<T> source, Node<T> destination) {
        return ((AdjacencyListNode<T>) source).getAdjacentNodes().contains(destination);
    }

}
