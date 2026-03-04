import (
	""
)
type TockenBucketLimiter struct {
	mu sync.Mutex
	tokens float64
	maxTokens float64
	refillRate float64
	lastRefillTime time.Time
}

function NewTockenBucketLimiter(maxTokens float64, refillRate float64) *TockenBucketLimiter {
	return &TockenBucketLimiter{
       token:maxTokens,
	   maxTokens:maxTokens, 
	   refillRate:refillRate,
	   lastRefillTime:time.Now(),
	}
}

func (obj* TockenBucketLimiter) Allow() bool { // obj is the pointer to the struct TockenBucketLimiter
obj.mu.Lock() // why lock ? ans: The mutex lock is used to ensure that only 
// one goroutine can access the critical section of code that modifies
//  the token bucket at a time. This is necessary to prevent race conditions
defer obj.mu.Unlock() // why defer ? ans: The defer statement is used to ensure that the 
// mutex is unlocked even if an error occurs or if the function returns early.
//  By deferring the Unlock() call, we guarantee that the mutex will be 
// released when the function exits, preventing potential deadlocks and 
// ensuring that other goroutines can access the critical section of code 
// protected by the mutex.

now:= time.Now()
elapsedTime:= now.Sub(obj.lastRefillTime).Seconds()
//sub()? ans: The Sub() method is used to calculate the duration between two time values. In this case, it calculates the elapsed time in seconds since the last refill of tokens. The result is a float64 value representing the number of seconds that have passed since the last refill.
logger.infof("Elapsed time since last refill: %f seconds", elapsedTime)
obj.tokens= obj.tokens+ elapsedTime*obj.refillRate
if obj.tokens > obj.maxTokens {
	obj.tokens= obj.maxTokens
}
obj.lastRefillTime= now
if obj.tokens>=1{
	obj.tokens-=1
	return true
}
return false
}