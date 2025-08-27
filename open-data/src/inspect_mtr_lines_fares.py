from pathlib import Path
import pandas as pd

CURDIR = Path(__file__).parent.absolute()
data_file = CURDIR.parent / "data" / \
    "MTR routes, fares and barrier-free facilities" / "mtr_lines_fares.csv"


def compute_mtr_mean_fare(df: pd.DataFrame, key: str):
    data = df[df[key] != 0]  # Remove costs that are zero
    return data[key].mean()


def main():
    df = pd.read_csv(data_file)
    item_keys = {
        "Adult Octopus Fare": "OCT_ADT_FARE",
        "Adult Single Journey Ticket Fare": "SINGLE_ADT_FARE",
    }
    for item, key in item_keys.items():
        mean_fare = compute_mtr_mean_fare(df, key)
        print(f"{item} mean: {mean_fare}")


if __name__ == "__main__":
    main()
