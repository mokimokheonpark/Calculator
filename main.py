while True:

    equation = input()

    if equation in ["exit", "Exit", "EXIT"]:
        break

    if equation.count("+") + equation.count("*") + equation.count("/") == 1:
        if "+" in equation:
            operator = "+"
        elif "*" in equation:
            operator = "*"
        else:
            operator = "/"
        operatorIndex = equation.index(operator)
    elif "+" not in equation and "*" not in equation and "/" not in equation:
        operator = "-"
        if equation.count("-") == 1 or (equation.count("-") == 2 and "--" in equation):
            operatorIndex = equation.index("-")
        elif equation.count("-") == 3 or (equation.count("-") == 2 and "--" not in equation):
            operatorIndex = equation.index("-", 2)
        else:
            print("Error: Invalid Number of Operators")
            continue
    else:
        print("Error: Invalid Number of Operators")
        continue

    left = equation[:operatorIndex]
    right = equation[operatorIndex+1:]

    if "." not in left and "." not in right:
        try:
            leftNumber = int(equation[:operatorIndex])
            rightNumber = int(equation[operatorIndex+1:])
        except:
            print("Error: Invalid Number(s)")
            continue
    else:
        try:
            leftNumber = float(equation[:operatorIndex])
            rightNumber = float(equation[operatorIndex+1:])
        except:
            print("Error: Invalid Number(s)")
            continue

    if operator == "+":
        print(leftNumber + rightNumber)
    elif operator == "-":
        print(leftNumber - rightNumber)
    elif operator == "*":
        print(leftNumber * rightNumber)
    elif rightNumber == 0:
        print("Error: Division By Zero")
    else:
        print(leftNumber / rightNumber)