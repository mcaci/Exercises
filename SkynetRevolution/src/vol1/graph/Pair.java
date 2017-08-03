package vol1.graph;

/**
 * Created by mcaci on 8/2/17.
 */
public class Pair<T> {
    private final T one;
    private final T two;

    Pair(T source, T destination) {
        this.one = source;
        this.two = destination;
    }

    public T getOne() {
        return one;
    }

    public T getTwo() {
        return two;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Pair<?> pair = (Pair<?>) o;

        if (one != null ? !one.equals(pair.one) : pair.one != null) return false;
        return two != null ? two.equals(pair.two) : pair.two == null;
    }

    @Override
    public int hashCode() {
        int result = one != null ? one.hashCode() : 0;
        result = 31 * result + (two != null ? two.hashCode() : 0);
        return result;
    }
}
