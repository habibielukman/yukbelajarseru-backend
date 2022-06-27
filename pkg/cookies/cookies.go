package cookies

import "net/http"

func DeleteCookie(w http.ResponseWriter, r *http.Request, cookieName string) {
	c, err := r.Cookie("authenticated_id")
	if err != nil {
		println("keluar tapi belum masuk")
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
}
func CloseCookie(c *http.Cookie) *http.Cookie {
	return c
}
