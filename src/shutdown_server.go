package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// ポート番号とシグナルファイルのパスを受け取る
	port := flag.String("port", "11451", "Port number to listen")
	signalFile := flag.String("signal-file", "/tmp/shutdown_signal", "Path to the shutdown signal file")
	flag.Parse()

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		// シグナルファイルを作成してBashに通知
		if err := os.WriteFile(*signalFile, []byte("shutdown"), 0644); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("signal fileの作成に失敗しました😞"))
			fmt.Println("signal fileの作成に失敗しました😞")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("シャットダウンします😊"))
		fmt.Println("シャットダウンします😊")
	})
	http.ListenAndServe(":"+*port, nil)
}
