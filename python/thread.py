#coding=utf-8

import Queue
import threading
import time

exitFlag = 0

class MyThread(threading.Thread):
	"""docstring for ClassName"""
	def __init__(self,  threadID, name, q):
		threading.Thread.__init__(self)
		self.threadID = threadID
		self.name = name
		self.q = q
	def run(self):
		print "Starting " + self.name
		process_data(self.name, self.q)
		print "Exiting " + self.name

def process_data(threadName, q):
	while  not exitFlag:
		queueLock.acquire()
		if not q.empty():
			data = q.get()			
			print "%s processing %s" % (threadName, data)
		queueLock.release()
		time.sleep(1)

threadList = ["Thread-1", "Thread-2", "Thread-3"]
nameList = ["One", "Two", "Three", "Four", "Five"]
queueLock = threading.Lock()
workQueue = Queue.Queue(10)
threads = []
threadID = 1

# 创建新线程
for tName in threadList:
	thread = MyThread(threadID, tName, workQueue)
	thread.start()
	threads.append(thread)
	threadID+=1

# 填充队列
for tName in nameList:
	queueLock.acquire()
	workQueue.put(tName)
	queueLock.release()

# 等待队列清空
while not workQueue.empty():
	pass

# 通知线程是时候退出
exitFlag = 1

# 等待所有线程完成
for t in threads:
	t.join()

print "Exiting Main Thread"	