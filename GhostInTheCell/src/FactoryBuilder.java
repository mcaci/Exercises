/**
 * Created by nikiforos on 17/09/17.
 */
public class FactoryBuilder {
    private final Factory toBuild = new Factory();

    public void addData(String data) {
        String[] splitData = data.split(" ");
        this.addId(splitData[0]).addOwner(splitData[2]).addCyborgQuantity(splitData[3]).addCyborgProduced(splitData[4]);
    }

    private FactoryBuilder addId(String id) {
        toBuild.setId(id);
        return this;
    }

    private FactoryBuilder addOwner(String owner) {
        toBuild.setOwner(owner);
        return this;
    }

    private FactoryBuilder addCyborgQuantity(String cyborgQuantity) {
        toBuild.setCyborgQuantity(Integer.parseInt(cyborgQuantity));
        return this;
    }

    private FactoryBuilder addCyborgProduced(String cyborgProduced) {
        toBuild.setCyborgProduced(Integer.parseInt(cyborgProduced));
        return this;
    }

    public Factory build() {
        return toBuild;
    }
}
