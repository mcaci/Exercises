import java.util.HashMap;
import java.util.Map;

/**
 * Created by nikiforos on 17/09/17.
 */
public class DistancesMapBuilder {
    private final Map<Integer, Map<Integer, Integer>> originDestinationDistanceMap = new HashMap<>();

    public Map<Integer, Map<Integer, Integer>> build() {
        return new HashMap<>(originDestinationDistanceMap);
    }

    public void addDistancePair(String distancePair) {
        String[] values = distancePair.split(" ");
        Map<Integer, Integer> valueMap = new HashMap<>();
        valueMap.put(Integer.parseInt(values[1]), Integer.parseInt(values[2]));
        originDestinationDistanceMap.merge(Integer.parseInt(values[0]), valueMap, (prev, next) -> {
            prev.putAll(next);
            return prev;
        });
    }
}
