package kata.potter;

import java.util.ArrayList;
import java.util.List;

public class BasketPriceComputer {

  public double calculatePrice(Basket basket) {
    int numberOfDifferentBooks = basket.getNumberOfDifferentBooks();
    float price = 8F;
    if (numberOfDifferentBooks > 1) {
      float discount = getDiscount(numberOfDifferentBooks);
      return numberOfDifferentBooks * price * discount;
    }
    return basket.getSize() * price;
  }

  private float getDiscount(int numberOfDifferentBooks) {
    float discount = 1;
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

  public float bruteForce(Basket basket) {
    List<Float> possiblePrices = new ArrayList();
    float price = 0;
    for (int i = basket.getNumberOfDifferentBooks(); i >= 1; i--) {
      Basket copyBasket = basket.copy();
      for (int j = 0; j < 5; j++) {
        if (copyBasket.books[j] > 0) {
          copyBasket.books[j]--;
        }
      }
      price = 8 * i * getDiscount(i);

    }
    return getMinimumPrice(possiblePrices);
  }

  private float getMinimumPrice(List<Float> possiblePrices) {
    // TODO Auto-generated method stub
    return 0;
  }

}
