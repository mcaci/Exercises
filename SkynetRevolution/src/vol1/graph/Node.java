package vol1.graph;

/**
 * Created by mcaci on 8/2/17.
 */
public class Node {

    private final int id;
    private boolean gateway;

    public Node(int id) {
        this.id = id;
        this.gateway = false;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Node node = (Node) o;

        return id == node.id;
    }

    @Override
    public int hashCode() {
        return id;
    }

    public boolean isGateway() {
        return gateway;
    }

    public void setGateway(boolean gateway) {
        this.gateway = gateway;
    }
}
