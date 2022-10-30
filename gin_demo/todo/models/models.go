package models

import "strings"

var todoList = []*Todo{
	{1, "todo1", "早上8点吃早餐", 1666973280, 1666973310, 1666974093, 1666974093},
	{2, "todo2", "中午12点吃午饭", 1667016000, 1667017800, 1666974093, 1666974093},
	{3, "todo3", "下午6点吃晚饭", 1667037600, 1667039400, 1666974093, 1666974093},
	{4, "todo4", "晚上8点去健身", 1667044800, 1667046600, 1666974093, 1666974093},
}

type Todo struct {
	Id         int64
	Title      string
	Content    string
	StartTime  int64
	EndTime    int64
	CreateTime int64
	UpdateTime int64
}

func GetList(limit, offset int, title, content string, startTimeStart, startTimeEnd, endTimeStart, endTimeEnd int64) []*Todo {
	var result []*Todo
	for _, todo := range todoList {
		titleCompare := true
		contentCompare := true
		startTimeStartCompare := true
		startTimeEndCompare := true
		endTimeStartCompare := true
		endTimeEndCompare := true
		if title != "" {
			titleCompare = strings.Contains(todo.Title, title)
		}
		if content != "" {
			contentCompare = strings.Contains(todo.Content, content)
		}
		if startTimeStart != 0 {
			startTimeStartCompare = todo.StartTime > startTimeStart
		}
		if startTimeEnd != 0 {
			startTimeEndCompare = todo.StartTime < startTimeEnd
		}
		if endTimeStart != 0 {
			startTimeStartCompare = todo.EndTime > endTimeStart
		}
		if endTimeEnd != 0 {
			startTimeEndCompare = todo.EndTime < endTimeEnd
		}

		if titleCompare &&
			contentCompare &&
			startTimeStartCompare &&
			startTimeEndCompare &&
			endTimeStartCompare &&
			endTimeEndCompare {
			result = append(result, todo)
		}
	}
	return result[offset : offset+limit]
}

func GetById(id int64) *Todo {
	for _, todo := range todoList {
		if todo.Id != id {
			continue
		}
		return todo
	}
	return nil
}

func Add(todo *Todo) {
	todoList = append(todoList, todo)
}

func EditById(id int64, todoNew *Todo) {
	for _, todo := range todoList {
		if todo.Id != id {
			continue
		}
		if todoNew.Title != "" {
			todo.Title = todoNew.Title
		}
		if todoNew.Content != "" {
			todo.Content = todoNew.Content
		}
		if todoNew.StartTime != 0 {
			todo.StartTime = todoNew.StartTime
		}
		if todoNew.EndTime != 0 {
			todo.EndTime = todoNew.EndTime
		}
	}
}

func DeleteById(id int64) {
	var deleteIndex int
	for i, todo := range todoList {
		if todo.Id != id {
			continue
		}
		deleteIndex = i
		break
	}
	todoList = append(todoList[:deleteIndex], todoList[deleteIndex+1:]...)
}
