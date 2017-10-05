package kata.potter;

public class BasketPriceComputer {

  public double calculatePrice(Basket basket) {
    int numberOfDifferentBooks = basket.getNumberOfDifferentBooks();
    float price = 8F;
    if (numberOfDifferentBooks > 1) {
      double discount = getDiscount(numberOfDifferentBooks);
      return numberOfDifferentBooks * price * discount;
    }
    return basket.getSize() * price;
  }

  private double getDiscount(int numberOfDifferentBooks) {
    double discount = 1;
    if (numberOfDifferentBooks == 2) {
      discount = 0.95F;
    }
    else if (numberOfDifferentBooks == 3) {
      discount = 0.9F;
    }
    else if (numberOfDifferentBooks == 4) {
      discount = 0.8F;
    }
    else if (numberOfDifferentBooks == 5) {
      discount = 0.75F;
    }
    return discount;
  }

}
