package kata.potter;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

/**
 * Unit test for simple App.
 */
public class BasketPriceComputerTest
{
  private BasketCreator basketBuilder = new BasketCreator();

  @Test
  public void testPriceOfEmptyBasket() {
    testPriceComputed(this.basketBuilder.create(0, 0, 0, 0, 0), 0);
  }

  @Test
  public void testWithOneBook() {
    testPriceComputed(this.basketBuilder.create(1, 0, 0, 0, 0), 8);
  }

  @Test
  public void testTwoSameBooks() {
    testPriceComputed(this.basketBuilder.create(0, 0, 2, 0, 0), 16);
  }

  @Test
  public void testFiveSameBooks() {
    testPriceComputed(this.basketBuilder.create(0, 0, 0, 5, 0), 40);
  }

  @Test
  public void testOneBundleOfTwoBooks() {
    testPriceComputed(this.basketBuilder.create(1, 1, 0, 0, 0), 16 * 0.95);
  }

  @Test
  public void testOneBundleOfFiveBooks() {
    testPriceComputed(this.basketBuilder.create(1, 1, 1, 1, 1), 40 * 0.75);
  }

  @Test
  public void testTwoBundlesOfOneAndFiveBooks() {
      testPriceComputed(this.basketBuilder.create(1, 1, 1, 2, 1), 40 * 0.75 + 8);
  }

  @Test
  public void testTwoBundlesOfThreeAndFiveBooks() {
      testPriceComputed(this.basketBuilder.create(1, 2, 2, 2, 1), 2 * 32 * 0.8);
  }

  @Test
  public void testThreeFullBundlesAndTwoBundlesOfFour() {
    testPriceComputed(this.basketBuilder.create(5, 5, 4, 5, 4), (3 * (8 * 5 * 0.75) + 2 * (8 * 4 * 0.8)));
  }


  private void testPriceComputed(Basket basket, double expectedPrice) {
    BasketPriceComputer basketPriceComputer = new BasketPriceComputer();
    double price = basketPriceComputer.calculateBasketPrice(basket);
    assertEquals(expectedPrice, price, 0.1);
  }
}
