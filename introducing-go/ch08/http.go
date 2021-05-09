package main

import "net/http"

func main() {
	http.Handle(
		"/Users/telemachus/Documents/git-repos/telemachus.me/build/",
		http.StripPrefix(
			"/Users/telemachus/Documents/git-repos/telemachus.me/build/",
			http.FileServer(http.Dir("/Users/telemachus/Documents/git-repos/telemachus.me/build/")),
		),
	)
	http.ListenAndServe(":8080", nil)
}
