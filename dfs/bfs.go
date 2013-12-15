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
	
	bfs(adj, "s");

}

func bfs(G map[string][]string, s string) {
	
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
	 q.PushFront(s)
	 

	 for(q.Len()!=0) { // the loop iterates asl long as there remain grey vertices
		 u := q.Back().(string)
		 // fmt.Println(u)
		 vertices := G[u]
		 
		 for _, v := range vertices {
			 if(color[v]=="w") { //undiscovered vertices
				 color[v]="g" //discovered
				 d[v]=d[u]+1 //increment distance
				 pi[v]=u // u is recorded as the parent
				 q.PushFront(v) //add to tail of Q to be searched later
				 
			 }
			 
		 }
		 
		 q.PopBack()
 		 
		 color[u]="b"
		 
		 fmt.Println("q: ", q)
	 }
	 
	 
	 fmt.Println("colors", color)
 	 fmt.Println("distances", d)
	 fmt.Println("pi", pi)
	 
	 fmt.Println("done")
	
}




