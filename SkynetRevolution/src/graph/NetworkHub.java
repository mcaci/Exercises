package graph;

/**
 * Created by mcaci on 8/3/17.
 */
public class NetworkHub {

    final int index;
    boolean isGateway;

    public NetworkHub(int index) {
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

        NetworkHub that = (NetworkHub) o;

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
