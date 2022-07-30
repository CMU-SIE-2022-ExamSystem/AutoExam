def single_choice(solution, answer):
    for key in solution.keys():
        lower_str1 = answer[key][0].lower()
        for i in range(len(solution[key])):
            lower_str2 = solution[key][i].lower()
            if lower_str1 == lower_str2:
                return 1.0
        return 0.0