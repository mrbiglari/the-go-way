## Testing in Go

Unit testing is a critical discipline in software engineering, aimed at verifying the correctness of individual units of code in isolation. In Go (Golang), the built-in `testing` package provides a lightweight and effective mechanism for writing and executing unit tests. While this standard package supports basic assertions and test functions, it lacks more advanced features such as expressive assertions, mock generation, and behavioural verification — which are essential for testing components that depend on interfaces or external systems.

To bridge this gap, the [**`testify`**](https://github.com/stretchr/testify) library provides an extended toolkit for Go testing. It includes assertion utilities and a powerful mocking framework that integrates seamlessly with Go’s type system.

This document outlines best practices and technical strategies for unit testing and mocking in Go using the standard `testing` package and 3rd party `testify/mock` package.

## Standard Library: `testing`

Go's `testing` package provides the foundation for unit testing. A test function in Go is defined with the following signature:

```go
func TestFunctionName(t *testing.T) {
    // test logic
}
```

Assertions are typically implemented manually, using standard control structures:

```go
if result != expected {
    t.Errorf("expected %v, got %v", expected, result)
}
```

While effective for basic test scenarios, these manual assertions can become verbose and less readable in more complex tests. Additionally, `testing` does not provide native support for mocking dependencies.

### Subtests with `t.Run`

Subtests enable developers to organise related test logic into independently executable units within a single test function. Using `t.Run`, each subtest receives its own `*testing.T` context, which allows for clearer test output, better failure isolation, and the ability to selectively run specific cases using `go test -run`.

#### Example

```go
func TestStringOperations(t *testing.T) {
	t.Run("ToUpper", func(t *testing.T) {
		result := ToUpper("go")
		if result != "GO" {
			t.Errorf("expected GO, got %s", result)
		}
	})

	t.Run("TrimSpace", func(t *testing.T) {
		result := TrimSpace("  go  ")
		if result != "go" {
			t.Errorf("expected 'go', got '%s'", result)
		}
	})
}
```

### Table-Driven Testing

Table-driven testing is a common pattern in Go for validating multiple scenarios with minimal duplication. It involves defining a slice of test cases and iterating through them using a loop. This pattern becomes especially effective when combined with subtests to isolate and label each case.

#### Example

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"zero values", 0, 0, 0},
		{"negative numbers", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}
```

## Enhanced Testing with `testify`

The `testify` library introduces the following capabilities:

- Rich **assertion** functions (e.g., `assert.Equal`, `assert.NoError`)
- A **mocking framework** (`testify/mock`) that supports:
  - Method expectations and return values
  - Argument matching
  - Call and call count assertions

### Creating and Using Mocks with `testify`

To mock a Go interface with `testify`, one must define a custom mock struct that embeds `mock.Mock` and implements the interface explicitly. Go does not support dynamic or runtime mock generation as seen in languages like Java or Python.

#### Example Interface and Implementation

```go
type State int

const (
	Happy State = iota
	Hungry
	Sad
	Sleepy
)

type ControllableCharacter interface {
	GetState() State
	SetState(state State)
}

type Character struct {
	Mood State
}

func (c *Character) GetState() State { return c.Mood }
func (c *Character) SetState(state State) { c.Mood = state }
```

#### Corresponding Mock

```go
type MockCharacter struct {
	mock.Mock
}

func (mc *MockCharacter) GetState() State {
	args := mc.Called()
	return args.Get(0).(State)
}

func (mc *MockCharacter) SetState(state State) {
	mc.Called(state)
}
```

### Behavioural Verification

The `testify/mock` library provides multiple ways to verify method calls:

| Purpose                          | Method                                  |
| -------------------------------- | --------------------------------------- |
| Verify method was called         | `AssertCalled(t, "Method", args...)`    |
| Verify method was **not** called | `AssertNotCalled(t, "Method", args...)` |
| Verify exact call count          | `AssertNumberOfCalls(t, "Method", n)`   |
| Verify all expected calls        | `AssertExpectations(t)`                 |

#### Example Test

```go
func TestMakesHappyIfHungry(t *testing.T) {
	mockChar := new(MockCharacter)
	mockChar.On("GetState").Return(Hungry)
	mockChar.On("SetState", Happy).Return()

	game := &Game{cc: mockChar}
	game.FeedPlayer()

	mockChar.AssertCalled(t, "SetState", Happy)
	mockChar.AssertNumberOfCalls(t, "GetState", 1)
	mockChar.AssertExpectations(t)
}
```

<details>
<summary>⭕ Critique of Go </summary>
Unlike mocking frameworks in Java or C#, mocking in Go requires passing method names as *strings* when setting up expectations and assertions.

### Setting Expectations with testify/mock

In `testify`, expectations are declared by calling `On(...)` with the method name as a string, followed by the expected arguments:

```go
mockCharacter.On("SetState", Happy).Return()
```

Assertions on method calls are made in a similar way:

```go
mockCharacter.AssertCalled(t, "SetState", Happy)
```

This string-based approach introduces a potential _maintenance risk_: method names are not validated by the compiler. If the method name changes in the interface but not in the test, the test will not fail at compile time. Instead, it may panic at runtime or silently produce incorrect results.

### Why Go Mocks Use Strings Instead of Method References

Although Go supports reflection, higher-order functions, and function values, it lacks several features that enable more expressive mocking in other languages. Go does not provide a way to extract the method name from a function value, nor does it support automatic discovery of method invocations at runtime. Let’s say we write:

```go
fn := myStruct.SetState
```

Now we have a function that takes a `State`. But Go has no way to say:

> "This function is the `SetState` method on `*MyStruct`"

There's no metadata attached to `fn` — it’s just a behaviour.

### Language Limitations That Impact Mocking in Go

| Limitation                    | Description                                                            |
| ----------------------------- | ---------------------------------------------------------------------- |
| No expression trees           | Go cannot introspect lambdas or function bodies to detect method calls |
| No runtime proxies            | Go lacks a mechanism to dynamically wrap interfaces at runtime         |
| Function values lose identity | Passing a method as a function does not preserve its name or receiver  |
| No annotations or metadata    | There is no way to tag or inspect method semantics as metadata         |

### How This Differs from Java and C#

#### Java (Mockery)

In Java, mocking frameworks like Mockery and Mockito allow developers to define expectations using method calls directly:

```java
context.checking(new Expectations() {{
    oneOf(mockedObject).setState(HAPPY);
}});
```

These frameworks use bytecode instrumentation and runtime reflection to record the method being invoked and verify it later. There is no need to pass method names as strings. The mocking framework automatically extracts method names, arguments, and types during execution, resulting in more maintainable and type-safe test code.

#### C# (Moq)

In C#, the Moq library supports expression-based mocking:

```csharp
mock.Setup(m => m.SetState(Happy));
mock.Verify(m => m.SetState(Happy), Times.Once);
```

In this case, `m => m.SetState(Happy)` is not just a function. It is an expression that the framework can inspect to extract the method name and parameters. This enables compile-time validation and reduces the risk of incorrect test definitions.

### Why Go Follows a Different Philosophy

Go was designed with a preference for simplicity, clarity, and explicit behaviour. It avoids advanced metaprogramming features such as runtime proxies, expression trees, and annotations. These design choices limit the flexibility of mocking libraries but result in more predictable and readable code.

As a result, mocking in Go with `testify/mock` is more manual:

- The developer must write mock implementations explicitly
- Method names must be passed as strings
- Expectations are not validated at compile time

While this approach is less automated, it aligns with Go's philosophy of transparency and minimalism.

---

---

</details>

### Mock Flexibility and Limitations

Unlike some object-oriented mocking frameworks (e.g., Mockito for Java), `testify/mock` is **strict by default**. Any method called during a test must have a corresponding `.On(...).Return(...)` definition. Failure to do so results in a runtime panic.

There is **no built-in concept of "allow" or "ignore"** for method calls. To permit a method to be called without asserting it, the developer must still define an expectation via `.On(...)` and omit any assert functions. There is also no global toggle for "strict" vs. "lenient" mocking modes.

This approach, while more verbose, aligns with Go's emphasis on **explicitness and clarity**.

### Recommendations

- Always define `.On(...).Return(...)` for any method that might be called on a mock.
- Use `AssertExpectations(t)` at the end of your test to ensure no declared expectation was missed.
- For methods that might be called but do not affect test correctness, declare them but do not assert on them.
- Consider using tools like `mockery` or `moq` to generate mock implementations for large interfaces.

## Summary

Unit testing in Go is effectively supported by the standard `testing` package, but for advanced use cases — particularly those involving dependencies and interfaces — the `testify` library provides essential extensions. Although `testify/mock` enforces strict mocking by default, this behaviour encourages more intentional and reliable test design. Through careful use of expectations and assertions, developers can write expressive, behaviourally accurate tests that are idiomatic to Go’s design philosophy.
