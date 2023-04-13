package protocol

const (
	csTypeInvalid          = 0
	csTypeUtf8mb4GeneralCi = 45 // default, case insensitivity
	csTypeUtf8mb4Bin       = 46 // case sensitive
	csTypeBinary           = 63
	csTypeCollationFree    = 100
	csTypeMax              = 101
)

type ObCollationType struct {
	value int
}

func (t *ObCollationType) ToString() string {
	var collationTypeStr string
	switch t.value {
	case csTypeUtf8mb4GeneralCi:
		collationTypeStr = "csTypeUtf8mb4GeneralCi"
	case csTypeUtf8mb4Bin:
		collationTypeStr = "csTypeUtf8mb4Bin"
	case csTypeBinary:
		collationTypeStr = "csTypeBinary"
	case csTypeCollationFree:
		collationTypeStr = "csTypeCollationFree"
	case csTypeMax:
		collationTypeStr = "csTypeMax"
	default:
		collationTypeStr = "csTypeInvalid"
	}
	return "ObCollationType{" +
		"collationType:" + collationTypeStr +
		"}"
}

func NewObCollationType(value int) ObCollationType {
	switch value {
	case csTypeInvalid:
		return ObCollationType{csTypeInvalid}
	case csTypeUtf8mb4GeneralCi:
		return ObCollationType{csTypeUtf8mb4GeneralCi}
	case csTypeUtf8mb4Bin:
		return ObCollationType{csTypeUtf8mb4Bin}
	case csTypeBinary:
		return ObCollationType{csTypeBinary}
	case csTypeCollationFree:
		return ObCollationType{csTypeCollationFree}
	case csTypeMax:
		return ObCollationType{csTypeMax}
	default:
		return ObCollationType{csTypeInvalid}
	}
}

const (
	csLevelExplicit  = 0
	csLevelNone      = 1
	csLevelImplicit  = 2
	csLevelSysConst  = 3
	csLevelCoercible = 4
	csLevelNumeric   = 5
	csLevelIgnorable = 6
	csLevelInvalid   = 127
)

type ObCollationLevel struct {
	value int
}

func (l *ObCollationLevel) ToString() string {
	var collationLevelStr string
	switch l.value {
	case csLevelExplicit:
		collationLevelStr = "csLevelExplicit"
	case csLevelNone:
		collationLevelStr = "csLevelNone"
	case csLevelImplicit:
		collationLevelStr = "csLevelImplicit"
	case csLevelSysConst:
		collationLevelStr = "csLevelSysConst"
	case csLevelCoercible:
		collationLevelStr = "csLevelCoercible"
	case csLevelNumeric:
		collationLevelStr = "csLevelNumeric"
	case csLevelIgnorable:
		collationLevelStr = "csLevelIgnorable"
	default:
		collationLevelStr = "csLevelInvalid"
	}
	return "ObCollationLevel{" +
		"collationLevel:" + collationLevelStr +
		"}"
}

func newObCollationLevel(value int) ObCollationLevel {
	switch value {
	case csLevelExplicit:
		return ObCollationLevel{csLevelExplicit}
	case csLevelNone:
		return ObCollationLevel{csLevelNone}
	case csLevelImplicit:
		return ObCollationLevel{csLevelImplicit}
	case csLevelSysConst:
		return ObCollationLevel{csLevelSysConst}
	case csLevelCoercible:
		return ObCollationLevel{csLevelCoercible}
	case csLevelNumeric:
		return ObCollationLevel{csLevelNumeric}
	case csLevelIgnorable:
		return ObCollationLevel{csLevelIgnorable}
	default:
		return ObCollationLevel{csLevelInvalid}
	}
}
