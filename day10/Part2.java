import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

public class Part2 {

    final int[] adapterList;

    HashMap<Integer, Long> knownBranches = new HashMap<>();

    Part2(int[] adapterList) {
        this.adapterList = adapterList;
    }

    private int validAdapterIndexes(int adapterIndex) {
      int curIndex = adapterIndex;

      for (;curIndex < this.adapterList.length && this.adapterList[curIndex] <= this.adapterList[adapterIndex] + 3;) {
        curIndex++;
      }

      return curIndex;

    }

    private long countBranches(int adapter) {

      if (adapter == this.adapterList.length - 1) { // Adapter is at end of list
        return 1; // End of branch
      }

      if (this.knownBranches.containsKey(adapter)) { // Branch has already been cached
        return this.knownBranches.get(adapter);
      }

      long ans = 0;

      for (int newAdapter = adapter + 1; newAdapter < this.validAdapterIndexes(adapter); newAdapter++) {
        ans += this.countBranches(newAdapter);
      }

      this.knownBranches.put(adapter, ans);

      return ans;
    }

    public static void main(String[] args) {

      int[] input = Helpers.stringArrayToIntArray(Helpers.readFile("input"), true);

      var part2 = new Part2(input);

      System.out.println(part2.countBranches(0));
    }
}
