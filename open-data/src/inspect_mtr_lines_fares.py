import json
from pathlib import Path
import pandas as pd

CURDIR = Path(__file__).parent.absolute()
data_file = CURDIR.parent / "data" / \
    "MTR routes, fares and barrier-free facilities" / "mtr_lines_fares.csv"


def mtr_fare_stats(df: pd.DataFrame, key: str):
    data = df[df[key] != 0][key]  # Remove costs that are zero
    return {
        'mean': data.mean(),
        'median': data.median(),
        'standard_deviation': data.std(),
    }


def main():
    df = pd.read_csv(data_file)
    item_keys = {
        "Adult Octopus Fare": "OCT_ADT_FARE",
        "Adult Single Journey Ticket Fare": "SINGLE_ADT_FARE",
    }
    for item, key in item_keys.items():
        stats = mtr_fare_stats(df, key)
        print(f"[{item}]\n{json.dumps(stats, ensure_ascii=False)}")


if __name__ == "__main__":
    main()
