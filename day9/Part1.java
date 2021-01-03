import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Scanner;
import java.util.List;

public class Part1 {

  private static String[] readFile(String fileName) {

    List<String> data = new ArrayList<String>();

    File parsedFile = new File(fileName);

    try {
      Scanner fileScanner = new Scanner(parsedFile);

      while (fileScanner.hasNextLine()) {

        data.add(fileScanner.nextLine());

      }

      fileScanner.close();

    } catch (FileNotFoundException err) {
      System.out.println(err);
      err.getStackTrace();
    }

    String[] result = new String[data.size()];

    return data.toArray(result);

  }

  private boolean preambleContainsSum(int target) {

    // Oh no, runs in O^2 time
    for (Integer i : this.currentPreamble) {
      for (Integer j : this.currentPreamble) {
        if (i + j == target) {
          return true;
        }
      }
    }
    return false;
  }

  private void moveToNext(int num) {
    this.currentPreamble.removeFirst();
    this.currentPreamble.addLast(num);
  }

  private LinkedList<Integer> currentPreamble = new LinkedList<>();

  public Part1(String[] preambleString) {

    for (String i : preambleString) {
      this.currentPreamble.add(Integer.parseInt(i));
    }

  }

  public static void main(String[] args) {

    String[] input = readFile("input");
    
    Part1 part1 = new Part1(Arrays.copyOfRange(input, 0, 25));

    // Cut out the values that were already used
    input = Arrays.copyOfRange(input, 25, input.length);

    for (String i : input) {

      // Note: Literally errors out when trying to parse integers > 32bit limit
      if (!part1.preambleContainsSum(Integer.parseInt(i))) {
        System.out.println(i);
      }

      part1.moveToNext(Integer.parseInt(i));
    }

  }
}
