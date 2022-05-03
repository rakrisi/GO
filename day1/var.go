package main
import "fmt"
import "math"
//var name string ="radhe"
func add (x , y ,z int) (a,b int) {
	a=x+y
	b=z
	return
}

func main(){
	action := func() {
		const ii =12
		ss := fmt.Sprintf("jello 6")
		name,age := " radhe", 12
		//age=17
		class := int8(12.000)
		a,b := add(5,3,6)
		fmt.Println("radhe ", name ,age , class,ii , ss ,math.Pi , a,b)
	}
	action()

}