#coding=utf-8
print ("helllo print 你好，世界");
if True:
	print("true")
else :
	print("false")
	print("false2")
print("dddd")

str = "test 123 "
print str*3
print str[-2:-5]
print str +"456"

x = "100"
y=int(x, 16)
print y
y = oct(10)
print y

y = 128
if (y > 1000) or  (y < 900):
	print(y)


#name =raw_input("input a name")
#print name

for letter in 'Python':     # First Example
   print 'Current Letter :', letter

fruits = ['banana', 'apple',  'mango']

for index in range(len(fruits)):
   print 'Current fruit :', fruits[index]

for num in range(10,20):
	print num


i = 2
while(i < 100):
   j = 2
   while(j <= (i/j)):
      if not(i%j): break
      j = j + 1
   if (j > i/j) : print i, " 是素数"
   i = i + 1

print "Good bye!"


str = "dddldldeojkjdskjfksdjfd"
print str.capitalize()
print str.center(100, "2")

str = str.encode('base64','strict');

print "Encoded String: " + str;
print "Decoded String: " + str.decode('base64','strict')

for x in range(2,10,3) :
   print x

obj = open("1.txt", "a+", -1)
print "file = ", obj.name, obj.mode, obj.closed
obj.write("heheheheheehehehehheehehehe")
obj.flush()
obj.close()

import os
print os.getcwd()

list = ['physics', 'chemistry', 1997, 2000];

print "Value available at index 1: "
print list[1];
list[1] = 2001;
print "New value available at index 1 : "
print list[1];
print max(list)
print min(list)
list.append("ddddd")
print list;


dict = {'Name': 'Zara', 'Age': 7};
print "Variable Type : %s" %  type (dict)

import time
loc = time.localtime(time.time())
print loc.tm_year, loc.tm_hour
print time.asctime(loc)

import calendar
print calendar.month(2015, 8)

def procedure():
    time.sleep(0.5)

# measure process time
t0 = time.clock()
print t0, "first process time"
procedure()
print time.clock() -t0, "seconds process time"

# measure wall time
t0 = time.time()
procedure()
print time.time() - t0, "seconds wall time"

import thread

# 为线程定义一个函数
def print_time1( threadName, delay):
   count = 0
   while count < 5:
      time.sleep(delay)
      count += 1
      print "%s: %s" % ( threadName, time.ctime(time.time()) )

# 创建两个线程
#try:
 #  thread.start_new_thread( print_time1, ("Thread-1", 2, ) )
 #  thread.start_new_thread( print_time1, ("Thread-2", 4, ) )
#except:
 #  print "Error: unable to start thread"


import threading
exitFlag = 0
class myThread(threading.Thread):
   def __init__(self, hreadID, name, counter):
      threading.Thread.__init__(self)
      self.hreadID = hreadID
      self.name = name
      self.counter = counter

   def run(self):
      print "Starting " + self.name
      threadLock.acquire()
      print_time(self.name, self.counter, 5)
      threadLock.release()
      print "Exiting " + self.name

def print_time(threadName, delay, counter):
      while counter:
        if exitFlag:
         thread.exit()
        time.sleep(delay)
        print "%s: %s" % (threadName, time.ctime(time.time()))
        counter -= 1  

threadLock = threading.Lock()
threads = []

thread1 = myThread(1, "Thread-3", 1)
thread2 = myThread(2, "Thread-4", 1)

# 开启线程
thread1.start()
thread2.start()

# 添加线程到线程列表
threads.append(thread1)
threads.append(thread2)

# 等待所有线程完成
for t in threads:
   print t.getName()
   t.join()
print "Exiting Main Thread"