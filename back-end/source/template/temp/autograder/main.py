# coding:utf-8
import os
import yaml
import json
import autograders


def file_path_test(path):
    try:
        f =open(path)
        f.close()
        return True
    except IOError:
        return False


curPath = os.path.dirname(os.path.realpath(__file__))
yamlPath = os.path.join(curPath, "config.yaml")
sulPath = os.path.join(curPath, "solution.json")
ansPath = os.path.join(curPath, "answer.json")

if file_path_test(yamlPath) == False:
    print("Configuration is not found.")
    exit(-1)

if file_path_test(sulPath) == False:
    print("Solution is not found.")
    exit(-1)

if file_path_test(ansPath) == False:
    print("Answer is not found.")
    exit(-1)

with open(yamlPath, 'r', encoding='utf-8') as f:
    config = f.read()
configuration = yaml.load(config,Loader=yaml.FullLoader)
problems = configuration['problems']
# print(configuration)

with open(sulPath, 'r', encoding='utf-8') as sul_file:
    sul_data = json.load(sul_file)
# print(sul_data)

with open(ansPath, 'r', encoding='utf-8') as ans_file:
    ans_data = json.load(ans_file)
# print(ans_data)

score_detail = {}
type_detail = {}
for id in range(len(problems)):
    score_detail.update({problems[id]['name']:problems[id]['max_score']})
    type_detail.update({problems[id]['name']:problems[id]['type']})

score_result = {}
for key in sul_data.keys():
    for sub_key in sul_data[key]:
        try:   
            module = getattr(autograders, type_detail[sub_key])
            func = getattr(module, type_detail[sub_key])
        except:
            print("Autograder for " + type_detail[sub_key]  + " is not found.")
            exit(-1)
        rate = func(sul_data[key][sub_key], ans_data[key][sub_key])
        score_result.update({sub_key: score_detail[sub_key] * rate})

js = json.dumps({"scores": score_result})
print(js)