import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.Scanner;
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
public class Player {

    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int factoryCount = in.nextInt(); // the number of factories
        int linkCount = in.nextInt(); // the number of links between factories
        for (int i = 0; i < linkCount; i++) {
            int factory1 = in.nextInt();
            int factory2 = in.nextInt();
            int distance = in.nextInt();
        }

        // game loop
        while (true) {
            List<Factory> factories = new ArrayList<>();

            int entityCount = in.nextInt(); // the number of entities (e.g. factories and troops)
            for (int i = 0; i < entityCount; i++) {
                int entityId = in.nextInt();
                String entityType = in.next();
                int arg1 = in.nextInt();
                int arg2 = in.nextInt();
                int arg3 = in.nextInt();
                int arg4 = in.nextInt();
                int arg5 = in.nextInt();
                if(entityType.equalsIgnoreCase("FACTORY")) {
                    StringBuilder sb = new StringBuilder();
                    sb.
                            append(entityId).append(" ").
                            append(entityType).append(" ").
                            append(arg1).append(" ").
                            append(arg2).append(" ").
                            append(arg3).append(" ").
                            append(arg4).append(" ").
                            append(arg5);
                    FactoryBuilder fb = new FactoryBuilder();
                    fb.addData(sb.toString());
                    factories.add(fb.build());
                }
            }

            // Write an action using System.out.println()
            // To debug: System.err.println("Debug messages...");
            final Optional<Factory> firstMyFactory = factories.stream().parallel().filter(factory -> "1".equalsIgnoreCase(factory.getOwner())).filter(factory -> 10 < factory.getCyborgQuantity()).max((f1, f2) -> f1.getCyborgQuantity() - f2.getCyborgQuantity());
            final Optional<Factory> firstOtherFactory = factories.stream().parallel().filter(factory -> !"1".equalsIgnoreCase(factory.getOwner())).filter(factory -> factory.getCyborgProduced() > 0).findAny();

            // for(int i = 0; i < factoryCount; i++) {
            //     System.err.println(factories.get(i));
            // }

            // System.err.println(firstMyFactory.get().getOwner());
            // System.err.println(firstMyFactory.get().getId());

            final String separator = " ";
            String command = "WAIT";
            if(firstMyFactory.isPresent() && firstOtherFactory.isPresent()) {
                command = "MOVE" + separator + firstMyFactory.get().getId() + separator +
                        firstOtherFactory.get().getId() + separator + firstMyFactory.get().getCyborgQuantity() / 3;
            }

            // Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
            System.out.println(command);
        }
    }
}