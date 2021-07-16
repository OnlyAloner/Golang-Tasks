package main

type Task struct {
	ID int
	Assignee,
	Title,
	Deadline string
	Done bool
}

type TaskManager struct {
	tasks []Task
}

func NewTaskList() *TaskManager {
	Temp := new(TaskManager)
	Temp.tasks = make([]Task, 0)
	return Temp

}

func (This *TaskManager) create(newTask Task) *Task {

	for ind, elem := range This.tasks {
		if elem.ID == newTask.ID {
			return &This.tasks[ind]
		}
	}

	This.tasks = append(This.tasks, newTask)
	return &This.tasks[len(This.tasks)-1]
	//we returning the pointer to element in list that equal to newTask
}

func (This *TaskManager) update(updTask Task) *Task {
	for ind, elem := range This.tasks {
		if elem.ID == updTask.ID {

			This.tasks[ind] = updTask
			return &This.tasks[ind]

		}
	}

	return &updTask // if we find the element we change him and return pointer to him, otherwise we return pointer to updTask
}

func (This *TaskManager) delete(delTask Task) *Task {
	for ind, elem := range This.tasks {
		if elem.ID == delTask.ID {

			This.tasks = append(This.tasks[:ind], This.tasks[ind+1:]...)
			return &delTask

		}
	}

	return &delTask // If task not finded or he was deleted we will return the pointer to delTask
}

func (This *TaskManager) MakeDone(ID int) *Task {
	for ind, elem := range This.tasks {
		if elem.ID == ID {

			This.tasks[ind].Done = true
			return &This.tasks[ind]

		}
	}

	return &Task{} // If we find we return the pointer to element, otherwise pointer to empty Task
}

func (This *TaskManager) Get(ID int) *Task {
	for ind, elem := range This.tasks {
		if elem.ID == ID {

			return &This.tasks[ind]

		}
	}

	return &Task{} // If we find we return the pointer to element, otherwise pointer to empty Task
}

func (This *TaskManager) GetAll() *[]Task {
	return &This.tasks
}
