package kata.potter;

/**
 * Created by nikiforos on 12/10/17.
 */
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
