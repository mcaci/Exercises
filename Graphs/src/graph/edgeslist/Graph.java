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
        Node<T> node = new Node<>(nodeContent);
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    public void add(T sourceNodeContent, T destinationNodeContent) {
        add(sourceNodeContent,destinationNodeContent,false);
    }

    public void add(T sourceNodeContent, T destinationNodeContent, boolean isBidirectional) {
        add(sourceNodeContent);
        add(destinationNodeContent);
        edges.add(new Edge<>(sourceNodeContent, destinationNodeContent));
        if(isBidirectional){
            edges.add(new Edge<>(destinationNodeContent, sourceNodeContent));
        }
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
