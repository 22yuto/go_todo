package models

import (
	"database/sql"
)

// Todo構造体
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// database/sqlパッケージのDB型
type TodoModel struct {
	DB *sql.DB
}

// 新規TodoModelのインスタンスを作成
func NewTodoModel(DB *sql.DB) *TodoModel {
	// TodoModel型のポインタを返却...初期設定を行い同じインスタンスを共有する
	return &TodoModel{DB: DB}
}

// 全件取得
func (m *TodoModel) All() ([]Todo, error) {
	rows, err := m.DB.Query("SELECT id, task FROM todos")

	// errにエラーがあるかチェック エラーの場合nilとエラーを返却
	if err != nil {
		return nil, err
	}

	// DBのクエリ結果を閉じる...後続の処理をスムーズにするため
	defer rows.Close()

	// スライス
	var todos []Todo

	// Next()...クエリの結果を1行ずつループ処理
	for rows.Next() {
		var todo Todo

		// 現在の行のidとtaskの値をtodoにスキャン
		// Scan...クエリ結果のカラムの値を変数に割り当てるメソッド
		if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
			// スキャン中にエラーが発生したらnilとエラーを返す
			return nil, err
		}
		// todosスライスにtodo追加
		todos = append(todos, todo)
	}

	return todos, nil
}

// 新規作成
func (m *TodoModel) Insert(task string) (int, error) {
	// taskをinsert
	result, err := m.DB.Exec("INSERT INTO todos (task) VALUES (?)", task)
	if err != nil {
		return 0, err
	}
	// 直近に挿入された行のIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
