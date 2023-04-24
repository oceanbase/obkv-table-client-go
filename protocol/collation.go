package protocol

const (
	CsTypeInvalid          = 0
	CsTypeUtf8mb4GeneralCi = 45 // default, case insensitivity
	CsTypeUtf8mb4Bin       = 46 // case sensitive
	CsTypeBinary           = 63
	CsTypeCollationFree    = 100
	CsTypeMax              = 101
)

type ObCollationType struct {
	value int
}

func (t *ObCollationType) Value() int {
	return t.value
}

func (t *ObCollationType) String() string {
	var collationTypeStr string
	switch t.value {
	case CsTypeUtf8mb4GeneralCi:
		collationTypeStr = "CsTypeUtf8mb4GeneralCi"
	case CsTypeUtf8mb4Bin:
		collationTypeStr = "CsTypeUtf8mb4Bin"
	case CsTypeBinary:
		collationTypeStr = "CsTypeBinary"
	case CsTypeCollationFree:
		collationTypeStr = "CsTypeCollationFree"
	case CsTypeMax:
		collationTypeStr = "CsTypeMax"
	default:
		collationTypeStr = "CsTypeInvalid"
	}
	return "ObCollationType{" +
		"collationType:" + collationTypeStr +
		"}"
}

func NewObCollationType(value int) ObCollationType {
	switch value {
	case CsTypeInvalid:
		return ObCollationType{CsTypeInvalid}
	case CsTypeUtf8mb4GeneralCi:
		return ObCollationType{CsTypeUtf8mb4GeneralCi}
	case CsTypeUtf8mb4Bin:
		return ObCollationType{CsTypeUtf8mb4Bin}
	case CsTypeBinary:
		return ObCollationType{CsTypeBinary}
	case CsTypeCollationFree:
		return ObCollationType{CsTypeCollationFree}
	case CsTypeMax:
		return ObCollationType{CsTypeMax}
	default:
		return ObCollationType{CsTypeInvalid}
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

func (l *ObCollationLevel) String() string {
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
