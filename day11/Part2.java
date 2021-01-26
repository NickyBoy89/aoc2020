import java.util.ArrayList;
import java.util.Arrays;

class Part2 {

  private static int numCanSee(int rowIndex, int colIndex, char[][] input) {

    int count = 0;

    for (int check = rowIndex - 1; check < rowIndex + 2; check++) {
      for (int seat = colIndex - 1; seat < colIndex + 2; seat++) {

        int curRow = check;
        int curCol = seat;

        if (check < 0 || seat < 0 || check >= input.length || seat >= input[check].length || (check == rowIndex && seat == colIndex)) {
          continue;
        }

        while(curRow >= 0 && curCol >= 0 && curRow < input.length && curCol < input[curRow].length) {
          if (input[curRow][curCol] != '.') {
            if (input[curRow][curCol] == '#') {
              count++;
            }
            break;
          } else {
            curRow -= rowIndex - check;
            curCol -= colIndex - seat;
          }
        }
      }
    }

    return count;

  }

  public static void main(String[] args) {
    String[] inputData = Helpers.readFile("input");

    char[][] input = new char[inputData.length][inputData[0].length()];

    for (int cnt = 0; cnt < inputData.length; cnt++) {
      input[cnt] = inputData[cnt].toCharArray();
    }

    char[][] previousLayout = new char[input.length][input[0].length]; // Assumes that all rows will be of equal width

    ArrayList<String[]> changeIndexes = new ArrayList<>();

    while (!Helpers.deepEquals(input, previousLayout)) {
      previousLayout = Helpers.deepCopy(input);

      for (int row = 0; row < input.length; row++) {
        for (int col = 0; col < input[row].length; col++) {
          if (input[row][col] == 'L') {
            if (numCanSee(row, col, input) == 0) {
              changeIndexes.add(new String[]{"#", String.valueOf(row), String.valueOf(col)});
            }
          } else if (input[row][col] == '#') {
            if (numCanSee(row, col, input) >= 5) {
              changeIndexes.add(new String[]{"L", String.valueOf(row), String.valueOf(col)});
            }
          }
        }
      }

      for (String[] change : changeIndexes) {
        input[Integer.parseInt(change[1])][Integer.parseInt(change[2])] = change[0].toCharArray()[0];
      }
      changeIndexes.clear();

      // System.out.println("Shifted");

    }

    Helpers.prettyPrint(input);

    int count = 0;

    for (char[] i : input) {
      for (char j = 0; j < i.length; j++) {
        if (i[j] == '#') {
          count++;
        }
      }
    }

    System.out.println(count);

  }
}
