import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

/**
 * Created by mcaci on 9/8/17.
 */
public class FactoryMapTest {

    private static final String[] FACTORY_INPUT = {"1 FACTORY 1 12 1", "2 FACTORY 1 13 2", "3 FACTORY -1 8 3", "4 FACTORY 0 10 2"};

    @Test
    public void testFactoryDataOwnerForFirstFactory() throws Exception {
        FactoryBuilder factoryBuilder = new FactoryBuilder();
        factoryBuilder.addData(FACTORY_INPUT[0]);
        Factory factory = factoryBuilder.build();
        assertEquals("1", factory.getOwner());
    }

    @Test
    public void testFactoryDataOwnerForThirdFactory() throws Exception {
        FactoryBuilder factoryBuilder = new FactoryBuilder();
        factoryBuilder.addData(FACTORY_INPUT[2]);
        Factory factory = factoryBuilder.build();
        assertEquals("-1", factory.getOwner());
    }
    @Test
    public void testFactoryDataCyborgQuantityForSecondFactory() throws Exception {
        FactoryBuilder factoryBuilder = new FactoryBuilder();
        factoryBuilder.addData(FACTORY_INPUT[1]);
        Factory factory = factoryBuilder.build();
        assertEquals(13, factory.getCyborgQuantity());
    }
    @Test
    public void testFactoryDataCyborgProductionForFourthFactory() throws Exception {
        FactoryBuilder factoryBuilder = new FactoryBuilder();
        factoryBuilder.addData(FACTORY_INPUT[3]);
        Factory factory = factoryBuilder.build();
        assertEquals(2, factory.getCyborgProduced());
    }

}
