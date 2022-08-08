package graphql

import "strconv"

// OptionType represents the logic of graphql query construction
type OptionType string

const (
	// optionTypeOperationName is private because it's option is built-in and unique
	optionTypeOperationName      OptionType = "operation_name"
	OptionTypeOperationDirective OptionType = "operation_directive"

	optionTypePersistentOperationName    OptionType = "persistent_query_name"
	optionTypePersistentOperationVersion OptionType = "persistent_query_version"
)

// Option abstracts an extra render interface for the query string
// They are optional parts. By default GraphQL queries can request data without them
type Option interface {
	// Type returns the supported type of the renderer
	// available types: operation_name and operation_directive
	Type() OptionType
	// String returns the query component string
	String() string
}

// operationNameOption represents the operation name render component
type operationNameOption struct {
	name string
}

func (ono operationNameOption) Type() OptionType {
	return optionTypeOperationName
}

func (ono operationNameOption) String() string {
	return ono.name
}

// OperationName creates the operation name option
func OperationName(name string) Option {
	return operationNameOption{name}
}

type persistentQueryOptionName struct {
	name string
}

func (pon persistentQueryOptionName) Type() OptionType {
	return optionTypePersistentOperationName
}

func (pon persistentQueryOptionName) String() string {
	return pon.name
}

// PersistentQueryName creates the persistent operation name option
func PersistentQueryName(name string) Option {
	return persistentQueryOptionName{name}
}

type persistentQueryOptionVersion struct {
	version string
}

func (pon persistentQueryOptionVersion) Type() OptionType {
	return optionTypePersistentOperationVersion
}

func (pon persistentQueryOptionVersion) String() string {
	return pon.version
}

// PersistentQueryVersion creates the persistent operation version option
func PersistentQueryVersion(version int) Option {
	return persistentQueryOptionVersion{strconv.Itoa(version)}
}
