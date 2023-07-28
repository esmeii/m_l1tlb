import os
import pandas as pd

# Get the list of all CSV files in the current directory
csv_files = [file for file in os.listdir() if file.endswith('.csv')]

for file in csv_files:
    # Read the CSV file into a DataFrame
    df = pd.read_csv(file)

    # Filter rows with 'L1VTLB' in the 'where' column
    filtered_df = df[df[' where'].str.contains('L1VTLB', na=False)]

    # Filter rows with 'what' column values as 'hit' or 'miss'
    filtered_df = filtered_df[filtered_df[' what'].isin([' hit', ' miss'])]

    # Convert 'value' column to numeric type (in case it's not already)
    filtered_df[' value'] = pd.to_numeric(filtered_df[' value'])

    # Calculate the sum of 'value' column in the filtered DataFrame
    total_value = filtered_df[' value'].sum()

    # Filter rows with 'L1VTLB' in the 'where' column and 'hit' in the 'what' column
    hit_filtered_df = df[(df[' where'].str.contains('L1VTLB', na=False)) & (df[' what'] == ' hit')]

    # Convert 'value' column to numeric type (in case it's not already)
    hit_filtered_df[' value'] = pd.to_numeric(hit_filtered_df[' value'])

    # Calculate the sum of 'value' column in the filtered DataFrame
    total_hit_value = hit_filtered_df[' value'].sum()

    # Calculate and print the hit ratio
    if total_value > 0:
        hit_ratio = total_hit_value / total_value
        print(f"File: {file}, Hit Ratio: {hit_ratio:.2f}")
    else:
        print(f"File: {file}, No hits recorded.")

