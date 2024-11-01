import argparse
import serial
import time

# 解析命令列參數
parser = argparse.ArgumentParser()
parser.add_argument("-p", "--port", help="指定要連接的串口", default="/dev/ttys059")
args = parser.parse_args()

port = args.port
baudrate = 9600  # 根據實際設備調整
timeout = 1      # 1秒的timeout設定

try:
    # 建立串口連接
    ser = serial.Serial(port, baudrate=baudrate, timeout=timeout)
    print(f"已連接到 {port}")
    
    # 不斷讀取資料
    while True:
        # 讀取從設備傳回的資料
        if ser.in_waiting > 0:
            data = ser.readline().decode("utf-8").strip()  # 讀取並移除多餘的空白字符
            print(f"接收到的資料: {data}")
        else:
            print("尚未接收到資料，等待中...")
        
        # 延遲一段時間，避免占用過多資源
        time.sleep(1)

except serial.SerialException as e:
    print(f"串口錯誤: {e}")

finally:
    # 確保在結束程式時關閉連接
    if 'ser' in locals() and ser.is_open:
        ser.close()
        print("已關閉串口連接")
