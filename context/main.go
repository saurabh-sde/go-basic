package main

import (
	"context"
	"fmt"
	"time"
)

// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

type ConfigKey string

type Config struct {
	APIKey    string
	Timeout   int
	DebugMode bool
}

/*
Context used to pass deadlines, cancellation and other values across different scopes without depending on func parameters
- context.TODO()
- context.Backgroud()
- context.WithValue(parentctx, key, value)
- context.WithDeadline(parentCtx, deadline) -> returns cancel() func along with ctx to call ctx.Done() channel
- context.WithTimeout(parentCtx, duration)
*/
func main() {
	ctx := context.Background() // generally used as default instead of using context.TODO()

	// Create a Config instance
	appConfig := Config{
		APIKey:    "your_api_key",
		Timeout:   10,
		DebugMode: true,
	}
	// ctx with Value to be used to store configs, user id etc
	ctx = context.WithValue(ctx, ConfigKey("config"), appConfig)
	fmt.Printf("context: %+v\ncontext value: %+v\n", ctx, ctx.Value(ConfigKey("config")))

	// ctx with deadline of 2s
	deadline := time.Now().Add(2 * time.Second)
	fmt.Println("deadline: ", deadline)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	// Simulate some work that takes 10 seconds but deadline is less
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Task completed.")
	case <-ctx.Done():
		// This case is executed when the context is canceled due to the deadline
		fmt.Println("Task canceled due to deadline.")
	}
}

/*
select statement allows you to wait on multiple channels and proceed with the case that becomes ready first

	case: time.After(10 * time.Second)
	The time.After function returns a channel that sends a single value after the specified duration (in this case, 10 seconds).
	The <-time.After(10 * time.Second) part waits for this channel to produce a value (after 10 seconds).
	If the 10 seconds elapse before the ctx.Done() case becomes ready, the code within this case will execute, indicating that the task has completed within the specified time.

	case:  <-ctx.Done()
	ctx.Done() returns a channel that's closed when the context is canceled or its deadline is exceeded.
	The <-ctx.Done() part waits for this channel to be closed.
	If the context is canceled before the 10 seconds elapse, the code within this case will execute, indicating that the task was canceled due to the deadline being exceeded.
	This select statement essentially says: "Whichever case is ready first, execute the corresponding code." If the task completes within 10 seconds, the first case is executed; otherwise, if the context is canceled before that, the second case is executed.

- It's a common pattern in Go for handling concurrent operations with timeouts or deadlines. In this case, it helps in handling the completion of a task within a specified time limit.

- Context in Real-World Scenarios
Context in Golang is widely used in various real-world scenarios. Let’s explore some practical examples where context plays a crucial role.

- 	Example: Context in Microservices
	In a microservices architecture, each service often relies on various external dependencies and communicates with other services. Context can be used to propagate important information, such as authentication tokens, request metadata, or tracing identifiers, throughout the service interactions.

- 	Example: Context in Web Servers
	Web servers handle multiple concurrent requests, and context helps manage the lifecycle of each request. Context can be used to set timeouts, propagate cancellation signals, and pass request-specific values to the different layers of a web server application.

	Example: Context in Test Suites
	When writing test suites, context can be utilized to manage test timeouts, control test-specific configurations, and enable graceful termination of tests. Context allows tests to be canceled or skipped based on certain conditions, enhancing test control and flexibility.
*/
