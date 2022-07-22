def multiple_blank(solution, answer):
    tmp = []
    for key in solution.keys():
        lower_str1 = answer[key][0].lower()
        lower_str2 = solution[key][0].lower()
        if lower_str1 == lower_str2:
            tmp.append(True)
        else:
            tmp.append(False)
    return sum(tmp)/len(tmp)