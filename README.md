# JSDoc Inserter

Usefull to insert a repeatetive JSDoc line in all JSDoc blocks of a file.

```sh
❯ cat test_replace.js
/**
 * @param {number} a - The first number
 * @param {number} b - The second number
 * @returns {number} The sum of the two numbers
 */
function add(a, b) {
    return a + b;
}
❯ go run main.go test_replace.js "@description Hello World"
/**
 * @param {number} a - The first number
 * @param {number} b - The second number
 * @returns {number} The sum of the two numbers
 * @description Hello World
 */
function add(a, b) {
    return a + b;
}
❯ cat test_replace_new.js
/**
 * @param {number} a - The first number
 * @param {number} b - The second number
 * @returns {number} The sum of the two numbers
 * @description Hello World
 */
function add(a, b) {
    return a + b;
}
```

## Usage
```
go run main.go <filename.js> <line_to_add>
```