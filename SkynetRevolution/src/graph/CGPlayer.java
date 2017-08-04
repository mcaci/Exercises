package graph;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.LinkedList;
import java.util.List;
import java.util.Scanner;

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
public class CGPlayer {

    public static void main(String args[]) throws FileNotFoundException {
        Graph<CGNetworkHub> skynetNetwork = new Graph<>();

        Scanner in = new Scanner(System.in);
        int N = in.nextInt(); // the total number of nodes in the level, including the gateways
        int L = in.nextInt(); // the number of links
        int E = in.nextInt(); // the number of exit gateways
        for (int i = 0; i < L; i++) {
            int N1 = in.nextInt(); // N1 and N2 defines a link between these nodes
            int N2 = in.nextInt();
            skynetNetwork.add(new CGNetworkHub(N1), new CGNetworkHub(N2), true);
        }
        for (int i = 0; i < E; i++) {
            int EI = in.nextInt(); // the index of a gateway node
            for (Node<CGNetworkHub> networkHubNode : skynetNetwork.getNodes()) {
                CGNetworkHub networkHub = networkHubNode.getContent();
                if(networkHub.index == EI) {
                    networkHub.setAsGateway();
                    break;
                }
            }
        }

        AdjacencyListNode<CGNetworkHub> currentNode = null;
        // game loop
        while (true) {
            int SI = in.nextInt(); // The index of the node on which the Skynet agent is positioned this turn
            currentNode = (AdjacencyListNode<CGNetworkHub>) skynetNetwork.getNode(new CGNetworkHub(SI));
            System.err.println("Current node: " + currentNode.getContent().index);
            // Write an action using System.out.println()
            // To debug: System.err.println("Debug messages...");

            Node<CGNetworkHub> nodeToCut = null;
            for (Node<CGNetworkHub> networkHubNode : currentNode.getAdjacentNodes()) {
                if (networkHubNode.getContent().isGateway) {
                    nodeToCut = networkHubNode;
                    break;
                }
            }
            if(nodeToCut == null) {
                nodeToCut = currentNode.getAdjacentNodes().get(0);
            }
            skynetNetwork.removeEdge(currentNode, nodeToCut);

            // Example: 0 1 are the indices of the nodes you wish to sever the link between
            System.out.println(currentNode.getContent().index + " " + nodeToCut.getContent().index);

        }
    }
}

class Graph<T> {

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
        AdjacencyListNode<T> sourceNode = (AdjacencyListNode<T>) this.getNode(source.getContent());
        AdjacencyListNode<T> destinationNode = (AdjacencyListNode<T>) this.getNode(destination.getContent());
        sourceNode.getAdjacentNodes().add(destinationNode);
        if (isBidirectional) {
            destinationNode.getAdjacentNodes().add(sourceNode);
        }
    }

    public void add(T sourceContent, T destinationContent, boolean isBidirectional) {
        add(new AdjacencyListNode<>(sourceContent), new AdjacencyListNode<>(destinationContent), isBidirectional);
    }


    public void removeEdge(AdjacencyListNode<T> source, Node<T> destination) {
        source.getAdjacentNodes().removeIf(e -> e.equals(destination));
    }

    public List<Node<T>> getNodes() {
        return nodes;
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

class Node<T> {

    private final T content;

    public Node(T content) {
        this.content = content;
    }

    public T getContent() {
        return content;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Node<?> node = (Node<?>) o;

        return content != null ? content.equals(node.content) : node.content == null;
    }

    @Override
    public int hashCode() {
        return content != null ? content.hashCode() : 0;
    }
}

class AdjacencyListNode<T> extends Node<T> {

    private final List<Node<T>> adjacentNodes;

    public AdjacencyListNode(T t) {
        super(t);
        adjacentNodes = new LinkedList<>();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        return super.equals(o);
    }

    @Override
    public int hashCode() {
        int result = super.hashCode();
        result = 31 * result + (adjacentNodes != null ? adjacentNodes.hashCode() : 0);
        return result;
    }

    public List<Node<T>> getAdjacentNodes() {
        return adjacentNodes;
    }
}

class CGNetworkHub {

    final int index;
    boolean isGateway;

    public CGNetworkHub(int index) {
        this.index = index;
        this.isGateway = false;
    }

    public void setAsGateway() {
        isGateway = true;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        CGNetworkHub that = (CGNetworkHub) o;

        if (index != that.index) return false;
        return isGateway == that.isGateway;
    }

    @Override
    public int hashCode() {
        int result = index;
        result = 31 * result + (isGateway ? 1 : 0);
        return result;
    }
}
