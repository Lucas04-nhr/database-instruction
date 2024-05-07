import cx_Oracle

conn = cx_Oracle.connect('system/wtsgyh1972@localhost:1521/FREE')
cursor = conn.cursor()

cursor.execute("SELECT * FROM v$version")
rows = cursor.fetchall()  # 得到所有数据集
for row in rows:
    print("%s, %s, %s, %s" % (row[0], row[1], row[2], row[3]))  # python3以上版本中print()要加括号用了

print("Number of rows returned: %d" % cursor.rowcount)

cursor.execute("SELECT * FROM v$version")
while (True):
    row = cursor.fetchone()  # 逐行得到数据集
    if row == None:
        break
    print("%s, %s, %s, %s" % (row[0], row[1], row[2], row[3]))

print("Number of rows returned: %d" % cursor.rowcount)