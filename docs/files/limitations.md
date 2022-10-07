## Termination rules

The supported rules are well known as `termination rule` because the `THEN` clause do nothing. 
Only returns a `string` that it's passed to the activation handler to identify which rule has been activated.

## Fact inferation 

The fact inferation is a useful feature of all rule engines. But this mechanism hides a recursive call to
the evaluation method over the same ruleset and sometimes this behavior becomes in a performance problem. In order to 
mitigate this possible issue the `context` exposes a method to apply a hard limit to the recursion (`ctx.WithMaxIterations(10)`).

This must be set at context creation:

```go
rs, err := createRuleset()
if err != nil {
    panic(fmt.Sprintf("rulse cannot be created. Error=%s", err))
}

// Once that you have the ruleset context, you can inject it everywhere.
ctx := rs.Context()
ctx.WithMaxIterations(10) // recursion limit (1)
```

1. Recursion hard limit set by user. Default value: **`const maxIterations = 100`**

Also, the same rule cannot be activated twice if a fact updates matchs again with the previous activated rule to avoid run an infinite loop.