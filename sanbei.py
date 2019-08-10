import requests
from bs4 import BeautifulSoup
import pymysql

headers = {
    "user-agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36"
}

def getCategory():
    urlAddress = "https://www.shanbay.com/wordbook/category/10/"
    r = requests.get(urlAddress, headers = headers)
    soup = BeautifulSoup(r.text,"html.parser",from_encoding="utf-8")
    ulnode = soup.find("ul", class_ = "nav nav-tabs nav-stacked")
    rlist = []
    for item in ulnode.find_all("a"):
        rlist.append({"name":item.get_text(), "href":item["href"]})
    return rlist

def getBooks(hrefUrl):
    urlAddress = "https://www.shanbay.com" + hrefUrl
    r = requests.get(urlAddress, headers = headers)
    soup = BeautifulSoup(r.text,"html.parser",from_encoding="utf-8")
    booksNode = soup.find("div", class_ = "span8")
    rlist = []
    for item in booksNode.find_all("div", class_ = "wordbook-title"):
        item = item.find("a")
        rlist.append({"name":item.get_text(), "href":item["href"]})
    return rlist

def getBookUnit(hrefUrl):
    urlAddress = "https://www.shanbay.com" + hrefUrl
    r = requests.get(urlAddress, headers = headers)
    soup = BeautifulSoup(r.text,"html.parser",from_encoding="utf-8")
    node = soup.find("div", id = "wordbook-wordlist-container", class_ = "well")
    rlist = []
    for item in node.find_all("a"):
        rlist.append({"name":item.get_text(), "href":item["href"]})
    return rlist

def getBookUnitWords(hrefUrl):
    rList = []
    page = 1
    while True:
        urlAddress = "https://www.shanbay.com" + hrefUrl +"?page=" + str(page)
        r = requests.get(urlAddress, headers = headers)
        soup = BeautifulSoup(r.text,"html.parser",from_encoding="utf-8")
        node = soup.find("div", class_ = "span8")
        if len(node.find_all("tr")) == 0:
            break
        for item in node.find_all("tr"):
            rList.append({"name": item.find("strong").get_text(), "value":item.find("td", class_ = "span10").get_text()})
        page += 1
    return rList

def insertWord(cursor, category, book, unit, name, value):
    value = value.replace("'", "\\'")
    value = value.replace("\\n", "#")
    sqls = "INSERT INTO words(`category`, `book`, `unit`, `word`, `value`) VALUES('%s', '%s', '%s', '%s', '%s')" % (category, book, unit, name.replace("'", "\\'"), value)
    try:
        cursor.execute(sqls)
    except:
        print("错误的sql:", sqls)
    
def scanAllWords(cursor):
    for item in getCategory():
        for itembook in getBooks(item["href"]):
            for itemUnit in getBookUnit(itembook["href"]):
                print("正在抓取：", item["name"], "  ", itembook["name"], "  ", itemUnit["name"])
                for word in getBookUnitWords(itemUnit["href"]):
                    insertWord(cursor, item["name"], itembook["name"], itemUnit["name"], word["name"], word["value"])


db = pymysql.connect(host = "localhost", port = 3306, user = "root", password = "root", db = "enwords", charset='utf8')
scanAllWords(db.cursor())
db.close()
