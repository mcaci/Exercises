package kata.potter;

import java.util.Objects;

public class Book {
    private final int bookNumber;

    public Book(int bookNumber) {
        this.bookNumber = bookNumber;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Book book = (Book) o;
        return bookNumber == book.bookNumber;
    }

    @Override
    public int hashCode() {
        return Objects.hash(bookNumber);
    }

    @Override
    public String toString() {
        return "Book{" +
                "bookNumber=" + bookNumber +
                '}';
    }
}
