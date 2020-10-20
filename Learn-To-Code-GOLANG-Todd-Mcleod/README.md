Notes for future reference.

Padding and Architectural alignment 
- If we are prioritising performance, we should name our fields in the struct from the largest to the smallest. EG.. ( int64 , int64, int32, float32, bool )

-In most cases go for readability first in the arrangement of your struct

Unfurling a slice

- If we are required to pass in a []int into a function that only takes in int as the arguments we can use the ... notation to unfurl the slices and pass them in.
-EG....     func sum ( values . . . int  ) int ,  x :=  [ ] int { 2, 3, 4, 5, 6}
-We can solve it using this way : sum ( x . . . ) 
-This way the values will be able to pass into the function for calculation
