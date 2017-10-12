package kata.potter;

import java.util.List;
import java.util.Optional;
import java.util.function.Predicate;

public class BundleAccumulator {

    public void bundleAccumulator(List<BookBundle> bookBundles, BookBundle bundle) {
        Optional<BookBundle> availableBundle = getBookBundleWhereToAddBooks(bookBundles, bundle);
        if (availableBundle.isPresent()) {
            availableBundle.get().addAllBooksIn(bundle);
        } else {
            bookBundles.add(bundle);
        }
    }

    private Optional<BookBundle> getBookBundleWhereToAddBooks(List<BookBundle> bookBundles, BookBundle bundle) {
        int bundleSize = bundle.size();

        Optional<BookBundle> result = Optional.empty();
        for (BundleValueOrder bundleValueOrder : BundleValueOrder.values()) {
            int targetSize = bundleValueOrder.getBundleSize();
            Predicate<BookBundle> onPresence = currentBundle -> currentBundle.contains(bundle);
            Predicate<BookBundle> onSize = bookBundle -> bookBundle.size() + bundleSize == targetSize;
            result = bookBundles.stream().filter(onPresence.negate().and(onSize)).findFirst();
            if (result.isPresent()) {
                break;
            }
        }
        return result;
    }

    public enum BundleValueOrder {
        BUNDLE_OF_4(4), BUNDLE_OF_5(5), BUNDLE_OF_3(3), BUNDLE_OF_2(2);

        private final int bundleSize;

        BundleValueOrder(int bundleSize) {
            this.bundleSize = bundleSize;
        }

        public int getBundleSize() {
            return bundleSize;
        }
    }
}
