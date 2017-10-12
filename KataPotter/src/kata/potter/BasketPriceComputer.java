package kata.potter;

import java.util.ArrayList;
import java.util.List;
import java.util.function.ToDoubleFunction;

public class BasketPriceComputer {

    private static final double PRICE_PER_UNIT = 8;
    private static final ToDoubleFunction<BookBundle> BUNDLE_PRICE_FUNCTION = bundle -> PRICE_PER_UNIT * bundle.size() * bundle.getDiscount();

    public double calculateBasketPrice(Basket basket) {
        BundleAccumulator accumulator = new BundleAccumulator();
        List<BookBundle> bundles = basket.stream().map(BookBundle::new).
                collect(ArrayList::new, accumulator::bundleAccumulator, ArrayList::addAll);
        return bundles.stream().mapToDouble(BUNDLE_PRICE_FUNCTION).sum();
    }

}
