# Exact TSP Solver (TSP ORACLE) [Go] <sup>v0.1.5</sup>

---

A high-performance, exact solver for the Traveling Salesman Problem (TSP) implemented in Go. 
Utilizes an intelligent Branch and Bound algorithm with adaptive thresholding to 
find the globally optimal solution for small to medium-sized TSP instances.

---

![GitHub top language](https://img.shields.io/github/languages/top/smartlegionlab/exact-tsp-solver)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/smartlegionlab/exact-tsp-solver)](https://github.com/smartlegionlab/exact-tsp-solver/)
[![GitHub](https://img.shields.io/github/license/smartlegionlab/exact-tsp-solver)](https://github.com/smartlegionlab/exact-tsp-solver/blob/master/LICENSE)
[![GitHub Repo stars](https://img.shields.io/github/stars/smartlegionlab/exact-tsp-solver?style=social)](https://github.com/smartlegionlab/exact-tsp-solver/)
[![GitHub watchers](https://img.shields.io/github/watchers/smartlegionlab/exact-tsp-solver?style=social)](https://github.com/smartlegionlab/exact-tsp-solver/)
[![GitHub forks](https://img.shields.io/github/forks/smartlegionlab/exact-tsp-solver?style=social)](https://github.com/smartlegionlab/exact-tsp-solver/)

---

**ATTENTION!** используйте `tsp_oracle/v2/tsp_oracle_v2.go` или `tsp_oracle/v3/tsp_oracle_v3.go`

---

> **Disclaimer:** The TSP is NP-Hard. This solver is designed for **educational and research purposes** 
> and is practical for instances up to ~20-35 points on standard hardware. 
> For larger instances, consider heuristic approaches.

## 🧠 What is "Branch and Bound" with Adaptive Thresholding?

This isn't a naive brute-force search. It's a sophisticated method that:
1.  **Sets a smart initial threshold** based on a quick heuristic solution (Greedy + 2-opt).
2.  **Uses a lower bound** (based on Minimum Spanning Tree) to prune futile branches of the search tree.
3.  **Dynamically adjusts the threshold** every iteration: lowers it if a better solution is found, or raises it if the search space is too constrained.
4.  **Employs neighborhood ordering** by checking nearest neighbors first to maximize pruning effectiveness.

This combination makes it significantly faster than a pure brute-force approach for finding provably optimal solutions.

## ⚡ Features

*   **✅ Provably Optimal:** Finds the absolute shortest possible route, not an approximation.
*   **🧩 Adaptive Algorithm:** Smart Branch and Bound with dynamic threshold adjustment for efficient search.
*   **📊 Built-in Benchmarking:** Compare optimal vs. heuristic solutions and see the performance gap.
*   **🛡️ Production-Grade Go:** Compiled to a single binary, static typing, no external dependencies.
*   **📈 Progress Tracking:** Real-time console output showing search speed and number of paths evaluated.

## 🚀 Usage

```bash
git clone https://github.com/smartlegionlab/exact-tsp-solver.git
cd exact-tsp-solver
```

The solver generates random points by default for demonstration. You can run it simply by:

v2 (stable):

```bash
go run tsp_oracle/v2/tsp_oracle_v2.go -n 10 -seed 123
```

**ATTENTION!** `tsp_oracle/v1/tsp_oracle_v1.go` This is v1 version, not recommended for use, serves only as an example for comparison with `tsp_oracle_v2.go`

v1 (old):

```bash
go run tsp_oracle/v1/tsp_oracle_v1.go -n 10 -seed 123
```

### Command Line Flags
| Flag | Default | Description |
| :--- | :--- | :--- |
| `-n` | `10` | Number of points to generate. |
| `-seed` | `42` | Random seed for reproducible point generation. |


## 🧪 Use Cases

*   **Academic Research:** Validate the performance of heuristic algorithms against the true optimum on small graphs.
*   **Education:** Perfect for teaching and understanding the Branch and Bound technique and the complexity of NP-Hard problems.
*   **Algorithmic Competitions:** As a reference solver for small test cases.
*   **Basis for Hybrid Solvers:** The optimal solver can be used in large-scale solvers that break the problem into smaller, solvable sub-problems.

## 🔗 See Also

- **[Smart TSP Solver](https://github.com/smartlegionlab/smart-tsp-solver)** - My Python library featuring advanced heuristics (`Dynamic Gravity`, `Angular Radial`) for solving *large* TSP instances where finding the exact optimum is impractical.
- **Smart TSP Oracle:** [smart-tsp-oracle](https://github.com/smartlegionlab/smart-tsp-oracle) - A high-performance, exact solver for the Traveling Salesman Problem (TSP) implemented in Go. Utilizes an intelligent Branch and Bound algorithm with adaptive thresholding to find the globally optimal solution for small to medium-sized TSP instances.
- **Smart TSP Benchmark** - [Smart TSP Benchmark](https://github.com/smartlegionlab/smart-tsp-benchmark) is a professional algorithm testing infrastructure with customizable scenarios and detailed metrics.

## 👨‍💻 Author

[**A.A. Suvorov**](https://github.com/smartlegionlab/)

*   Passionate about pushing the boundaries of algorithmic optimization.
*   This solver was developed to bridge the gap between theoretical computer science and practical implementation.

## 📜 License & Disclaimer

BSD 3-Clause License

Copyright (c) 2025, Alexander Suvorov

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
    AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
    IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
    DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
    FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
    DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
    SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
    CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
    OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

---

## 📊 Sample Output

`go run tsp_oracle/v2/tsp_oracle_v2.go -n 20 -seed 123`

```
==================================================
🚀 TSP ORACLE v2 - 20 POINTS
🔢 SEED: 123
==================================================

📍 Coordinates of points:
   Dot 0: (581.32, 26.22)
   Dot 1: (249.97, 615.85)
   Dot 2: (229.48, 636.50)
   Dot 3: (64.10, 448.80)
   Dot 4: (805.03, 382.43)
   Dot 5: (510.44, 503.08)
   Dot 6: (17.24, 930.66)
   Dot 7: (272.70, 204.95)
   Dot 8: (272.45, 156.08)
   Dot 9: (857.67, 846.22)
   Dot 10: (113.95, 530.24)
   Dot 11: (874.01, 291.98)
   Dot 12: (201.23, 548.09)
   Dot 13: (774.84, 533.88)
   Dot 14: (740.95, 317.05)
   Dot 15: (554.59, 325.05)
   Dot 16: (910.44, 311.79)
   Dot 17: (473.72, 932.95)
   Dot 18: (163.15, 236.74)
   Dot 19: (299.14, 485.05)
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 4322.31
2. Launching adaptive search...
   🎯 We start the search from 3890.08 (90.0%)
   🔍 Threshold: 3890.08 (90.0%)... ✗ cut off (63ms)
   🔍 Threshold: 4162.38 (96.3%)... ✓ found: 3975.71 (195ms)
   🔍 Threshold: 3697.41 (85.5%)... ✗ cut off (13ms)
   🏆 The optimum has been found.: 3975.71


📊 RESULTS:
==================================================
Number of points: 20
Seed: 123
Total possible paths: so many
Checked paths: 0
Execution time: 0.27 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 4322.309878
Optimal length: 3975.712587
Improvement: 346.597291 (8.019%)

Greedy way: [3 10 12 1 2 6 17 9 13 4 16 11 14 15 5 19 18 7 8 0]
The optimal path: [0 15 5 14 11 16 4 13 9 17 6 2 1 19 12 10 3 18 7 8]

💾 The results are saved in tsp_result_n20_seed123.txt
```

`go run tsp_oracle/v3/tsp_oracle_v3.go -n 30 -seed 1222`


```
⚠️  WARNING: for 30 points there will be approximately so many permutations
This may take a considerable amount of time.
Continue? (y/n): y
==================================================
🚀 TSP SOLVER (ORACLE v2) - 30 POINTS
🔢 SEED: 1222
==================================================

📍 Coordinates of points:
   Dot 0: (521.84, 998.65)
   Dot 1: (239.34, 374.98)
   Dot 2: (552.09, 440.40)
   Dot 3: (538.79, 510.33)
   Dot 4: (726.68, 971.58)
   Dot 5: (197.46, 506.00)
   Dot 6: (499.57, 105.44)
   Dot 7: (742.56, 269.45)
   Dot 8: (333.14, 347.61)
   Dot 9: (490.19, 146.89)
   Dot 10: (752.69, 175.08)
   Dot 11: (395.82, 967.65)
   Dot 12: (71.30, 843.90)
   Dot 13: (703.35, 764.47)
   Dot 14: (342.65, 498.74)
   Dot 15: (467.98, 545.24)
   Dot 16: (515.63, 415.79)
   Dot 17: (720.26, 874.17)
   Dot 18: (525.53, 686.15)
   Dot 19: (429.43, 105.12)
   Dot 20: (218.80, 919.85)
   Dot 21: (234.23, 65.48)
   Dot 22: (127.28, 688.89)
   Dot 23: (740.72, 901.09)
   Dot 24: (708.35, 971.53)
   Dot 25: (755.63, 123.60)
   Dot 26: (168.76, 324.78)
   Dot 27: (573.87, 134.81)
   Dot 28: (52.38, 309.08)
   Dot 29: (946.31, 544.56)
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 4478.33
2. Launching adaptive search...
   🎯 We start the search from 4030.50 (90.0%)
   🔍 Threshold: 4030.50 (90.0%)... ✗ cut off (142ms)
Checked: 309 paths | Speed: 69/sec | Time: 4s✓ found: 4180.48 (4.626s)
   🔍 Threshold: 3887.85 (86.8%)... ✗ cut off (17ms)
   🏆 The optimum has been found.: 4180.48


📊 RESULTS:
==================================================
Number of points: 30
Seed: 1222
Total possible paths: so many
Checked paths: 0
Execution time: 4.79 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 4478.332474
Optimal length: 4180.479698
Improvement: 297.852776 (6.651%)

Greedy way: [1 8 26 28 5 14 16 2 3 15 18 22 12 20 11 0 24 4 23 17 13 29 7 10 25 27 6 9 19 21]
The optimal path: [0 11 20 12 22 5 14 15 18 3 2 16 8 1 26 28 21 19 9 6 27 25 10 7 29 13 17 23 4 24]

💾 The results are saved in tsp_result_n30_seed1222.txt
```

`go run tsp_oracle/v3/tsp_oracle_v3.go -n 35 -seed 1222`

```
⚠️  WARNING: for 35 points there will be approximately so many permutations
This may take a considerable amount of time.
Continue? (y/n): y
==================================================
🚀 TSP SOLVER (ORACLE v2) - 35 POINTS
🔢 SEED: 1222
==================================================

📍 Coordinates of points:
   Dot 0: (521.84, 998.65)
   Dot 1: (239.34, 374.98)
   Dot 2: (552.09, 440.40)
   Dot 3: (538.79, 510.33)
   Dot 4: (726.68, 971.58)
   Dot 5: (197.46, 506.00)
   Dot 6: (499.57, 105.44)
   Dot 7: (742.56, 269.45)
   Dot 8: (333.14, 347.61)
   Dot 9: (490.19, 146.89)
   Dot 10: (752.69, 175.08)
   Dot 11: (395.82, 967.65)
   Dot 12: (71.30, 843.90)
   Dot 13: (703.35, 764.47)
   Dot 14: (342.65, 498.74)
   Dot 15: (467.98, 545.24)
   Dot 16: (515.63, 415.79)
   Dot 17: (720.26, 874.17)
   Dot 18: (525.53, 686.15)
   Dot 19: (429.43, 105.12)
   Dot 20: (218.80, 919.85)
   Dot 21: (234.23, 65.48)
   Dot 22: (127.28, 688.89)
   Dot 23: (740.72, 901.09)
   Dot 24: (708.35, 971.53)
   Dot 25: (755.63, 123.60)
   Dot 26: (168.76, 324.78)
   Dot 27: (573.87, 134.81)
   Dot 28: (52.38, 309.08)
   Dot 29: (946.31, 544.56)
   Dot 30: (767.87, 392.41)
   Dot 31: (731.39, 336.48)
   Dot 32: (968.27, 300.19)
   Dot 33: (164.72, 314.39)
   Dot 34: (163.18, 117.21)
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 4682.18
2. Launching adaptive search...
   🎯 We start the search from 4213.96 (90.0%)
   🔍 Threshold: 4213.96 (90.0%)... ✗ cut off (57ms)
Checked: 151 paths | Speed: 39/sec | Time: 4s✓ found: 4446.24 (4.301s)
   🔍 Threshold: 4135.00 (88.3%)... ✗ cut off (18ms)
   🏆 The optimum has been found.: 4446.24


📊 RESULTS:
==================================================
Number of points: 35
Seed: 1222
Total possible paths: so many
Checked paths: 0
Execution time: 4.38 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 4682.180857
Optimal length: 4446.237222
Improvement: 235.943636 (5.039%)

Greedy way: [2 16 3 15 18 22 12 20 11 0 24 4 23 17 13 29 32 30 31 7 10 25 27 6 9 19 21 34 28 33 26 1 5 14 8]
The optimal path: [0 11 20 12 22 5 14 8 1 26 33 28 34 21 19 9 6 27 25 10 7 32 29 30 31 16 2 3 15 18 13 17 23 4 24]

💾 The results are saved in tsp_result_n35_seed1222.txt
```