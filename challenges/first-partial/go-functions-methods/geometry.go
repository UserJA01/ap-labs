package main

import "math"
import "fmt"
import "os"
import "strconv"
import "math/rand"
import "time"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

type Line struct{p1, p2 Point}

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
func Min(x, y float64) float64 {
    if x < y {
    return x
    }
    return y
}

func Max(x, y float64) float64 {
    if x > y {
    return x
    }
    return y
}

func WithLine(a Line,p Point) bool {   //check whether p is on the line or not
   if(p.X <= Max(a.p1.X, a.p2.X) && p.X <= Min(a.p1.X, a.p2.X) &&
      (p.Y <= Max(a.p1.Y, a.p2.Y) && p.Y <= Min(a.p1.Y, a.p2.Y))){
        return true
    }
   return false
}

func direction(a,b,c Point) int{
    val := (b.Y - a.Y)*(c.X-b.X)-(b.X-a.X)*(c.Y-b.Y)
    
    if(val == 0){
        return 0
    }else if(val < 0){
        return 2
    }
    
    return 1
    //return 15
}

func intersects(a,b Line) bool{
    dir1 := direction(a.p1, a.p2, b.p1)
    dir2 := direction(a.p1, a.p2, b.p2)
    dir3 := direction(b.p1, b.p2, a.p1)
    dir4 := direction(b.p1, b.p2, a.p2)
    
	/*
    fmt.Println("1: ", dir1)
    fmt.Println("1: ", dir2)
    fmt.Println("1: ", dir3)
    fmt.Println("1: ", dir4)
    */
    
    if dir1 != dir2 && dir3 != dir4{
      return true;
    }
    if dir1==0 && WithLine(a, b.p1){
      return true;
    }
    if dir2==0 && WithLine(a, b.p2){
      return true;
    }
    if dir3==0 && WithLine(b, a.p1) {
      return true;
    }
    if dir4==0 && WithLine(b, a.p2){
      return true; 
    }
    return false;
}

func main() {
    //Path[0] = Point{0, 0}
    //Path[1] = Point{5, 5}
    rand.Seed(time.Now().UnixNano())
    Path := []Point{}
    sides, err := strconv.Atoi(os.Args[1])
    if err == nil{}

	if sides < 3{
		fmt.Println("You need a minimum of 3 sides for a figure")
		//return 0

		os.Exit(3)
	}
    for x := 0; x < sides; x++ {
        
        r := -100 + rand.Float64()*(100+100)
        r2 := -100 + rand.Float64()*(100+100)
        //fmt.Println("#", x , " Random " , r , " HAH ", r2)
        //fmt.Println("#", x , " Random " , r , " ", r2)
        n := Point{X: r, Y: r2}
        if(x >= 2){
            a := Line{p1: Path[x-1], p2: Path[x-2]}
            for WithLine(a, n){
                r = -100 + rand.Float64()*(100+100)
                r2 = -100 + rand.Float64()*(100+100)
                
                n = Point{X: r, Y: r2}
            }            
        }

        if(x >= 3){
            flag := true
            for flag {
                t1 := Line{p1: Path[x-1], p2: n}
                current := Line{p1: Path[x-2], p2: Path[x-3]}
                for y := 0; y < sides-3; y ++ {
                    current = Line{p1: Path[x-(y+2)],p2: Path[x-(y+3)]}
                    if intersects(t1, current){
                        r = -100 + rand.Float64()*(100+100)
                        r2 = -100 + rand.Float64()*(100+100) 
                        n = Point{X: r, Y: r2}
                        continue
                    }
                }
                flag = false
            }
        }

        //fmt.Println("#",x," Random " , r , " HAH ", r2)
        //fmt.Println("#", x , " Random " , r , " ", r2)
        Path = append(Path, n)
    }
    /*
	go run geometry.go 4
	- Generating a [4] sides figure
	- Figure's vertices
		- (  -3,   1)
		- (   2,   3)
		- (   0,   0)
		- (-1.5,-2.5)
	- Figure's Perimeter
   		- 5.38 + 3.60 + 3.92 + 3.80 = 16.70
	*/
    
    //fmt.Println("Num ")
	fmt.Println("- Generating a [" , sides, "] sides figure")
	fmt.Println("- Figure's vertices")
	for x:= 0; x < sides; x++ {
		/*k := Path[x].X
		q := Path[x].Y
		fmt.Printf("%.2f", k , " %.2f\n", q)*/
		k := Path[x].X
		fmt.Printf("	- ( %.2f,", k)
		fmt.Printf(" %.2f)\n", k)
	}
    fmt.Println("- Figure's Perimeter")
    sum := 0.0
	fmt.Print("	- ")
    for x := 1; x < sides; x++ {
		temp := Path[x-1].Distance(Path[x])
        sum += temp
		
		fmt.Printf("%.2f + ", temp)
		if(x < sides-1){
			//fmt.Printf("	- %.2f + ", temp)
			//fmt.Printf("%.2f + ", temp)
		}else{
			//fmt.Print(" = ")
			//fmt.Printf("%.2f", temp)
		}
    }
	temp := Path[0].Distance(Path[sides-1])
	sum += temp
	fmt.Printf("%.2f", temp)

    fmt.Println(" = ", sum)


    
}
