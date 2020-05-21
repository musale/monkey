# Monkey

Monkey is a C-like interpreter for the `monkey language` that is built in the book ["Writing an interpreter in Go"](https://interpreterbook.com) by Thorsten Ball.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have knowledge of programming and can do so in Go.
- You have Go v1.7+ installed.

## Language Features

- [x] Function literals.
- [x] First class and Higher Order Functions.
- [x] Closures.
- [x] Variable Bindings.
- [x] Prefix and Infix operations.
- [x] Builtin Data types;
  - [x] Booleans.
  - [x] Strings.
  - [x] Hashes.
  - [x] Integers.
  - [x] Arrays.
- Builtin functions
  - [x] `len` for length of iterables.
  - [x] `first` for first element in an array.
  - [x] `last` for the last element in an array.
  - [x] `rest` for the elements in an array except the first.
  - [x] `push` to add a new element to the end of the array.

## Personal TODOs:

- Read source code input from files.
- Support Unicode characters.
- Support floating point numbers.
- Postfix operations [?].
- Improve Syntax errors and stack errors.
- Add support for multiline strings.
- Improve string lexing i.e. check for opening and closing double quotes.

## Installing monkey

To install monkey, follow these steps:

- Clone this repository and then:

  `$ cd monkey`

- And then build a binary

  `$ go build cmd/* -o monkey`

- Finally run the binary

  `$ ./monkey`

## Using monkey

To use monkey, follow these steps:

```
$ ./monkey
Hello mus!
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
>> 1 + 1
3
>> let newAdder = fn(x) { fn(y) { x + y } };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
```

## Contributing to monkey

To contribute to monkey, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull a request.

## Contact

If you want to contact me you can reach me at martinmshale@gmail.com.

## License

This project uses the following license: [MIT](/LICENSE).
