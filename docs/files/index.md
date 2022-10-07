# Goldfish-RE

<span class="dsMainColor">_Reactive_</span> and <span class="dsMainColor">_Embeddable_</span> rules engine library written in pure Go!. 

This rules engine has been thought to trigger automatically an event each time that
a condition from your ruleset has been satisfied by some updated fact into a context.

The evaluation algorithm is RETE-based with focus on evaluation and memory using a `Trie` struct to improve its performance.

## Data types
The rule engines expose different data types to work with:

 - **String**: This is a well known `string` type
 - **Number**: A number is a representation of an integer value. In this case is a wrapper for a `int64` data type
 - **Float**: Floats are also numeric types. They represent the decimal numbers implemented as `float64`
 - **Boolean**: The boolean are useful to assert a condition as `true` or `false`
 - **Date**: Represents a `time.Time` data type


## Key concepts

 - **Fact**: A fact is a representatioon of a variable that stores a value.
    - Each fact represents an object's attribute, like `User.birthday`
    - Each fact only can be defined as one of the supported types: `String` `Number` `Float` `Boolean` `Date`
 - **Term**: A term represents one side of a condition, could be a fact or could be a discrete value.
 - **Condition**: The condition is a boolean sentence which after its evaluation the result can be `true` or `false`.
    - Ex. `User.birthday > 1984-10-22`
    - A condition must have a Fact term as left term.
        - Not valid: `2000 >= Trip.miles`
        - Valid: `Trip.miles <= 2000`
 - **Rule**: The rule is a collection of conditions under the operator `All` or `Any`
    - All: The rule will be valid (true) when all contained conditions are `true` after its evaluation
    - Any: The rule will be valid (true) when at least one contained condition is `true` after its evaluation
 - **RuleSet**: The ruleset is a group of rules defined by the user. The ruleset exposes a context were the facts are registered and evaluated on each rule into the ruleset when a change happens
 - **Context**: A context groups the facts involved on the rule activation that belongs to a ruleset. 


