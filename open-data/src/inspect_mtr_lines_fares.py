import json
from pathlib import Path
import pandas as pd

CURDIR = Path(__file__).parent.absolute()
data_file = CURDIR.parent / "data" / \
    "MTR routes, fares and barrier-free facilities" / "mtr_lines_fares.csv"


def mtr_stations(df: pd.DataFrame):
    return df['SRC_STATION_NAME'].nunique()


def mtr_fare_stats(df: pd.DataFrame, key: str):
    data = df[df[key] != 0][key]  # Remove costs that are zero
    return {
        'mean': data.mean(),
        'median': data.median(),
        'standard_deviation': data.std(),
    }


def main():
    df = pd.read_csv(data_file)

    # Count MTR Stations
    print("Number of MTR stations:", mtr_stations(df))

    # Fare Statistics
    item_keys = {
        "Adult Octopus Fare": "OCT_ADT_FARE",
        "Adult Single Journey Ticket Fare": "SINGLE_ADT_FARE",
    }
    fare_statistics = {
        item: mtr_fare_stats(df, key)
        for item, key in item_keys.items()
    }
    print("Fare Statistics:",
          json.dumps(fare_statistics, ensure_ascii=False, indent=2))


if __name__ == "__main__":
    main()
