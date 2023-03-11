// January: 31 days
// February: 28 days in a common year, 29 days in a leap year
// March: 31 days
// April: 30 days
// May: 31 days
// June: 30 days
// July: 31 days
// August: 31 days
// September: 30 days
// October: 31 days
// November: 30 days
// December: 31 days

class Solution {
    public int numberOfDays(int year, int month) {
        if (month == 1
        || month == 3
        || month == 5
        || month == 7
        || month == 8
        || month == 10
        || month == 12) {
            return 31;
        }

        if (month == 2) {
            if (year % 4 != 0) return 28;
            else if (year % 100 != 0) return 29;
            else if (year % 400 != 0) return 28;
            return 29;
        }

        return 30;
    }

}

