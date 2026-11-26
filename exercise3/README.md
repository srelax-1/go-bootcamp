# Exercise 3: Student Grading Problem

**Labels:** Beginner | **Estimated Time:** 1 Day

## Overview
Solve the student grading problem in Go by writing both the implementation and unit tests based on the starter repository: [one2nc/student-grading-go](https://github.com/one2nc/student-grading-go).

> **Note:** This problem is used as an offline exercise during our hiring process. For the bootcamp, you do not need to submit your solution—complete the exercise and use the checklist below to self-evaluate your code.

## What You Will Learn
* Go basics & syntax
* File reading & line-by-line CSV parsing
* Structs and methods (`Stringer` interface)
* Loops, slices, and maps

---

## Problem Statement

The `grades.csv` file contains records for 30 students across various universities (with the first row acting as the header). Each record includes:
* `FirstName`: Student's first name
* `LastName`: Student's last name
* `University`: University name
* `Test1`, `Test2`, `Test3`, `Test4`: Individual test scores

### Grading Rules
1. **Final Score:** Calculated as the average of all four test scores.
   * *Example:* If `Test1=50`, `Test2=60`, `Test3=65`, and `Test4=45`, the final score is `(50 + 60 + 65 + 45) / 4 = 55`.
2. **Grade Assignment:**
   * `< 35` $\rightarrow$ **F** (Failed)
   * `>= 35` and `< 50` $\rightarrow$ **C**
   * `>= 50` and `< 70` $\rightarrow$ **B**
   * `>= 70` $\rightarrow$ **A**

### Core Tasks
Your implementation must parse `grades.csv` into a `[]student` slice and solve the following tasks (verified by the test suite in `main_test.go`):
1. Calculate each student's final score and grade.
2. Find the overall topper student.
3. Find the topper student for each university.

---

## Constraints & Requirements
* **Standard Library Only:** Do not use any third-party CSV parsing libraries. Read the file line by line and parse each row into a `Student` struct.
* **Validation:** Perform only basic validation checks on the CSV. You may assume input data types and values in the CSV are valid and well-formed.
* **Testing:** All existing and newly written test cases must pass.

---

## Self-Evaluation & Code Review Checklist

Use the guidelines below to review and refine your solution:

### Go Idioms & Architecture
- [ ] **Resource Management:** Are you closing the CSV file properly after use with a `defer` block?
- [ ] **Code Reuse:** Are you reusing `findOverallTopper` inside `findTopperPerUniversity` after grouping students by university?
- [ ] **Data Structures:** Are you using `map[string]studentStat` in the `findTopperPerUniversity` method?
- [ ] **Stringer Implementation:** Have you implemented the `String()` method (`fmt.Stringer`) for the `Student` struct so it prints cleanly with `fmt.Println`? *(Bonus / nice-to-have)*

### Naming & Style Conventions
- [ ] **Verb-Based Methods:** Do your function and method names start with verbs (e.g., `parseCSV`, `calculateGrade`, `findOverallTopper`)?
- [ ] **Concise Naming:** Are variable and method names short and idiomatic to Go?
- [ ] **No Type Names in Variables:** Have you avoided embedding data types in variable names (e.g., using `name` instead of `strName`)?
- [ ] **Left-Indentation:** Is your code left-indented as much as possible (e.g., using early returns to avoid unnecessary `else` blocks)?

### Error Handling & Testing
- [ ] **Pure Functions:** Do functions return values rather than printing directly to the console to minimize side effects?
- [ ] **Error Propagation:** Do functions follow the *handle or declare* rule by returning `(value, error)` pairs?
- [ ] **Edge-Case Testing:** For functions returning collections (arrays, slices, maps), have you written test cases covering **0**, **1**, and **many** elements?
