package kata.potter;

public class BasketPriceComputer {

  private static final double PRICE_PER_UNIT = 8;

  public double calculatePrice(Basket basket) {
    return basket.size() * PRICE_PER_UNIT;
  }

}
