package main

import "net/http"

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		proto := r.Proto
		method := r.Method
		uri := r.RequestURI

		app.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

		next.ServeHTTP(w, r)
	})
}

// commonHeaders is a middleware that sets some useful common headers to each request.
func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CSP headers are used to restrict where the resources for your web page (e.g. JavaScript, images, fonts etc) can be loaded from.
		// Setting a strict CSP policy helps prevent a variety of cross-site scripting, clickjacking, and other code-injection attacks.
		// In our case, the header tells the browser that it’s OK to load fonts from fonts.gstatic.com, stylesheets from fonts.googleapis.com
		// and self (our own origin), and then everything else only from self. Inline JavaScript is blocked by default.
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		// "Referrer-Policy" is used to control what information is included in a Referer header when a user navigates away from your web page.
		// In our case, we’ll set the value to "origin-when-cross-origin", which means that the full URL will be included for same-origin requests,
		// but for all other requests information like the URL path and any query string values will be stripped out.
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// "nosniff" instructs browsers to not MIME-type sniff the content-type of the response, which in turn helps to prevent content-sniffing attacks.
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// "deny" is used to help prevent clickjacking attacks in older browsers that don’t support CSP headers.
		w.Header().Set("X-Frame-Options", "deny")

		// is used to disable the blocking of cross-site scripting attacks. Previously it was good practice to set this header to
		// X-XSS-Protection: 1; mode=block, but when you’re using CSP headers like we are the recommendation is to disable it.
		w.Header().Set("X-XSS-Protection", "0")

		// A custom header.
		w.Header().Add("Server", "Go")

		// Execute the next handler.
		next.ServeHTTP(w, r)
	})
}
