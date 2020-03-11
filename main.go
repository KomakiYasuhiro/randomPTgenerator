package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	_ "myrepo/pokemonPT/statik"

	"github.com/rakyll/statik/fs"
)

// Pokemon struct is
// Name : ポケモン名
// No : 図鑑番号
// Evoling : 進化するかどうか
type Pokemon struct {
	Name    string `json:"name"`
	No      int    `json:"no"`
	Evoling bool   `json:"evoling"`
}

func allKeys(m map[int]bool) []int {
	i := 0
	result := make([]int, len(m))
	for key := range m {
		result[i] = key
		i++
	}
	return result
}

func pickup(min int, max int, num int) []int {
	numRange := max - min

	selected := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	for counter := 0; counter < num; {
		n := rand.Intn(numRange) + min
		if selected[n] == false {
			selected[n] = true
			counter++
		}
	}
	keys := allKeys(selected)
	// ソートしたくない場合は以下1行をコメントアウト
	//sort.Sort(sort.IntSlice(keys))
	return keys
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	f, err := statikFS.Open("/pokemon.json")
	if err != nil {
		panic(err)
	}

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//ポケモンjsonのデータが入る配列
	var pokemon []Pokemon
	json.Unmarshal(raw, &pokemon)
	//6体をランダムで抽出
	noList := pickup(0, 400, 6)
	for i, no := range noList {
		fmt.Printf("%d体目\n 図鑑No : %d\n ポケモン名 : %s\n", i+1, pokemon[no].No, pokemon[no].Name)
	}
}
