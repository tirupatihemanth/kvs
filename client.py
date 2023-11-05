import requests
import time
import threading
import sys

HOST, PORT = 'http://localhost', 80
NUM_REQUESTS = int(sys.argv[1])  # Number of requests to send per client
NUM_CLIENTS = int(sys.argv[2])  # Number of concurrent clients

def client_thread(client_id):
    session = requests.Session()

    # Start time
    start_time = time.time()

    for i in range(NUM_REQUESTS):
        key = f"key-{client_id}-{i}"
        value = f"value-{client_id}-{i}"

        # PUT Request
        data = {'Key': key, 'Val': value}
        response = session.put(f"{HOST}:{PORT}", headers=data)

        # GET request
        response = session.get(f"{HOST}:{PORT}", headers=data)
        #print(response.text)
        # DEL request
        response = session.delete(f"{HOST}:{PORT}", headers=data)
        #print(response.text)
    # End time
    end_time = time.time()

    elapsed_time = end_time - start_time
    throughput = NUM_REQUESTS * 3 / elapsed_time  # Multiply by 3 for PUT, GET, DEL
    latency = elapsed_time / (NUM_REQUESTS * 3)  # Average latency per request

    print(f"Client-{client_id} | Throughput: {throughput:.2f} req/s | Average Latency: {latency:.6f} seconds")

if __name__ == "__main__":
    clients = []
    for i in range(NUM_CLIENTS):
        t = threading.Thread(target=client_thread, args=(i,))
        clients.append(t)
        t.start()

    for t in clients:
        t.join()

    print("Testing completed.")
