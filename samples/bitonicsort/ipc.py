import pandas as pd

# Read the CSV file into a pandas DataFrame
df = pd.read_csv('4baseline.csv')

# Filter rows containing cu_IPC values
cu_ipc_rows = df[df[' what'] == ' cu_IPC']

# Extract cu_IPC values and convert to a list of floats
cu_ipc_values = []
for row in cu_ipc_rows[' value']:
    if isinstance(row, float):
        cu_ipc_values.append([row])  # If it's a float, create a list with that single value
    else:
        cu_ipc_values.append([float(v) for v in row.split()])

# Calculate the average of cu_IPC values
average_cu_ipc = sum(pd.Series([v for sublist in cu_ipc_values for v in sublist])) / len(pd.Series([v for sublist in cu_ipc_values for v in sublist]))

print("cu_IPC values:")
print(cu_ipc_values)
print("Average cu_IPC:", average_cu_ipc)

