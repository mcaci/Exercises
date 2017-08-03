package graph;

import java.util.LinkedList;
import java.util.List;

/**
 * Created by mcaci on 8/2/17.
 */
public class Graph {


    private List<Edge> edges;
    private List<Node> nodes;

    public Graph() {
        edges = new LinkedList<>();
        nodes = new LinkedList<>();
    }

    public void add(Edge edge) {
        edges.add(edge);
        edges.add(edge.reverse());
        addNode(edge.getOne());
        addNode(edge.getTwo());
    }

    private void addNode(Node node) {
        if (!nodes.contains(node)) {
            nodes.add(node);
        }
    }

    public List<Edge> getEdges() {
        return edges;
    }

    public void removeEdge(Edge edge) {
        edges.removeIf(e -> e.equals(edge));
    }

    public List<Node> getNodes() {
        return nodes;
    }

    public boolean connectionExistsBetween(Node one, Node two) {
        return edges.contains(new Edge(one, two));
    }
}
