package main
import "testing"
func TestSums(t *testing.T){
	/*total := Sum(10,10)
	if total != 20{
		t.Errorf("Sum wad incorrect, got %v, want %v", total, 20)
	}*/
	tables := []struct {
		a int 
		b int
		n int
	}{
		{1,2,3},
		{25,23,48},
		{2,-1,1},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum wad incorrect, got %v, want %v", total, item.n)
		}
	}
}

func TestMax(t *testing.T){
	tables := []struct{
		a int 
		b int 
		n int
	}{
		{4,1,4},
		{5,1,5},
		{10,2,10},
		{1,10,10},
	}
	for _,item:=range tables{
		max:=GetMax(item.a,item.b)
		if max!=item.n{
			t.Errorf("Max was incorrect, got %v, want %v",item.n,max)
		}
	}
}

func TestFib(t *testing.T){
	tables := []struct{
		a int
		n int
	}{
		{1,1},
		{8,21},
		{50,12586269025},
	}
	for _,item :=range tables {
		fib:=Fibonacci(item.a)
		if fib!= item.n {
			t.Errorf("Fibonacci was incorrect, got %v, want %v",fib,item.n)
		}
	}
}