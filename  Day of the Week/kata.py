# math equation by Tomohiko Sakamoto
class Solution:
    def dayOfTheWeek(self, d: int, m: int, y: int) -> str:
        t = [ 0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4 ]

        # if month is less than 3 reduce year by 1
        if (m < 3) :
            y = y - 1

        f = (y + y // 4 - y // 100 + y // 400 + t[m - 1] + d) % 7

        weekDay = ""
        if f == 6:
            weekDay = "Saturday"
        elif f == 0:
            weekDay = "Sunday"
        elif f == 1:
            weekDay = "Monday"
        elif f == 2:
            weekDay = "Tuesday"
        elif f == 3:
            weekDay = "Wednesday"
        elif f == 4:
            weekDay = "Thursday"
        elif f == 5:
            weekDay = "Friday"

        return weekDay
        