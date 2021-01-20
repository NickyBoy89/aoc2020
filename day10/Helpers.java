import java.util.ArrayList;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

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

    public static int sum(int[] nums) {

      int counter = 0;

      for (int num : nums) {
          counter += num;
      }

      return counter;
    }
}