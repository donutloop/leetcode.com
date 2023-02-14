import java.time.LocalTime;

class Solution {
    public boolean haveConflict(String[] event1, String[] event2) {
        var eventAStart = LocalTime.parse(event1[0]);
        var eventAEnd = LocalTime.parse(event1[1]);
        var eventBStart = LocalTime.parse(event2[0]);
        var eventBEnd = LocalTime.parse(event2[1]);

        return eventAEnd.equals(eventBStart)
        || (eventAStart.equals(eventBEnd))
        || (eventBStart.isAfter(eventAStart) && eventBStart.isBefore(eventAEnd))
        || (eventBEnd.isAfter(eventAStart) && eventBEnd.isBefore(eventAEnd))
        || (eventAStart.isAfter(eventBStart) && eventAStart.isBefore(eventBEnd))
        || (eventAEnd.isAfter(eventBStart) && eventAEnd.isBefore(eventBEnd));
    }
}