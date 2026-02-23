# List and Tuples in Python
student = ["John", "Emily", "Michael", "Sarah", "David"]
print(student[0])  # Output: John
student.append("Jessica")
print(student)  # Output: ['John', 'Emily', 'Michael', 'Sarah', 'David', 'Jessica']
student.insert(2, "Daniel")
print(student)  # Output: ['John', 'Emily', 'Daniel', 'Michael', 'Sarah', 'David', 'Jessica']
student.remove("Emily") 
print(student)  # Output: ['John', 'Daniel', 'Michael', 'Sarah', 'David', 'Jessica']
print(student[1:4])  # Output: ['Daniel', 'Michael', 'Sarah']
print(student[:3])  # Output: ['John', 'Daniel', 'Michael']     
print(student[3:])  # Output: ['Sarah', 'David', 'Jessica']
print(len(student))  # Output: 6
print(student.index("David"))  # Output: 4
print(student.count("Michael"))  # Output: 1
student.sort()
print(student)  # Output: ['Daniel', 'David', 'Jessica', 'John', 'Michael', 'Sarah']
student.reverse()
print(student)  # Output: ['Sarah', 'Michael', 'John', 'Jessica', 'David', 'Daniel']
#tuple
fruits = ("apple", "banana", "orange")  
print(fruits[0])  # Output: apple
print(len(fruits))  # Output: 3 
print(fruits.index("banana"))  # Output: 1
print(fruits.count("orange"))  # Output: 1

