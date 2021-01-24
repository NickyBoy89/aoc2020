import java.util.ArrayList;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.util.Arrays;
import java.util.Iterator;

public class Helpers {
    public static String[] readFile(String fileName) {
      ArrayList<String> data = new ArrayList<>();

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

      data.add("0");

      data.add(String.valueOf(Integer.valueOf(findMax(data)) + 3));

      String[] result = new String[data.size()];

      return data.toArray(result);
    }

    public static int[] stringArrayToIntArray(String[] strArray) {

        int[] result = new int[strArray.length];
        for (int i = 0; i < strArray.length; i++) {
            result[i] = Integer.parseInt(strArray[i]);
        }
        return result;
    }

    public static int[] stringArrayToIntArray(String[] strArray, boolean sorted) {
      int[] data = stringArrayToIntArray(strArray);

      if (sorted) {
        Arrays.sort(data);
      }

      return data;
    }

    public static int sum(int[] nums) {

      int counter = 0;

      for (int num : nums) {
          counter += num;
      }

      return counter;
    }

    public static int findMax(ArrayList<String> arr) {
      int max = 0;

      for (String cur : arr) {

        int curNum = Integer.parseInt(cur);

        if (curNum > max) {
          max = curNum;
        }
      }

      return max;
    }

    public static int findMax(int[] arr) {

      int max = 0;

      for (int i : arr) {
        if (i > max) {
          max = i;
        }
      }

      return max;
    }
}
