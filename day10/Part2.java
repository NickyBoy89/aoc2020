import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

public class Part2 {

    final int[] adapterList;

    Part2(int[] adapterList) {
        this.adapterList = adapterList;
    }

    private int[] validAdapters(int rating) {

        ArrayList<Integer> valid = new ArrayList<>();

        for (int i = 0; i < this.adapterList.length && this.adapterList[i] <= rating + 3; i++) { // Makes sure that all adapters are at most 3 jolts higher than the target
            if (this.adapterList[i] > rating && this.adapterList[i] != rating) { // Filters out all adapters that are lower than the current adapter's rating
                valid.add(adapterList[i]);
            }
        }

        int[] result = new int[valid.size()];
        for (int i = 0; i < valid.size(); i++) {
            result[i] = valid.get(i);
        }

        return result;

    }

    private int countBranches(int adapter, int total) {

        int[] valid = validAdapters(adapter);

        if (valid.length == 0) {
            return total + 1; // Accounts for end of branch
        } else if (valid.length > 1) {
            return this.countAllBranches(valid, total);
        }

        return countBranches(valid[0], total);
    }

    private int countAllBranches(int[] adapters, int total) {

        int counter = 0;

        for (int adapter : adapters) {
            counter += countBranches(adapter, total);
        }

        return counter;
    }

    public static void main(String[] args) {

        int[] input = Helpers.stringArrayToIntArray(Helpers.readFile("input"));

        Arrays.sort(input);

        Part2 part2 = new Part2(input);

        System.out.println(part2.countBranches(0, 0));

    }
}
