package kata.potter;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class BasketPriceComputer {

  private static final double PRICE_PER_UNIT = 8;

  public double calculateBasketPrice(Basket basket) {
    List<BookBundle> bundles = basket.stream().map(book -> new BookBundle(book)).
            collect(ArrayList::new, this::mergeBundle, ArrayList::addAll);
    System.out.println(bundles);
    return bundles.stream().mapToDouble(this::calculateBundlePrice).sum();
  }

  private double calculateBundlePrice(BookBundle bundle) {
    return bundle.size() * PRICE_PER_UNIT * bundle.getDiscount();
  }

  private void mergeBundle(List<BookBundle> bookBundles, BookBundle bundle) {

    Map<Integer, List<BookBundle>> bundlesBySize = bookBundles.stream().collect(Collectors.groupingBy(BookBundle::size));
    boolean isBookBundlesEmpty = isBookBundlesEmpty(bundlesBySize);
    if (isBookBundlesEmpty) {
      bookBundles.add(bundle);
    }
    else {
      boolean hasMerged = false;
      for (Integer n : bundlesBySize.keySet()) {
        List<BookBundle> bundlesWithSizeN = bundlesBySize.get(n);
        for (BookBundle currentBundle : bundlesWithSizeN) {
          if (!currentBundle.contains(bundle)) {
            currentBundle.addAllBooksIn(bundle);
            hasMerged = true;
            break;
          }
        }
        if(hasMerged) {break;}
      }
      if (!hasMerged) {
        bookBundles.add(bundle);
      }
    }

//    boolean hasMerged = false;
//    for (BookBundle currentBundle : bookBundles) {
//      if (!currentBundle.contains(bundle)) {
//        currentBundle.addAllBooksIn(bundle);
//        hasMerged = true;
//        break;
//      }
//    }
//      if(!hasMerged) {
//        bookBundles.add(bundle);
//      }
//  }
  }

  private boolean isBookBundlesEmpty(Map<Integer, List<BookBundle>> bundlesBySize) {
    boolean isBookBundlesEmpty = bundlesBySize.isEmpty();
    if(!isBookBundlesEmpty) {
      isBookBundlesEmpty = bundlesBySize.keySet().size() == 1 && bundlesBySize.keySet().contains(0);
    }
    return isBookBundlesEmpty;
  }
}
