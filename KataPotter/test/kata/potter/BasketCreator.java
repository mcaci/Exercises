package kata.potter;

public class BasketCreator {

  public Basket create(int... quantityOfBooks) {
    Basket basket = new Basket();
    for (int i = 0; i < quantityOfBooks.length; i++) {
      int bookId = i + 1;
      int bookQuantity = quantityOfBooks[i];
      for (int j = 0; j < bookQuantity; j++) {
        basket.add(bookId);
      }
    }
    return basket;
  }

}
