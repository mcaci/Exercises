package graph.adjacencylist;

import graph.Node;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/2/17.
 */
public class Graph<T> {

    private List<Node<T>> nodes;

    public Graph() {
        nodes = new LinkedList<>();
    }

    public void add(AdjacencyListNode<T> node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    public void add(AdjacencyListNode<T> source, AdjacencyListNode<T> destination, boolean isBidirectional) {
        add(source);
        add(destination);
        AdjacencyListNode<T> sourceNode = (AdjacencyListNode<T>)this.getNode(source.getContent());
        AdjacencyListNode<T> destinationNode = (AdjacencyListNode<T>)this.getNode(destination.getContent());
        sourceNode.getAdjacentNodes().add(destinationNode);
        if(isBidirectional){
            destinationNode.getAdjacentNodes().add(sourceNode);
        }
    }

    public void add(AdjacencyListNode<T> source, AdjacencyListNode<T> destination) {
        add(source,destination,false);
    }


    public void removeEdge(AdjacencyListNode<T> source, AdjacencyListNode<T> destination) {
        source.getAdjacentNodes().removeIf(e -> e.equals(destination));
    }

    public List<Node<T>> getNodes() {
        return nodes;
    }

    public boolean connectionExistsBetween(AdjacencyListNode<T> source, AdjacencyListNode<T> destination) {
        return source.getAdjacentNodes().contains(destination);
    }

    public Node<T> getNode(T content) {
        Node<T> nodeWithOne = null;
        for (Node<T> node : this.getNodes()) {
            if (node.getContent().equals(content)) {
                nodeWithOne = node;
                break;
            }
        }
        return nodeWithOne;
    }
}
