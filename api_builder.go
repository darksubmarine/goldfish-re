package goldfish_re

// builder_ builder struct that works as entrypoint for the library users
type builder_ struct{}

// Builder builder_ constructor
func Builder() *builder_ {
	return &builder_{}
}

// Ruleset returns a new rulesetBuilder
func (b *builder_) Ruleset() *rulesetBuilder { return newRulesetBuilder() }

// Rule returns a new ruleBuilder
func (b *builder_) Rule() *ruleBuilder { return newRuleBuilder() }

// condition returns a new conditionBuilder
func (b *builder_) condition() *conditionBuilder { return newConditionBuilder() }

// StringCondition returns a new stringConditionBuilder
func (b *builder_) StringCondition() *stringConditionBuilder { return newStringConditionBuilder() }

// NumberCondition returns a new numberConditionBuilder
func (b *builder_) NumberCondition() *numberConditionBuilder { return newNumberConditionBuilder() }

// FloatCondition returns a new floatConditionBuilder
func (b *builder_) FloatCondition() *floatConditionBuilder { return newFloatConditionBuilder() }

// BooleanCondition returns a new booleanConditionBuilder
func (b *builder_) BooleanCondition() *booleanConditionBuilder { return newBooleanConditionBuilder() }

// DateCondition returns a new dateConditionBuilder
func (b *builder_) DateCondition() *dateConditionBuilder { return newDateConditionBuilder() }
