package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
/*
* Point
*/
type Point struct {
	X float64
	Y float64
	ClusterId int
}

func (a *Point) Distance(b Point) float64 {
	xx := math.Pow(float64(a.X-b.X), 2)
	yy := math.Pow(float64(a.Y-b.Y), 2)
	return math.Sqrt(xx + yy)
}

func (a Point) String() string {
	return fmt.Sprintf("{%v, %v}",a.X, a.Y)
}

/*
* Cluster
*/
type Cluster struct {
	Center Point
	Points []Point
}

func (cluster *Cluster) Calculate_centroid(){
	var x, y float64
	var clusterCount = len(cluster.Points)

	for i := 0; i < clusterCount; i++ {
		x = x + cluster.Points[i].X
		y = y + cluster.Points[i].Y
	}
	cluster.Points = []Point{}
	cluster.Center = Point{x / float64(clusterCount), y / float64(clusterCount), -1}
}

func (cluster Cluster) String() string {
	return fmt.Sprintf("Cluster centroid %v --> %v \n",cluster.Center, cluster.Points)
}

/*
* K-mean
*/

/*
*Step 1: Init k clusters
*/
func initClusters(dataset []Point,k int) (clusters []Cluster) {
    rand.Seed(time.Now().UnixNano())
    randlist := make(map[int]int)
    for len(randlist)<k{
        randlist[rand.Intn(len(dataset))] +=1
    }
    for i  := range randlist {
        clusters = append(clusters, Cluster{ dataset[i], []Point{}})
    }
    return
}

/*
*Step 2: Assign objects to their closest cluster center according to the Euclidean distance function.
*/

func assignClusters(dataset []Point, k int, clusters []Cluster) (hasChanged bool)  {
	hasChanged = false
	for i := 0; i < len(dataset); i++ {
		var minDist float64 = dataset[i].Distance(clusters[0].Center)
		var updatedClusterIndex int
		for j := 0; j < len(clusters); j++ {
			tmpDist := dataset[i].Distance(clusters[j].Center)
			if tmpDist < minDist {
				minDist = tmpDist
				updatedClusterIndex = j
			}
		}

		if dataset[i].ClusterId != updatedClusterIndex {
			hasChanged = true
		}
		dataset[i].ClusterId = updatedClusterIndex
		clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[i])
	}
	return hasChanged
}

/*
*Step 3: Calculate the centroid of all objects in each cluster.
*/
func renewCentroid(clusters []Cluster) {
	for i := 0; i < len(clusters); i++ {
		clusters[i].Calculate_centroid()
	}
}

/*
*Run K-mean
*/
func Kmean(dataset []Point, k int) []Cluster {
	clusters := initClusters(dataset, k)
	for assignClusters(dataset, k, clusters) {
			renewCentroid(clusters)
	}
	return clusters
}

func main() {
	var dataset []Point
	dataset = append(dataset,Point{X:1,Y:1})
	dataset = append(dataset,Point{X:1,Y:2})
	dataset = append(dataset,Point{X:2,Y:1})
	dataset = append(dataset,Point{X:2,Y:2})

	dataset = append(dataset,Point{X:4,Y:4})
	dataset = append(dataset,Point{X:4,Y:5})
	dataset = append(dataset,Point{X:5,Y:4})
	dataset = append(dataset,Point{X:5,Y:5})

	dataset = append(dataset,Point{X:7,Y:2})
	dataset = append(dataset,Point{X:7,Y:3})
	dataset = append(dataset,Point{X:8,Y:2})
	dataset = append(dataset,Point{X:8,Y:3})
	clusters := Kmean(dataset, 3)

	fmt.Println(clusters)
}
