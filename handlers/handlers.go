package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/ereminiu/spb-subway/dijkstra"
	. "github.com/ereminiu/spb-subway/models"
	"github.com/ereminiu/spb-subway/utils"
	"github.com/gin-gonic/gin"
)

var n int // amount of vertices

var gr [][]*Edge
var stations map[string]Station
var names []string
var idx map[string]int
var path []string

// min - duration of optimal path, route - stations of this path
var mincost int
var route []int

func FindRouteHandler(c *gin.Context) {
	from, to := getIdx(utils.EnToRu(c.Request.FormValue("from"))), getIdx(utils.EnToRu(c.Request.FormValue("to")))

	mincost, route = dijkstra.Dijkstra(gr, from, to)
	mincost += 2

	c.Redirect(http.StatusFound, "/showroute")
}

func getIdx(s string) int {
	mn := int(1e9 + 228)
	closest := ""

	for _, name := range names {
		dist := utils.EditDistance(s, name)
		if dist < mn {
			closest = name
			mn = dist
		}
	}

	if idx[closest] == 0 {
		panic("Can't find name " + s)
	}

	return idx[closest] - 1
}

func ShowRouteHandler(c *gin.Context) {
	inRoute := make(map[string]bool)
	for _, i := range route {
		inRoute[names[i]] = true
	}

	normalStations, transStations := make(map[string]Station), make(map[string]Station)
	normalLines, transLines := make([]Line, 0), make([]Line, 0)
	normalLables, transLables := make([]Lable, 0), make([]Lable, 0)

	for st_id, from := range stations {
		// add lables for the stations
		lable := Lable{X: from.X + from.Dx, Y: from.Y + from.Dy, Name: from.Name}
		switch from.Shift {
		case "left":
			lable.X -= 12

		case "right":
			lable.X += 65

		case "down":
			lable.X += 10
			lable.Y += 10

		case "up":
			lable.X += 10
			lable.Y -= 10
		}

		// if station not in route, it should be transparent
		if inRoute[from.Name] {
			normalStations[st_id] = from
			normalLables = append(normalLables, lable)
		} else {
			transStations[st_id] = from
			transLables = append(transLables, lable)
		}

		for _, next := range from.Adj {
			to := stations[next]
			line := Line{X1: from.X, X2: to.X, Y1: from.Y, Y2: to.Y, Color: from.Line}
			if inRoute[from.Name] && inRoute[to.Name] {
				normalLines = append(normalLines, line)
			} else {
				transLines = append(transLines, line)
			}
		}
	}

	fmt.Println(normalLables)
	fmt.Println(transLables)

	c.HTML(http.StatusOK, "route.tmpl", gin.H{
		"mincost":   mincost,
		"nstations": normalStations,
		"tstations": transStations,
		"nlines":    normalLines,
		"tlines":    transLines,
		"nlables":   normalLables,
		"tlables":   transLables,
	})
}

func HomeHandler(c *gin.Context) {
	// получаем станции из json
	data, err := ioutil.ReadFile("stations.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(data, &stations)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(stations)

	// строим граф из станций
	lines := make([]Line, 0)
	lables := make([]Lable, 0)
	names = make([]string, 0)
	idx = make(map[string]int)

	// get names of stations
	for _, from := range stations {
		names = append(names, from.Name)
	}

	// sort names
	sort.Strings(names)
	for i, name := range names {
		idx[name] = i + 1
	}
	fmt.Println(idx)

	// init gr
	n = len(names)
	gr = make([][]*Edge, n)
	for i := 0; i < n; i++ {
		gr[i] = make([]*Edge, 0)
	}

	for _, from := range stations {
		// add text shift
		switch from.Shift {
		case "left":
			lables = append(lables, Lable{X: from.X - 12, Y: from.Y, Name: from.Name})

		case "right":
			lables = append(lables, Lable{X: from.X + 65, Y: from.Y, Name: from.Name})

		case "down":
			lables = append(lables, Lable{X: from.X + 10, Y: from.Y + 10, Name: from.Name})

		case "up":
			lables = append(lables, Lable{X: from.X + 10, Y: from.Y - 10, Name: from.Name})
		}

		lables[len(lables)-1].X += from.Dx
		lables[len(lables)-1].Y += from.Dy

		// build lines between stations
		// and init graph
		for i, next := range from.Adj {
			to := stations[next]
			lines = append(lines, Line{X1: from.X, X2: to.X, Y1: from.Y, Y2: to.Y, Color: from.Line})

			// add edge between From and To
			u, v := idx[from.Name]-1, idx[to.Name]-1
			gr[u] = append(gr[u], NewEdge(v, from.Dist[i]))
			gr[v] = append(gr[v], NewEdge(u, from.Dist[i]))
		}
	}

	// отрисовываем метро
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"stations": stations,
		"lines":    lines,
		"lables":   lables,
	})
}
