def multiple_choice(solution, answer):
    for key in solution.keys():
        if len(answer[key]) > len(solution[key]):
            return 0.0
        else:
            tmp = []
            for i in range(len(answer[key])):
                if solution[key].find(answer[key][i]) == -1:
                    return 0.0
                else:
                    tmp.append(True)
            return float(sum(tmp)/len(solution[key]))