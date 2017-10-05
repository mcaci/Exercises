package kata.potter;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

/**
 * Unit test for simple App.
 */
public class BasketPriceComputerTest
{

  Basket basket = new Basket();
  BasketPriceComputer basketPriceComputer = new BasketPriceComputer();

  @Test
  public void testPriceOfEmptyBasket() {
    testPriceComputed(0F);
  }

  @Test
  public void testWithOneBook() {
    addBooksToBasket(1, 0, 0, 0, 0);
    testPriceComputed(8F);
  }

  @Test
  public void testTwoSameBooks() {
    addBooksToBasket(2, 0, 0, 0, 0);
    testPriceComputed(16F);
  }

  @Test
  public void testTwoDifferentBooks() {
    addBooksToBasket(1, 1, 0, 0, 0);
    testPriceComputed(16F * 0.95);
  }

  @Test
  public void testThreeDifferentBooks() {
    addBooksToBasket(1, 1, 1, 0, 0);
    testPriceComputed(24F * 0.9);
  }

  @Test
  public void testThreeSameBooks() {
    addBooksToBasket(3, 0, 0, 0, 0);
    testPriceComputed(24F);
  }

  private void addBooksToBasket(int quantityOfBook1, int quantityOfBook2, int quantityOfBook3, int quantityOfBook4,
      int quantityOfBook5) {
    addBooksToBasket(1, quantityOfBook1);
    addBooksToBasket(2, quantityOfBook2);
    addBooksToBasket(3, quantityOfBook3);
    addBooksToBasket(4, quantityOfBook4);
    addBooksToBasket(5, quantityOfBook5);
  }

  private void addBooksToBasket(int bookId, int bookQuantity) {
    for (int i = 0; i < bookQuantity; i++) {
      basket.add(bookId);
    }
  }

  private void testPriceComputed(double expectedPrice) {
    double price = basketPriceComputer.calculatePrice(basket);
    assertEquals(expectedPrice, price, 0.1F);
  }

}
