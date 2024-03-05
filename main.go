package mai

import (
	"database/sql"
	"fmt"
	"go_todo_api/controllers"
	"go_todo_api/models"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// モデルとコントローラのインスタンス作成
// ルーターの設定、サーバー起動
func main() {
	// MySQL接続開始
	db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// NewTodoModelにDB接続オブジェクトを渡す
	todoModel := models.NewTodoModel(db)
	// HTTPリクエストに応じてTodoの操作を制御するコントローラ生成
	todoHandler := controllers.NewTodoController(todoModel)

	// HTTPリクエストのルーティング生成
	router := mux.NewRouter()

	// /todosパスに対するルータに関連付け
	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

	// port8080でWebサーバを起動
	fmt.Println("Server starting at :8080")
	// エラーはプログラム終了
	log.Fatal(http.ListenAndServe(":8080", router))
}
