package kata.potter;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

public class Basket {

  List<Book> books = new ArrayList<>();

  public void add(int bookNumber) {
    this.books.add(new Book(bookNumber));
  }

  public int size() {
    return this.books.size();
  }

  public Book get(int index) {
    return books.get(index);
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
