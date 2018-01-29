### Limits

- Step Function Input : 32768 chars
- Lambda Synchronous Invocation Input : 6MB
- Lambda Asynchronous Invocation Input : 128KB

### Test Data

largest_non_compressed_stepfn.json - 32KB
largest_non_compressed_stepfn.json.gz - 607B

massive.json.10k.gz - original size : 1MB
massive.json.20k.gz - original size : 1.9MB
massive.json.30k.gz - original size : 2.8MB
massive.json.40k.gz - original size : 3.6MB
massive.json.50k.gz - original size : 4.5MB

## Results

### Step Function Max Size Input

No Compression	GZIP	JSON Optimisation	GZIP Fastest	LZ4
1265.48	1400.1	1189.2	1149.16	1286.23
2679.67	1385.39	1224.03	1185.19	1253.35
1448.92	1198.44	1213.78	1123.25	2270.49
1323.7	1355.37	1154.82	1153.13	2500.85
1324.87	1262.99	1202.77	1204.1	1262.69
1288.28	1229.45	1165.8	1162.22	1242.38
1389.03	1249.64	1221.47	1149.26	1250.79
1433.56	1215.51	1185.99	1119.26	1284.05
1337.96	1202.63	1280.92	1195.88	1205.02
1418.96	1399.72	1216.25	1191.44	1235.35
1491.043	1289.924	1205.503	1163.289	1479.12

### 30K Input

GZIP	JSON Optimisation	GZIP Fastest	LZ4
11467.62	6493.14	6666.27	7964.93
11895.34	6443.97	6412.57	7960.27
11321.56	6582.05	6500.98	7928.75
11681.86	6681.03	6409.2	7646.82
11465.53	5504.28	6537.33	7760.93
11804.2	6601.9	5365.73	8086.45
11608.79	6596.48	6433.47	7635.67
11773.75	6499.25	5416.71	7970.21
10640.79	6710.29	6475.18	7798.95
11478.13	6408.65	5334.84	7772.6
11513.757	6452.104	6155.228	7852.558

## Observations

### Step Function Max Size Input

- 14.464% faster when compressed using gzip
- 21.1782% faster when gzipping and optimising JSON
- 24.6958% faster when gziping (fastest) and optimising JSON
- 0.802852% faster when compressed using lz4 optimising JSON

### 30K Input

- 56.3475% faster when gzipping
- 60.6546% faster when gzipping and optimising JSON
- 37.81% faster when compressed using lz4 optimising JSON

### Thoughts

- The bottle neck is the data transfer between steps in a step function. 
- The results indicate that the more compressed the data, the faster the execution time.
- lz4 sacrifices speed for compression size and is slower than gzipped data by quite a margin.
- Optimising json marshaling/unmarshaling improves performance but not as much as gzip. 

