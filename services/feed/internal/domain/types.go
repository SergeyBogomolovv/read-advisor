package domain

type PreferenceType string

const (
	PreferenceTypeDefault PreferenceType = "default"
	PreferenceTypeLiked   PreferenceType = "liked"
	PreferenceTypeSaved   PreferenceType = "saved"
)

type Preference string

const (
	PreferenceAuthor    Preference = "author"
	PreferenceCategory  Preference = "category"
	PreferencePublisher Preference = "publisher"
)
