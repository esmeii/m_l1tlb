import os
import csv

benchmarks = ['aes', 'bfs', 'bitonicsort', 'concurrentkernel', 'fastwalshtransform', 'fir', 'floydwarshall', 'im2col', 'kmeans', 'matrixmultiplication', 'matrixtranspose', 'memcopy', 'pagerank', 'relu', 'stencil2d', 'xor']

result = []
csv_name = "baseline.csv"
output_csv_name = "baseline_cache.csv"

for benchmark in benchmarks:
    csv_path = os.path.join(benchmark, csv_name)

    stats = {
        "L1VCache": {"read-hit": 0, "read-miss": 0},
        "L2": {"read-hit": 0, "read-miss": 0},
        "DRAM": {"read_trans_count": 0, "write_trans_count": 0},
    }

    with open(csv_path, 'r') as f:
        reader = csv.reader(f)
        next(reader) # skip header

        for row in reader:
            row = [x.strip() for x in row]

            for key in stats:
                if key in row[1] and (row[-2] in stats[key] or row[-2].startswith("trans_count")):
                    stats[key][row[-2]] += float(row[-1])

    l1_hit_ratio = stats["L1VCache"]["read-hit"] / (stats["L1VCache"]["read-hit"] + stats["L1VCache"]["read-miss"]) if (stats["L1VCache"]["read-hit"] + stats["L1VCache"]["read-miss"]) != 0 else 0
    l2_hit_ratio = stats["L2"]["read-hit"] / (stats["L2"]["read-hit"] + stats["L2"]["read-miss"]) if (stats["L2"]["read-hit"] + stats["L2"]["read-miss"]) != 0 else 0
    dram_hit_ratio = stats["DRAM"]["read_trans_count"] / (stats["DRAM"]["read_trans_count"] + stats["DRAM"]["write_trans_count"]) if (stats["DRAM"]["read_trans_count"] + stats["DRAM"]["write_trans_count"]) != 0 else 0

    result.append([benchmark, l1_hit_ratio, l2_hit_ratio, dram_hit_ratio])

with open(output_csv_name, 'w') as f:
    writer = csv.writer(f)
    writer.writerow(['Benchmark', 'L1VCache HIT ratio', 'L2 HIT ratio', 'DRAM HIT ratio'])
    writer.writerows(result)

