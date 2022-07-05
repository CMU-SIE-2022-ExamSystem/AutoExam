def single_blank(answer, solution):
    for key in solution.keys():
        lower_str1 = answer[key].lower()
        lower_str2 = solution[key].lower()
        if lower_str1 == lower_str2:
            return 1.0
        else:
            return 0.0