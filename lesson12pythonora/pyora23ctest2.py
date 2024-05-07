import cx_Oracle

class pytestforora:

    def __init__(self, username, passwd, portNO):
        self.portNO = portNO
        self.username = username

    def oratest(self):
        user = 'system'
        password = 'wtsgyh1972'
        host = 'localhost:1521'
        service_name = 'FREE'
        conn_str = f"{self.username}/{password}@{host}/{service_name}"  # ('system/system@localhost:1521/FREE')
        connect = cx_Oracle.connect(conn_str)
        cursor = connect.cursor()
        print("测试Python连接Oracle 23c:")
        cursor.execute('select * from dual')
        # if cursor != null
        print(cursor.fetchall())
        print("看看当前Oracle的版本号：")
        all = cursor.execute('select * from v$version')
        print(all.fetchall())


pytestforora("system", "wtsgyh1972", "").oratest()
