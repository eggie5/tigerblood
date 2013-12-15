package main

import "fmt"
import "./queue"

func main() {

	adj := map[string][]string{
	    "r": {"s", "v"},
		"s": {"r", "w"},
		"t": {"u", "w", "x"},
		"u": {"t", "y"},
		"v": {"r"},
	    "w": {"s", "t", "x"},
		"x": {"t", "w", "y"},
		"y": {"u", "x"},
	}
	

	fmt.Println(adj)
	
	pi := bfs(adj, "s");
	
	print_path(pi, "s", "y")

}

func bfs(G map[string][]string, s string) (_pi map[string]string) {
	
	color := map[string]string{}
	d := map[string]int{} //distance
	pi := map[string]string{} //parents
	
	// color all verts white
	for _, v := range G { 
		// fmt.Println(k)
		for _, val := range v {
			color[val] = "w"
			d[val]=-1
			// pi[val]=
		}
	 }
	 
	 //start w/ s
	 color[s]="g"
	 d[s]=0
	 //pi[s]=-1 //source has no predecessor
	 
	 q := queue.New()
	 q.PushBack(s)
	 

	 for(q.Len()!=0) { // the loop iterates asl long as there remain grey vertices
		 u := q.Front().(string)
		 // fmt.Println(u)
		 vertices := G[u]
		 
		 for _, v := range vertices {
			 if(color[v]=="w") { //undiscovered vertices
				 color[v]="g" //discovered
				 d[v]=d[u]+1 //increment distance
				 pi[v]=u // u is recorded as the parent
				 q.PushBack(v) //add to tail of Q to be searched later
				 
			 }
			 
		 }
		 
		 q.PopFront()
 		 
		 color[u]="b"
		 
		 fmt.Println("q: ", q)

	 }
	 
	 
	 fmt.Println("colors", color)
 	 fmt.Println("distances", d)
	 fmt.Println("pi", pi)
	 
	 fmt.Println("done")
	 
	 return pi
	
}


//BFS runs in O(V) setup time + O(E) were E is the number of edges. The Queue operations are in theory O(1) so
//BFS runs in linear time


//the BFS routine populated populated a pi hash which holds the GFS tree and in theory can be traversed to form
// the shortest path to a given vertex. Here is a recursive algo to print the shortest path

func print_path(pi map[string]string, s string, v string)() {
	
	if(v==s) {
		fmt.Println("s") //base case of recursion
		//end recursion 
	} else {
		print_path(pi, s, pi[v])
		fmt.Println(v)
	}
}





