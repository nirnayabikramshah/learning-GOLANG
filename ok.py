num1=float(input("Enter the first number = "))
num2=float(input("Enter the second number =  "))
operator = input("Enter the operator =  ")
if operator =="-":
    print(num1-num2)
elif operator =="+":
    print(num1+num2)
elif operator == "*":
    print(num1*num2)
elif operator == "/":
    print(num1/num2)
    if num2 ==0:
        print("Can't divide my 0")
elif operator =="Exit" or "exit":
    exit()
elif num1 or num2 =="":
    print("Please enter both numbers")
else:
    print('please type a operator')


