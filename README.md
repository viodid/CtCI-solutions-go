# Cracking the Coding Interview Solutions in Go

This repository contains comprehensive solutions to problems from "Cracking the Coding Interview" 6th Edition by Gayle Laakmann McDowell, implemented in Go with extensive test suites.

## 📚 Project Scope

**This repository is a work in progress** that aims to provide complete solutions to all problems from "Cracking the Coding Interview" 6th Edition. The repository currently includes:

- **Chapter 1: Arrays and Strings** - Problems 1.1 through 1.9 (Complete)
- **Chapter 2: Linked Lists** - Problems 2.1 through 2.7 (Complete)
- **Additional chapters** - Will be added progressively

### What This Repository Will Include (When Complete):
- ✅ All 189 problems from the book implemented in Go
- ✅ Comprehensive test suites for every problem
- ✅ Multiple solution approaches where applicable
- ✅ Time and space complexity analysis
- ✅ Clean, well-documented code following Go best practices

### Current Status:
- **Chapters 1-2**: Fully implemented with tests
- **Remaining chapters**: Work in progress

## Project Structure

The repository is organized by book chapters, with each chapter containing individual problem solutions:

```
CtCI-solutions/
├── ch1/           # Chapter 1: Arrays and Strings
│   ├── 1_1.go     # Problem 1.1: Is Unique
│   ├── 1_1_test.go
│   ├── 1_2.go     # Problem 1.2: Check Permutation
│   ├── 1_2_test.go
│   ├── ...
│   └── go.mod
├── ch2/           # Chapter 2: Linked Lists
│   ├── 2_1.go     # Problem 2.1: Remove Dups
│   ├── 2_1_test.go
│   ├── ...
│   ├── go.mod
│   └── ll/        # Linked List utilities
│       └── linkedlist.go
└── README.md
```

### Chapter Organization

Each chapter is structured as a separate Go module with:
- **Problem solutions**: `{chapter}_{problem}.go` (e.g., `1_1.go`)
- **Test files**: `{chapter}_{problem}_test.go` (e.g., `1_1_test.go`)
- **Utility packages**: Supporting data structures and helper functions
- **Module file**: `go.mod` for dependency management

## 🚀 Getting Started

### Prerequisites

- Go 1.18+ (required for generics support)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/viodid/CtCI-solutions.git
cd CtCI-solutions
```

2. Navigate to any chapter directory:
```bash
cd ch1
```

3. Run tests for all problems in the chapter:
```bash
go test -v
```

## 🧪 Running Tests

### Run All Tests in a Chapter
```bash
cd ch1
go test -v
```

### Run Tests for a Specific Problem
```bash
cd ch1
go test -v -run TestIsUnique
```

### Run Tests with Coverage
```bash
cd ch1
go test -cover
```

### Run All Tests in All Chapters
From the root directory:
```bash
for dir in ch*/; do
    echo "Testing $dir"
    cd "$dir" && go test -v && cd ..
done
```

## 📋 Problem Categories

### Currently Available:

#### Chapter 1: Arrays and Strings (Complete)
Problems 1.1 through 1.9 involving string manipulation, array operations, and character handling.

#### Chapter 2: Linked Lists (Complete)  
Problems 2.1 through 2.7 involving linked list operations, including a custom generic linked list implementation with operations like Add, Remove, Search, and Deep Copy.

### Coming Soon:
- **Chapter 3**: Stacks and Queues
- **Chapter 4**: Trees and Graphs  
- **Chapter 5**: Bit Manipulation
- **Chapter 6**: Math and Logic Puzzles
- **Chapter 7**: Object-Oriented Design
- **Chapter 8**: Recursion and Dynamic Programming
- **Chapter 9**: System Design and Scalability
- **Chapter 10**: Sorting and Searching
- **Chapter 11**: Testing
- **Additional chapters** through Chapter 17

## Code Quality Features

### Multiple Solution Approaches
Many problems include multiple implementations with different time/space trade-offs:
- Brute force vs. optimized solutions
- Different data structures for various constraints
- Memory-efficient vs. time-efficient approaches

### Comprehensive Testing
Each solution includes:
- **Table-driven tests** following Go best practices
- **Edge cases** testing (empty inputs, single elements, etc.)
- **Multiple test scenarios** covering various input patterns

### Documentation
- **Time complexity** analysis for each solution
- **Space complexity** analysis for each solution
- **Clear comments** explaining algorithmic approaches
- **Function documentation** following Go conventions

## Development

### Adding New Solutions

1. Create solution file: `ch{X}/{X}_{Y}.go`
2. Create test file: `ch{X}/{X}_{Y}_test.go`
3. Include complexity analysis in comments
4. Add comprehensive test cases
5. Follow Go naming conventions

### Test Structure
All tests follow Go's table-driven testing pattern with comprehensive test cases covering:
- Normal input scenarios
- Edge cases (empty inputs, single elements)
- Boundary conditions
- Invalid inputs where applicable

## Book Reference

These solutions correspond to **"Cracking the Coding Interview" 6th Edition** by Gayle Laakmann McDowell. Each problem number matches the book's organization:

- **1.1**: Is Unique
- **1.2**: Check Permutation
- **1.3**: URLify
- **2.1**: Remove Dups
- **2.2**: Return Kth to Last
- And so on...

## 🤝 Contributing

Contributions are welcome! Please feel free to:

1. Add solutions for missing problems
2. Improve existing solutions
3. Add more test cases
4. Optimize algorithms
5. Fix bugs or improve documentation

### Contribution Guidelines
- Follow Go best practices and formatting (`go fmt`)
- Include comprehensive tests
- Document time/space complexity
- Maintain consistent code style
- Update README if adding new chapters

## License

This project is open source and available under the [MIT License](LICENSE).

