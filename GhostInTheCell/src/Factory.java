/**
 * Created by nikiforos on 17/09/17.
 */
public class Factory {
    private String owner;
    private int cyborgQuantity;
    private int cyborgProduced;
    private String id;

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }

    public void setCyborgQuantity(int cyborgQuantity) {
        this.cyborgQuantity = cyborgQuantity;
    }

    public int getCyborgQuantity() {
        return cyborgQuantity;
    }

    public void setCyborgProduced(int cyborgProduced) {
        this.cyborgProduced = cyborgProduced;
    }

    public int getCyborgProduced() {
        return cyborgProduced;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    @Override
    public String toString() {
        return "Factory{" +
                "owner='" + owner + '\'' +
                ", cyborgQuantity=" + cyborgQuantity +
                ", cyborgProduced=" + cyborgProduced +
                ", id='" + id + '\'' +
                '}';
    }
}
