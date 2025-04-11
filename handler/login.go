package handler

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func getLogin(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postLogint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "無効なリクエストメソッドです", http.StatusMethodNotAllowed)
		return
	}

	// フォームの値を取得
	name := r.FormValue("name")
	message := r.FormValue("message")

	// 確認のためログに出力
	log.Printf("受け取ったデータ - 名前: %s, メッセージ: %s\n", name, message)

	// レスポンスとして表示
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := "<h1>フォームの送信が完了しました！</h1>" +
		"<p>名前: " + template.HTMLEscapeString(name) + "</p>" +
		"<p>メッセージ: " + template.HTMLEscapeString(message) + "</p>" +
		`<p><a href="/">戻る</a></p>` +
		`<p>token:test1234</p>`

	w.Write([]byte(response))
}
