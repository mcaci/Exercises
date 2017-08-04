package graph;

import graph.edgeslist.Edge;

import java.util.Collection;

/**
 * Created by mcaci on 8/4/17.
 */
public interface Graph<T> {

    void addNode(T node);
    void addEdge(T source, T destination);

    Collection<Node<T>> getNodes();
    Collection<Edge<T>> getEdges();

    void removeNode(Node<T> node);
    void removeEdge(Node<T> source, Node<T> destination);

    boolean connectionExists(Node<T> source, Node<T> destination);

    default Node<T> getNode(T content) {
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
