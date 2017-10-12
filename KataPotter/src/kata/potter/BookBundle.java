package kata.potter;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Objects;

public class BookBundle {

    private final Collection<Book> bundle = new ArrayList<>();

    public BookBundle(Book book) {
        this.getBundle().add(book);
    }

    public void addAllBooksIn(BookBundle otherBundle) {
        this.getBundle().addAll(otherBundle.getBundle());
    }

    public boolean contains(BookBundle otherBundle) {
        return this.getBundle().stream().anyMatch(book -> otherBundle.getBundle().contains(book));
    }

    public int size() {
        return this.getBundle().size();
    }

    public double getDiscount() {
        return BookBundleDiscount.getDiscount(size());
    }

    private Collection<Book> getBundle() {
        return bundle;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        BookBundle bundle1 = (BookBundle) o;
        return Objects.equals(bundle, bundle1.bundle);
    }

    @Override
    public int hashCode() {
        return Objects.hash(bundle);
    }

    @Override
    public String toString() {
        return "BookBundle{" +
                "bundle=" + bundle +
                '}';
    }
}
