# Introduction to the customized autograder
## Pre-requirement
- A *.py* file that will be able to run on **python 3.9** environment.
- The **grader name**, **file name**, and the **'main' function name** should be exactly the same. This name should follow the rule of the module name in python like *customized_grader*.
## Input and Output
- The input is two *python dictionaries*, structured as follows. The number of keys in the dictionary depends on how many blanks this grader has. For example, if it is a grader for a single-blank question, there will be only one key; and if it is a grader for a multi-blank question, there will be multiple keys. 
```json
{
    "blank1": ["string"],
    "blank2": ["string"],
    "blank3": ["string"]
}
```
- The first input dictionary is **solution**. The value corresponding to each key is an array of strings. Each element of the string array is the solution to this blank, and the number of elements depends on the number of solutions. Usually solutions will be input by the instructor when creating the question.
- The second input dictionary is **answer**. The value corresponding to each key is an array of strings. Each string array will have only one element, which is the answer submitted by the student.
- The return value of the 'main' function should be **greater than or equal to 0.0** and **less than or equal to 1.0**. A return value of 1 means full score, 0 means no score, and intermediate values mean partial score.
```python
def customized_grader(solution, answer):
    # operations
    # ...
    return 1.0
```

## Validation
- Testing the grader in the Autoexam system requires entering answer and solution.
- When doing the grader test in the Autoexam system, *print* some information can help you judge whether the logic of the grader you wrote is correct.