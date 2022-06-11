import argparse

from preparation import (
    load_target,
    load_cities_dataset,
    preprocess_cities_dataset,
    normalize_train_test,
    export_to_json
)
from model import fit_predict_catboost, fit_predict_regression, prepare_preds


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Loyal city prediction")
    parser.add_argument("-t", "--target", help="Path to target file")
    parser.add_argument("-c", "--cities", help="Path to cities dataset")
    parser.add_argument("-p", "--prediction", help="Path to prediction file")

    output = parser.parse_args()
    try:
        target = load_target(output.target)
        cities_df = load_cities_dataset(output.cities)

        train, y_train, test = preprocess_cities_dataset(cities_df, target)
        train_normed, test_normed = normalize_train_test(train, test)

        y_pred_catboost = fit_predict_catboost(train, y_train, test)
        y_pred_regression = fit_predict_regression(train_normed, y_train, test_normed)
        y_pred_total = prepare_preds(y_pred_catboost, y_pred_regression, test.index)

        data = {
            "Status": "Ok",
            "Message": None,
            "Data": y_pred_total
        }
    except Exception as e:
        data = {
            "Status": "Error",
            "Message": str(e),
            "Data": {}
        }

    export_to_json(data, output.prediction)
