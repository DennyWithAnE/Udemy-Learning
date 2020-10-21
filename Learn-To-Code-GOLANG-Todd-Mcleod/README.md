Notes for future reference.

Padding and Architectural alignment:
- If we are prioritising performance, we should name our fields in the struct from the largest to the smallest. EG.. ( int64 , int64, int32, float32, bool )

-In most cases go for readability first in the arrangement of your struct

Unfurling a slice:

- If we are required to pass in a []int into a function that only takes in int as the arguments we can use the ... notation to unfurl the slices and pass them in.
-EG....     func sum ( values . . . int  ) int ,  x :=  [ ] int { 2, 3, 4, 5, 6}
-We can solve it using this way : sum ( x . . . ) 
-This way the values will be able to pass into the function for calculation

Pointers:
- we use pointers when there is a huge chunk of data needed to be passed into a certain area, instead of passing all the data in we can use pointer to pass in the address where the data is being stored to optimise the performance of the system
- we could also use pointer to change the data at a certain memory address 
- such examples in an application would be when a user wants to change their details...

*** To Dereference a struct we use the (*value).field 
As an exception, if the type x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is a shorthand for (*x).f

JSON Marshaling and Unmarshaling 
- an example of Marshaling:  https://play.golang.org/p/h3YCTfpJXV6
- when you want to take data into your program and turn it into JSON and assign it to a variable we would use Marshal
- an exmaple of Unmarshaling:    https://play.golang.org/p/Ahenqx7OQRs
- similarly if you have data coming into your program and you need to turn it into GO data structures
we would use UnMarshal
Websites to convert JSON to GO : https://mholt.github.io/json-to-go/
 
