import pandas as pd
import os

def calculate_average_cu_ipc(file_path):
    df = pd.read_csv(file_path)
    cu_ipc_rows = df[df[' what'] == ' miss']
    cu_ipc_values = []
    for row in cu_ipc_rows[' value']:
        if isinstance(row, float):
            cu_ipc_values.append([row])
        else:
            cu_ipc_values.append([float(v) for v in row.split()])
    average_cu_ipc = sum(pd.Series([v for sublist in cu_ipc_values for v in sublist])) / len(pd.Series([v for sublist in cu_ipc_values for v in sublist]))
    return average_cu_ipc

def process_csv_files_in_directory():
    files = [file for file in os.listdir('.') if file.endswith('.csv')]
    for file in files:
        average_cu_ipc = calculate_average_cu_ipc(file)
        print(f"File: {file}, Average cu_IPC: {average_cu_ipc}")

if __name__ == "__main__":
    process_csv_files_in_directory()

