import cx_Oracle

#sure can connect and create table and insert items: WSPDB
#remember to use the right user/passwd and PDB name. and open the PDB!to: 4 WSPDB			  READ WRITE NO
conn = cx_Oracle.connect('biodbuser/biodb123@localhost:1521/WSPDB')
cursor = conn.cursor()

cursor.execute("CREATE TABLE INSERTTEST(ID INT, C1 VARCHAR(50), C2 VARCHAR(50), C3 VARCHAR(50))")

cursor.execute("INSERT INTO INSERTTEST (ID, C1, C2, C3)VALUES(1213412, 'asdfa', 'ewewe', 'sfjgsfg')")
cursor.execute("INSERT INTO INSERTTEST (ID, C1, C2, C3)VALUES(12341, 'ashdfh', 'shhsdfh', 'sghs')")
cursor.execute("INSERT INTO INSERTTEST (ID, C1, C2, C3)VALUES(123451235, 'werwerw', 'asdfaf', 'awew')")
conn.commit()   ## 这里一定要commit才行，要不然数据是不会插入的

cursor.close()
conn.close()
