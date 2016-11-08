#!/usr/bin/python3

import psycopg2
import os


if __name__ == '__main__':
	
	PG_PORT = os.environ.get("PG_PORT")
	conn = None
	
	try:
     
		conn = psycopg2.connect(database="docker", user="docker", password="docker", host="127.0.0.1", port=PG_PORT)
		print( "Opened database successfully")

		cur = conn.cursor()

		cur.execute("SELECT * FROM company;")
		
		print(" ".join(["{:>12s}".format(i.name) for i in cur.description]))
		rows = cur.fetchall()

		for row in rows:
			print (row)

	except psycopg2.DatabaseError as e:
	    print ('Error {ERR}'.format(ERR=e))
	    sys.exit(1)
    
    
	finally:
	    
	    if conn:
	        conn.close()
