package kata.potter;

import java.util.ArrayList;
import java.util.List;

public class BasketPriceComputer {

  private static final double PRICE_PER_UNIT = 8;

  public double calculateBasketPrice(Basket basket) {
    List<BookBundle> bundles = basket.stream().map(book -> new BookBundle(book)).
            collect(ArrayList::new, this::mergeBundle, ArrayList::addAll);
    return bundles.stream().mapToDouble(this::calculateBundlePrice).sum();
  }

  private double calculateBundlePrice(BookBundle bundle) {
    return bundle.size() * PRICE_PER_UNIT * bundle.getDiscount();
  }

  private void mergeBundle(List<BookBundle> bookBundles, BookBundle bundle) {
      boolean hasMerged = merge(bookBundles, bundle);
      if(!hasMerged) {
        bookBundles.add(bundle);
      }
  }

  private boolean merge(List<BookBundle> bookBundles, BookBundle bundle) {
    boolean added = false;
    for (BookBundle currentBundle : bookBundles) {
      if (!currentBundle.contains(bundle)) {
        currentBundle.addAllBooksIn(bundle);
        added = true;
        break;
      }
    }
    return added;
  }

}
