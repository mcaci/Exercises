package kata.potter;

import java.util.List;
import java.util.Optional;
import java.util.function.Predicate;

public class BundleAccumulator {

    public void addBundle(List<BookBundle> bookBundles, BookBundle bundleToAdd) {
        final BundleChooser bundleChooser = new BundleChooser(bundleToAdd);
        Optional<BookBundle> availableBundle = bundleChooser.getBookBundleWhereToAddBooks(bookBundles);
        if (availableBundle.isPresent()) {
            availableBundle.get().addAllBooksIn(bundleToAdd);
        } else {
            bookBundles.add(bundleToAdd);
        }
    }

    private class BundleChooser {

        private int bundleToAddSize;
        private Predicate<BookBundle> bookBundlePresencePredicate;

        public BundleChooser(BookBundle toAdd) {
            this.bundleToAddSize = toAdd.size();
            this.bookBundlePresencePredicate = bundle -> bundle.contains(toAdd);
        }

        public Optional<BookBundle> getBookBundleWhereToAddBooks(List<BookBundle> bookBundles) {
            Optional<BookBundle> bundleWhereToAddBooks = Optional.empty();

            for (BundleValueOrder bundleValueOrder : BundleValueOrder.values()) {
                Predicate<BookBundle> onSize = getBookBundlePredicate(bundleValueOrder);
                bundleWhereToAddBooks = bookBundles.stream().filter(this.bookBundlePresencePredicate.negate().and(onSize)).findFirst();
                if (bundleWhereToAddBooks.isPresent()) { break; }
            }
            return bundleWhereToAddBooks;
        }

        private Predicate<BookBundle> getBookBundlePredicate(BundleValueOrder bundleValueOrder) {
            return bookBundle -> bookBundle.size() + this.bundleToAddSize == bundleValueOrder.getBundleSize();
        }
    }

}
