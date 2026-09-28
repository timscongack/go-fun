package main

import "fmt"

/* This is the first program I wrote in go, just implementing
the battleship search_destroy function as a solver first.
*/


var grid = [10][10]byte{
      {'.', '.', '.', 'C', 'C', 'C', 'C', 'C', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', 'B', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', 'B', '.', '.', '.', '.', '.', 'S', 'S', 'S'},
      {'.', 'B', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', 'B', '.', '.', 'R', 'R', 'R', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', 'D', 'D', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
}


var my_grid = [10][10]byte{
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
      {'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
}

func main()  {



}

func search_destroy(x int, y int, grid byte) {

  switch {
    // 46 is .
    case  byte[x][y] <> 46:
      fmt.Println("Hit! " + string(grid[x][y]))
      //update fog
      //pic next hit candidate
      //search_destroy call
    case bytep[x][y] = 46:
      //update fog
      //get next coord (does binary search)

}

funct update_fog(x int, y int, grid byte) {


}

func get_next_coord() {
//find candidate coords and returns 
//logic on grid checking
}
