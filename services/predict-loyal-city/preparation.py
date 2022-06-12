import json
import pandas as pd

from consts import (
    TRAIN_COLUMNS,
    COLUMNS_TO_NORMALIZE,
    TRANSACTION_COLUMNS,
    DEBIT_NAME,
    CREDIT_NAME,
)


def load_transactions_and_get_target(path, product_type: str = "debit"):
    df = pd.read_csv(path, usecols=TRANSACTION_COLUMNS)
    if product_type == "debit":
        mask = df["product_category_name"] == DEBIT_NAME
    elif product_type == "credit":
        mask = df["product_category_name"] == CREDIT_NAME
    else:
        raise KeyError

    target = df[mask].groupby("city")["client_id"].nunique()
    return target.rename("target")


def load_cities_dataset(path):
    cities_df = pd.read_csv(path, sep=";", usecols=TRAIN_COLUMNS)
    cities_df = cities_df[cities_df["year"] == 2019]
    cities_df = cities_df.drop(columns=["year"])
    cities_df = cities_df.set_index("settlement")
    return cities_df


def preprocess_cities_dataset(cities_df, target, norm_target: bool = True):
    # assign target column to data
    cities_df = cities_df.merge(
        target,
        left_index=True,
        right_index=True,
        how="left",
    )

    for col in COLUMNS_TO_NORMALIZE:
        cities_df[col] = cities_df[col] / cities_df["population"]

    if norm_target:
        cities_df["target"] = cities_df["target"] / cities_df["population"]

    test_mask = cities_df["target"].isna()
    train = cities_df[~test_mask]
    test = cities_df[test_mask]

    y_train = train.pop("target")
    test.pop("target")

    return train, y_train, test


def normalize_train_test(train, test):
    means = train.mean()
    stds = train.std()
    train_normed = (train - means) / stds
    test_normed = (test - means) / stds

    train_normed.fillna(0, inplace=True)
    test_normed.fillna(0, inplace=True)

    return train_normed, test_normed


def export_to_json(dct, filepath):
    with open(filepath, "w") as fp:
        json.dump(dct, fp, ensure_ascii=False)
