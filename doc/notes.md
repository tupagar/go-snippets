- Array and a **slice** both share the same storage space.
- length of an array is always fixed but length of slice can keep changing
- `make([]int, 50, 100)` >> `make([]T, length, capacity)`
    - is an inbuilt function which helps creating slices. this example creates an array of size 100 and returns a slice of length 50 from it

- tag is an optional string literal following a field declaration. `Example: myfield [10]uint32 "my tag name"`
    - this tag becomes an attribute of each field in the corresponding field declaration

- `**veriadic**` >>> it can be a final function parameter with ... prefixed type. `Example: func(<other params>, field ...int) bool`

- all types implement empty **interface** `interface{}`
- a type can implement more than one interfaces
- can embed one interface inside another interface:

    `    type abc interface {... some methods...}  
        type xyz interface { abc; somemethod()}
    `
    - cannot embed itself or any interface causing loop recursively

- `map[keyType]valueType` >>> keyType should implement operations == and !=
    - can create instance of map using make. Example: `make(map[string]int, 100)` >>> size can still grow beyong 100

- **Channel type:**
    - channel provides a mechanism for concurrently executing functions to communicate to each other using send and receive 
    - examples: 
        - `chan T     // used to send and receive values of type T`
        - `chan<- int  // to send ints`
        - `<-chan      // to receive ints`
    - channel can be made of channel! (chan<- chan int  //channel for sending channel of integers!!!)
    - can create instance by `make(chan, int, 100)`    
    - if capacity == 0 channel is called unbuffered, both sender and receiver should be ready same time
    - else can fill until the buffer is not full
    - FIFO queues

- Exporting the identifiers:
    - The identifier is exported if it starts with Capital letter and is defined at package level

- **IOTA constant:**
    - 
    