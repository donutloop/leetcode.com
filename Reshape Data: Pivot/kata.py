import pandas as pd

def pivotTable(weather: pd.DataFrame) -> pd.DataFrame:
    weather = weather.pivot_table(index='month', columns='city', values='temperature', aggfunc='max')
    return weather