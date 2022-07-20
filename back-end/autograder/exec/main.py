# coding:utf-8
import os
import sys
import json
import autograders
import re


def is_number(num):
    pattern = re.compile(r'^(-?\d+)(\.\d+)?$')
    result = pattern.match(num)
    if result:
        return True
    else:
        return False


def file_path_test(path):
    try:
        f =open(path)
        f.close()
        return True
    except IOError:
        return False


curPath = os.path.dirname(os.path.realpath(__file__))
question_type = str(sys.argv[1])
sulPath = os.path.join(curPath, "solution.json")
ansPath = os.path.join(curPath, "answer.json")

if file_path_test(sulPath) == False:
    print("Solution is not found.")
    exit(1)

if file_path_test(ansPath) == False:
    print("Answer is not found.")
    exit(1)

with open(sulPath, 'r', encoding='utf-8') as sul_file:
    sul_data = json.load(sul_file)
# print(sul_data)

with open(ansPath, 'r', encoding='utf-8') as ans_file:
    ans_data = json.load(ans_file)
# print(ans_data)

for key in sul_data:
    try:   
        module = getattr(autograders, question_type)
        func = getattr(module, question_type)
    except:
        print("Autograder for " + question_type  + " is not found.")
        exit(1)
    rate = func(sul_data[key], ans_data[key])

if isinstance(rate,int) or isinstance(rate,float):
    if is_number(str(rate)):
        if rate >= 0 and rate <= 1:
            print("According to the test example you provided, the score for this question is " + str(rate*100) + " out of 100.")
            exit(0)
        else:
            print("The return value of the function must be greater than or equal to 0.0 and less than or equal to 1.0.")
            exit(1)
    else:
        print("The return value of the function must be greater than or equal to 0.0 and less than or equal to 1.0.")
        exit(1)
else:
    print("The return value of the function must be an integer or a floating point number.")
    exit(1)