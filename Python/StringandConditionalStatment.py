#string
print("Hello World")
name = "Alice"
print("My name is", name)
print(len(name))  # Output: 5
print(name.upper())  # Output: ALICE   
print(name.lower())  # Output: alice
print(name[0])  # Output: A
print(name + " Smith")  # Output: Alice Smith   
print(name.count("i")) # Output: 1
name = name.replace("Alice", "Bob")
print(name)  # Output: Bob
print(name.split("i"))  # Output: ['Al', 'ce']
print(name.startswith("A"))  # Output: True
print(name.endswith("e"))  # Output: True
#slice
print(name[1:4])  # Output: lice
print(name[:4])  # Output: Alic
print(name[1:])  # Output: lice


