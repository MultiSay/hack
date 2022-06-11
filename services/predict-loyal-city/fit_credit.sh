#!/bin/bash

TARGET_PATH="data/target_credit.csv"
CITY_PATH="data/cities.csv"
PREDICTION_PATH="predictions/prediction_credit.json"

python main.py -t $TARGET_PATH -c $CITY_PATH -p $PREDICTION_PATH