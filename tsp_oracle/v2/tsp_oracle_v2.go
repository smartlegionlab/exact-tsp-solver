// Copyright © 2025 Alexander Suvorov. All rights reserved.
package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Point struct {
	X, Y float64
}

type Edge struct {
	Weight float64
	From   int
	To     int
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

type TSPSolver struct {
	numPoints         int
	points            []Point
	distanceMatrix    [][]float64
	bestPath          []int
	bestDistance      float64
	visitedCount      int64
	totalPermutations int64
	seed              int64
	localRand         *rand.Rand
	mstCache          map[string]float64
	nearestNeighbors  [][]int
	lastUpdate        time.Time
	startTime         time.Time
	greedyPath        []int
	greedyDistance    float64
}

func NewTSPSolver(numPoints int, seed int64) *TSPSolver {
	solver := &TSPSolver{
		numPoints: numPoints,
		seed:      seed,
		mstCache:  make(map[string]float64),
	}
	source := rand.NewSource(seed)
	solver.localRand = rand.New(source)
	solver.points = solver.generateRandomPoints()
	solver.distanceMatrix = solver.calculateDistanceMatrix()
	solver.totalPermutations = calculateTotalPermutations(numPoints)
	solver.precomputeNearestNeighbors()
	return solver
}

func (t *TSPSolver) generateRandomPoints() []Point {
	points := make([]Point, t.numPoints)
	for i := 0; i < t.numPoints; i++ {
		points[i] = Point{
			X: t.localRand.Float64() * 1000,
			Y: t.localRand.Float64() * 1000,
		}
	}
	return points
}

func (t *TSPSolver) calculateDistanceMatrix() [][]float64 {
	matrix := make([][]float64, t.numPoints)
	for i := range matrix {
		matrix[i] = make([]float64, t.numPoints)
	}

	for i := 0; i < t.numPoints; i++ {
		for j := i + 1; j < t.numPoints; j++ {
			dx := t.points[i].X - t.points[j].X
			dy := t.points[i].Y - t.points[j].Y
			distance := math.Sqrt(dx*dx + dy*dy)
			matrix[i][j] = distance
			matrix[j][i] = distance
		}
	}
	return matrix
}

func (t *TSPSolver) precomputeNearestNeighbors() {
	t.nearestNeighbors = make([][]int, t.numPoints)
	for i := 0; i < t.numPoints; i++ {
		neighbors := make([]int, t.numPoints)
		for j := 0; j < t.numPoints; j++ {
			neighbors[j] = j
		}
		sort.Slice(neighbors, func(k, l int) bool {
			return t.distanceMatrix[i][neighbors[k]] < t.distanceMatrix[i][neighbors[l]]
		})
		if len(neighbors) > 10 {
			t.nearestNeighbors[i] = neighbors[1:11]
		} else {
			t.nearestNeighbors[i] = neighbors[1:]
		}
	}
}

func (t *TSPSolver) calculatePathDistance(path []int) float64 {
	total := 0.0
	for i := 0; i < len(path)-1; i++ {
		total += t.distanceMatrix[path[i]][path[i+1]]
	}
	total += t.distanceMatrix[path[len(path)-1]][path[0]]
	return total
}

func (t *TSPSolver) calculatePartialDistance(path []int) float64 {
	total := 0.0
	for i := 0; i < len(path)-1; i++ {
		total += t.distanceMatrix[path[i]][path[i+1]]
	}
	return total
}

func reverseSegment(path []int, i, j int) {
	for i < j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}
}

func (t *TSPSolver) twoOpt(path []int) ([]int, float64) {
	if len(path) < 4 {
		return path, t.calculatePathDistance(path)
	}

	improved := true
	currentPath := make([]int, len(path))
	copy(currentPath, path)
	currentDistance := t.calculatePathDistance(currentPath)
	n := len(currentPath)

	maxIterations := 1000
	iteration := 0

	for improved && iteration < maxIterations {
		iteration++
		improved = false
		bestDelta := 0.0
		bestI, bestJ := -1, -1

		for i := 1; i < n-2; i++ {
			for j := i + 1; j < n-1; j++ {
				delta := -t.distanceMatrix[currentPath[i-1]][currentPath[i]] -
					t.distanceMatrix[currentPath[j]][currentPath[j+1]] +
					t.distanceMatrix[currentPath[i-1]][currentPath[j]] +
					t.distanceMatrix[currentPath[i]][currentPath[j+1]]

				if delta < bestDelta {
					bestDelta = delta
					bestI = i
					bestJ = j
					improved = true
				}
			}
		}

		if improved {
			reverseSegment(currentPath, bestI, bestJ)
			currentDistance += bestDelta
		}
	}
	return currentPath, currentDistance
}

func (t *TSPSolver) multiStartGreedy() ([]int, float64) {
	bestPath := []int{}
	bestDist := math.MaxFloat64

	for startPoint := 0; startPoint < t.numPoints && startPoint < 5; startPoint++ {
		path := []int{startPoint}
		unvisited := make(map[int]bool)
		for i := 0; i < t.numPoints; i++ {
			if i != startPoint {
				unvisited[i] = true
			}
		}

		for len(unvisited) > 0 {
			current := path[len(path)-1]
			var nextPoint int
			minDist := math.MaxFloat64

			for _, neighbor := range t.nearestNeighbors[current] {
				if unvisited[neighbor] {
					if dist := t.distanceMatrix[current][neighbor]; dist < minDist {
						minDist = dist
						nextPoint = neighbor
					}
				}
			}

			if minDist == math.MaxFloat64 {
				for point := range unvisited {
					if dist := t.distanceMatrix[current][point]; dist < minDist {
						minDist = dist
						nextPoint = point
					}
				}
			}

			path = append(path, nextPoint)
			delete(unvisited, nextPoint)
		}

		distance := t.calculatePathDistance(path)
		if distance < bestDist {
			bestDist = distance
			bestPath = path
		}
	}

	improvedPath, improvedDist := t.twoOpt(bestPath)
	return improvedPath, improvedDist
}

func (t *TSPSolver) calculateMSTKruskal(points []int) float64 {
	if len(points) <= 1 {
		return 0.0
	}

	sort.Ints(points)
	key := fmt.Sprintf("%v", points)
	if cached, exists := t.mstCache[key]; exists {
		return cached
	}

	edges := make([]Edge, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			from, to := points[i], points[j]
			edges = append(edges, Edge{
				Weight: t.distanceMatrix[from][to],
				From:   from,
				To:     to,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	pointToIndex := make(map[int]int)
	for idx, point := range points {
		pointToIndex[point] = idx
	}

	uf := NewUnionFind(len(points))
	mstLength := 0.0
	edgesUsed := 0

	for _, edge := range edges {
		if edgesUsed == len(points)-1 {
			break
		}

		idxFrom := pointToIndex[edge.From]
		idxTo := pointToIndex[edge.To]

		if uf.Find(idxFrom) != uf.Find(idxTo) {
			uf.Union(idxFrom, idxTo)
			mstLength += edge.Weight
			edgesUsed++
		}
	}

	t.mstCache[key] = mstLength
	return mstLength
}

func (t *TSPSolver) minConnectionToPath(unvisited []int, path []int) float64 {
	if len(unvisited) == 0 {
		return 0.0
	}

	start, end := path[0], path[len(path)-1]
	minStart, minEnd := math.MaxFloat64, math.MaxFloat64

	for _, point := range unvisited {
		if dist := t.distanceMatrix[start][point]; dist < minStart {
			minStart = dist
		}
		if dist := t.distanceMatrix[end][point]; dist < minEnd {
			minEnd = dist
		}
	}

	return minStart + minEnd
}

func (t *TSPSolver) calculateLowerBound(currentPath []int, visited []bool) float64 {
	currentLength := t.calculatePartialDistance(currentPath)

	unvisited := make([]int, 0)
	for i := 0; i < t.numPoints; i++ {
		if !visited[i] {
			unvisited = append(unvisited, i)
		}
	}

	if len(unvisited) == 0 {
		return currentLength + t.distanceMatrix[currentPath[len(currentPath)-1]][currentPath[0]]
	}

	mstLength := t.calculateMSTKruskal(unvisited)
	connectionLength := t.minConnectionToPath(unvisited, currentPath)

	return currentLength + mstLength + connectionLength
}

func (t *TSPSolver) printProgress() {
	elapsed := time.Since(t.startTime)
	pathsPerSec := float64(t.visitedCount) / elapsed.Seconds()

	fmt.Printf("\rChecked: %s paths | Speed: %.0f/sec | Time: %v",
		formatLargeNumber(t.visitedCount), pathsPerSec, elapsed.Round(time.Second))
}

func (t *TSPSolver) adaptiveSearch() ([]int, float64) {
	greedyPath, greedyDist := t.multiStartGreedy()
	t.greedyPath = greedyPath
	t.greedyDistance = greedyDist

	currentThreshold := greedyDist * 0.90
	step := 0.07
	bestPath := greedyPath
	bestDist := greedyDist
	foundAny := false

	fmt.Printf("   🎯 We start the search from %.2f (90.0%%)\n", currentThreshold)

	for iteration := 0; iteration < 200; iteration++ {
		t.bestPath = make([]int, 0)
		t.bestDistance = currentThreshold
		t.visitedCount = 0
		t.mstCache = make(map[string]float64)

		visited := make([]bool, t.numPoints)
		visited[0] = true

		fmt.Printf("   🔍 Threshold: %.2f (%.1f%%)... ", currentThreshold, (currentThreshold/greedyDist)*100)

		startTime := time.Now()
		t.bruteForceRecursive([]int{0}, 0.0, visited)
		searchTime := time.Since(startTime)

		if t.bestDistance < currentThreshold {
			fmt.Printf("✓ found: %.2f (%v)\n", t.bestDistance, searchTime.Round(time.Millisecond))
			bestPath = make([]int, len(t.bestPath))
			copy(bestPath, t.bestPath)
			bestDist = t.bestDistance
			foundAny = true

			currentThreshold = bestDist * (1.0 - step)
		} else {
			fmt.Printf("✗ cut off (%v)\n", searchTime.Round(time.Millisecond))

			if foundAny {
				fmt.Printf("   🏆 The optimum has been found.: %.2f\n", bestDist)
				return bestPath, bestDist
			} else {
				currentThreshold = currentThreshold * (1.0 + step)

				if currentThreshold >= greedyDist {
					fmt.Printf("   ⚠️  No better solutions than greedy found\n")
					return greedyPath, greedyDist
				}
			}
		}
	}

	fmt.Printf("   🏆 Best found: %.2f\n", bestDist)
	return bestPath, bestDist
}

func (t *TSPSolver) bruteForceRecursive(currentPath []int, currentDistance float64, visited []bool) {
	lowerBound := t.calculateLowerBound(currentPath, visited)
	if lowerBound >= t.bestDistance {
		return
	}

	if len(currentPath) == t.numPoints {
		finalDistance := currentDistance + t.distanceMatrix[currentPath[len(currentPath)-1]][currentPath[0]]
		t.visitedCount++

		if time.Since(t.lastUpdate) > 500*time.Millisecond {
			t.printProgress()
			t.lastUpdate = time.Now()
		}

		if finalDistance < t.bestDistance {
			t.bestDistance = finalDistance
			t.bestPath = make([]int, len(currentPath))
			copy(t.bestPath, currentPath)
		}
		return
	}

	lastPoint := currentPath[len(currentPath)-1]
	pointsToTry := make([]int, 0)

	for _, neighbor := range t.nearestNeighbors[lastPoint] {
		if !visited[neighbor] {
			pointsToTry = append(pointsToTry, neighbor)
		}
	}

	for i := 0; i < t.numPoints; i++ {
		if !visited[i] {
			found := false
			for _, pt := range pointsToTry {
				if pt == i {
					found = true
					break
				}
			}
			if !found {
				pointsToTry = append(pointsToTry, i)
			}
		}
	}

	sort.Slice(pointsToTry, func(i, j int) bool {
		return t.distanceMatrix[lastPoint][pointsToTry[i]] < t.distanceMatrix[lastPoint][pointsToTry[j]]
	})

	for _, nextPoint := range pointsToTry {
		newDistance := currentDistance + t.distanceMatrix[lastPoint][nextPoint]

		if newDistance >= t.bestDistance {
			continue
		}

		visited[nextPoint] = true
		newPath := make([]int, len(currentPath))
		copy(newPath, currentPath)
		newPath = append(newPath, nextPoint)

		t.bruteForceRecursive(newPath, newDistance, visited)
		visited[nextPoint] = false
	}
}

func (t *TSPSolver) bruteForce() ([]int, float64, time.Duration) {
	t.bestPath = []int{}
	t.bestDistance = math.MaxFloat64
	t.visitedCount = 0
	t.mstCache = make(map[string]float64)
	t.lastUpdate = time.Now()
	t.startTime = time.Now()

	fmt.Println("1. Launching the multi-start greedy algorithm...")
	t.greedyPath, t.greedyDistance = t.multiStartGreedy()
	fmt.Printf("   ✅ Multi-start greedy + 2-opt: length = %.2f\n", t.greedyDistance)

	fmt.Println("2. Launching adaptive search...")
	optimalPath, optimalDist := t.adaptiveSearch()

	fmt.Println()
	elapsed := time.Since(t.startTime)

	return optimalPath, optimalDist, elapsed
}

func calculateTotalPermutations(n int) int64 {
	if n <= 1 {
		return 1
	}

	if n > 20 {
		return math.MaxInt64
	}

	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func formatLargeNumber(n int64) string {
	if n < 0 {
		return "so many"
	}
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	} else if n < 1000000 {
		return fmt.Sprintf("%.1f thousand.", float64(n)/1000)
	} else if n < 1000000000 {
		return fmt.Sprintf("%.1f million.", float64(n)/1000000)
	} else if n < 1000000000000 {
		return fmt.Sprintf("%.1f billion.", float64(n)/1000000000)
	} else {
		return "so many"
	}
}

func main() {
	numPoints := flag.Int("n", 10, "Number of points")
	seed := flag.Int64("seed", 42, "Seed for generating random points")
	flag.Parse()

	if *numPoints < 3 {
		fmt.Println("Error: Minimum 3 dots required")
		os.Exit(1)
	}

	if *numPoints > 25 {
		totalPerms := calculateTotalPermutations(*numPoints - 1)
		fmt.Printf("⚠️  WARNING: for %d points there will be approximately %s permutations\n",
			*numPoints, formatLargeNumber(totalPerms))
		fmt.Printf("This may take a considerable amount of time.\n")
		fmt.Print("Continue? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Cancelled by user")
			os.Exit(0)
		}
	}

	solver := NewTSPSolver(*numPoints, *seed)

	fmt.Println("==================================================")
	fmt.Printf("🚀 TSP SOLVER (ORACLE v2) - %d POINTS\n", *numPoints)
	fmt.Printf("🔢 SEED: %d\n", *seed)
	fmt.Println("==================================================")

	fmt.Println("\n📍 Coordinates of points:")
	for i, point := range solver.points {
		fmt.Printf("   Dot %d: (%.2f, %.2f)\n", i, point.X, point.Y)
	}

	brutePath, bruteDistance, elapsed := solver.bruteForce()

	fmt.Println("\n📊 RESULTS:")
	fmt.Println("==================================================")
	fmt.Printf("Number of points: %d\n", *numPoints)
	fmt.Printf("Seed: %d\n", *seed)
	fmt.Printf("Total possible paths: %s\n", formatLargeNumber(solver.totalPermutations))
	fmt.Printf("Checked paths: %s\n", formatLargeNumber(solver.visitedCount))
	fmt.Printf("Execution time: %.2f seconds\n", elapsed.Seconds())
	if elapsed.Seconds() > 0 {
		fmt.Printf("Speed: %.0f paths/sec\n", float64(solver.visitedCount)/elapsed.Seconds())
	}
	fmt.Printf("Greedy + 2-opt: %.6f\n", solver.greedyDistance)
	fmt.Printf("Optimal length: %.6f\n", bruteDistance)
	improvement := solver.greedyDistance - bruteDistance
	improvementPercent := (improvement / solver.greedyDistance) * 100
	fmt.Printf("Improvement: %.6f (%.3f%%)\n", improvement, improvementPercent)

	fmt.Printf("\nGreedy way: %v\n", solver.greedyPath)
	fmt.Printf("The optimal path: %v\n", brutePath)

	filename := fmt.Sprintf("tsp_result_n%d_seed%d.txt", *numPoints, *seed)
	file, err := os.Create(filename)
	if err == nil {
		defer file.Close()
		fmt.Fprintf(file, "SEED: %d\n", *seed)
		fmt.Fprintf(file, "Points:\n")
		for i, p := range solver.points {
			fmt.Fprintf(file, "%d: (%.6f, %.6f)\n", i, p.X, p.Y)
		}
		fmt.Fprintf(file, "Greedy + 2-opt: %.6f\n", solver.greedyDistance)
		fmt.Fprintf(file, "Optimal: %.6f\n", bruteDistance)
		fmt.Fprintf(file, "Improvement: %.6f (%.3f%%)\n", improvement, improvementPercent)
		fmt.Fprintf(file, "Greedy path: %v\n", solver.greedyPath)
		fmt.Fprintf(file, "Optimal path: %v\n", brutePath)
		fmt.Fprintf(file, "Time: %.2f seconds\n", elapsed.Seconds())
		fmt.Fprintf(file, "Paths checked: %d\n", solver.visitedCount)
		fmt.Fprintf(file, "Total paths: %d\n", solver.totalPermutations)
		fmt.Printf("\n💾 The results are saved in %s\n", filename)
	} else {
		fmt.Printf("\n❌ Error saving file: %v\n", err)
	}
}
