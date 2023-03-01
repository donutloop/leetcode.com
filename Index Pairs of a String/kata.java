class Solution {
    public int[][] indexPairs(String text, String[] words) {
        var list = new ArrayList<ArrayList<Integer>>();
        for (var i = 0; i < text.length(); i++) {
            for (var word: words) {
                if ((i+word.length()) <= text.length()) {
                    if (text.substring(i, i+word.length()).equals(word)) {
                        final var index = i;
                        list.add(new ArrayList<Integer>(){{
                            add(index);
                            add(index+word.length()-1);
                        }});
                    }
                }
            }
        }

        Collections.sort(list, new Comparator<ArrayList<Integer>>() {
            public int compare(ArrayList<Integer> a, ArrayList<Integer> b) {
                if (a.get(0) == b.get(0)) {
                    if (a.get(1) == b.get(1)) {
                        return 0;
                    } else if (a.get(1) < b.get(1)) {
                        return -1;
                    } else {
                        return 1;
                    }
                } else if (a.get(0) < b.get(0)) {
                        return -1;
                } else {
                    return 1;
                }
            }
        });


        var res = new int[list.size()][2];
        var i = 0;
        for (var element : list) {
            res[i] = new int[]{element.get(0), element.get(1)};
            i++;
        }

        return res;
    }

}