## Overview

It's my solution for the problem 1.2 "Іmplement minimum and maximum values". This is the app that takes 5 numbers from cmd argumnets and returns maximum and minimum possible sum that can be calculated by summing precisely 4 of 5 elements. Also it prints min and max sum to stdout.
It was unit tested by autogenerated unit test.

## How to run

### App

Clone the repository:

    git clone https://git.foxminded.ua/foxstudent106264/task-1.2.git

Run app(instead of 1 2 3 4 5 you can put your numbers as argument for an array):

    go run min_max.go 1 2 3 4 5

Or if you want to run the tests:

    go test

To run benchmark:

    go test -bench=. -benchmem

### Using as a library

`go get` it:

    go get git.foxminded.ua/foxstudent106264/task-1.2

And import to your project to use the function:

    import "git.foxminded.ua/foxstudent106264/task-1.2"
