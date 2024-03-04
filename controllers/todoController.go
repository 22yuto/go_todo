package controllers

import (
	"encoding/json"
	"go_todo_api/models"
	"net/http"
)

type TodoController struct {
	Model *models.TodoModel
}

// 初期化・インスタンス生成関数
// TodoModelを持つTodoControllerのインスタンスを作成しそのポインタを返却
func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{Model: m}
}

// (h *TodoController) レシーバー...構造体に関連する関数を定義
func (h *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// models/todo.goのAll関数を使ってデータ取得
	todos, err := h.Model.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo // models.Todo型
	// リクエストからJSONを読み取り、todo変数にデコード
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// models/todo.goのInsert関数を使って挿入
	id, err := h.Model.Insert(todo.Task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo.ID = id
	json.NewEncoder(w).Encode(todo)
}
