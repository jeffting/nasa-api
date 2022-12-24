import requests
import sys
 

def getLastTenDays():
    res = requests.get("http://localhost:8081/v1/rover/images")
    print(res.json())

def getSpecificDay(date):
    url = "http://localhost:8081/v1/rover/images/day?earth_date=" + date
    res = requests.get(url)
    print(res.json())


args = sys.argv
if len(args) == 2 and args[1] == "lastTenDays":
    getLastTenDays()
if len(args) == 3 and args[1] == "specificDay":
    date = args[2]
    getSpecificDay(date)



