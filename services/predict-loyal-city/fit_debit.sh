#!/bin/bash

TARGET_PATH="data/transactions.csv"
CITY_PATH="data/cities.csv"
PREDICTION_PATH="predictions/prediction_debit.json"
PRODUCT_TYPE="debit"

python main.py -t $TARGET_PATH -c $CITY_PATH -p $PREDICTION_PATH -pt $PRODUCT_TYPE