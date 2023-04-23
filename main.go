package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Transfer-Encoding", "chunked")
		w.WriteHeader(http.StatusOK)

		content := `<html>
<head>
	<title>Horizon</title>
</head>
<body>
	<h1>Hello World</h1>
</body>
</html>`

		w.Write([]byte(content))
	})

	http.ListenAndServe(":8080", nil)
}