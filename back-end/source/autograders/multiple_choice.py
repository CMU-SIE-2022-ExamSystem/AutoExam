def multiple_choice(solution, answer):
    for key in solution.keys():
        if len(answer[key][0]) > len(solution[key][0]):
            return 0
        else:
            tmp = []
            for i in range(len(answer[key][0])):
                if solution[key][0].find(answer[key][0][i]) == -1:
                    return 0
                else:
                    tmp.append(True)
            return sum(tmp)/len(solution[key][0])