package kata.potter;

import java.util.ArrayList;
import java.util.Collection;
import java.util.stream.Stream;

public class Basket {

  Collection<Book> books = new ArrayList<>();

  public void add(int bookNumber) {
    this.books.add(new Book(bookNumber));
  }

  public Stream<Book> stream() {
    return books.stream();
  }

  @Override
  public String toString() {
    return "Basket{" +
            "books=" + books +
            '}';
  }
}
