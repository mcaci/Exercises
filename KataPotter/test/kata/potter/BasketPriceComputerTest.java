package kata.potter;

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

  @Test
  public void testOneBundleOfTwoBooks() {
    testPriceComputed(this.basketBuilder.build(1, 1, 0, 0, 0), 16F * 0.95);
  }

  @Test
  public void testOneBundleOfFiveBooks() {
    testPriceComputed(this.basketBuilder.build(1, 1, 1, 1, 1), 40F * 0.75);
  }

  private void testPriceComputed(Basket basket, double expectedPrice) {
    BasketPriceComputer basketPriceComputer = new BasketPriceComputer();
    double price = basketPriceComputer.calculateBasketPrice(basket);
    assertEquals(expectedPrice, price, 0.1F);
  }
}
