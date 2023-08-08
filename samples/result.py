import os
import csv

benchmarks = ['aes', 'bfs', 'bitonicsort', 'concurrentkernel', 'fastwalshtransform', 'fir', 'floydwarshall', 'im2col', 'kmeans', 'matrixmultiplication', 'matrixtranspose', 'memcopy', 'pagerank', 'relu', 'stencil2d', 'xor']

result = []
csv_name = "32setIpoly.csv"
for benchmark in benchmarks:
    csv_path = os.path.join(benchmark, csv_name)
    ipc_sum = 0
    counter_ipc = 0
    l2_hit_sum = 0
    l2_total_sum = 0

    with open(csv_path, 'r') as f:
        reader = csv.reader(f)
        next(reader) # skip header

        l1_hit_count = dict()
        l1_total_count = dict()
        l1_hit_ratios = dict()

        for row in reader:
            row = [x.strip() for x in row]

            if row[-2] == 'cu_IPC':
                ipc_sum += float(row[-1])
                counter_ipc += 1
            elif row[1] == 'GPU[1].L2TLB':
                if row[-2] == 'hit':
                    l2_hit_sum += float(row[-1])
                elif row[-2] == 'miss':
                    l2_total_sum = l2_hit_sum + float(row[-1])
            else:
                if row[1].startswith('GPU[1].SA['):
                    l1vtlb_key = row[1][row[1].find('L1VTLB['):]

                    if row[-2] == 'hit':
                        l1_hit_count[l1vtlb_key] = float(row[-1])
                    elif row[-2] == 'miss':
                        if l1vtlb_key not in l1_total_count:
                            l1_total_count[l1vtlb_key] = 0
                        l1_total_count[l1vtlb_key] += float(row[-1])

                        l1_hit_ratios[l1vtlb_key] = l1_hit_count[l1vtlb_key] / (l1_hit_count[l1vtlb_key] + l1_total_count[l1vtlb_key])

    ipc_avg = ipc_sum / counter_ipc if counter_ipc > 0 else 0

    l2_hit_ratio = l2_hit_sum / l2_total_sum if l2_total_sum != 0 else 0

    l1_hit_ratio_avg = sum(l1_hit_ratios.values()) / len(l1_hit_ratios) if l1_hit_ratios else 0

    result.append([benchmark, ipc_avg, l2_hit_ratio, l1_hit_ratio_avg])

with open(f'{csv_name}_result.csv', 'w') as f:
    writer = csv.writer(f)
    writer.writerow(['Benchmark', 'Average IPC', 'L2TLB HIT ratio', 'L1VTLB HIT ratio'])
    writer.writerows(result)

