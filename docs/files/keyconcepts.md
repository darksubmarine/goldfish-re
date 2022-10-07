
## Fact
A fact is a representatioon of a variable that stores a value.

 - Each fact represents an object's attribute, like `User.birthday`
 - Each fact only can be defined as one of the supported types: `String` `Number` `Float` `Boolean` `Date`

## Term
A term represents one side of a condition, could be a fact or could be a discrete value.

## Condition
The condition is a boolean sentence which after its evaluation the result can be `true` or `false`.

 - Ex. `User.birthday > 1984-10-22`
 - A condition must have a Fact term as left term.
    - Invalid: `2000 >= Trip.miles`
    - Valid: `Trip.miles <= 2000`

## Rule
The rule is a collection of conditions under the operator `All` or `Any`

 - All: The rule will be valid (true) when all contained conditions are `true` after its evaluation
 - Any: The rule will be valid (true) when at least one contained condition is `true` after its evaluation

## Ruleset
The ruleset is a group of rules defined by the user. 

The ruleset exposes a context where the facts are registered and evaluated on each rule into the ruleset when a change happens

## Context
A context groups the facts involved on the rule activation that belongs to a ruleset. 


