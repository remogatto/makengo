<tt>Makengo</tt> (make'n go) is a build program written in go and inspired to
ruby rake.

# Preamble

When I switched from ruby to go there was a thing I particularly
missed: a build program à la rake. So, I decided to make an attempt
trying to roughly reproduce rake's dsl in go and I ended up with what
follows.

Please note that this software is in an early alpha stage of
development. If you find it interesting feel free to send me feedbacks
or better to fork it ;)

# Quick start

    $ git clone 
    $ cd makengo
    $ make install

# The DSL

## Define a simple task without dependencies

    Task("Hello", func() { fmt.Println("Hello!") })

## Define two dependent tasks (order of definition doesn't matter)

    Task("Joe", func() { fmt.Println("Joe") })
    Task("Hello", func() { fmt.Println("Hello ") }).DependsOn("Joe")

## Add descriptions to tasks

    Describe("Print hello.")
    Task("Hello", func() { fmt.Println("Hello!") })

## Define a default task among a set of defined tasks (not yet implemented)

    Task("Hello", func() { fmt.Println("Hello!") })
    Default("Hello")

# Makengo file

Tasks are defined inside a file named Makengo and embraced by a init()
function.

## Makengo file example

    package main
    import ( "fmt" "makengo" )

    func init() {
            Describe("Print hello.")
            Task("Hello", func() { fmt.Println("Hello!") })
    }

# Tasks execution and command-line

Tasks are invoked using makengo executable:

    $ makengo # Run the default task if any
    $ makengo Hello # Output "Hello!"
    $ makengo Hello Joe # Output "Hello Joe"
    $ makengo -T # Show task descriptions

# Concurrency

I desired to look at one of go's neatest feature: goroutines. Concurrency 
is exploited following these rules:

1. Independent tasks run concurrently. 

2. If task1 depends on (task2, task3) and (task2, task3) are independent 
tasks then (task2, task3) run concurrently and task1 waits for (task2, 
task3) to finish their job.


