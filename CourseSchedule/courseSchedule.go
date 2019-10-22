package main

func CanFinish(numCourses int, prerequisites [][]int) bool {
	a := make(map[int][]int)
	dep := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		preCourse := prerequisites[i][1]
		dep[course]++

		if _, ok := a[preCourse]; ok {
			a[preCourse] = append(a[preCourse], course)
		} else {
			a[preCourse] = append([]int{}, course)
		}
	}

	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if dep[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		currCourse := q[0]
		nextCourses := a[currCourse]

		for i := 0; i < len(nextCourses); i++ {
			dep[nextCourses[i]]--

			if dep[nextCourses[i]] == 0 {
				q = append(q, nextCourses[i])
			}
		}
		q = q[1:]
	}

	for i := 0; i < numCourses; i++ {
		if dep[i] != 0 {
			return false
		}
	}

	return true
}
