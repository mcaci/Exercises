import org.junit.Test;

import java.util.*;
import java.util.stream.Collectors;

import static org.junit.Assert.assertNotNull;

/**
 * Created by mcaci on 9/8/17.
 */
public class Strategy1Test {

    private static final String[] DISTANCE_INPUT = {"1 2 20", "1 3 15", "2 3 18"};
    private static final String[] FACTORY_INPUT = {"1 FACTORY 1 12 1", "2 FACTORY 1 13 2", "3 FACTORY -1 8 3", "4 FACTORY 0 10 2"};

    @Test
    public void testFactoryDataOwnerForFirstFactory() throws Exception {
        final Map<Integer, Map<Integer, Integer>> distanceMap = mapDistances();
        List<Factory> factories = listFactories();

        final Optional<Factory> firstMyFactory = factories.stream().parallel().filter(factory -> "1".equalsIgnoreCase(factory.getOwner())).filter(factory -> 10 < factory.getCyborgQuantity()).max((f1, f2) -> f1.getCyborgQuantity() - f2.getCyborgQuantity());
        int factoryCount = 100;
        final Optional<Factory> firstOtherFactory = factories.stream().parallel().filter(factory -> !"1".equalsIgnoreCase(factory.getOwner())).filter(factory -> factory.getCyborgProduced() - (Math.round(Math.random())) > 0).findAny();

        final String separator = " ";
        String command = "MOVE" + separator + firstMyFactory.get().getId() + separator +
                firstOtherFactory.get().getId() + separator + "5";
        assertNotNull(command);


    }

    private Map<Integer, Map<Integer, Integer>> mapDistances() {
        DistancesMapBuilder distancesMapBuilder = new DistancesMapBuilder();
        Arrays.stream(DISTANCE_INPUT).forEach(distancesMapBuilder::addDistancePair);
        return distancesMapBuilder.build();
    }

    private List<Factory> listFactories() {
        List<Factory> factories = Arrays.stream(FACTORY_INPUT).map(input -> {
            FactoryBuilder factoryBuilder = new FactoryBuilder();
            factoryBuilder.addData(input);
            return factoryBuilder.build();
        }).collect(Collectors.toList());
        return factories;
    }

}
