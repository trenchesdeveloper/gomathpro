# Gomathpro

Gomathpro is a CLI tool for evaluating mathematical expressions with support for variables, exponents, factorials, and advanced functions.

## Installation

To install Gomathpro, run:

```bash
go install github.com/trenchesdeveloper/gomathpro
```
## Usage
### Basic Usage
Evaluate a mathematical expression:

```bash
gomathpro eval "5 + 3 * 2"

```

### Variables
Evaluate a mathematical expression with variables:

```bash
gomathpro eval "A = 5; B = 7; A + B"

```

### Advanced Expressions

Evaluate advanced mathematical expressions with exponents, factorials, and more:

```bash
gomathpro eval "A = 5; B = 7; (A + B) * 2 - fact(3)"
# Output: 18

```

### Functions

Use built-in functions like sqrt, sin, cos, log, and more:

```bash
gomathpro eval "sqrt(16) + sin(0) + log(10)"

```

Here’s a preview of the **Supported Features** table:

| Feature               | Syntax Example         | Description                                                                 |
|-----------------------|------------------------|-----------------------------------------------------------------------------|
| **Basic Arithmetic**  | `5 + 3`, `10 - 4`      | Addition, subtraction, multiplication, and division.                        |
| **Exponents**         | `2 ^ 3`                | Exponentiation (`2^3` = 8).                                                |
| **Factorials**        | `fact(5)`              | Factorial of a number (`fact(5)` = 120).                                    |
| **Square Root**       | `sqrt(16)`             | Square root of a number (`sqrt(16)` = 4).                                   |
| **Trigonometric**     | `sin(0)`, `cos(0)`     | Sine, cosine, and tangent functions.                                        |
| **Logarithm**         | `log(10)`, `log10(100)`| Natural logarithm (`log`) and base-10 logarithm (`log10`).                  |
| **Exponential**       | `exp(2)`               | Exponential function (`exp(2)` = 7.389).                                   |
| **Power**             | `pow(2, 3)`            | Power function (`pow(2, 3)` = 8).                                           |
| **Absolute Value**    | `abs(-5)`              | Absolute value of a number (`abs(-5)` = 5).                                 |
| **Ceiling**           | `ceil(3.2)`            | Round a number up to the nearest integer (`ceil(3.2)` = 4).                 |
| **Floor**             | `floor(3.8)`           | Round a number down to the nearest integer (`floor(3.8)` = 3).              |
| **Round**             | `round(3.5)`           | Round a number to the nearest integer (`round(3.5)` = 4).                   |
| **Minimum**           | `min(5, 10)`           | Minimum of two numbers (`min(5, 10)` = 5).                                  |
| **Maximum**           | `max(5, 10)`           | Maximum of two numbers (`max(5, 10)` = 10).                                 |
| **Variables**         | `A = 5; A + 3`         | Assign variables and use them in expressions.                               |
| **Multiple Statements**| `A = 5; B = 7; A + B`  | Evaluate multiple statements separated by semicolons.                       |
| **Polynomial Roots**  | `polynomial roots "x^2 - 3x + 2"` | Find the roots of a polynomial.                                           |
| **Polynomial Factorization** | `polynomial factorize "x^2 - 3x + 2"` | Factorize a polynomial into irreducible factors.                     |
| **Polynomial Interpolation** | `polynomial interpolate 1 2 3 4` | Interpolate a polynomial given a set of points.

---