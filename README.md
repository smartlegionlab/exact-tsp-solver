# Exact TSP Solver (TSP ORACLE) [Go] <sup>v0.1.1</sup>

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

*   **[smart-tsp-heuristics](https://github.com/smartlegionlab/smart-tsp-heuristics)** - My Python library featuring advanced heuristics (`Dynamic Gravity`, `Angular Radial`) for solving *large* TSP instances where finding the exact optimum is impractical.

## 👨‍💻 Author

[**A.A. Suvorov**](https://github.com/smartlegionlab/)

*   Passionate about pushing the boundaries of algorithmic optimization.
*   This solver was developed to bridge the gap between theoretical computer science and practical implementation.

## 📜 Licensing

This project uses a dual licensing system:

### 🆓 BSD 3-Clause License
- For non-commercial use
- For academic and research purposes
- For open-source projects

### 💼 Commercial License
- For commercial products and services
- For enterprises using the code in proprietary solutions
- For additional features and support

**To obtain a commercial license:** [smartlegiondev@gmail.com](mailto:smartlegiondev@gmail.com)

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

`go run tsp_oracle/v2/tsp_oracle_v2.go -n 30 -seed 123`

```
⚠️  WARNING: for 30 points there will be approximately so many permutations
This may take a considerable amount of time.
Continue? (y/n): y
==================================================
🚀 TSP ORACLE v2 - 30 POINTS
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
   Dot 20: (55.62, 636.14)
   Dot 21: (179.74, 520.60)
   Dot 22: (601.68, 494.52)
   Dot 23: (933.10, 18.90)
   Dot 24: (462.19, 321.68)
   Dot 25: (790.67, 418.69)
   Dot 26: (716.88, 986.36)
   Dot 27: (361.88, 489.54)
   Dot 28: (86.09, 83.12)
   Dot 29: (211.74, 899.58)
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 4893.87
2. Launching adaptive search...
   🎯 We start the search from 4404.48 (90.0%)
   🔍 Threshold: 4404.48 (90.0%)... ✗ cut off (34ms)
   🔍 Threshold: 4712.80 (96.3%)... ✗ cut off (2.056s)
   ⚠️  No better solutions than greedy found


📊 RESULTS:
==================================================
Number of points: 30
Seed: 123
Total possible paths: so many
Checked paths: 0
Execution time: 2.09 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 4893.870477
Optimal length: 4893.870477
Improvement: 0.000000 (0.000%)

Greedy way: [3 10 21 12 2 1 19 27 5 22 15 24 7 8 18 28 0 23 16 11 14 4 25 13 9 26 17 29 6 20]
The optimal path: [3 10 21 12 2 1 19 27 5 22 15 24 7 8 18 28 0 23 16 11 14 4 25 13 9 26 17 29 6 20]

💾 The results are saved in tsp_result_n30_seed123.txt
```

`go run tsp_oracle/v2/tsp_oracle_v2.go -n 35 -seed 123`

```
⚠️  WARNING: for 35 points there will be approximately so many permutations
This may take a considerable amount of time.
Continue? (y/n): y
==================================================
🚀 TSP ORACLE v2 - 35 POINTS
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
   Dot 20: (55.62, 636.14)
   Dot 21: (179.74, 520.60)
   Dot 22: (601.68, 494.52)
   Dot 23: (933.10, 18.90)
   Dot 24: (462.19, 321.68)
   Dot 25: (790.67, 418.69)
   Dot 26: (716.88, 986.36)
   Dot 27: (361.88, 489.54)
   Dot 28: (86.09, 83.12)
   Dot 29: (211.74, 899.58)
   Dot 30: (789.24, 304.27)
   Dot 31: (813.16, 337.54)
   Dot 32: (243.24, 247.72)
   Dot 33: (823.18, 644.94)
   Dot 34: (934.04, 394.11)
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 5444.44
2. Launching adaptive search...
   🎯 We start the search from 4899.99 (90.0%)
   🔍 Threshold: 4899.99 (90.0%)... ✗ cut off (6.895s)
Checked: 18 paths | Speed: 0/sec | Time: 1m6s✓ found: 5041.09 (1m24.956s)
   🔍 Threshold: 4688.21 (86.1%)... ✗ cut off (373ms)
   🏆 The optimum has been found.: 5041.09


📊 RESULTS:
==================================================
Number of points: 35
Seed: 123
Total possible paths: so many
Checked paths: 0
Execution time: 92.22 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 5444.435629
Optimal length: 5041.088739
Improvement: 403.346890 (7.408%)

Greedy way: [4 14 30 31 11 16 34 25 13 33 9 26 17 29 6 20 10 21 12 2 1 19 27 5 22 15 24 3 28 18 32 7 8 0 23]
The optimal path: [0 8 28 18 32 7 24 15 22 5 27 19 1 2 12 21 10 3 20 6 29 17 26 9 33 13 25 4 14 30 31 34 16 11 23]

💾 The results are saved in tsp_result_n35_seed123.txt
```

`go run tsp_oracle/v2/tsp_oracle_v2.go -n 7 -seed 123`

```
==================================================
🚀 TSP SOLVER (ORACLE v2) - 7 POINTS
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
1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy + 2-opt: length = 2852.11
2. Launching adaptive search...
   🎯 We start the search from 2566.90 (90.0%)
   🔍 Threshold: 2566.90 (90.0%)... ✓ found: 2566.65 (0s)
   🔍 Threshold: 2386.99 (83.7%)... ✗ cut off (0s)
   🏆 The optimum has been found.: 2566.65


📊 RESULTS:
==================================================
Number of points: 7
Seed: 123
Total possible paths: 5.0 thousand.
Checked paths: 0
Execution time: 0.00 seconds
Speed: 0 paths/sec
Greedy + 2-opt: 2852.110178
Optimal length: 2566.654781
Improvement: 285.455397 (10.009%)

Greedy way: [0 4 5 1 2 3 6]
The optimal path: [0 4 5 1 2 6 3]

💾 The results are saved in tsp_result_n7_seed123.txt
```

**ATTENTION!** `tsp_oracle/v1/tsp_oracle_v1.go` This is v1 version, not recommended for use, serves only as an example for comparison with v2

`go run tsp_oracle/v1/tsp_oracle_v1.go -n 20 -seed 123`

```
==================================================
🚀 TSP ORACLE v1 - 20 POINTS
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
   ✅ Multi-start greedy: length = 4352.89
Launching a search with optimizations...


📊 RESULTS:
==================================================
Number of points: 20
Seed: 123
Total possible paths: so many
Checked paths: 16
Lead time: 0.28 seconds
Speed: 58 paths/secMulti-start greedy: 4352.888734
The optimal path:   3975.712587
Improvement:          377.176147 (8.665%)

Multi-start greedy path: [18 7 8 15 5 19 12 1 2 10 3 6 17 9 13 4 14 11 16 0]
The optimal path:       [0 15 5 14 11 16 4 13 9 17 6 2 1 19 12 10 3 18 7 8]

💾 The results are saved in tsp_result_n20_seed123.txt
```

`go run tsp_oracle/v1/tsp_oracle_v1.go -n 30 -seed 123`

```
⚠️  WARNING: for 30 points there will be approximately so many permutations
This may take a considerable amount of time.
Continue? (y/n): y
==================================================
🚀 TSP ORACLE v1 - 30 POINTS
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
   Dot 20: (55.62, 636.14)
   Dot 21: (179.74, 520.60)
   Dot 22: (601.68, 494.52)
   Dot 23: (933.10, 18.90)
   Dot 24: (462.19, 321.68)
   Dot 25: (790.67, 418.69)
   Dot 26: (716.88, 986.36)
   Dot 27: (361.88, 489.54)
   Dot 28: (86.09, 83.12)
   Dot 29: (211.74, 899.58)

1. Launching the multi-start greedy algorithm...
   ✅ Multi-start greedy: length = 5137.63
Launching a search with optimizations...
Checked: 11 paths | Speed: 1 paths/sec | Time: 14s

📊 RESULTS:
==================================================
Number of points: 30
Seed: 123
Total possible paths: so many
Checked paths: 11
Lead time: 16.17 seconds
Speed: 1 paths/secMulti-start greedy: 5137.627462
The optimal path:   4859.916181
Improvement:          277.711281 (5.405%)

Multi-start greedy path: [19 27 5 22 15 24 7 8 18 28 3 10 21 12 1 2 20 6 29 17 26 9 13 25 4 14 11 16 23 0]
The optimal path:       [0 23 16 11 14 4 25 13 9 26 17 29 6 20 3 10 21 12 2 1 19 27 5 22 15 24 7 18 28 8]

💾 The results are saved in tsp_result_n30_seed123.txt
```

---

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