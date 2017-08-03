package graph.edgeslist;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/2/17.
 */
public class Graph<T> {


    private List<Edge<T>> edges;
    private List<Node<T>> nodes;

    public Graph() {
        edges = new LinkedList<>();
        nodes = new LinkedList<>();
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

    public List<Edge<T>> getEdges() {
        return edges;
    }

    public void removeEdge(Edge<T> edge) {
        edges.removeIf(e -> e.equals(edge));
    }

    public List<Node<T>> getNodes() {
        return nodes;
    }

    public boolean connectionExistsBetween(Node<T> one, Node<T> two) {
        return edges.contains(new Edge<>(one, two));
    }
}
