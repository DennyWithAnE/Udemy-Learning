Notes for future reference.

Padding and Architectural alignment 
- If we are prioritising performance, we should name our fields in the struct from the largest to the smallest. EG.. ( int64 , int64, int32, float32, bool )

-In most cases go for readability first in the arrangement of your struct

