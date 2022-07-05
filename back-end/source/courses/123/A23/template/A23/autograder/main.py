# coding:utf-8
import os
import yaml
import json

def file_path_test(path):
    try:
        f =open(yamlPath)
        f.close()
        return True
    except IOError:
        return False

def compare_str(normal_str1, normal_str2):
    lower_str1 = normal_str1.lower()
    lower_str2 = normal_str2.lower()
    if lower_str1 == lower_str2:
        # print("True")
        return True
    else:
        # print("False")
        return False

curPath = os.path.dirname(os.path.realpath(__file__))
yamlPath = os.path.join(curPath, "config.yaml")
sulPath = os.path.join(curPath, "solution.json")
ansPath = os.path.join(curPath, "answer.json")

if file_path_test(yamlPath) == False :
    print("Configuration is not found.")
    exit(-1)

if file_path_test(sulPath) == False :
    print("Solution is not found.")
    exit(-1)
    
if file_path_test(ansPath) == False :
    print("Answer is not found.")
    exit(-1)

with open(yamlPath, 'r', encoding='utf-8') as f:
    config = f.read()
configuration = yaml.load(config,Loader=yaml.FullLoader)
# print(configuration)

with open(sulPath, 'r', encoding='utf-8') as sul_file:
    sul_data = json.load(sul_file)
# print(sul_data)

with open(ansPath, 'r', encoding='utf-8') as ans_file:
    ans_data = json.load(ans_file)
# print(ans_data)

score_result = {}
for key in sul_data.keys():
    tmp_Q = {}
    for sub_key in sul_data[key]:
        if compare_str(sul_data[key][sub_key],ans_data[key][sub_key]):
            tmp_sub = {key + "_" + sub_key: configuration[key][sub_key]}
        else:
            tmp_sub = {key + "_" + sub_key: 0}
        score_result.update(tmp_sub)
        
        # tmp_Q.update(tmp_sub)
    # score_result.update({key:tmp_Q})

js = json.dumps({"scores": score_result})
print(js)