from itertools import zip_longest

from catboost import CatBoostRegressor
from sklearn.linear_model import Ridge


def fit_predict_catboost(X_train, y_train, X_test):
    model = CatBoostRegressor(depth=3)
    model.fit(X_train, y_train, verbose=False)

    y_pred = model.predict(X_test)
    return y_pred


def fit_predict_regression(X_train, y_train, X_test):
    model = Ridge(alpha=1000)
    model.fit(X_train, y_train)

    y_pred = model.predict(X_test)
    return y_pred


def _fillna(val, replace=0):
    if val is None:
        return replace
    else:
        return val


def prepare_preds(pred1, pred2, y_true, keys, product_type):
    y_pred_total = 0.5 * (pred1 + pred2)
    result = [
        {
            "id": i,
            "city": key,
            "predictClientIndex": y_p,
            "currentClientIndex": _fillna(y_t),
            "product": product_type
        }
        for i, (key, y_p, y_t) in enumerate(zip_longest(keys, y_pred_total, y_true))
    ]
    result_sorted = sorted(result, key=lambda y: y["predictClientIndex"], reverse=True)
    for i, item in enumerate(result_sorted):
        item.update({"position": i})

    return result_sorted
