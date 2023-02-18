class Solution {
    public String categorizeBox(int length, int width, int height, int mass) {
        var bulky = false;
        var heavy = false;

        // If the mass of the box is greater or equal to 100, it is "Heavy".
        if (mass >= 100) {
            heavy = true;
        }

        // The box is "Bulky" if any of the dimensions of the box is greater or equal to.
        var dimensionThreshold = (long) Math.pow(10.0, 4.0);
        if (length >= dimensionThreshold || width >= dimensionThreshold || height >= dimensionThreshold) {
            bulky = true;
        }

        //  The box is "Bulky" if the volume of the box is greater or equal to
        if (((long)length * (long) width * (long) height) >= (long) Math.pow(10.0, 9)) {
            bulky = true;
        }

        // If the box is both "Bulky" and "Heavy", then its category is "Both"
        if (bulky && heavy) {
            return "Both";
        // If the box is neither "Bulky" nor "Heavy", then its category is "Neither"
        }else if (!bulky && !heavy) {
            return "Neither";
        // If the box is "Bulky" but not "Heavy", then its category is "Bulky
        } else if (bulky && !heavy) {
            return "Bulky";
        }

        // If the box is "Heavy" but not "Bulky", then its category is "Heavy".
        return "Heavy";
    }
}