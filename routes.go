package main

import (
	"log"
	"encoding/json"
	"net/http"
	. "github.com/jstolp/pofadder-go/api"
	"fmt"
	"math"
	"strconv"
)

var leftBound int = 1;
var topBound int = 1;
var rightBound int = 0;
var botBound int = 0;

var edgeSnakeLimit int = 0;
var turn int = 0;
var move string = "down";
var nextMove string = "";
var prevMove string = "";
var headPos Coord;
var tailPos Coord;
var health int = 100;
var numOfStartingSnakes int = 1;
var numSnakesLeft int = 1;
var enemySnakes int = 0;
var foodPointList []Coord;
var endCicle bool = false;

/* heads: "beluga" "bendr" "dead" "evil" "fang" "pixel" "regular" "safe" "sand-worm" "shades" "silly" "smile" "tongue"
tails: "block-bum" "bolt" "curled" "fat-rattle" "freckled" "hook" "pixel" "regular" "round-bum" "sharp" "skinny" "small-rattle" */

func Start(res http.ResponseWriter, req *http.Request) {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	decoded := SnakeRequest{}
	err := DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

  headPos = decoded.You.Body[0]
  //foodPointList := decoded.Board.Food
  numOfStartingSnakes = len(decoded.Board.Snakes)
  topBound, leftBound = 1, 1; // Set NW bound, X, Y
	botBound, rightBound = decoded.Board.Height, decoded.Board.Width // SE corner X, Y
	
	log.Print("BOARD Size: TOP LEFT  NW Corner x:" + strconv.Itoa(topBound) + " , " + strconv.Itoa(leftBound))
	log.Print("BOARD Size: BOT RIGHT SE Corner x:" + strconv.Itoa(botBound) + "," + strconv.Itoa(rightBound))
	log.Print("Enemy Snakes: " + strconv.Itoa(numOfStartingSnakes - 1) + "\n\n")

	log.Print("Start Pos: " + strconv.Itoa(headPos.X) + "," + strconv.Itoa(headPos.Y))
	if(numOfStartingSnakes == 1) {
		log.Print("\n\n It's Gonna be a SOLO GAME \n")
	}
	/*
	 e19c41 - orange test 2
   00ff55 - green
   ff4f00 - orange test 1 -nee te rood
	*/
	respond(res, StartResponse{
		Color: "#e19c41",
		HeadType: "tongue",
		TailType: "curled",
	})
}


func Move(res http.ResponseWriter, req *http.Request) {
	nextMove = prevMove
	decoded := SnakeRequest{}
	err := DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	me := decoded.You
	headPos := getHeadPos(me)
	foodPointList = decoded.Board.Food
	health = me.Health
	myLength := len(me.Body)
	numSnakesLeft = len(decoded.Board.Snakes)
	enemySnakes = numSnakesLeft - 1
	turn = decoded.Turn

	respond(res, MoveResponse{
		Move: nextMove,
	})
}

// see if i can attach these methods to the struct Snake or something..
// func (target Snake) Head() Coord { return target.Body[0] }

func getHeadPos(target Snake) Coord {
	body := target.Body
  return body[0]
}
func getTailPos(target Snake) Coord {
	body := target.Body
  return body[len(body) - 1]
}


/* Dist to function in steps (int) */
func dist(a Coord, b Coord) int {
	return int(math.Abs(float64(b.X-a.X)) + math.Abs(float64(b.Y-a.Y)))
}

func GetBodies(snakes SnakesList) []Coord {
  list := make([]Coord, 0)
  for _, s := range snakes {
    list = append(list, s.Body...)
  }
  return list
}

// closestFoodPoint
func minDistFood(headPos Coord, food []Coord) Coord {
	min := food[0]
	for _, f := range food {
		if dist(min, headPos) < dist(f, headPos) {
			min = f
		}
	}
	return min
}

// just a testing function to dump a object../
func dd(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}

func Ping(res http.ResponseWriter, req *http.Request) {
	log.Print("PONG to a server ping... \n")
	return
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func End(res http.ResponseWriter, req *http.Request) {
	log.Print("The game has ended.... \n\n")
	return
}
