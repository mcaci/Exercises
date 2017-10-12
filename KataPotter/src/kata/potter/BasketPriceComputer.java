package kata.potter;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.function.Predicate;
import java.util.stream.Stream;

public class BasketPriceComputer {

  private static final double PRICE_PER_UNIT = 8;

    public double calculateBasketPrice(Basket basket) {
        List<BookBundle> bundles = basket.stream().map(BookBundle::new).
                collect(ArrayList::new, this::bundleAccumulator, ArrayList::addAll);
        System.out.println(bundles);
        return bundles.stream().mapToDouble(this::calculateBundlePrice).sum();
    }

    private double calculateBundlePrice(BookBundle bundle) {
      return bundle.size() * PRICE_PER_UNIT * bundle.getDiscount();
    }

    private void bundleAccumulator(List<BookBundle> bookBundles, BookBundle bundle) {
        Optional<BookBundle> availableBundle = getBookBundleWhereToAddBooks(bookBundles, bundle);
        if(availableBundle.isPresent()) {
            availableBundle.get().addAllBooksIn(bundle);
        }
        else {
            addNewBundleToList(bookBundles, bundle);
        }
    }

    private Optional<BookBundle> getBookBundleWhereToAddBooks(List<BookBundle> bookBundles, BookBundle bundle) {
        int bundleSize = bundle.size();

        Optional<BookBundle> result = Optional.empty();
        int[] order = {4, 5, 3, 2};
        for (int index : order) {
            Stream<BookBundle> bundles = getBookBundleStream(bookBundles, bundle);
            result = bundles.filter(getBookBundlePredicate(bundleSize, index)).findFirst();
            if(result.isPresent()) { break; }
        }
        return result;
    }

    private Predicate<BookBundle> getBookBundlePredicate(int bundleSize, int targetSize) {
        return bookBundle -> bookBundle.size() + bundleSize == targetSize;
    }

    private Stream<BookBundle> getBookBundleStream(List<BookBundle> bookBundles, BookBundle bundle) {
        Predicate<BookBundle> containsPredicate = currentBundle -> currentBundle.contains(bundle);
        return bookBundles.stream().filter(containsPredicate.negate());
    }

    private void addNewBundleToList(List<BookBundle> bookBundles, BookBundle bundle) {
        bookBundles.add(bundle);
    }

}
