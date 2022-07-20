def single_choice(answer, solution):
    for key in solution.keys():
        lower_str1 = answer[key][0].lower()
        lower_str2 = solution[key][0].lower()
        if lower_str1 == lower_str2:
            return 1.0
        else:
            return 0.0