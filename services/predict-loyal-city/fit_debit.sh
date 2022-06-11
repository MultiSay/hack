#!/bin/bash

TARGET_PATH="data/target_debit.csv"
CITY_PATH="data/cities.csv"
PREDICTION_PATH="predictions/prediction_debit.json"

python main.py -t $TARGET_PATH -c $CITY_PATH -p $PREDICTION_PATH