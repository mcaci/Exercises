package graph;

/**
 * Created by mcaci on 8/3/17.
 */
public class NetworkNode {

    final int index;
    boolean isGateway;

    public NetworkNode(int index) {
        this.index = index;
        this.isGateway = false;
    }

    public void setAsGateway() {
        isGateway = true;
    }
}
