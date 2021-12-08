import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class puzzle {
    public static void main(String[] args) throws FileNotFoundException {
        puzzle.puzzle1();
        puzzle.puzzle2();
    }

    public static void puzzle1() throws FileNotFoundException {
        File file = new File("day_8/input.txt");
        Scanner scanner = new Scanner(file);
        String line;
        int amt = 0;
        while (scanner.hasNext()) {
            line = scanner.nextLine();
            line = line.split("\\|")[1];
            for (String number : line.split(" ")) {
                if (number.length() == 0) continue;
                int l = number.length();
                if (l == 2 || l == 3 || l == 4 || l == 7) {
                    amt++;
                }
            }
        }
        System.out.println("Puzzle 1 : " + amt);
    }

    public static void puzzle2() {
        System.out.println("not yet implemented");
    }
}