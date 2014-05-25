package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/paulsmith/gogeos/geos"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	serveHttp()
}

func testCoordinatePresence(x, y, polygyon string) bool {
	fx, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return false
	}

	fy, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return false
	}

	point, err := geos.NewPoint(geos.NewCoord(fx, fy))
	if err != nil {
		return false
	}

	boundary, err := geos.FromWKT(polygyon)
	if err != nil {
		return false
	}

	result, err := boundary.Contains(point)
	if err != nil {
		return false
	}

	return result
}

func serveHttp() {
	m := martini.Classic()
	m.Use(render.Renderer())

	// UI Calls
	m.Get("/", homeHandler)
	m.Post("/", testCoordinatesHandler)
	m.NotFound(notFoundHandler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), m); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("home").ParseFiles("views/layout.html", "views/home/index.html"))
	t.ExecuteTemplate(w, "layout", "Home")
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("404").ParseFiles("views/layout.html", "views/404.html"))
	t.ExecuteTemplate(w, "layout", "Whoops!")
}

func testCoordinatesHandler(req *http.Request, r render.Render) {
	req.ParseForm()
	x := req.FormValue("point-x")
	y := req.FormValue("point-y")
	polygon := req.FormValue("polygon-wkt")
	contained := testCoordinatePresence(x, y, polygon)

	r.JSON(http.StatusOK,
		map[string]interface{}{
			"contained": contained,
			"timestamp": time.Now().UTC().Unix(),
		})
}
