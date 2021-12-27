package kata

"""
# Employee info
class Employee:
    def __init__(self, id: int, importance: int, subordinates: List[int]):
        # It's the unique id of each node.
        # unique id of this employee
        self.id = id
        # the importance value of this employee
        self.importance = importance
        # the id of direct subordinates
        self.subordinates = subordinates
"""

class Solution:
    def getImportance(self, employees: List['Employee'], id: int) -> int:
        
        employeesResolver = {}
        for employee in employees:
            employeesResolver[employee.id] = employee
        
        employee = employeesResolver[id]
        
        return self.getSubImportance(employeesResolver, employee)
        
    def getSubImportance(self, employees, employee) -> int:
        sum = employee.importance
          
        for subEmployeeID in employee.subordinates:
            sum = sum + employees[subEmployeeID].importance
            if employees[subEmployeeID].subordinates:
                for subSubEmployeeID in employees[subEmployeeID].subordinates:
                    sum = sum + self.getSubImportance(employees, employees[subSubEmployeeID])
        return sum        
