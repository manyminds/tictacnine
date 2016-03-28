package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/manyminds/tictacnine/backend/game"
)

type server struct {
	handler http.Handler
}

type data struct {
	sync.Mutex
	player game.Player
	g      game.Game
}

var dataContainer = data{}

func initGame(w http.ResponseWriter, r *http.Request) {
	dataContainer.Lock()
	defer dataContainer.Unlock()
	strength := r.FormValue("aiStrength")
	level, err := strconv.Atoi(strength)
	if err != nil {
		level = 5
	}

	log.Printf("Initializing AI with Level: %d\n", level)

	dataContainer.g = game.NewTicTacNineGame()
	dataContainer.player = game.NewAIPlayer(game.MoveCircle, level)
	data := map[string]interface{}{
		"success": true,
		"level":   level,
	}
	respondWith(data, w)
}

func getTurn(w http.ResponseWriter, r *http.Request) {
	dataContainer.Lock()
	defer dataContainer.Unlock()
	data := map[string]interface{}{
		"nextTurn": dataContainer.g.Board().GetNextTurn().String(),
	}
	respondWith(data, w)
}

func getAiMove(w http.ResponseWriter, r *http.Request) {
	dataContainer.Lock()
	defer dataContainer.Unlock()
	m := dataContainer.player.NextMove(dataContainer.g.Board())
	fx, fy := m.Field()
	px, py := m.Position()
	dataContainer.g.Board().PutStone(fx, fy, px, py, game.MoveCircle)

	result := map[string]interface{}{}

	if dataContainer.g.Board().HasWinner() {
		result["winner"] = dataContainer.g.Board().GetWinner().String()
	}

	result["x"] = fx*3 + px
	result["y"] = fy*3 + py

	respondWith(result, w)
}

func putStone(w http.ResponseWriter, r *http.Request) {
	dataContainer.Lock()
	defer dataContainer.Unlock()
	x := r.FormValue("x")
	posX, err := strconv.Atoi(x)
	if err != nil {
		respondWith(map[string]interface{}{"error": "x not int"}, w)
		return
	}

	y := r.FormValue("y")
	posY, err := strconv.Atoi(y)
	if err != nil {
		respondWith(map[string]interface{}{"error": "y not int"}, w)
		return
	}

	moveToPlay := game.MoveCross
	move := r.FormValue("move")
	if move == "o" {
		moveToPlay = game.MoveCircle
	}

	tx := posX % 3
	ty := posY % 3
	fx := posX / 3
	fy := posY / 3

	dataContainer.g.Board().PutStone(fx, fy, tx, ty, moveToPlay)
}

func respondWith(data interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//NewServer returns a new configured http handler
func NewServer(dist string) http.Handler {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(dist))
	mux.Handle("/", fileHandler)
	mux.HandleFunc("/ai/init", initGame)
	mux.HandleFunc("/ai/getTurn", getTurn)
	mux.HandleFunc("/ai/putStone", putStone)
	mux.HandleFunc("/ai/getStone", getAiMove)
	s := server{}
	s.handler = mux
	return &s
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}
