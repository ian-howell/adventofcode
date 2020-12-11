import java.util.*;

public class main {
    public static void main(String[] args) {
        int totalYeses = 0;
        Scanner in = new Scanner(System.in);
        ArrayList<String> group = getGroup(in);
        while (group.size() > 0) {
            totalYeses += countYeses(group);
            group = getGroup(in);
        }
        System.out.println(totalYeses);
    }

    public static ArrayList<String> getGroup(Scanner in) {
        ArrayList<String> people = new ArrayList<>();
        try {
            String line = in.nextLine();
            while (line.length() > 0) {
                people.add(line);
                line = in.nextLine();
            }
        } catch (NoSuchElementException e) {
            // NoSuchElementException is expected at end of input
        }
        return people;
    }

    public static int countYeses(ArrayList<String> group) {
        Set<Character> yeses = new HashSet<>();
        for (String person : group) {
            for (int i = 0; i < person.length(); i++) {
                yeses.add(person.charAt(i));
            }
        }
        return yeses.size();
    }

}
