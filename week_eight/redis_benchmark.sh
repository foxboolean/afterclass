OUTFILE="benchmark_result.log"
# 10 20 50 100 200 1k 5k 字节 value 大小
for size in 10 20 50 100 200 1000 5000;
do
  # 并发 6，指定请求 10000 次，
	redis-benchmark -h 127.0.0.1 -p 6379 -c 6 -d $size -n 10000 -t get >> $OUTFILE
	redis-benchmark -h 127.0.0.1 -p 6379 -c 6 -d $size -n 10000 -t set >> $OUTFILE
done