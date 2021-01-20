import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Scanner;
import java.util.List;

public class Part1 {

  private static String[] readFile(String fileName) {

    List<String> data = new ArrayList<>();

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

  private static int[] stringArrayToIntArray(String[] strArray) {
    int[] result = new int[strArray.length];
    for (int i = 0; i < strArray.length; i++) {
      result[i] = Integer.parseInt(strArray[i]);
    }
    return result;
  }

  public static void main(String[] args) {

    int[] input = stringArrayToIntArray(readFile("input"));

    Arrays.sort(input);

    HashMap<Integer, Integer> differences = new HashMap<>();

    int prev = 0;

    for (int i : input) { // Create adapter chain

      int curDiff = i - prev;

      if (differences.containsKey(curDiff)) {
        differences.replace(curDiff, differences.get(curDiff) + 1);
      } else {
        differences.put(curDiff, 1);
      }

      prev = i;
    }

    // Account for the charger being three jolts higher than the highest adapter
    differences.replace(3, differences.get(3) + 1);

    System.out.println(differences);

    System.out.println(differences.get(1) * differences.get(3));

  }
}
