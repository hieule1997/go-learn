package main 

import("fmt"
		"net/http")

func Index_page(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<p>adasd</p>")
	fmt.Fprintf(w, "<a href='/about'>chuyen sang about </a>")
}
func Index_about(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>about ne</h1>")
	fmt.Fprintf(w, "<a href='/'>chuyen sang trang chu </a>")
}
func main()  {
	http.HandleFunc("/",Index_page)
	http.HandleFunc("/about",Index_about)
	http.ListenAndServe(":8000",nil)
}
