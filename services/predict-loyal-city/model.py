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


def prepare_preds(pred1, pred2, keys):
    y_pred_total = 0.5 * (pred1 + pred2)
    result = [
        {"id": i, "city": key, "predictScore": y}
        for i, (key, y) in enumerate(zip(keys, y_pred_total))
    ]
    result_sorted = sorted(result, key=lambda y: y["predictScore"], reverse=True)
    for i, item in enumerate(result_sorted):
        item.update({"position": i})

    return result_sorted
