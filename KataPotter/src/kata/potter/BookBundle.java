package kata.potter;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class BookBundle {

    private final List<Book> bundle = new ArrayList<>();

    public BookBundle() {}

    public BookBundle(Book book) {
        this.getBundle().add(book);
    }

    public List<Book> getBundle() {
        return bundle;
    }

    public void add(Book book) {
        this.getBundle().add(book);
    }

    public void addAllBooksIn(BookBundle otherBundle) {
        this.getBundle().addAll(otherBundle.getBundle());
    }

    public boolean contains(BookBundle otherBundle) {
        boolean otherBundleIsContained = false;
        for (Book book : this.bundle) {
            otherBundleIsContained = otherBundle.getBundle().contains(book);
            if(otherBundleIsContained) {break;}
        }
        return otherBundleIsContained;
    }

    public int size() {
        return this.getBundle().size();
    }

    public double getDiscount() {
        return BookBundleDiscount.getDiscount(size());
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
