import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Scanner;

import java.util.List;

public class Part2 {

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

  private boolean preambleContainsSum(Long target) {

    // Oh no, runs in O^2 time
    for (Long i : this.currentPreamble) {
      for (Long j : this.currentPreamble) {
        if (i + j == target) {
          return true;
        }
      }
    }
    return false;
  }

  private Long[] findContiguousSet(Long target, String[] input) {

    long shortest = 0L;
    long longest = 0L;
    long currentSum = 0L;

      for (int i = 0; i < input.length; i++) {

        for (String j : Arrays.copyOfRange(input, i, input.length)) {

            if (Long.valueOf(j) < shortest || shortest == 0) {
                shortest = Long.valueOf(j);
            }

            if (Long.valueOf(j) > longest) {
                longest = Long.valueOf(j);
            }

            if (currentSum + Long.valueOf(j) == target && Long.valueOf(j) != target) {
                return new Long[] {shortest, longest};
            } else if (currentSum + Long.valueOf(j) > target) {
                shortest = 0;
                longest = 0;
                currentSum = 0;
                break;
            }

            currentSum += Long.valueOf(j);
            // System.out.println(currentSum);
        }

      }

      throw new Error("No continous set found in range");

  }

  private void moveToNext(Long num) {

    // Removes the first element of the list, and adds the specified one to the end
    this.currentPreamble.removeFirst();
    this.currentPreamble.addLast(num);
  }

  private LinkedList<Long> currentPreamble;

  public Part2(String[] preambleString) {

    this.currentPreamble = new LinkedList<>();

    for (String i : preambleString) {
      this.currentPreamble.add(Long.parseLong(i));
    }

  }

  public static void main(String[] args) {

    String[] input = readFile("input");
    
    Part2 part2 = new Part2(Arrays.copyOfRange(input, 0, 25));

    // Cut out the values that were already used
    String[] workingSet = Arrays.copyOfRange(input, 25, input.length);

    for (String i : workingSet) {

      if (!part2.preambleContainsSum(Long.parseLong(i))) {
        Long[] result =  part2.findContiguousSet(Long.valueOf(i), input);

        System.out.printf("%d + %d: %d\n", result[0], result[1], result[0] + result[1]);
      }

      part2.moveToNext(Long.parseLong(i));
    }

  }
}
