import socket
import time
import threading

HOST, PORT = '127.0.0.1', 8080
NUM_REQUESTS = 1000  # Number of requests to send per client
NUM_CLIENTS = 10    # Number of concurrent clients

def client_thread(client_id):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        
        # Start time
        start_time = time.time()
        
        for i in range(NUM_REQUESTS):

            method = 'PUT'
            key = f"key-{client_id}-{i}"
            value = f"value-{client_id}-{i}"

            # PUT Request
            request = f"{method} / HTTP/1.1\r\n"
            request += f"Host: {HOST}:{PORT}\r\n"
            request += f"Key: {key}\r\n"
            request += f"Val: {value}\r\n"
            request += "\r\n"
            
            s.sendall(request.encode())
            s.recv(1024)
            
            # GET request
            method = 'GET'
            request = f"{method} / HTTP/1.1\r\n"
            request += f"Host: {HOST}:{PORT}\r\n"
            request += f"Key: {key}\r\n"
            request += f"Val: {value}\r\n"
            request += "\r\n"
            
            s.sendall(request.encode())
            s.recv(1024)

            # DEL request
            method = 'DEL'
            request = f"{method} / HTTP/1.1\r\n"
            request += f"Host: {HOST}:{PORT}\r\n"
            request += f"Key: {key}\r\n"
            request += f"Val: {value}\r\n"
            request += "\r\n"
            
            s.sendall(request.encode())
            s.recv(1024)
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