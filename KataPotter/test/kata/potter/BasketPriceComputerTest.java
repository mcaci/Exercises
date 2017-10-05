package kata.potter;

import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

/**
 * Unit test for simple App.
 */
public class BasketPriceComputerTest
{
  private BasketBuilder basketBuilder = new BasketBuilder();

  @Test
  public void testPriceOfEmptyBasket() {
    testPriceComputed(this.basketBuilder.build(0, 0, 0, 0, 0), 0F);
  }

  @Test
  public void testWithOneBook() {
    testPriceComputed(this.basketBuilder.build(1, 0, 0, 0, 0), 8F);
  }

  @Test
  public void testTwoSameBooks() {
    testPriceComputed(this.basketBuilder.build(0, 0, 2, 0, 0), 16F);
  }

  @Test
  public void testFiveSameBooks() {
    testPriceComputed(this.basketBuilder.build(0, 0, 0, 5, 0), 40F);
  }

  @Ignore
  @Test
  public void testTwoDifferentBooks() {
    testPriceComputed(this.basketBuilder.build(1, 1, 0, 0, 0), 16F * 0.95);
  }

  @Ignore
  @Test
  public void testThreeDifferentBooks() {
    testPriceComputed(this.basketBuilder.build(1, 1, 1, 0, 0), 24F * 0.9);
  }

  private void testPriceComputed(Basket basket, double expectedPrice) {
    BasketPriceComputer basketPriceComputer = new BasketPriceComputer();
    double price = basketPriceComputer.calculatePrice(basket);
    assertEquals(expectedPrice, price, 0.1F);
  }
}
