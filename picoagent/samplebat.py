import json
import time

if __name__ == '__main__':
    for i in range(10):
        d = {"schema": "picoagent_sample:0.0.1", "time": time.time(), "number": i}
        print(json.dumps(d))
        time.sleep(1)