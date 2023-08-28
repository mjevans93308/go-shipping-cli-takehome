## Platform Science Go Project
### Problem Prompt
This is just the prompt given via PDF. If you are already familiar with it, feel free to skip to [Getting Started](#getting-started).

Our sales team has just struck a deal with Acme Inc to become the exclusive provider for routing their product shipments via 3rd party trucking fleets. The catch is that we can only route one shipment to one driver per day.

Each day we get the list of shipment destinations that are available for us to offer to drivers in our network. Fortunately our team of highly trained data scientists have developed a mathematical model for determining which drivers are best suited to deliver each shipment.
With that hard work done, now all we have to do is implement a program that assigns each shipment destination to a given driver while maximizing the total suitability of all shipments to all drivers.

The top-secret algorithm is:
- If the length of the shipment's destination street name is even, the base suitability score (SS) is the number of vowels in the driver’s
name multiplied by 1.5.
- If the length of the shipment's destination street name is odd, the base SS is the number of consonants in the driver’s name multiplied
by 1.
- If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, the SS is increased by 50% above the base SS.

For example, if provided a driver file with Daniel Davidson on one line and an address file with 44 Fake Dr., San Diego, CA 92122 on a line, that pairing’s suitability score would be 9 .

Write an application in the language of your choice that assigns shipment destinations to drivers in a way that maximizes the total SS over the set of drivers. Each driver can only have one shipment and each shipment can only be offered to one driver. Your program should run on the command line and take as input two newline separated files, the first containing the street addresses of the shipment destinations and the second containing the names of the drivers. The output should be the total SS and a matching between shipment destinations and drivers. You do not need to worry about malformed input, but you should certainly handle both upper and lower case names.

Deliverable Your app:

May make use of any existing open source libraries Send us:

The full source code, including any code written which is not part of the normal program run (e.g. build scripts)
Clear instructions on how to build/run the app

Please provide any deliverable and instructions using a public Github (or similar) repository as several people will need to inspect the solution

Evaluation

The point of the exercise is for us to see:

- Code craftsmanship
- How you think about and solve a problem
- How you explain the approach you took and the assumptions you made

We will especially consider:
- Code organisation
- Code readability
- Quality of instructions

### Getting Started
From the project's root directory, run the following commands to build the go executable and then run the CLI itself. Two sample files have been supplied in the `examples` folder for quick running:
```shell
go build -o platsci

./platsci process -n examples/names.txt -a examples/addresses.txt [-o output.txt]
```

To run all unit tests:
```shell
go test -v ./...
```


To see usage for the CLI:
```shell
./platsci -h
./platsci process -h
```

### CLI Format
There are 2 required parameters [`-n | --name`, `-a | --addresses`] that correspond to the files you are using to hold the names and addresses that we want to map to a Suitability Score.

This is 1 non-required parameter `[-o | --output]` that corresponds to the name of the file you would like the results saved to. This filename must have a `.txt` extension for now, and will have save the data in a newline-deliminated format. If this parameter is skipped, the data is printed out to the console.

Example output:
```shell
Data:
1. Charlotte Robinson: 654 Kiwi Lane, Detroit, MI, 48201 - 16.500000
2. Robert Thompson: 210 Grape Drive, Orlando, FL, 32801 - 15.000000
3. Daniel Moore: 456 Elm Avenue, Los Angeles, CA, 90001 - 13.500000
4. Isabella Lewis: 876 Birch Court, Denver, CO, 80201 - 13.500000
5. Amelia Young: 987 Papaya Avenue, Portland, OR, 97201 - 13.500000
6. Thomas Jackson: 890 Spruce Road, Atlanta, GA, 30301 - 13.500000
```

### Next steps
Here are some follow up ideas for where to take this project in the future
- support other file output types, like `.csv`
- investigate whether using Go's new structured log [`slog`](https://go.dev/blog/slog) makes sense here, so we can get rid of the `zap` external logging library
- add integration tests to test file upload as well as our processing algorithm
- add [race condition tests](https://go.dev/doc/articles/race_detector), always a good idea when any concurrency is in play