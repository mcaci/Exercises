package graph.adjacencylist;

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

    private void add(Node<T> node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    public void add(Node<T> source, Node<T> destination, boolean isBidirectional) {
        add(source);
        add(destination);
        source.getAdjacentNodes().add(destination);
        if(isBidirectional){
            destination.getAdjacentNodes().add(source);
        }
    }

    public void add(Node<T> source, Node<T> destination) {
        add(source,destination,false);
    }

    public void add(T sourceNodeContent, T destinationNodeContent) {
        add(sourceNodeContent,destinationNodeContent,false);
    }


    public void add(T sourceNodeContent, T destinationNodeContent, boolean isBidirectional) {
        Node<T> source = new Node<>(sourceNodeContent);
        Node<T> destination = new Node<>(destinationNodeContent);
        add(source);
        add(destination);
        source.getAdjacentNodes().add(destination);
        if(isBidirectional){
            destination.getAdjacentNodes().add(source);
        }
    }

    public void removeEdge(Node<T> source, Node<T> destination) {
        source.getAdjacentNodes().removeIf(e -> e.equals(destination));
    }

    public List<Node<T>> getNodes() {
        return nodes;
    }

    public boolean connectionExistsBetween(Node<T> source, Node<T> destination) {
        return source.getAdjacentNodes().contains(destination);
    }
}
