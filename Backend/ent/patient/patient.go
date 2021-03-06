// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCardID holds the string denoting the card_id field in the database.
	FieldCardID = "card_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAllergic holds the string denoting the allergic field in the database.
	FieldAllergic = "allergic"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"

	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "gender"
	// EdgeTitle holds the string denoting the title edge name in mutations.
	EdgeTitle = "title"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// GenderTable is the table the holds the gender relation/edge.
	GenderTable = "patients"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the gender relation/edge.
	GenderColumn = "gender_id"
	// TitleTable is the table the holds the title relation/edge.
	TitleTable = "patients"
	// TitleInverseTable is the table name for the Title entity.
	// It exists in this package in order to avoid circular dependency with the "title" package.
	TitleInverseTable = "titles"
	// TitleColumn is the table column denoting the title relation/edge.
	TitleColumn = "title_id"
	// JobTable is the table the holds the job relation/edge.
	JobTable = "patients"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldCardID,
	FieldFirstName,
	FieldLastName,
	FieldAllergic,
	FieldAddress,
	FieldAge,
	FieldBirthday,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Patient type.
var ForeignKeys = []string{
	"gender_id",
	"job_id",
	"title_id",
}

var (
	// CardIDValidator is a validator for the "Card_id" field. It is called by the builders before save.
	CardIDValidator func(string) error
	// FirstNameValidator is a validator for the "First_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "Last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// AllergicValidator is a validator for the "Allergic" field. It is called by the builders before save.
	AllergicValidator func(string) error
	// AddressValidator is a validator for the "Address" field. It is called by the builders before save.
	AddressValidator func(string) error
)
