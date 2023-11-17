import multiprocessing
import requests
import time
import random
import multiprocessing
import requests
import time
import random
import sys
#from hashring import HashRing
#BASE_URLS = ['http://127.0.0.1:8080', 'http://127.0.0.1:8081', 'http://127.0.0.1:8082']
#BASE_URLS = ['http://127.0.0.1:8080', 'http://127.0.0.1:8081']
#BASE_URLS = ['http://127.0.0.1:8085']
#ring = HashRing(BASE_URLS)


def kv_store_client(operation, key, value=None):
    # Determine which server to use based on the key
    server_url = 'http://127.0.0.1:80'
    start_time = time.time()
    try:
        if operation == 'set':
            response = requests.put(f'{server_url}', headers={'Key':key, 'Val': value})
        elif operation == 'get':
            response = requests.get(f'{server_url}', headers={'Key':key})
        else:
            raise ValueError('Invalid operation')
        end_time = time.time()
        return end_time - start_time
    except Exception as e:
        print(f'Error performing {operation} on {key}: {e}')
        return None

def worker(num_operations, latencies_queue, operation, process_index):
    for i in range(num_operations):
        key = f'key{process_index}_{i}'
        value = f'value{process_index}_{i}' if operation == 'set' else None
        latency = kv_store_client(operation, key, value)
        if latency is not None:
            latencies_queue.put(latency)

def benchmark(num_operations, num_processes):
    manager = multiprocessing.Manager()
    latencies_queue = manager.Queue()

    set_processes = [
        multiprocessing.Process(target=worker, args=(num_operations, latencies_queue, 'set', i))
        for i in range(num_processes)
    ]
    get_processes = [
        multiprocessing.Process(target=worker, args=(num_operations, latencies_queue, 'get', i))
        for i in range(num_processes)
    ]
    start_time = time.time()

    for p in set_processes:
        p.start()
    for p in set_processes:
        p.join()

    for p in get_processes:
        p.start()
    for p in get_processes:
        p.join()

    total_time = time.time() - start_time
    total_operations = num_operations * num_processes * 2  # Each process does num_operations SET and GET
    total_latencies = []

    while not latencies_queue.empty():
        total_latencies.append(latencies_queue.get())

    average_latency = sum(total_latencies) / len(total_latencies)
    print(f'Total Latency: {sum(total_latencies):.2f} second')
    print(f'Length of latencies: {len(total_latencies):.0f} latencies')
    throughput = total_operations / total_time
    print(f'Total Operation: {total_operations:.0f} operations')
    print(f'Average Latency: {average_latency:.5f} seconds per operation')
    print(f'Throughput: {throughput:.2f} operations per second')
    print(f'Total Benchmark Time: {total_time:.2f} seconds')

if __name__ == '__main__':
    num_operations_per_process = int(sys.argv[1])  # Number of requests
    num_processes = int(sys.argv[2])  # Number of concurrent clients
    benchmark(num_operations_per_process, num_processes)
