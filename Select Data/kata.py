import pandas as pd

def selectData(students: pd.DataFrame) -> pd.DataFrame:
    students = students.query("student_id == 101")
    return students[['name', 'age']]