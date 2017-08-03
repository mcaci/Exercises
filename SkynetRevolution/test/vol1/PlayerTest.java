package vol1;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.*;

/**
 * Created by mcaci on 8/3/17.
 */
public class PlayerTest {

    private Player player;

    @Before
    public void setUp() throws Exception {
        player = new Player();
    }

    @Test
    void testCreateGraphFromInput() {
        player.createGraph();
    }

}