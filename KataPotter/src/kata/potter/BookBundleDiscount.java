package kata.potter;

public enum BookBundleDiscount {
  NO_BUNDLE(1), TWO_BUNDLE_BOOKS(0.95), THREE_BUNDLE_BOOKS(0.9), FOUR_BOOK_BUNDLE(0.8), FULL_BUNDLE(0.75);

  private final double discount;

  BookBundleDiscount(double discount) {
    this.discount = discount;
  }

  public double getDiscount() {
    return discount;
  }
}
