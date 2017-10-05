package kata.potter;

import java.util.ArrayList;
import java.util.List;

public class Basket {

  List<Book> books = new ArrayList<>();
  int[] booksArray = new int[5];

  public void add(int bookNumber) {
    this.books.add(new Book(bookNumber));
  }

  public int size() {
    return this.books.size();
  }
}
