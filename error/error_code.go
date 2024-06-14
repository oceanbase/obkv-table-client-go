/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package error

type ObErrorCode int32

const (
	ObSuccess                                                                     ObErrorCode = 0
	ObError                                                                       ObErrorCode = -4000
	ObObjTypeError                                                                ObErrorCode = -4001
	ObInvalidArgument                                                             ObErrorCode = -4002
	ObArrayOutOfRange                                                             ObErrorCode = -4003
	ObServerListenError                                                           ObErrorCode = -4004
	ObInitTwice                                                                   ObErrorCode = -4005
	ObNotInit                                                                     ObErrorCode = -4006
	ObNotSupported                                                                ObErrorCode = -4007
	ObIterEnd                                                                     ObErrorCode = -4008
	ObIoError                                                                     ObErrorCode = -4009
	ObErrorFuncVersion                                                            ObErrorCode = -4010
	ObTimeout                                                                     ObErrorCode = -4012
	ObAllocateMemoryFailed                                                        ObErrorCode = -4013
	ObInnerStatError                                                              ObErrorCode = -4014
	ObErrSys                                                                      ObErrorCode = -4015
	ObErrUnexpected                                                               ObErrorCode = -4016
	ObEntryExist                                                                  ObErrorCode = -4017
	ObEntryNotExist                                                               ObErrorCode = -4018
	ObSizeOverflow                                                                ObErrorCode = -4019
	ObRefNumNotZero                                                               ObErrorCode = -4020
	ObConflictValue                                                               ObErrorCode = -4021
	ObItemNotSetted                                                               ObErrorCode = -4022
	ObEagain                                                                      ObErrorCode = -4023
	ObBufNotEnough                                                                ObErrorCode = -4024
	ObReadNothing                                                                 ObErrorCode = -4026
	ObFileNotExist                                                                ObErrorCode = -4027
	ObDiscontinuousLog                                                            ObErrorCode = -4028
	ObSerializeError                                                              ObErrorCode = -4033
	ObDeserializeError                                                            ObErrorCode = -4034
	ObAioTimeout                                                                  ObErrorCode = -4035
	ObNeedRetry                                                                   ObErrorCode = -4036
	ObNotMaster                                                                   ObErrorCode = -4038
	ObDecryptFailed                                                               ObErrorCode = -4041
	ObNotTheObject                                                                ObErrorCode = -4050
	ObLastLogRuinned                                                              ObErrorCode = -4052
	ObInvalidError                                                                ObErrorCode = -4055
	ObDecimalOverflowWarn                                                         ObErrorCode = -4057
	ObEmptyRange                                                                  ObErrorCode = -4063
	ObDirNotExist                                                                 ObErrorCode = -4066
	ObInvalidData                                                                 ObErrorCode = -4070
	ObCanceled                                                                    ObErrorCode = -4072
	ObLogNotAlign                                                                 ObErrorCode = -4074
	ObNotImplement                                                                ObErrorCode = -4077
	ObDivisionByZero                                                              ObErrorCode = -4078
	ObExceedMemLimit                                                              ObErrorCode = -4080
	ObQueueOverflow                                                               ObErrorCode = -4085
	ObStartLogCursorInvalid                                                       ObErrorCode = -4099
	ObLockNotMatch                                                                ObErrorCode = -4100
	ObDeadLock                                                                    ObErrorCode = -4101
	ObChecksumError                                                               ObErrorCode = -4103
	ObInitFail                                                                    ObErrorCode = -4104
	ObRowkeyOrderError                                                            ObErrorCode = -4105
	ObPhysicChecksumError                                                         ObErrorCode = -4108
	ObStateNotMatch                                                               ObErrorCode = -4109
	ObInStopState                                                                 ObErrorCode = -4114
	ObLogNotClear                                                                 ObErrorCode = -4116
	ObFileAlreadyExist                                                            ObErrorCode = -4117
	ObUnknownPacket                                                               ObErrorCode = -4118
	ObRpcPacketTooLong                                                            ObErrorCode = -4119
	ObLogTooLarge                                                                 ObErrorCode = -4120
	ObRpcSendError                                                                ObErrorCode = -4121
	ObRpcPostError                                                                ObErrorCode = -4122
	ObLibeasyError                                                                ObErrorCode = -4123
	ObConnectError                                                                ObErrorCode = -4124
	ObRpcPacketInvalid                                                            ObErrorCode = -4128
	ObBadAddress                                                                  ObErrorCode = -4144
	ObErrMinValue                                                                 ObErrorCode = -4150
	ObErrMaxValue                                                                 ObErrorCode = -4151
	ObErrNullValue                                                                ObErrorCode = -4152
	ObResourceOut                                                                 ObErrorCode = -4153
	ObErrSqlClient                                                                ObErrorCode = -4154
	ObOperateOverflow                                                             ObErrorCode = -4157
	ObInvalidDateFormat                                                           ObErrorCode = -4158
	ObInvalidArgumentNum                                                          ObErrorCode = -4161
	ObEmptyResult                                                                 ObErrorCode = -4165
	ObLogInvalidModId                                                             ObErrorCode = -4168
	ObLogModuleUnknown                                                            ObErrorCode = -4169
	ObLogLevelInvalid                                                             ObErrorCode = -4170
	ObLogParserSyntaxErr                                                          ObErrorCode = -4171
	ObUnknownConnection                                                           ObErrorCode = -4174
	ObErrorOutOfRange                                                             ObErrorCode = -4175
	ObOpNotAllow                                                                  ObErrorCode = -4179
	ObErrAlreadyExists                                                            ObErrorCode = -4181
	ObSearchNotFound                                                              ObErrorCode = -4182
	ObItemNotMatch                                                                ObErrorCode = -4187
	ObInvalidDateFormatEnd                                                        ObErrorCode = -4190
	ObDdlTaskExecuteTooMuchTime                                                   ObErrorCode = -4192
	ObHashExist                                                                   ObErrorCode = -4200
	ObHashNotExist                                                                ObErrorCode = -4201
	ObHashGetTimeout                                                              ObErrorCode = -4204
	ObHashPlacementRetry                                                          ObErrorCode = -4205
	ObHashFull                                                                    ObErrorCode = -4206
	ObWaitNextTimeout                                                             ObErrorCode = -4208
	ObMajorFreezeNotFinished                                                      ObErrorCode = -4213
	ObInvalidDateValue                                                            ObErrorCode = -4219
	ObInactiveSqlClient                                                           ObErrorCode = -4220
	ObInactiveRpcProxy                                                            ObErrorCode = -4221
	ObIntervalWithMonth                                                           ObErrorCode = -4222
	ObTooManyDatetimeParts                                                        ObErrorCode = -4223
	ObDataOutOfRange                                                              ObErrorCode = -4224
	ObErrTruncatedWrongValueForField                                              ObErrorCode = -4226
	ObErrOutOfLowerBound                                                          ObErrorCode = -4233
	ObErrOutOfUpperBound                                                          ObErrorCode = -4234
	ObBadNullError                                                                ObErrorCode = -4235
	ObFileNotOpened                                                               ObErrorCode = -4243
	ObErrDataTruncated                                                            ObErrorCode = -4249
	ObNotRunning                                                                  ObErrorCode = -4250
	ObErrCompressDecompressData                                                   ObErrorCode = -4257
	ObErrIncorrectStringValue                                                     ObErrorCode = -4258
	ObDatetimeFunctionOverflow                                                    ObErrorCode = -4261
	ObErrDoubleTruncated                                                          ObErrorCode = -4262
	ObCacheFreeBlockNotEnough                                                     ObErrorCode = -4273
	ObLastLogNotComplete                                                          ObErrorCode = -4278
	ObUnexpectInternalError                                                       ObErrorCode = -4388
	ObErrTooMuchTime                                                              ObErrorCode = -4389
	ObErrThreadPanic                                                              ObErrorCode = -4396
	ObErrIntervalPartitionExist                                                   ObErrorCode = -4728
	ObErrIntervalPartitionError                                                   ObErrorCode = -4729
	ObFrozenInfoAlreadyExist                                                      ObErrorCode = -4744
	ObCreateStandbyTenantFailed                                                   ObErrorCode = -4765
	ObLsWaitingSafeDestroy                                                        ObErrorCode = -4766
	ObLsLockConflict                                                              ObErrorCode = -4768
	ObInvalidRootKey                                                              ObErrorCode = -4769
	ObErrParserSyntax                                                             ObErrorCode = -5006
	ObErrColumnNotFound                                                           ObErrorCode = -5031
	ObErrSysVariableUnknown                                                       ObErrorCode = -5044
	ObErrReadOnly                                                                 ObErrorCode = -5081
	ObIntegerPrecisionOverflow                                                    ObErrorCode = -5088
	ObDecimalPrecisionOverflow                                                    ObErrorCode = -5089
	ObNumericOverflow                                                             ObErrorCode = -5093
	ObErrSysConfigUnknown                                                         ObErrorCode = -5099
	ObInvalidArgumentForExtract                                                   ObErrorCode = -5106
	ObInvalidArgumentForIs                                                        ObErrorCode = -5107
	ObInvalidArgumentForLength                                                    ObErrorCode = -5108
	ObInvalidArgumentForSubstr                                                    ObErrorCode = -5109
	ObInvalidArgumentForTimeToUsec                                                ObErrorCode = -5110
	ObInvalidArgumentForUsecToTime                                                ObErrorCode = -5111
	ObInvalidNumeric                                                              ObErrorCode = -5114
	ObErrRegexpError                                                              ObErrorCode = -5115
	ObErrUnknownCharset                                                           ObErrorCode = -5142
	ObErrUnknownCollation                                                         ObErrorCode = -5143
	ObErrCollationMismatch                                                        ObErrorCode = -5144
	ObErrWrongValueForVar                                                         ObErrorCode = -5145
	ObTenantNotInServer                                                           ObErrorCode = -5150
	ObTenantNotExist                                                              ObErrorCode = -5157
	ObErrDataTooLong                                                              ObErrorCode = -5167
	ObErrWrongValueCountOnRow                                                     ObErrorCode = -5168
	ObCantAggregate_2collations                                                   ObErrorCode = -5177
	ObErrUnknownTimeZone                                                          ObErrorCode = -5192
	ObErrTooBigPrecision                                                          ObErrorCode = -5203
	ObErrMBiggerThanD                                                             ObErrorCode = -5204
	ObErrTruncatedWrongValue                                                      ObErrorCode = -5222
	ObErrWrongValue                                                               ObErrorCode = -5241
	ObErrUnexpectedTzTransition                                                   ObErrorCode = -5297
	ObErrInvalidTimezoneRegionId                                                  ObErrorCode = -5341
	ObErrInvalidHexNumber                                                         ObErrorCode = -5342
	ObErrFieldNotFoundInDatetimeOrInterval                                        ObErrorCode = -5352
	ObErrInvalidJsonText                                                          ObErrorCode = -5411
	ObErrInvalidJsonTextInParam                                                   ObErrorCode = -5412
	ObErrInvalidJsonBinaryData                                                    ObErrorCode = -5413
	ObErrInvalidJsonPath                                                          ObErrorCode = -5414
	ObErrInvalidJsonCharset                                                       ObErrorCode = -5415
	ObErrInvalidJsonCharsetInFunction                                             ObErrorCode = -5416
	ObErrInvalidTypeForJson                                                       ObErrorCode = -5417
	ObErrInvalidCastToJson                                                        ObErrorCode = -5418
	ObErrInvalidJsonPathCharset                                                   ObErrorCode = -5419
	ObErrInvalidJsonPathWildcard                                                  ObErrorCode = -5420
	ObErrJsonValueTooBig                                                          ObErrorCode = -5421
	ObErrJsonKeyTooBig                                                            ObErrorCode = -5422
	ObErrJsonUsedAsKey                                                            ObErrorCode = -5423
	ObErrJsonVacuousPath                                                          ObErrorCode = -5424
	ObErrJsonBadOneOrAllArg                                                       ObErrorCode = -5425
	ObErrNumericJsonValueOutOfRange                                               ObErrorCode = -5426
	ObErrInvalidJsonValueForCast                                                  ObErrorCode = -5427
	ObErrJsonOutOfDepth                                                           ObErrorCode = -5428
	ObErrJsonDocumentNullKey                                                      ObErrorCode = -5429
	ObErrBlobCantHaveDefault                                                      ObErrorCode = -5430
	ObErrInvalidJsonPathArrayCell                                                 ObErrorCode = -5431
	ObErrMissingJsonValue                                                         ObErrorCode = -5432
	ObErrMultipleJsonValues                                                       ObErrorCode = -5433
	ObInvalidArgumentForTimestampToScn                                            ObErrorCode = -5436
	ObInvalidArgumentForScnToTimestamp                                            ObErrorCode = -5437
	ObErrInvalidJsonType                                                          ObErrorCode = -5441
	ObErrJsonPathSyntaxError                                                      ObErrorCode = -5442
	ObErrJsonValueNoScalar                                                        ObErrorCode = -5444
	ObErrDuplicateKey                                                             ObErrorCode = -5453
	ObErrJsonPathExpressionSyntaxError                                            ObErrorCode = -5454
	ObErrNotIso8601Format                                                         ObErrorCode = -5472
	ObErrValueExceededMax                                                         ObErrorCode = -5473
	ObErrBoolNotConvertNumber                                                     ObErrorCode = -5475
	ObErrJsonKeyNotFound                                                          ObErrorCode = -5485
	ObErrYearConflictsWithJulianDate                                              ObErrorCode = -5629
	ObErrDayOfYearConflictsWithJulianDate                                         ObErrorCode = -5630
	ObErrMonthConflictsWithJulianDate                                             ObErrorCode = -5631
	ObErrDayOfMonthConflictsWithJulianDate                                        ObErrorCode = -5632
	ObErrDayOfWeekConflictsWithJulianDate                                         ObErrorCode = -5633
	ObErrHourConflictsWithSecondsInDay                                            ObErrorCode = -5634
	ObErrMinutesOfHourConflictsWithSecondsInDay                                   ObErrorCode = -5635
	ObErrSecondsOfMinuteConflictsWithSecondsInDay                                 ObErrorCode = -5636
	ObErrDateNotValidForMonthSpecified                                            ObErrorCode = -5637
	ObErrInputValueNotLongEnough                                                  ObErrorCode = -5638
	ObErrInvalidYearValue                                                         ObErrorCode = -5639
	ObErrInvalidQuarterValue                                                      ObErrorCode = -5640
	ObErrInvalidMonth                                                             ObErrorCode = -5641
	ObErrInvalidDayOfTheWeek                                                      ObErrorCode = -5642
	ObErrInvalidDayOfYearValue                                                    ObErrorCode = -5643
	ObErrInvalidHour12Value                                                       ObErrorCode = -5644
	ObErrInvalidHour24Value                                                       ObErrorCode = -5645
	ObErrInvalidMinutesValue                                                      ObErrorCode = -5646
	ObErrInvalidSecondsValue                                                      ObErrorCode = -5647
	ObErrInvalidSecondsInDayValue                                                 ObErrorCode = -5648
	ObErrInvalidJulianDateValue                                                   ObErrorCode = -5649
	ObErrAmOrPmRequired                                                           ObErrorCode = -5650
	ObErrBcOrAdRequired                                                           ObErrorCode = -5651
	ObErrFormatCodeAppearsTwice                                                   ObErrorCode = -5652
	ObErrDayOfWeekSpecifiedMoreThanOnce                                           ObErrorCode = -5653
	ObErrSignedYearPrecludesUseOfBcAd                                             ObErrorCode = -5654
	ObErrJulianDatePrecludesUseOfDayOfYear                                        ObErrorCode = -5655
	ObErrYearMayOnlyBeSpecifiedOnce                                               ObErrorCode = -5656
	ObErrHourMayOnlyBeSpecifiedOnce                                               ObErrorCode = -5657
	ObErrAmPmConflictsWithUseOfAmDotPmDot                                         ObErrorCode = -5658
	ObErrBcAdConflictWithUseOfBcDotAdDot                                          ObErrorCode = -5659
	ObErrMonthMayOnlyBeSpecifiedOnce                                              ObErrorCode = -5660
	ObErrDayOfWeekMayOnlyBeSpecifiedOnce                                          ObErrorCode = -5661
	ObErrFormatCodeCannotAppear                                                   ObErrorCode = -5662
	ObErrNonNumericCharacterValue                                                 ObErrorCode = -5663
	ObInvalidMeridianIndicatorUse                                                 ObErrorCode = -5664
	ObErrDayOfMonthRange                                                          ObErrorCode = -5667
	ObErrArgumentOutOfRange                                                       ObErrorCode = -5674
	ObErrIntervalInvalid                                                          ObErrorCode = -5676
	ObErrTheLeadingPrecisionOfTheIntervalIsTooSmall                               ObErrorCode = -5708
	ObErrInvalidTimeZoneHour                                                      ObErrorCode = -5709
	ObErrInvalidTimeZoneMinute                                                    ObErrorCode = -5710
	ObErrNotAValidTimeZone                                                        ObErrorCode = -5711
	ObErrDateFormatIsTooLongForInternalBuffer                                     ObErrorCode = -5712
	ObErrOperatorCannotBeUsedWithList                                             ObErrorCode = -5729
	ObInvalidRowId                                                                ObErrorCode = -5802
	ObErrNumericNotMatchFormatLength                                              ObErrorCode = -5873
	ObErrDatetimeIntervalInternalError                                            ObErrorCode = -5898
	ObErrDblinkRemoteEcode                                                        ObErrorCode = -5975
	ObErrDblinkNoLib                                                              ObErrorCode = -5976
	ObSwitchingToFollowerGracefully                                               ObErrorCode = -6202
	ObMaskSetNoNode                                                               ObErrorCode = -6203
	ObTransTimeout                                                                ObErrorCode = -6210
	ObTransKilled                                                                 ObErrorCode = -6211
	ObTransStmtTimeout                                                            ObErrorCode = -6212
	ObTransCtxNotExist                                                            ObErrorCode = -6213
	ObTransUnknown                                                                ObErrorCode = -6225
	ObErrReadOnlyTransaction                                                      ObErrorCode = -6226
	ObErrGisDifferentSrids                                                        ObErrorCode = -7201
	ObErrGisUnsupportedArgument                                                   ObErrorCode = -7202
	ObErrGisUnknownError                                                          ObErrorCode = -7203
	ObErrGisUnknownException                                                      ObErrorCode = -7204
	ObErrGisInvalidData                                                           ObErrorCode = -7205
	ObErrBoostGeometryEmptyInputException                                         ObErrorCode = -7206
	ObErrBoostGeometryCentroidException                                           ObErrorCode = -7207
	ObErrBoostGeometryOverlayInvalidInputException                                ObErrorCode = -7208
	ObErrBoostGeometryTurnInfoException                                           ObErrorCode = -7209
	ObErrBoostGeometrySelfIntersectionPointException                              ObErrorCode = -7210
	ObErrBoostGeometryUnknownException                                            ObErrorCode = -7211
	ObErrGisDataWrongEndianess                                                    ObErrorCode = -7212
	ObErrAlterOperationNotSupportedReasonGis                                      ObErrorCode = -7213
	ObErrBoostGeometryInconsistentTurnsException                                  ObErrorCode = -7214
	ObErrGisMaxPointsInGeometryOverflowed                                         ObErrorCode = -7215
	ObErrUnexpectedGeometryType                                                   ObErrorCode = -7216
	ObErrSrsParseError                                                            ObErrorCode = -7217
	ObErrSrsProjParameterMissing                                                  ObErrorCode = -7218
	ObErrWarnSrsNotFound                                                          ObErrorCode = -7219
	ObErrSrsNotCartesian                                                          ObErrorCode = -7220
	ObErrSrsNotCartesianUndefined                                                 ObErrorCode = -7221
	ObErrSrsNotFound                                                              ObErrorCode = -7222
	ObErrWarnSrsNotFoundAxisOrder                                                 ObErrorCode = -7223
	ObErrNotImplementedForGeographicSrs                                           ObErrorCode = -7224
	ObErrWrongSridForColumn                                                       ObErrorCode = -7225
	ObErrCannotAlterSridDueToIndex                                                ObErrorCode = -7226
	ObErrWarnUselessSpatialIndex                                                  ObErrorCode = -7227
	ObErrOnlyImplementedForSrid_0And_4326                                         ObErrorCode = -7228
	ObErrNotImplementedForCartesianSrs                                            ObErrorCode = -7229
	ObErrNotImplementedForProjectedSrs                                            ObErrorCode = -7230
	ObErrSrsMissingMandatoryAttribute                                             ObErrorCode = -7231
	ObErrSrsMultipleAttributeDefinitions                                          ObErrorCode = -7232
	ObErrSrsNameCantBeEmptyOrWhitespace                                           ObErrorCode = -7233
	ObErrSrsOrganizationCantBeEmptyOrWhitespace                                   ObErrorCode = -7234
	ObErrSrsIdAlreadyExists                                                       ObErrorCode = -7235
	ObErrWarnSrsIdAlreadyExists                                                   ObErrorCode = -7236
	ObErrCantModifySrid_0                                                         ObErrorCode = -7237
	ObErrWarnReservedSridRange                                                    ObErrorCode = -7238
	ObErrCantModifySrsUsedByColumn                                                ObErrorCode = -7239
	ObErrSrsInvalidCharacterInAttribute                                           ObErrorCode = -7240
	ObErrSrsAttributeStringTooLong                                                ObErrorCode = -7241
	ObErrSrsNotGeographic                                                         ObErrorCode = -7242
	ObErrPolygonTooLarge                                                          ObErrorCode = -7243
	ObErrSpatialUniqueIndex                                                       ObErrorCode = -7244
	ObErrGeometryParamLongitudeOutOfRange                                         ObErrorCode = -7246
	ObErrGeometryParamLatitudeOutOfRange                                          ObErrorCode = -7247
	ObErrSrsGeogcsInvalidAxes                                                     ObErrorCode = -7248
	ObErrSrsInvalidSemiMajorAxis                                                  ObErrorCode = -7249
	ObErrSrsInvalidInverseFlattening                                              ObErrorCode = -7250
	ObErrSrsInvalidAngularUnit                                                    ObErrorCode = -7251
	ObErrSrsInvalidPrimeMeridian                                                  ObErrorCode = -7252
	ObErrTransformSourceSrsNotSupported                                           ObErrorCode = -7253
	ObErrTransformTargetSrsNotSupported                                           ObErrorCode = -7254
	ObErrTransformSourceSrsMissingTowgs84                                         ObErrorCode = -7255
	ObErrTransformTargetSrsMissingTowgs84                                         ObErrorCode = -7256
	ObErrFunctionalIndexOnJsonOrGeometryFunction                                  ObErrorCode = -7257
	ObErrSpatialFunctionalIndex                                                   ObErrorCode = -7258
	ObErrGeometryInUnknownLengthUnit                                              ObErrorCode = -7259
	ObErrInvalidCastToGeometry                                                    ObErrorCode = -7260
	ObErrInvalidCastPolygonRingDirection                                          ObErrorCode = -7261
	ObErrGisDifferentSridsAggregation                                             ObErrorCode = -7262
	ObErrLongitudeOutOfRange                                                      ObErrorCode = -7263
	ObErrLatitudeOutOfRange                                                       ObErrorCode = -7264
	ObErrStdBadAllocError                                                         ObErrorCode = -7265
	ObErrStdDomainError                                                           ObErrorCode = -7266
	ObErrStdLengthError                                                           ObErrorCode = -7267
	ObErrStdInvalidArgument                                                       ObErrorCode = -7268
	ObErrStdOutOfRangeError                                                       ObErrorCode = -7269
	ObErrStdOverflowError                                                         ObErrorCode = -7270
	ObErrStdRangeError                                                            ObErrorCode = -7271
	ObErrStdUnderflowError                                                        ObErrorCode = -7272
	ObErrStdLogicError                                                            ObErrorCode = -7273
	ObErrStdRuntimeError                                                          ObErrorCode = -7274
	ObErrStdUnknownException                                                      ObErrorCode = -7275
	ObErrCantCreateGeometryObject                                                 ObErrorCode = -7276
	ObErrSridWrongUsage                                                           ObErrorCode = -7277
	ObErrIndexOrderWrongUsage                                                     ObErrorCode = -7278
	ObErrSpatialMustHaveGeomCol                                                   ObErrorCode = -7279
	ObErrSpatialCantHaveNull                                                      ObErrorCode = -7280
	ObErrIndexTypeNotSupportedForSpatialIndex                                     ObErrorCode = -7281
	ObErrUnitNotFound                                                             ObErrorCode = -7282
	ObErrInvalidOptionKeyValuePair                                                ObErrorCode = -7283
	ObErrNonpositiveRadius                                                        ObErrorCode = -7284
	ObErrSrsEmpty                                                                 ObErrorCode = -7285
	ObErrInvalidOptionKey                                                         ObErrorCode = -7286
	ObErrInvalidOptionValue                                                       ObErrorCode = -7287
	ObErrInvalidGeometryType                                                      ObErrorCode = -7288
	ObPacketClusterIdNotMatch                                                     ObErrorCode = -8004
	ObTenantIdNotMatch                                                            ObErrorCode = -8005
	ObUriError                                                                    ObErrorCode = -9001
	ObFinalMd5Error                                                               ObErrorCode = -9002
	ObOssError                                                                    ObErrorCode = -9003
	ObInitMd5Error                                                                ObErrorCode = -9004
	ObOutOfElement                                                                ObErrorCode = -9005
	ObUpdateMd5Error                                                              ObErrorCode = -9006
	ObFileLengthInvalid                                                           ObErrorCode = -9007
	ObBackupFileNotExist                                                          ObErrorCode = -9011
	ObInvalidBackupDest                                                           ObErrorCode = -9026
	ObCosError                                                                    ObErrorCode = -9060
	ObIoLimit                                                                     ObErrorCode = -9061
	ObBackupBackupReachCopyLimit                                                  ObErrorCode = -9062
	ObBackupIoProhibited                                                          ObErrorCode = -9063
	ObBackupPermissionDenied                                                      ObErrorCode = -9071
	ObEsiObsError                                                                 ObErrorCode = -9073
	ObBackupMetaIndexNotExist                                                     ObErrorCode = -9076
	ObBackupDeviceOutOfSpace                                                      ObErrorCode = -9082
	ObBackupPwriteOffsetNotMatch                                                  ObErrorCode = -9083
	ObBackupPwriteContentNotMatch                                                 ObErrorCode = -9084
	ObCloudObjectNotAppendable                                                    ObErrorCode = -9098
	ObRestoreTenantFailed                                                         ObErrorCode = -9099
	ObErrXmlParse                                                                 ObErrorCode = -9549
	ObErrXsltParse                                                                ObErrorCode = -9574
	ObPacketNotSent                                                               ObErrorCode = -4011
	ObPartialFailed                                                               ObErrorCode = -4025
	ObSchemaError                                                                 ObErrorCode = -4029
	ObTenantOutOfMem                                                              ObErrorCode = -4030
	ObUnknownObj                                                                  ObErrorCode = -4031
	ObNoMonitorData                                                               ObErrorCode = -4032
	ObTooManySstable                                                              ObErrorCode = -4037
	ObKilledByThrottling                                                          ObErrorCode = -4039
	ObUserNotExist                                                                ObErrorCode = -4042
	ObPasswordWrong                                                               ObErrorCode = -4043
	ObSkeyVersionWrong                                                            ObErrorCode = -4044
	ObPushdownStatusChanged                                                       ObErrorCode = -4045
	ObStorageSchemaInvalid                                                        ObErrorCode = -4046
	ObMediumCompactionInfoInvalid                                                 ObErrorCode = -4047
	ObNotRegistered                                                               ObErrorCode = -4048
	ObWaitqueueTimeout                                                            ObErrorCode = -4049
	ObAlreadyRegistered                                                           ObErrorCode = -4051
	ObNoCsSelected                                                                ObErrorCode = -4053
	ObNoTabletsCreated                                                            ObErrorCode = -4054
	ObDecimalUnlegalError                                                         ObErrorCode = -4058
	ObObjDivideError                                                              ObErrorCode = -4060
	ObNotADecimal                                                                 ObErrorCode = -4061
	ObDecimalPrecisionNotEqual                                                    ObErrorCode = -4062
	ObSessionKilled                                                               ObErrorCode = -4064
	ObLogNotSync                                                                  ObErrorCode = -4065
	ObSessionNotFound                                                             ObErrorCode = -4067
	ObInvalidLog                                                                  ObErrorCode = -4068
	ObAlreadyDone                                                                 ObErrorCode = -4071
	ObLogSrcChanged                                                               ObErrorCode = -4073
	ObLogMissing                                                                  ObErrorCode = -4075
	ObNeedWait                                                                    ObErrorCode = -4076
	ObResultUnknown                                                               ObErrorCode = -4081
	ObNoResult                                                                    ObErrorCode = -4084
	ObLogIdRangeNotContinuous                                                     ObErrorCode = -4090
	ObTermLagged                                                                  ObErrorCode = -4097
	ObTermNotMatch                                                                ObErrorCode = -4098
	ObPartialLog                                                                  ObErrorCode = -4102
	ObNotEnoughStore                                                              ObErrorCode = -4106
	ObBlockSwitched                                                               ObErrorCode = -4107
	ObReadZeroLog                                                                 ObErrorCode = -4110
	ObBlockNeedFreeze                                                             ObErrorCode = -4111
	ObBlockFrozen                                                                 ObErrorCode = -4112
	ObInFatalState                                                                ObErrorCode = -4113
	ObUpsMasterExists                                                             ObErrorCode = -4115
	ObNotFree                                                                     ObErrorCode = -4125
	ObInitSqlContextError                                                         ObErrorCode = -4126
	ObSkipInvalidRow                                                              ObErrorCode = -4127
	ObNoTablet                                                                    ObErrorCode = -4133
	ObSnapshotDiscarded                                                           ObErrorCode = -4138
	ObDataNotUptodate                                                             ObErrorCode = -4139
	ObRowModified                                                                 ObErrorCode = -4142
	ObVersionNotMatch                                                             ObErrorCode = -4143
	ObEnqueueFailed                                                               ObErrorCode = -4146
	ObInvalidConfig                                                               ObErrorCode = -4147
	ObStmtExpired                                                                 ObErrorCode = -4149
	ObMetaTableWithoutUseTable                                                    ObErrorCode = -4155
	ObDiscardPacket                                                               ObErrorCode = -4156
	ObPoolRegisteredFailed                                                        ObErrorCode = -4159
	ObPoolUnregisteredFailed                                                      ObErrorCode = -4160
	ObLeaseNotEnough                                                              ObErrorCode = -4162
	ObLeaseNotMatch                                                               ObErrorCode = -4163
	ObUpsSwitchNotHappen                                                          ObErrorCode = -4164
	ObCacheNotHit                                                                 ObErrorCode = -4166
	ObNestedLoopNotSupport                                                        ObErrorCode = -4167
	ObIndexOutOfRange                                                             ObErrorCode = -4172
	ObIntUnderflow                                                                ObErrorCode = -4173
	ObCacheShrinkFailed                                                           ObErrorCode = -4176
	ObOldSchemaVersion                                                            ObErrorCode = -4177
	ObReleaseSchemaError                                                          ObErrorCode = -4178
	ObNoEmptyEntry                                                                ObErrorCode = -4180
	ObBeyondTheRange                                                              ObErrorCode = -4183
	ObServerOutofDiskSpace                                                        ObErrorCode = -4184
	ObColumnGroupNotFound                                                         ObErrorCode = -4185
	ObCsCompressLibError                                                          ObErrorCode = -4186
	ObSchedulerTaskCntMismatch                                                    ObErrorCode = -4188
	ObInvalidMacroBlockType                                                       ObErrorCode = -4189
	ObPgIsRemoved                                                                 ObErrorCode = -4191
	ObPacketProcessed                                                             ObErrorCode = -4207
	ObLeaderNotExist                                                              ObErrorCode = -4209
	ObPrepareMajorFreezeFailed                                                    ObErrorCode = -4210
	ObCommitMajorFreezeFailed                                                     ObErrorCode = -4211
	ObAbortMajorFreezeFailed                                                      ObErrorCode = -4212
	ObPartitionNotLeader                                                          ObErrorCode = -4214
	ObWaitMajorFreezeResponseTimeout                                              ObErrorCode = -4215
	ObCurlError                                                                   ObErrorCode = -4216
	ObMajorFreezeNotAllow                                                         ObErrorCode = -4217
	ObPrepareFreezeFailed                                                         ObErrorCode = -4218
	ObPartitionNotExist                                                           ObErrorCode = -4225
	ObErrNoDefaultForField                                                        ObErrorCode = -4227
	ObErrFieldSpecifiedTwice                                                      ObErrorCode = -4228
	ObErrTooLongTableComment                                                      ObErrorCode = -4229
	ObErrTooLongFieldComment                                                      ObErrorCode = -4230
	ObErrTooLongIndexComment                                                      ObErrorCode = -4231
	ObNotFollower                                                                 ObErrorCode = -4232
	ObObconfigReturnError                                                         ObErrorCode = -4236
	ObObconfigAppnameMismatch                                                     ObErrorCode = -4237
	ObErrViewSelectDerived                                                        ObErrorCode = -4238
	ObCantMjPath                                                                  ObErrorCode = -4239
	ObErrNoJoinOrderGenerated                                                     ObErrorCode = -4240
	ObErrNoPathGenerated                                                          ObErrorCode = -4241
	ObErrWaitRemoteSchemaRefresh                                                  ObErrorCode = -4242
	ObTimerTaskHasScheduled                                                       ObErrorCode = -4244
	ObTimerTaskHasNotScheduled                                                    ObErrorCode = -4245
	ObParseDebugSyncError                                                         ObErrorCode = -4246
	ObUnknownDebugSyncPoint                                                       ObErrorCode = -4247
	ObErrInterrupted                                                              ObErrorCode = -4248
	ObInvalidPartition                                                            ObErrorCode = -4251
	ObErrTimeoutTruncated                                                         ObErrorCode = -4252
	ObErrTooLongTenantComment                                                     ObErrorCode = -4253
	ObErrNetPacketTooLarge                                                        ObErrorCode = -4254
	ObTraceDescNotExist                                                           ObErrorCode = -4255
	ObErrNoDefault                                                                ObErrorCode = -4256
	ObIsChangingLeader                                                            ObErrorCode = -4260
	ObMinorFreezeNotAllow                                                         ObErrorCode = -4263
	ObLogOutofDiskSpace                                                           ObErrorCode = -4264
	ObRpcConnectError                                                             ObErrorCode = -4265
	ObMinorMergeNotAllow                                                          ObErrorCode = -4266
	ObCacheInvalid                                                                ObErrorCode = -4267
	ObReachServerDataCopyInConcurrencyLimit                                       ObErrorCode = -4268
	ObWorkingPartitionExist                                                       ObErrorCode = -4269
	ObWorkingPartitionNotExist                                                    ObErrorCode = -4270
	ObLibeasyReachMemLimit                                                        ObErrorCode = -4271
	ObSyncWashMbTimeout                                                           ObErrorCode = -4274
	ObNotAllowMigrateIn                                                           ObErrorCode = -4275
	ObSchedulerTaskCntMistach                                                     ObErrorCode = -4276
	ObMissArgument                                                                ObErrorCode = -4277
	ObTableIsDeleted                                                              ObErrorCode = -4279
	ObVersionRangeNotContinues                                                    ObErrorCode = -4280
	ObInvalidIoBuffer                                                             ObErrorCode = -4281
	ObPartitionIsRemoved                                                          ObErrorCode = -4282
	ObGtsNotReady                                                                 ObErrorCode = -4283
	ObMajorSstableNotExist                                                        ObErrorCode = -4284
	ObVersionRangeDiscarded                                                       ObErrorCode = -4285
	ObMajorSstableHasMerged                                                       ObErrorCode = -4286
	ObMinorSstableRangeCross                                                      ObErrorCode = -4287
	ObMemtableCannotMinorMerge                                                    ObErrorCode = -4288
	ObTaskExist                                                                   ObErrorCode = -4289
	ObAllocateDiskSpaceFailed                                                     ObErrorCode = -4290
	ObCantFindUdf                                                                 ObErrorCode = -4291
	ObCantInitializeUdf                                                           ObErrorCode = -4292
	ObUdfNoPaths                                                                  ObErrorCode = -4293
	ObUdfExists                                                                   ObErrorCode = -4294
	ObCantOpenLibrary                                                             ObErrorCode = -4295
	ObCantFindDlEntry                                                             ObErrorCode = -4296
	ObObjectNameExist                                                             ObErrorCode = -4297
	ObObjectNameNotExist                                                          ObErrorCode = -4298
	ObErrDupArgument                                                              ObErrorCode = -4299
	ObErrInvalidSequenceName                                                      ObErrorCode = -4300
	ObErrDupMaxvalueSpec                                                          ObErrorCode = -4301
	ObErrDupMinvalueSpec                                                          ObErrorCode = -4302
	ObErrDupCycleSpec                                                             ObErrorCode = -4303
	ObErrDupCacheSpec                                                             ObErrorCode = -4304
	ObErrDupOrderSpec                                                             ObErrorCode = -4305
	ObErrConflMaxvalueSpec                                                        ObErrorCode = -4306
	ObErrConflMinvalueSpec                                                        ObErrorCode = -4307
	ObErrConflCycleSpec                                                           ObErrorCode = -4308
	ObErrConflCacheSpec                                                           ObErrorCode = -4309
	ObErrConflOrderSpec                                                           ObErrorCode = -4310
	ObErrAlterStartSeqNumberNotAllowed                                            ObErrorCode = -4311
	ObErrDupIncrementBySpec                                                       ObErrorCode = -4312
	ObErrDupStartWithSpec                                                         ObErrorCode = -4313
	ObErrRequireAlterSeqOption                                                    ObErrorCode = -4314
	ObErrSeqNotAllowedHere                                                        ObErrorCode = -4315
	ObErrSeqNotExist                                                              ObErrorCode = -4316
	ObErrSeqOptionMustBeInteger                                                   ObErrorCode = -4317
	ObErrSeqIncrementCanNotBeZero                                                 ObErrorCode = -4318
	ObErrSeqOptionExceedRange                                                     ObErrorCode = -4319
	ObErrMinvalueLargerThanMaxvalue                                               ObErrorCode = -4320
	ObErrSeqIncrementTooLarge                                                     ObErrorCode = -4321
	ObErrStartWithLessThanMinvalue                                                ObErrorCode = -4322
	ObErrMinvalueExceedCurrval                                                    ObErrorCode = -4323
	ObErrStartWithExceedMaxvalue                                                  ObErrorCode = -4324
	ObErrMaxvalueExceedCurrval                                                    ObErrorCode = -4325
	ObErrSeqCacheTooSmall                                                         ObErrorCode = -4326
	ObErrSeqOptionOutOfRange                                                      ObErrorCode = -4327
	ObErrSeqCacheTooLarge                                                         ObErrorCode = -4328
	ObErrSeqRequireMinvalue                                                       ObErrorCode = -4329
	ObErrSeqRequireMaxvalue                                                       ObErrorCode = -4330
	ObErrSeqNoLongerExist                                                         ObErrorCode = -4331
	ObErrSeqValueExceedLimit                                                      ObErrorCode = -4332
	ObErrDivisorIsZero                                                            ObErrorCode = -4333
	ObErrAesDecrypt                                                               ObErrorCode = -4334
	ObErrAesEncrypt                                                               ObErrorCode = -4335
	ObErrAesIvLength                                                              ObErrorCode = -4336
	ObStoreDirError                                                               ObErrorCode = -4337
	ObOpenTwice                                                                   ObErrorCode = -4338
	ObRaidSuperBlockNotMacth                                                      ObErrorCode = -4339
	ObNotOpen                                                                     ObErrorCode = -4340
	ObNotInService                                                                ObErrorCode = -4341
	ObRaidDiskNotNormal                                                           ObErrorCode = -4342
	ObTenantSchemaNotFull                                                         ObErrorCode = -4343
	ObInvalidQueryTimestamp                                                       ObErrorCode = -4344
	ObDirNotEmpty                                                                 ObErrorCode = -4345
	ObSchemaNotUptodate                                                           ObErrorCode = -4346
	ObRoleNotExist                                                                ObErrorCode = -4347
	ObRoleExist                                                                   ObErrorCode = -4348
	ObPrivDup                                                                     ObErrorCode = -4349
	ObKeystoreExist                                                               ObErrorCode = -4350
	ObKeystoreNotExist                                                            ObErrorCode = -4351
	ObKeystoreWrongPassword                                                       ObErrorCode = -4352
	ObTablespaceExist                                                             ObErrorCode = -4353
	ObTablespaceNotExist                                                          ObErrorCode = -4354
	ObTablespaceDeleteNotEmpty                                                    ObErrorCode = -4355
	ObFloatPrecisionOutRange                                                      ObErrorCode = -4356
	ObNumericPrecisionOutRange                                                    ObErrorCode = -4357
	ObNumericScaleOutRange                                                        ObErrorCode = -4358
	ObKeystoreNotOpen                                                             ObErrorCode = -4359
	ObKeystoreOpenNoMasterKey                                                     ObErrorCode = -4360
	ObSlogReachMaxConcurrency                                                     ObErrorCode = -4361
	ObErrByAccessOrSessionClauseNotAllowedForNoaudit                              ObErrorCode = -4362
	ObErrAuditingTheObjectIsNotSupported                                          ObErrorCode = -4363
	ObErrDdlStatementCannotBeAuditedWithBySessionSpecified                        ObErrorCode = -4364
	ObErrNotValidPassword                                                         ObErrorCode = -4365
	ObErrMustChangePassword                                                       ObErrorCode = -4366
	ObOversizeNeedRetry                                                           ObErrorCode = -4367
	ObObconfigClusterNotExist                                                     ObErrorCode = -4368
	ObErrGetMasterKey                                                             ObErrorCode = -4369
	ObErrTdeMethod                                                                ObErrorCode = -4370
	ObKmsServerConnectError                                                       ObErrorCode = -4371
	ObKmsServerIsBusy                                                             ObErrorCode = -4372
	ObKmsServerUpdateKeyConflict                                                  ObErrorCode = -4373
	ObErrValueLargerThanAllowed                                                   ObErrorCode = -4374
	ObDiskError                                                                   ObErrorCode = -4375
	ObUnimplementedFeature                                                        ObErrorCode = -4376
	ObErrDefensiveCheck                                                           ObErrorCode = -4377
	ObClusterNameHashConflict                                                     ObErrorCode = -4378
	ObHeapTableExausted                                                           ObErrorCode = -4379
	ObErrIndexKeyNotFound                                                         ObErrorCode = -4380
	ObUnsupportedDeprecatedFeature                                                ObErrorCode = -4381
	ObErrDupRestartSpec                                                           ObErrorCode = -4382
	ObGtiNotReady                                                                 ObErrorCode = -4383
	ObStackOverflow                                                               ObErrorCode = -4385
	ObNotAllowRemovingLeader                                                      ObErrorCode = -4386
	ObNeedSwitchConsumerGroup                                                     ObErrorCode = -4387
	ObErrRemoteSchemaNotFull                                                      ObErrorCode = -4390
	ObDdlSstableRangeCross                                                        ObErrorCode = -4391
	ObDiskHung                                                                    ObErrorCode = -4392
	ObErrObserverStart                                                            ObErrorCode = -4393
	ObErrObserverStop                                                             ObErrorCode = -4394
	ObErrObserviceStart                                                           ObErrorCode = -4395
	ObEncodingEstSizeOverflow                                                     ObErrorCode = -4397
	ObInvalidSubPartitionType                                                     ObErrorCode = -4398
	ObErrUnexpectedUnitStatus                                                     ObErrorCode = -4399
	ObAutoincCacheNotEqual                                                        ObErrorCode = -4400
	ObImportNotInServer                                                           ObErrorCode = -4505
	ObConvertError                                                                ObErrorCode = -4507
	ObBypassTimeout                                                               ObErrorCode = -4510
	ObRsStateNotAllow                                                             ObErrorCode = -4512
	ObNoReplicaValid                                                              ObErrorCode = -4515
	ObNoNeedUpdate                                                                ObErrorCode = -4517
	ObCacheTimeout                                                                ObErrorCode = -4518
	ObIterStop                                                                    ObErrorCode = -4519
	ObZoneAlreadyMaster                                                           ObErrorCode = -4523
	ObIpPortIsNotSlaveZone                                                        ObErrorCode = -4524
	ObZoneIsNotSlave                                                              ObErrorCode = -4525
	ObZoneIsNotMaster                                                             ObErrorCode = -4526
	ObConfigNotSync                                                               ObErrorCode = -4527
	ObIpPortIsNotZone                                                             ObErrorCode = -4528
	ObMasterZoneNotExist                                                          ObErrorCode = -4529
	ObZoneInfoNotExist                                                            ObErrorCode = -4530
	ObGetZoneMasterUpsFailed                                                      ObErrorCode = -4531
	ObMultipleMasterZonesExist                                                    ObErrorCode = -4532
	ObIndexingZoneInvalid                                                         ObErrorCode = -4533
	ObRootTableRangeNotExist                                                      ObErrorCode = -4537
	ObRootMigrateConcurrencyFull                                                  ObErrorCode = -4538
	ObRootMigrateInfoNotFound                                                     ObErrorCode = -4539
	ObNotDataLoadTable                                                            ObErrorCode = -4540
	ObDataLoadTableDuplicated                                                     ObErrorCode = -4541
	ObRootTableIdExist                                                            ObErrorCode = -4542
	ObIndexTimeout                                                                ObErrorCode = -4543
	ObRootNotIntegrated                                                           ObErrorCode = -4544
	ObIndexIneligible                                                             ObErrorCode = -4545
	ObRebalanceExecTimeout                                                        ObErrorCode = -4546
	ObMergeNotStarted                                                             ObErrorCode = -4547
	ObMergeAlreadyStarted                                                         ObErrorCode = -4548
	ObRootserviceExist                                                            ObErrorCode = -4549
	ObRsShutdown                                                                  ObErrorCode = -4550
	ObServerMigrateInDenied                                                       ObErrorCode = -4551
	ObRebalanceTaskCantExec                                                       ObErrorCode = -4552
	ObPartitionCntReachRootserverLimit                                            ObErrorCode = -4553
	ObRebalanceTaskNotInProgress                                                  ObErrorCode = -4554
	ObDataSourceNotExist                                                          ObErrorCode = -4600
	ObDataSourceTableNotExist                                                     ObErrorCode = -4601
	ObDataSourceRangeNotExist                                                     ObErrorCode = -4602
	ObDataSourceDataNotExist                                                      ObErrorCode = -4603
	ObDataSourceSysError                                                          ObErrorCode = -4604
	ObDataSourceTimeout                                                           ObErrorCode = -4605
	ObDataSourceConcurrencyFull                                                   ObErrorCode = -4606
	ObDataSourceWrongUriFormat                                                    ObErrorCode = -4607
	ObSstableVersionUnequal                                                       ObErrorCode = -4608
	ObUpsRenewLeaseNotAllowed                                                     ObErrorCode = -4609
	ObUpsCountOverLimit                                                           ObErrorCode = -4610
	ObNoUpsMajority                                                               ObErrorCode = -4611
	ObIndexCountReachTheLimit                                                     ObErrorCode = -4613
	ObTaskExpired                                                                 ObErrorCode = -4614
	ObTablegroupNotEmpty                                                          ObErrorCode = -4615
	ObInvalidServerStatus                                                         ObErrorCode = -4620
	ObWaitElecLeaderTimeout                                                       ObErrorCode = -4621
	ObWaitAllRsOnlineTimeout                                                      ObErrorCode = -4622
	ObAllReplicasOnMergeZone                                                      ObErrorCode = -4623
	ObMachineResourceNotEnough                                                    ObErrorCode = -4624
	ObNotServerCanHoldSoftly                                                      ObErrorCode = -4625
	ObResourcePoolAlreadyGranted                                                  ObErrorCode = -4626
	ObServerAlreadyDeleted                                                        ObErrorCode = -4628
	ObServerNotDeleting                                                           ObErrorCode = -4629
	ObServerNotInWhiteList                                                        ObErrorCode = -4630
	ObServerZoneNotMatch                                                          ObErrorCode = -4631
	ObOverZoneNumLimit                                                            ObErrorCode = -4632
	ObZoneStatusNotMatch                                                          ObErrorCode = -4633
	ObResourceUnitIsReferenced                                                    ObErrorCode = -4634
	ObDifferentPrimaryZone                                                        ObErrorCode = -4636
	ObServerNotActive                                                             ObErrorCode = -4637
	ObRsNotMaster                                                                 ObErrorCode = -4638
	ObCandidateListError                                                          ObErrorCode = -4639
	ObPartitionZoneDuplicated                                                     ObErrorCode = -4640
	ObZoneDuplicated                                                              ObErrorCode = -4641
	ObNotAllZoneActive                                                            ObErrorCode = -4642
	ObPrimaryZoneNotInZoneList                                                    ObErrorCode = -4643
	ObReplicaNumNotMatch                                                          ObErrorCode = -4644
	ObZoneListPoolListNotMatch                                                    ObErrorCode = -4645
	ObInvalidTenantName                                                           ObErrorCode = -4646
	ObEmptyResourcePoolList                                                       ObErrorCode = -4647
	ObResourceUnitNotExist                                                        ObErrorCode = -4648
	ObResourceUnitExist                                                           ObErrorCode = -4649
	ObResourcePoolNotExist                                                        ObErrorCode = -4650
	ObResourcePoolExist                                                           ObErrorCode = -4651
	ObWaitLeaderSwitchTimeout                                                     ObErrorCode = -4652
	ObLocationNotExist                                                            ObErrorCode = -4653
	ObLocationLeaderNotExist                                                      ObErrorCode = -4654
	ObZoneNotActive                                                               ObErrorCode = -4655
	ObUnitNumOverServerCount                                                      ObErrorCode = -4656
	ObPoolServerIntersect                                                         ObErrorCode = -4657
	ObNotSingleResourcePool                                                       ObErrorCode = -4658
	ObResourceUnitValueBelowLimit                                                 ObErrorCode = -4659
	ObStopServerInMultipleZones                                                   ObErrorCode = -4660
	ObSessionEntryExist                                                           ObErrorCode = -4661
	ObGotSignalAborting                                                           ObErrorCode = -4662
	ObServerNotAlive                                                              ObErrorCode = -4663
	ObGetLocationTimeOut                                                          ObErrorCode = -4664
	ObUnitIsMigrating                                                             ObErrorCode = -4665
	ObClusterNoMatch                                                              ObErrorCode = -4666
	ObCheckZoneMergeOrder                                                         ObErrorCode = -4667
	ObErrZoneNotEmpty                                                             ObErrorCode = -4668
	ObDifferentLocality                                                           ObErrorCode = -4669
	ObEmptyLocality                                                               ObErrorCode = -4670
	ObFullReplicaNumNotEnough                                                     ObErrorCode = -4671
	ObReplicaNumNotEnough                                                         ObErrorCode = -4672
	ObDataSourceNotValid                                                          ObErrorCode = -4673
	ObRunJobNotSuccess                                                            ObErrorCode = -4674
	ObNoNeedRebuild                                                               ObErrorCode = -4675
	ObNeedRemoveUnneedTable                                                       ObErrorCode = -4676
	ObNoNeedMerge                                                                 ObErrorCode = -4677
	ObConflictOption                                                              ObErrorCode = -4678
	ObDuplicateOption                                                             ObErrorCode = -4679
	ObInvalidOption                                                               ObErrorCode = -4680
	ObRpcNeedReconnect                                                            ObErrorCode = -4681
	ObCannotCopyMajorSstable                                                      ObErrorCode = -4682
	ObSrcDoNotAllowedMigrate                                                      ObErrorCode = -4683
	ObTooManyTenantPartitionsError                                                ObErrorCode = -4684
	ObActiveMemtbaleNotExsit                                                      ObErrorCode = -4685
	ObUseDupFollowAfterDml                                                        ObErrorCode = -4686
	ObNoDiskNeedRebuild                                                           ObErrorCode = -4687
	ObStandbyReadOnly                                                             ObErrorCode = -4688
	ObInvaldWebServiceContent                                                     ObErrorCode = -4689
	ObPrimaryClusterExist                                                         ObErrorCode = -4690
	ObArrayBindingSwitchIterator                                                  ObErrorCode = -4691
	ObErrStandbyClusterNotEmpty                                                   ObErrorCode = -4692
	ObNotPrimaryCluster                                                           ObErrorCode = -4693
	ObErrCheckDropColumnFailed                                                    ObErrorCode = -4694
	ObNotStandbyCluster                                                           ObErrorCode = -4695
	ObClusterVersionNotCompatible                                                 ObErrorCode = -4696
	ObWaitTransTableMergeTimeout                                                  ObErrorCode = -4697
	ObSkipRenewLocationByRpc                                                      ObErrorCode = -4698
	ObRenewLocationByRpcFailed                                                    ObErrorCode = -4699
	ObClusterIdNoMatch                                                            ObErrorCode = -4700
	ObErrParamInvalid                                                             ObErrorCode = -4701
	ObErrResObjAlreadyExist                                                       ObErrorCode = -4702
	ObErrResPlanNotExist                                                          ObErrorCode = -4703
	ObErrPercentageOutOfRange                                                     ObErrorCode = -4704
	ObErrPlanDirectiveNotExist                                                    ObErrorCode = -4705
	ObErrPlanDirectiveAlreadyExist                                                ObErrorCode = -4706
	ObErrInvalidPlanDirectiveName                                                 ObErrorCode = -4707
	ObFailoverNotAllow                                                            ObErrorCode = -4708
	ObAddClusterNotAllowed                                                        ObErrorCode = -4709
	ObErrConsumerGroupNotExist                                                    ObErrorCode = -4710
	ObClusterNotAccessible                                                        ObErrorCode = -4711
	ObTenantResourceUnitExist                                                     ObErrorCode = -4712
	ObErrDropTruncatePartitionRebuildIndex                                        ObErrorCode = -4713
	ObErrAtlerTableIllegalFk                                                      ObErrorCode = -4714
	ObErrNoResourceManagerPrivilege                                               ObErrorCode = -4715
	ObLeaderCoordinatorNeedRetry                                                  ObErrorCode = -4716
	ObRebalanceTaskNeedRetry                                                      ObErrorCode = -4717
	ObErrResMgrPlanNotExist                                                       ObErrorCode = -4718
	ObLsNotExist                                                                  ObErrorCode = -4719
	ObTooManyTenantLs                                                             ObErrorCode = -4720
	ObLsLocationNotExist                                                          ObErrorCode = -4721
	ObLsLocationLeaderNotExist                                                    ObErrorCode = -4722
	ObMappingBetweenTabletAndLsNotExist                                           ObErrorCode = -4723
	ObTabletExist                                                                 ObErrorCode = -4724
	ObTabletNotExist                                                              ObErrorCode = -4725
	ObErrStandbyStatus                                                            ObErrorCode = -4726
	ObLsNeedRevoke                                                                ObErrorCode = -4727
	ObErrLastPartitionInTheRangeSectionCannotBeDropped                            ObErrorCode = -4730
	ObErrSetIntervalIsNotLegalOnThisTable                                         ObErrorCode = -4731
	ObCheckClusterStatus                                                          ObErrorCode = -4732
	ObZoneResourceNotEnough                                                       ObErrorCode = -4733
	ObZoneServerNotEnough                                                         ObErrorCode = -4734
	ObSstableNotExist                                                             ObErrorCode = -4735
	ObResourceUnitValueInvalid                                                    ObErrorCode = -4736
	ObLsExist                                                                     ObErrorCode = -4737
	ObDeviceExist                                                                 ObErrorCode = -4738
	ObDeviceNotExist                                                              ObErrorCode = -4739
	ObLsReplicaTaskResultUncertain                                                ObErrorCode = -4740
	ObWaitReplayTimeout                                                           ObErrorCode = -4741
	ObWaitTabletReadyTimeout                                                      ObErrorCode = -4742
	ObFreezeServiceEpochMismatch                                                  ObErrorCode = -4743
	ObDeleteServerNotAllowed                                                      ObErrorCode = -4745
	ObPacketStatusUnknown                                                         ObErrorCode = -4746
	ObArbitrationServiceNotExist                                                  ObErrorCode = -4747
	ObArbitrationServiceAlreadyExist                                              ObErrorCode = -4748
	ObUnexpectedTabletStatus                                                      ObErrorCode = -4749
	ObInvalidTableStore                                                           ObErrorCode = -4750
	ObWaitDegrationTimeout                                                        ObErrorCode = -4751
	ObErrRootserviceStart                                                         ObErrorCode = -4752
	ObErrRootserviceStop                                                          ObErrorCode = -4753
	ObErrRootInspection                                                           ObErrorCode = -4754
	ObErrRootserviceThreadHung                                                    ObErrorCode = -4755
	ObMigrateNotCompatible                                                        ObErrorCode = -4756
	ObClusterInfoMaybeRemained                                                    ObErrorCode = -4757
	ObArbitrationInfoQueryFailed                                                  ObErrorCode = -4758
	ObIgnoreErrAccessVirtualTable                                                 ObErrorCode = -4759
	ObLsOffline                                                                   ObErrorCode = -4760
	ObLsIsDeleted                                                                 ObErrorCode = -4761
	ObSkipCheckingLsStatus                                                        ObErrorCode = -4762
	ObErrUseRowIdForUpdate                                                        ObErrorCode = -4763
	ObErrUnknownSetOption                                                         ObErrorCode = -4764
	ObLsNotLeader                                                                 ObErrorCode = -4767
	ObErrParserInit                                                               ObErrorCode = -5000
	ObErrParseSql                                                                 ObErrorCode = -5001
	ObErrResolveSql                                                               ObErrorCode = -5002
	ObErrGenPlan                                                                  ObErrorCode = -5003
	ObErrColumnSize                                                               ObErrorCode = -5007
	ObErrColumnDuplicate                                                          ObErrorCode = -5008
	ObErrOperatorUnknown                                                          ObErrorCode = -5010
	ObErrStarDuplicate                                                            ObErrorCode = -5011
	ObErrIllegalId                                                                ObErrorCode = -5012
	ObErrIllegalValue                                                             ObErrorCode = -5014
	ObErrColumnAmbiguous                                                          ObErrorCode = -5015
	ObErrLogicalPlanFaild                                                         ObErrorCode = -5016
	ObErrSchemaUnset                                                              ObErrorCode = -5017
	ObErrIllegalName                                                              ObErrorCode = -5018
	ObTableNotExist                                                               ObErrorCode = -5019
	ObErrTableExist                                                               ObErrorCode = -5020
	ObErrExprUnknown                                                              ObErrorCode = -5022
	ObErrIllegalType                                                              ObErrorCode = -5023
	ObErrPrimaryKeyDuplicate                                                      ObErrorCode = -5024
	ObErrKeyNameDuplicate                                                         ObErrorCode = -5025
	ObErrCreatetimeDuplicate                                                      ObErrorCode = -5026
	ObErrModifytimeDuplicate                                                      ObErrorCode = -5027
	ObErrIllegalIndex                                                             ObErrorCode = -5028
	ObErrInvalidSchema                                                            ObErrorCode = -5029
	ObErrInsertNullRowkey                                                         ObErrorCode = -5030
	ObErrDeleteNullRowkey                                                         ObErrorCode = -5032
	ObErrUserEmpty                                                                ObErrorCode = -5034
	ObErrUserNotExist                                                             ObErrorCode = -5035
	ObErrNoPrivilege                                                              ObErrorCode = -5036
	ObErrNoAvailablePrivilegeEntry                                                ObErrorCode = -5037
	ObErrWrongPassword                                                            ObErrorCode = -5038
	ObErrUserIsLocked                                                             ObErrorCode = -5039
	ObErrUpdateRowkeyColumn                                                       ObErrorCode = -5040
	ObErrUpdateJoinColumn                                                         ObErrorCode = -5041
	ObErrInvalidColumnNum                                                         ObErrorCode = -5042
	ObErrPrepareStmtNotFound                                                      ObErrorCode = -5043
	ObErrOlderPrivilegeVersion                                                    ObErrorCode = -5046
	ObErrLackOfRowkeyCol                                                          ObErrorCode = -5047
	ObErrUserExist                                                                ObErrorCode = -5050
	ObErrPasswordEmpty                                                            ObErrorCode = -5051
	ObErrGrantPrivilegesToCreateTable                                             ObErrorCode = -5052
	ObErrWrongDynamicParam                                                        ObErrorCode = -5053
	ObErrParamSize                                                                ObErrorCode = -5054
	ObErrFunctionUnknown                                                          ObErrorCode = -5055
	ObErrCreatModifyTimeColumn                                                    ObErrorCode = -5056
	ObErrModifyPrimaryKey                                                         ObErrorCode = -5057
	ObErrParamDuplicate                                                           ObErrorCode = -5058
	ObErrTooManySessions                                                          ObErrorCode = -5059
	ObErrTooManyPs                                                                ObErrorCode = -5061
	ObErrHintUnknown                                                              ObErrorCode = -5063
	ObErrWhenUnsatisfied                                                          ObErrorCode = -5064
	ObErrQueryInterrupted                                                         ObErrorCode = -5065
	ObErrSessionInterrupted                                                       ObErrorCode = -5066
	ObErrUnknownSessionId                                                         ObErrorCode = -5067
	ObErrProtocolNotRecognize                                                     ObErrorCode = -5068
	ObErrWriteAuthError                                                           ObErrorCode = -5069
	ObErrParseJoinInfo                                                            ObErrorCode = -5070
	ObErrAlterIndexColumn                                                         ObErrorCode = -5071
	ObErrModifyIndexTable                                                         ObErrorCode = -5072
	ObErrIndexUnavailable                                                         ObErrorCode = -5073
	ObErrNopValue                                                                 ObErrorCode = -5074
	ObErrPsTooManyParam                                                           ObErrorCode = -5080
	ObErrInvalidTypeForOp                                                         ObErrorCode = -5083
	ObErrCastVarcharToBool                                                        ObErrorCode = -5084
	ObErrCastVarcharToNumber                                                      ObErrorCode = -5085
	ObErrCastVarcharToTime                                                        ObErrorCode = -5086
	ObErrCastNumberOverflow                                                       ObErrorCode = -5087
	ObSchemaNumberPrecisionOverflow                                               ObErrorCode = -5090
	ObSchemaNumberScaleOverflow                                                   ObErrorCode = -5091
	ObErrIndexUnknown                                                             ObErrorCode = -5092
	ObErrTooManyJoinTables                                                        ObErrorCode = -5094
	ObErrDdlOnRemoteDatabase                                                      ObErrorCode = -5095
	ObErrMissingKeyword                                                           ObErrorCode = -5096
	ObErrDatabaseLinkExpected                                                     ObErrorCode = -5097
	ObErrVarcharTooLong                                                           ObErrorCode = -5098
	ObErrLocalVariable                                                            ObErrorCode = -5100
	ObErrGlobalVariable                                                           ObErrorCode = -5101
	ObErrVariableIsReadonly                                                       ObErrorCode = -5102
	ObErrIncorrectGlobalLocalVar                                                  ObErrorCode = -5103
	ObErrExpireInfoTooLong                                                        ObErrorCode = -5104
	ObErrExpireCondTooLong                                                        ObErrorCode = -5105
	ObErrUserVariableUnknown                                                      ObErrorCode = -5112
	ObIllegalUsageOfMergingFrozenTime                                             ObErrorCode = -5113
	ObSqlLogOpSetchildOverflow                                                    ObErrorCode = -5116
	ObSqlExplainFailed                                                            ObErrorCode = -5117
	ObSqlOptCopyOpFailed                                                          ObErrorCode = -5118
	ObSqlOptGenPlanFalied                                                         ObErrorCode = -5119
	ObSqlOptCreateRawexprFailed                                                   ObErrorCode = -5120
	ObSqlOptJoinOrderFailed                                                       ObErrorCode = -5121
	ObSqlOptError                                                                 ObErrorCode = -5122
	ObErrOciInitTimezone                                                          ObErrorCode = -5123
	ObErrZlibData                                                                 ObErrorCode = -5124
	ObErrDblinkSessionKilled                                                      ObErrorCode = -5125
	ObSqlResolverNoMemory                                                         ObErrorCode = -5130
	ObSqlDmlOnly                                                                  ObErrorCode = -5131
	ObErrNoGrant                                                                  ObErrorCode = -5133
	ObErrNoDbSelected                                                             ObErrorCode = -5134
	ObSqlPcOverflow                                                               ObErrorCode = -5135
	ObSqlPcPlanDuplicate                                                          ObErrorCode = -5136
	ObSqlPcPlanExpire                                                             ObErrorCode = -5137
	ObSqlPcNotExist                                                               ObErrorCode = -5138
	ObSqlParamsLimit                                                              ObErrorCode = -5139
	ObSqlPcPlanSizeLimit                                                          ObErrorCode = -5140
	ObUnknownPartition                                                            ObErrorCode = -5146
	ObPartitionNotMatch                                                           ObErrorCode = -5147
	ObErPasswdLength                                                              ObErrorCode = -5148
	ObErrInsertInnerJoinColumn                                                    ObErrorCode = -5149
	ObTablegroupNotExist                                                          ObErrorCode = -5151
	ObSubQueryTooManyRow                                                          ObErrorCode = -5153
	ObErrBadDatabase                                                              ObErrorCode = -5154
	ObCannotUser                                                                  ObErrorCode = -5155
	ObTenantExist                                                                 ObErrorCode = -5156
	ObDatabaseExist                                                               ObErrorCode = -5158
	ObTablegroupExist                                                             ObErrorCode = -5159
	ObErrInvalidTenantName                                                        ObErrorCode = -5160
	ObEmptyTenant                                                                 ObErrorCode = -5161
	ObWrongDbName                                                                 ObErrorCode = -5162
	ObWrongTableName                                                              ObErrorCode = -5163
	ObWrongColumnName                                                             ObErrorCode = -5164
	ObErrColumnSpec                                                               ObErrorCode = -5165
	ObErrDbDropExists                                                             ObErrorCode = -5166
	ObErrCreateUserWithGrant                                                      ObErrorCode = -5169
	ObErrNoDbPrivilege                                                            ObErrorCode = -5170
	ObErrNoTablePrivilege                                                         ObErrorCode = -5171
	ObInvalidOnUpdate                                                             ObErrorCode = -5172
	ObInvalidDefault                                                              ObErrorCode = -5173
	ObErrUpdateTableUsed                                                          ObErrorCode = -5174
	ObErrCoulumnValueNotMatch                                                     ObErrorCode = -5175
	ObErrInvalidGroupFuncUse                                                      ObErrorCode = -5176
	ObErrFieldTypeNotAllowedAsPartitionField                                      ObErrorCode = -5178
	ObErrTooLongIdent                                                             ObErrorCode = -5179
	ObErrWrongTypeForVar                                                          ObErrorCode = -5180
	ObWrongUserNameLength                                                         ObErrorCode = -5181
	ObErrPrivUsage                                                                ObErrorCode = -5182
	ObIllegalGrantForTable                                                        ObErrorCode = -5183
	ObErrReachAutoincMax                                                          ObErrorCode = -5184
	ObErrNoTablesUsed                                                             ObErrorCode = -5185
	ObCantRemoveAllFields                                                         ObErrorCode = -5187
	ObTooManyPartitionsError                                                      ObErrorCode = -5188
	ObNoPartsError                                                                ObErrorCode = -5189
	ObWrongSubKey                                                                 ObErrorCode = -5190
	ObKeyPart_0                                                                   ObErrorCode = -5191
	ObErrWrongAutoKey                                                             ObErrorCode = -5193
	ObErrTooManyKeys                                                              ObErrorCode = -5194
	ObErrTooManyRowkeyColumns                                                     ObErrorCode = -5195
	ObErrTooLongKeyLength                                                         ObErrorCode = -5196
	ObErrTooManyColumns                                                           ObErrorCode = -5197
	ObErrTooLongColumnLength                                                      ObErrorCode = -5198
	ObErrTooBigRowsize                                                            ObErrorCode = -5199
	ObErrUnknownTable                                                             ObErrorCode = -5200
	ObErrBadTable                                                                 ObErrorCode = -5201
	ObErrTooBigScale                                                              ObErrorCode = -5202
	ObErrTooBigDisplaywidth                                                       ObErrorCode = -5205
	ObWrongGroupField                                                             ObErrorCode = -5206
	ObNonUniqError                                                                ObErrorCode = -5207
	ObErrNonuniqTable                                                             ObErrorCode = -5208
	ObErrCantDropFieldOrKey                                                       ObErrorCode = -5209
	ObErrMultiplePriKey                                                           ObErrorCode = -5210
	ObErrKeyColumnDoesNotExits                                                    ObErrorCode = -5211
	ObErrAutoPartitionKey                                                         ObErrorCode = -5212
	ObErrCantUseOptionHere                                                        ObErrorCode = -5213
	ObErrWrongObject                                                              ObErrorCode = -5214
	ObErrOnRename                                                                 ObErrorCode = -5215
	ObErrWrongKeyColumn                                                           ObErrorCode = -5216
	ObErrBadFieldError                                                            ObErrorCode = -5217
	ObErrWrongFieldWithGroup                                                      ObErrorCode = -5218
	ObErrCantChangeTxCharacteristics                                              ObErrorCode = -5219
	ObErrCantExecuteInReadOnlyTransaction                                         ObErrorCode = -5220
	ObErrMixOfGroupFuncAndFields                                                  ObErrorCode = -5221
	ObErrWrongIdentName                                                           ObErrorCode = -5223
	ObWrongNameForIndex                                                           ObErrorCode = -5224
	ObIllegalReference                                                            ObErrorCode = -5225
	ObReachMemoryLimit                                                            ObErrorCode = -5226
	ObErrPasswordFormat                                                           ObErrorCode = -5227
	ObErrNonUpdatableTable                                                        ObErrorCode = -5228
	ObErrWarnDataOutOfRange                                                       ObErrorCode = -5229
	ObErrWrongExprInPartitionFuncError                                            ObErrorCode = -5230
	ObErrViewInvalid                                                              ObErrorCode = -5231
	ObErrOptionPreventsStatement                                                  ObErrorCode = -5233
	ObErrDbReadOnly                                                               ObErrorCode = -5234
	ObErrTableReadOnly                                                            ObErrorCode = -5235
	ObErrLockOrActiveTransaction                                                  ObErrorCode = -5236
	ObErrSameNamePartitionField                                                   ObErrorCode = -5237
	ObErrTablenameNotAllowedHere                                                  ObErrorCode = -5238
	ObErrViewRecursive                                                            ObErrorCode = -5239
	ObErrQualifier                                                                ObErrorCode = -5240
	ObErrViewWrongList                                                            ObErrorCode = -5242
	ObSysVarsMaybeDiffVersion                                                     ObErrorCode = -5243
	ObErrAutoIncrementConflict                                                    ObErrorCode = -5244
	ObErrTaskSkipped                                                              ObErrorCode = -5245
	ObErrNameBecomesEmpty                                                         ObErrorCode = -5246
	ObErrRemovedSpaces                                                            ObErrorCode = -5247
	ObWarnAddAutoincrementColumn                                                  ObErrorCode = -5248
	ObWarnChamgeNullAttribute                                                     ObErrorCode = -5249
	ObErrInvalidCharacterString                                                   ObErrorCode = -5250
	ObErrKillDenied                                                               ObErrorCode = -5251
	ObErrColumnDefinitionAmbiguous                                                ObErrorCode = -5252
	ObErrEmptyQuery                                                               ObErrorCode = -5253
	ObErrCutValueGroupConcat                                                      ObErrorCode = -5254
	ObErrFieldNotFoundPart                                                        ObErrorCode = -5255
	ObErrPrimaryCantHaveNull                                                      ObErrorCode = -5256
	ObErrPartitionFuncNotAllowedError                                             ObErrorCode = -5257
	ObErrInvalidBlockSize                                                         ObErrorCode = -5258
	ObErrUnknownStorageEngine                                                     ObErrorCode = -5259
	ObErrTenantIsLocked                                                           ObErrorCode = -5260
	ObEerUniqueKeyNeedAllFieldsInPf                                               ObErrorCode = -5261
	ObErrPartitionFunctionIsNotAllowed                                            ObErrorCode = -5262
	ObErrAggregateOrderForUnion                                                   ObErrorCode = -5263
	ObErrOutlineExist                                                             ObErrorCode = -5264
	ObOutlineNotExist                                                             ObErrorCode = -5265
	ObWarnOptionBelowLimit                                                        ObErrorCode = -5266
	ObInvalidOutline                                                              ObErrorCode = -5267
	ObReachMaxConcurrentNum                                                       ObErrorCode = -5268
	ObErrOperationOnRecycleObject                                                 ObErrorCode = -5269
	ObErrObjectNotInRecyclebin                                                    ObErrorCode = -5270
	ObErrConCountError                                                            ObErrorCode = -5271
	ObErrOutlineContentExist                                                      ObErrorCode = -5272
	ObErrOutlineMaxConcurrentExist                                                ObErrorCode = -5273
	ObErrValuesIsNotIntTypeError                                                  ObErrorCode = -5274
	ObErrWrongTypeColumnValueError                                                ObErrorCode = -5275
	ObErrPartitionColumnListError                                                 ObErrorCode = -5276
	ObErrTooManyValuesError                                                       ObErrorCode = -5277
	ObErrPartitionValueError                                                      ObErrorCode = -5278
	ObErrPartitionIntervalError                                                   ObErrorCode = -5279
	ObErrSameNamePartition                                                        ObErrorCode = -5280
	ObErrRangeNotIncreasingError                                                  ObErrorCode = -5281
	ObErrParsePartitionRange                                                      ObErrorCode = -5282
	ObErrUniqueKeyNeedAllFieldsInPf                                               ObErrorCode = -5283
	ObNoPartitionForGivenValue                                                    ObErrorCode = -5284
	ObEerNullInValuesLessThan                                                     ObErrorCode = -5285
	ObErrPartitionConstDomainError                                                ObErrorCode = -5286
	ObErrTooManyPartitionFuncFields                                               ObErrorCode = -5287
	ObErrBadFtColumn                                                              ObErrorCode = -5288
	ObErrKeyDoesNotExists                                                         ObErrorCode = -5289
	ObNonDefaultValueForGeneratedColumn                                           ObErrorCode = -5290
	ObErrBadCtxcatColumn                                                          ObErrorCode = -5291
	ObErrUnsupportedActionOnGeneratedColumn                                       ObErrorCode = -5292
	ObErrDependentByGeneratedColumn                                               ObErrorCode = -5293
	ObErrTooManyRows                                                              ObErrorCode = -5294
	ObWrongFieldTerminators                                                       ObErrorCode = -5295
	ObNoReadableReplica                                                           ObErrorCode = -5296
	ObErrSynonymExist                                                             ObErrorCode = -5298
	ObSynonymNotExist                                                             ObErrorCode = -5299
	ObErrMissOrderByExpr                                                          ObErrorCode = -5300
	ObErrNotConstExpr                                                             ObErrorCode = -5301
	ObErrPartitionMgmtOnNonpartitioned                                            ObErrorCode = -5302
	ObErrDropPartitionNonExistent                                                 ObErrorCode = -5303
	ObErrPartitionMgmtOnTwopartTable                                              ObErrorCode = -5304
	ObErrOnlyOnRangeListPartition                                                 ObErrorCode = -5305
	ObErrDropLastPartition                                                        ObErrorCode = -5306
	ObErrParallelServersTargetNotEnough                                           ObErrorCode = -5307
	ObErrIgnoreUserHostName                                                       ObErrorCode = -5308
	ObIgnoreSqlInRestore                                                          ObErrorCode = -5309
	ObErrTemporaryTableWithPartition                                              ObErrorCode = -5310
	ObErrInvalidColumnId                                                          ObErrorCode = -5311
	ObSyncDdlDuplicate                                                            ObErrorCode = -5312
	ObSyncDdlError                                                                ObErrorCode = -5313
	ObErrRowIsReferenced                                                          ObErrorCode = -5314
	ObErrNoReferencedRow                                                          ObErrorCode = -5315
	ObErrFuncResultTooLarge                                                       ObErrorCode = -5316
	ObErrCannotAddForeign                                                         ObErrorCode = -5317
	ObErrWrongFkDef                                                               ObErrorCode = -5318
	ObErrInvalidChildColumnLengthFk                                               ObErrorCode = -5319
	ObErrAlterColumnFk                                                            ObErrorCode = -5320
	ObErrConnectByRequired                                                        ObErrorCode = -5321
	ObErrInvalidPseudoColumnPlace                                                 ObErrorCode = -5322
	ObErrNocycleRequired                                                          ObErrorCode = -5323
	ObErrConnectByLoop                                                            ObErrorCode = -5324
	ObErrInvalidSiblings                                                          ObErrorCode = -5325
	ObErrInvalidSeparator                                                         ObErrorCode = -5326
	ObErrInvalidSynonymName                                                       ObErrorCode = -5327
	ObErrLoopOfSynonym                                                            ObErrorCode = -5328
	ObErrSynonymSameAsObject                                                      ObErrorCode = -5329
	ObErrSynonymTranslationInvalid                                                ObErrorCode = -5330
	ObErrExistObject                                                              ObErrorCode = -5331
	ObErrIllegalValueForType                                                      ObErrorCode = -5332
	ObErTooLongSetEnumValue                                                       ObErrorCode = -5333
	ObErDuplicatedValueInType                                                     ObErrorCode = -5334
	ObErTooBigEnum                                                                ObErrorCode = -5335
	ObErrTooBigSet                                                                ObErrorCode = -5336
	ObErrWrongRowId                                                               ObErrorCode = -5337
	ObErrInvalidWindowFunctionPlace                                               ObErrorCode = -5338
	ObErrParsePartitionList                                                       ObErrorCode = -5339
	ObErrMultipleDefConstInListPart                                               ObErrorCode = -5340
	ObErrWrongFuncArgumentsType                                                   ObErrorCode = -5343
	ObErrMultiUpdateKeyConflict                                                   ObErrorCode = -5344
	ObErrInsufficientPxWorker                                                     ObErrorCode = -5345
	ObErrForUpdateExprNotAllowed                                                  ObErrorCode = -5346
	ObErrWinFuncArgNotInPartitionBy                                               ObErrorCode = -5347
	ObErrTooLongStringInConcat                                                    ObErrorCode = -5348
	ObErrWrongTimestampLtzColumnValueError                                        ObErrorCode = -5349
	ObErrUpdCausePartChange                                                       ObErrorCode = -5350
	ObErrInvalidTypeForArgument                                                   ObErrorCode = -5351
	ObErrAddPartBounNotInc                                                        ObErrorCode = -5353
	ObErrDataTooLongInPartCheck                                                   ObErrorCode = -5354
	ObErrWrongTypeColumnValueV2Error                                              ObErrorCode = -5355
	ObCantAggregate_3collations                                                   ObErrorCode = -5356
	ObCantAggregateNcollations                                                    ObErrorCode = -5357
	ObErrDuplicatedUniqueKey                                                      ObErrorCode = -5358
	ObDoubleOverflow                                                              ObErrorCode = -5359
	ObErrNoSysPrivilege                                                           ObErrorCode = -5360
	ObErrNoLoginPrivilege                                                         ObErrorCode = -5361
	ObErrCannotRevokePrivilegesYouDidNotGrant                                     ObErrorCode = -5362
	ObErrSystemPrivilegesNotGrantedTo                                             ObErrorCode = -5363
	ObErrOnlySelectAndAlterPrivilegesAreValidForSequences                         ObErrorCode = -5364
	ObErrExecutePrivilegeNotAllowedForTables                                      ObErrorCode = -5365
	ObErrOnlyExecuteAndDebugPrivilegesAreValidForProcedures                       ObErrorCode = -5366
	ObErrOnlyExecuteDebugAndUnderPrivilegesAreValidForTypes                       ObErrorCode = -5367
	ObErrAdminOptionNotGrantedForRole                                             ObErrorCode = -5368
	ObErrUserOrRoleDoesNotExist                                                   ObErrorCode = -5369
	ObErrMissingOnKeyword                                                         ObErrorCode = -5370
	ObErrNoGrantOption                                                            ObErrorCode = -5371
	ObErrAlterIndexAndExecuteNotAllowedForViews                                   ObErrorCode = -5372
	ObErrCircularRoleGrantDetected                                                ObErrorCode = -5373
	ObErrInvalidPrivilegeOnDirectories                                            ObErrorCode = -5374
	ObErrDirectoryAccessDenied                                                    ObErrorCode = -5375
	ObErrMissingOrInvalidRoleName                                                 ObErrorCode = -5376
	ObErrRoleNotGrantedOrDoesNotExist                                             ObErrorCode = -5377
	ObErrDefaultRoleNotGrantedToUser                                              ObErrorCode = -5378
	ObErrRoleNotGrantedTo                                                         ObErrorCode = -5379
	ObErrCannotGrantToARoleWithGrantOption                                        ObErrorCode = -5380
	ObErrDuplicateUsernameInList                                                  ObErrorCode = -5381
	ObErrCannotGrantStringToARole                                                 ObErrorCode = -5382
	ObErrCascadeConstraintsMustBeSpecifiedToPerformThisRevoke                     ObErrorCode = -5383
	ObErrYouMayNotRevokePrivilegesFromYourself                                    ObErrorCode = -5384
	ObErrMissErrLogMandatoryColumn                                                ObErrorCode = -5385
	ObTableDefinitionChanged                                                      ObErrorCode = -5386
	ObErrObjectStringDoesNotExist                                                 ObErrorCode = -5400
	ObErrResultantDataTypeOfVirtualColumnIsNotSupported                           ObErrorCode = -5401
	ObErrGetStackedDiagnostics                                                    ObErrorCode = -5402
	ObDdlSchemaVersionNotMatch                                                    ObErrorCode = -5403
	ObErrColumnGroupDuplicate                                                     ObErrorCode = -5404
	ObErrReservedSyntax                                                           ObErrorCode = -5405
	ObErrInvalidParamToProcedure                                                  ObErrorCode = -5406
	ObErrWrongParametersToNativeFct                                               ObErrorCode = -5407
	ObErrCteMaxRecursionDepth                                                     ObErrorCode = -5408
	ObDuplicateObjectNameExist                                                    ObErrorCode = -5409
	ObErrRefreshSchemaTooLong                                                     ObErrorCode = -5410
	ObSqlRetrySpm                                                                 ObErrorCode = -5434
	ObOutlineNotReproducible                                                      ObErrorCode = -5435
	ObEerWindowNoChildPartitioning                                                ObErrorCode = -5438
	ObEerWindowNoInheritFrame                                                     ObErrorCode = -5439
	ObEerWindowNoRedefineOrderBy                                                  ObErrorCode = -5440
	ObErrInvalidDataTypeReturning                                                 ObErrorCode = -5443
	ObErrJsonValueNoValue                                                         ObErrorCode = -5445
	ObErrDefaultValueNotLiteral                                                   ObErrorCode = -5446
	ObErrJsonSyntaxError                                                          ObErrorCode = -5447
	ObErrJsonEqualOutsidePredicate                                                ObErrorCode = -5448
	ObErrWithoutArrWrapper                                                        ObErrorCode = -5449
	ObErrJsonPatchInvalid                                                         ObErrorCode = -5450
	ObErrOrderSiblingsByNotAllowed                                                ObErrorCode = -5451
	ObErrLobTypeNotSorting                                                        ObErrorCode = -5452
	ObErrJsonIllegalZeroLengthIdentifierError                                     ObErrorCode = -5455
	ObErrNoValueInPassing                                                         ObErrorCode = -5456
	ObErrInvalidColumnSpe                                                         ObErrorCode = -5457
	ObErrInputJsonNotBeNull                                                       ObErrorCode = -5458
	ObErrInvalidDataType                                                          ObErrorCode = -5459
	ObErrInvalidClause                                                            ObErrorCode = -5460
	ObErrInvalidCmpOp                                                             ObErrorCode = -5461
	ObErrInvalidInput                                                             ObErrorCode = -5462
	ObErrEmptyInputToJsonOperator                                                 ObErrorCode = -5463
	ObErrAdditionalIsJson                                                         ObErrorCode = -5464
	ObErrFunctionInvalidState                                                     ObErrorCode = -5465
	ObErrMissValue                                                                ObErrorCode = -5466
	ObErrDifferentTypeSelected                                                    ObErrorCode = -5467
	ObErrNoValueSelected                                                          ObErrorCode = -5468
	ObErrNonTextRetNotsupport                                                     ObErrorCode = -5469
	ObErrPlJsontypeUsage                                                          ObErrorCode = -5470
	ObErrNullInput                                                                ObErrorCode = -5471
	ObErrDefaultValueNotMatch                                                     ObErrorCode = -5474
	ObErrConversionFail                                                           ObErrorCode = -5476
	ObErrNotObjRef                                                                ObErrorCode = -5477
	ObErrUnsupportTruncateType                                                    ObErrorCode = -5478
	ObErrUnimplementJsonFeature                                                   ObErrorCode = -5479
	ObErrUsageKeyword                                                             ObErrorCode = -5480
	ObErrInputJsonTable                                                           ObErrorCode = -5481
	ObErrBoolCastNumber                                                           ObErrorCode = -5482
	ObErrNestedPathDisjunct                                                       ObErrorCode = -5483
	ObErrInvalidVariableInJsonPath                                                ObErrorCode = -5484
	ObErrInvalidDefaultValueProvided                                              ObErrorCode = -5486
	ObErrPathExpressionNotLiteral                                                 ObErrorCode = -5487
	ObErrInvalidArgumentForJsonCall                                               ObErrorCode = -5488
	ObErrSchemaHistoryEmpty                                                       ObErrorCode = -5489
	ObErrTableNameNotInList                                                       ObErrorCode = -5490
	ObErrDefaultNotAtLastInListPart                                               ObErrorCode = -5491
	ObErrMysqlCharacterSetMismatch                                                ObErrorCode = -5492
	ObErrRenamePartitionNameDuplicate                                             ObErrorCode = -5493
	ObErrRenameSubpartitionNameDuplicate                                          ObErrorCode = -5494
	ObErrInvalidWaitInterval                                                      ObErrorCode = -5495
	ObErrFunctionalIndexRefAutoIncrement                                          ObErrorCode = -5496
	ObErrDependentByFunctionalIndex                                               ObErrorCode = -5497
	ObErrFunctionalIndexOnLob                                                     ObErrorCode = -5498
	ObErrFunctionalIndexOnField                                                   ObErrorCode = -5499
	ObErrGencolLegitCheckFailed                                                   ObErrorCode = -5500
	ObErrGroupingFuncWithoutGroupBy                                               ObErrorCode = -5501
	ObErrDependentByPartitionFunc                                                 ObErrorCode = -5502
	ObErrViewSelectContainInto                                                    ObErrorCode = -5503
	ObErrDefaultNotAllowed                                                        ObErrorCode = -5504
	ObErrModifyRealcolToGencol                                                    ObErrorCode = -5505
	ObErrModifyTypeOfGencol                                                       ObErrorCode = -5506
	ObErrWindowFrameIllegal                                                       ObErrorCode = -5507
	ObErrWindowRangeFrameTemporalType                                             ObErrorCode = -5508
	ObErrWindowRangeFrameNumericType                                              ObErrorCode = -5509
	ObErrWindowRangeBoundNotConstant                                              ObErrorCode = -5510
	ObErrDefaultForModifyingViews                                                 ObErrorCode = -5511
	ObErrFkColumnNotNull                                                          ObErrorCode = -5512
	ObErrUnsupportedFkSetNullOnGeneratedColumn                                    ObErrorCode = -5513
	ObJsonProcessingError                                                         ObErrorCode = -5514
	ObErrTableWithoutAlias                                                        ObErrorCode = -5515
	ObErrDeprecatedSyntax                                                         ObErrorCode = -5516
	ObErrSpAlreadyExists                                                          ObErrorCode = -5541
	ObErrSpDoesNotExist                                                           ObErrorCode = -5542
	ObErrSpUndeclaredVar                                                          ObErrorCode = -5543
	ObErrSpUndeclaredType                                                         ObErrorCode = -5544
	ObErrSpCondMismatch                                                           ObErrorCode = -5545
	ObErrSpLilabelMismatch                                                        ObErrorCode = -5546
	ObErrSpCursorMismatch                                                         ObErrorCode = -5547
	ObErrSpDupParam                                                               ObErrorCode = -5548
	ObErrSpDupVar                                                                 ObErrorCode = -5549
	ObErrSpDupType                                                                ObErrorCode = -5550
	ObErrSpDupCondition                                                           ObErrorCode = -5551
	ObErrSpDupLabel                                                               ObErrorCode = -5552
	ObErrSpDupCursor                                                              ObErrorCode = -5553
	ObErrSpInvalidFetchArg                                                        ObErrorCode = -5554
	ObErrSpWrongArgNum                                                            ObErrorCode = -5555
	ObErrSpUnhandledException                                                     ObErrorCode = -5556
	ObErrSpBadConditionType                                                       ObErrorCode = -5557
	ObErrPackageAlreadyExists                                                     ObErrorCode = -5558
	ObErrPackageDoseNotExist                                                      ObErrorCode = -5559
	ObEerUnknownStmtHandler                                                       ObErrorCode = -5560
	ObErrInvalidWindowFuncUse                                                     ObErrorCode = -5561
	ObErrConstraintDuplicate                                                      ObErrorCode = -5562
	ObErrContraintNotFound                                                        ObErrorCode = -5563
	ObErrAlterTableAlterDuplicatedIndex                                           ObErrorCode = -5564
	ObEerInvalidArgumentForLogarithm                                              ObErrorCode = -5565
	ObErrReorganizeOutsideRange                                                   ObErrorCode = -5566
	ObErSpRecursionLimit                                                          ObErrorCode = -5567
	ObErUnsupportedPs                                                             ObErrorCode = -5568
	ObErStmtNotAllowedInSfOrTrg                                                   ObErrorCode = -5569
	ObErSpNoRecursion                                                             ObErrorCode = -5570
	ObErSpCaseNotFound                                                            ObErrorCode = -5571
	ObErrInvalidSplitCount                                                        ObErrorCode = -5572
	ObErrInvalidSplitGrammar                                                      ObErrorCode = -5573
	ObErrMissValues                                                               ObErrorCode = -5574
	ObErrMissAtValues                                                             ObErrorCode = -5575
	ObErCommitNotAllowedInSfOrTrg                                                 ObErrorCode = -5576
	ObPcGetLocationError                                                          ObErrorCode = -5577
	ObPcLockConflict                                                              ObErrorCode = -5578
	ObErSpNoRetset                                                                ObErrorCode = -5579
	ObErSpNoreturnend                                                             ObErrorCode = -5580
	ObErrSpDupHandler                                                             ObErrorCode = -5581
	ObErSpNoRecursiveCreate                                                       ObErrorCode = -5582
	ObErSpBadreturn                                                               ObErrorCode = -5583
	ObErSpBadCursorSelect                                                         ObErrorCode = -5584
	ObErSpBadSqlstate                                                             ObErrorCode = -5585
	ObErSpVarcondAfterCurshndlr                                                   ObErrorCode = -5586
	ObErSpCursorAfterHandler                                                      ObErrorCode = -5587
	ObErSpWrongName                                                               ObErrorCode = -5588
	ObErSpCursorAlreadyOpen                                                       ObErrorCode = -5589
	ObErSpCursorNotOpen                                                           ObErrorCode = -5590
	ObErSpCantSetAutocommit                                                       ObErrorCode = -5591
	ObErSpNotVarArg                                                               ObErrorCode = -5592
	ObErSpLilabelMismatch                                                         ObErrorCode = -5593
	ObErrTruncateIllegalFk                                                        ObErrorCode = -5594
	ObErrDupKey                                                                   ObErrorCode = -5595
	ObErInvalidUseOfNull                                                          ObErrorCode = -5596
	ObErrSplitListLessValue                                                       ObErrorCode = -5597
	ObErrAddPartitionToDefaultList                                                ObErrorCode = -5598
	ObErrSplitIntoOnePartition                                                    ObErrorCode = -5599
	ObErrNoTenantPrivilege                                                        ObErrorCode = -5600
	ObErrInvalidPercentage                                                        ObErrorCode = -5601
	ObErrCollectHistogram                                                         ObErrorCode = -5602
	ObErTempTableInUse                                                            ObErrorCode = -5603
	ObErrInvalidNlsParameterString                                                ObErrorCode = -5604
	ObErrDatetimeIntervalPrecisionOutOfRange                                      ObErrorCode = -5605
	ObErrCmdNotProperlyEnded                                                      ObErrorCode = -5607
	ObErrInvalidNumberFormatModel                                                 ObErrorCode = -5608
	ObWarnNonAsciiSeparatorNotImplemented                                         ObErrorCode = -5609
	ObWarnAmbiguousFieldTerm                                                      ObErrorCode = -5610
	ObWarnTooFewRecords                                                           ObErrorCode = -5611
	ObWarnTooManyRecords                                                          ObErrorCode = -5612
	ObErrTooManyValues                                                            ObErrorCode = -5613
	ObErrNotEnoughValues                                                          ObErrorCode = -5614
	ObErrMoreThanOneRow                                                           ObErrorCode = -5615
	ObErrNotSubQuery                                                              ObErrorCode = -5616
	ObInappropriateInto                                                           ObErrorCode = -5617
	ObErrTableIsReferenced                                                        ObErrorCode = -5618
	ObErrQualifierExistsForUsingColumn                                            ObErrorCode = -5619
	ObErrOuterJoinNested                                                          ObErrorCode = -5620
	ObErrMultiOuterJoinTable                                                      ObErrorCode = -5621
	ObErrOuterJoinOnCorrelationColumn                                             ObErrorCode = -5622
	ObErrOuterJoinAmbiguous                                                       ObErrorCode = -5623
	ObErrOuterJoinWithSubQuery                                                    ObErrorCode = -5624
	ObErrOuterJoinWithAnsiJoin                                                    ObErrorCode = -5625
	ObErrOuterJoinNotAllowed                                                      ObErrorCode = -5626
	ObSchemaEagain                                                                ObErrorCode = -5627
	ObErrZeroLenCol                                                               ObErrorCode = -5628
	ObErrInvalidCharFollowingEscapeChar                                           ObErrorCode = -5665
	ObErrInvalidEscapeCharLength                                                  ObErrorCode = -5666
	ObErrNotSelectedExpr                                                          ObErrorCode = -5668
	ObErrUkPkDuplicate                                                            ObErrorCode = -5671
	ObErrColumnListAlreadyIndexed                                                 ObErrorCode = -5672
	ObErrBushyTreeNotSupported                                                    ObErrorCode = -5673
	ObErrOrderByItemNotInSelectList                                               ObErrorCode = -5675
	ObErrNumericOrValueError                                                      ObErrorCode = -5677
	ObErrConstraintNameDuplicate                                                  ObErrorCode = -5678
	ObErrOnlyHaveInvisibleColInTable                                              ObErrorCode = -5679
	ObErrInvisibleColOnUnsupportedTableType                                       ObErrorCode = -5680
	ObErrModifyColVisibilityCombinedWithOtherOption                               ObErrorCode = -5681
	ObErrModifyColVisibilityBySysUser                                             ObErrorCode = -5682
	ObErrTooManyArgsForFun                                                        ObErrorCode = -5683
	ObPxSqlNeedRetry                                                              ObErrorCode = -5684
	ObTenantHasBeenDropped                                                        ObErrorCode = -5685
	ObErrExtractFieldInvalid                                                      ObErrorCode = -5686
	ObErrPackageCompileError                                                      ObErrorCode = -5687
	ObErrSpEmptyBlock                                                             ObErrorCode = -5688
	ObArrayBindingRollback                                                        ObErrorCode = -5689
	ObErrInvalidSubQueryUse                                                       ObErrorCode = -5690
	ObErrDateOrSysVarCannotInCheckCst                                             ObErrorCode = -5691
	ObErrNonexistentConstraint                                                    ObErrorCode = -5692
	ObErrCheckConstraintViolated                                                  ObErrorCode = -5693
	ObErrGroupFuncNotAllowed                                                      ObErrorCode = -5694
	ObErrPolicyStringNotFound                                                     ObErrorCode = -5695
	ObErrInvalidLabelString                                                       ObErrorCode = -5696
	ObErrUndefinedCompartmentStringForPolicyString                                ObErrorCode = -5697
	ObErrUndefinedLevelStringForPolicyString                                      ObErrorCode = -5698
	ObErrUndefinedGroupStringForPolicyString                                      ObErrorCode = -5699
	ObErrLbacError                                                                ObErrorCode = -5700
	ObErrPolicyRoleAlreadyExistsForPolicyString                                   ObErrorCode = -5701
	ObErrNullOrInvalidUserLabel                                                   ObErrorCode = -5702
	ObErrAddIndex                                                                 ObErrorCode = -5703
	ObErrProfileStringDoesNotExist                                                ObErrorCode = -5704
	ObErrInvalidResourceLimit                                                     ObErrorCode = -5705
	ObErrProfileStringAlreadyExists                                               ObErrorCode = -5706
	ObErrProfileStringHasUsersAssigned                                            ObErrorCode = -5707
	ObErrAddCheckConstraintViolated                                               ObErrorCode = -5713
	ObErrIllegalViewUpdate                                                        ObErrorCode = -5714
	ObErrVirtualColNotAllowed                                                     ObErrorCode = -5715
	ObErrOViewMultiupdate                                                         ObErrorCode = -5716
	ObErrNonInsertableTable                                                       ObErrorCode = -5717
	ObErrViewMultiupdate                                                          ObErrorCode = -5718
	ObErrNonupdateableColumn                                                      ObErrorCode = -5719
	ObErrViewDeleteMergeView                                                      ObErrorCode = -5720
	ObErrODeleteViewNonKeyPreserved                                               ObErrorCode = -5721
	ObErrOUpdateViewNonKeyPreserved                                               ObErrorCode = -5722
	ObErrModifyReadOnlyView                                                       ObErrorCode = -5723
	ObErrInvalidInitransValue                                                     ObErrorCode = -5724
	ObErrInvalidMaxtransValue                                                     ObErrorCode = -5725
	ObErrInvalidPctfreeOrPctusedValue                                             ObErrorCode = -5726
	ObErrProxyReroute                                                             ObErrorCode = -5727
	ObErrIllegalArgumentForFunction                                               ObErrorCode = -5728
	ObErrInvalidSamplingRange                                                     ObErrorCode = -5730
	ObErrSpecifyDatabaseNotAllowed                                                ObErrorCode = -5731
	ObErrStmtTriggerWithWhenClause                                                ObErrorCode = -5732
	ObErrTriggerNotExist                                                          ObErrorCode = -5733
	ObErrTriggerAlreadyExist                                                      ObErrorCode = -5734
	ObErrTriggerExistOnOtherTable                                                 ObErrorCode = -5735
	ObErrSignaledInParallelQueryServer                                            ObErrorCode = -5736
	ObErrCteIllegalQueryName                                                      ObErrorCode = -5737
	ObErrCteUnsupportedColumnAliasing                                             ObErrorCode = -5738
	ObErrUnsupportedUseOfCte                                                      ObErrorCode = -5739
	ObErrCteColumnNumberNotMatch                                                  ObErrorCode = -5740
	ObErrNeedColumnAliasListInRecursiveCte                                        ObErrorCode = -5741
	ObErrNeedUnionAllInRecursiveCte                                               ObErrorCode = -5742
	ObErrNeedOnlyTwoBranchInRecursiveCte                                          ObErrorCode = -5743
	ObErrNeedReferenceItselfDirectlyInRecursiveCte                                ObErrorCode = -5744
	ObErrNeedInitBranchInRecursiveCte                                             ObErrorCode = -5745
	ObErrCycleFoundInRecursiveCte                                                 ObErrorCode = -5746
	ObErrCteReachMaxLevelRecursion                                                ObErrorCode = -5747
	ObErrCteIllegalSearchPseudoName                                               ObErrorCode = -5748
	ObErrCteIllegalCycleNonCycleValue                                             ObErrorCode = -5749
	ObErrCteIllegalCyclePseudoName                                                ObErrorCode = -5750
	ObErrCteColumnAliasDuplicate                                                  ObErrorCode = -5751
	ObErrCteIllegalSearchCycleClause                                              ObErrorCode = -5752
	ObErrCteDuplicateCycleNonCycleValue                                           ObErrorCode = -5753
	ObErrCteDuplicateSeqNameCycleColumn                                           ObErrorCode = -5754
	ObErrCteDuplicateNameInSearchClause                                           ObErrorCode = -5755
	ObErrCteDuplicateNameInCycleClause                                            ObErrorCode = -5756
	ObErrCteIllegalColumnInCycleClause                                            ObErrorCode = -5757
	ObErrCteIllegalRecursiveBranch                                                ObErrorCode = -5758
	ObErrIllegalJoinInRecursiveCte                                                ObErrorCode = -5759
	ObErrCteNeedColumnAliasList                                                   ObErrorCode = -5760
	ObErrCteIllegalColumnInSerachCaluse                                           ObErrorCode = -5761
	ObErrCteRecursiveQueryNameReferencedMoreThanOnce                              ObErrorCode = -5762
	ObErrCbyPseudoColumnNotAllowed                                                ObErrorCode = -5763
	ObErrCbyLoop                                                                  ObErrorCode = -5764
	ObErrCbyJoinNotAllowed                                                        ObErrorCode = -5765
	ObErrCbyConnectByRequired                                                     ObErrorCode = -5766
	ObErrCbyConnectByPathNotAllowed                                               ObErrorCode = -5768
	ObErrCbyConnectByPathIllegalParam                                             ObErrorCode = -5769
	ObErrCbyConnectByPathInvalidSeparator                                         ObErrorCode = -5770
	ObErrCbyConnectByRootIllegalUsed                                              ObErrorCode = -5771
	ObErrCbyOrederSiblingsByNotAllowed                                            ObErrorCode = -5772
	ObErrCbyNocycleRequired                                                       ObErrorCode = -5773
	ObErrNotEnoughArgsForFun                                                      ObErrorCode = -5774
	ObErrPrepareStmtChecksum                                                      ObErrorCode = -5777
	ObErrEnableNonexistentConstraint                                              ObErrorCode = -5778
	ObErrDisableNonexistentConstraint                                             ObErrorCode = -5779
	ObErrDowngradeDop                                                             ObErrorCode = -5780
	ObErrDowngradeParallelMaxServers                                              ObErrorCode = -5781
	ObErrOrphanedChildRecordExists                                                ObErrorCode = -5785
	ObErrColCheckCstReferAnotherCol                                               ObErrorCode = -5786
	ObBatchedMultiStmtRollback                                                    ObErrorCode = -5787
	ObErrForUpdateSelectViewCannot                                                ObErrorCode = -5788
	ObErrPolicyWithCheckOptionViolation                                           ObErrorCode = -5789
	ObErrPolicyAlreadyAppliedToTable                                              ObErrorCode = -5790
	ObErrMutatingTableOperation                                                   ObErrorCode = -5791
	ObErrModifyOrDropMultiColumnConstraint                                        ObErrorCode = -5792
	ObErrDropParentKeyColumn                                                      ObErrorCode = -5793
	ObAutoincServiceBusy                                                          ObErrorCode = -5794
	ObErrConstraintConstraintDisableValidate                                      ObErrorCode = -5795
	ObErrAutonomousTransactionRollback                                            ObErrorCode = -5796
	ObOrderbyClauseNotAllowed                                                     ObErrorCode = -5797
	ObDistinctNotAllowed                                                          ObErrorCode = -5798
	ObErrAssignUserVariableNotAllowed                                             ObErrorCode = -5799
	ObErrModifyNonexistentConstraint                                              ObErrorCode = -5800
	ObErrSpExceptionHandleIllegal                                                 ObErrorCode = -5801
	ObErrInvalidInsertColumn                                                      ObErrorCode = -5803
	ObIncorrectUseOfOperator                                                      ObErrorCode = -5804
	ObErrNonConstExprIsNotAllowedForPivotUnpivotValues                            ObErrorCode = -5805
	ObErrExpectAggregateFunctionInsidePivotOperation                              ObErrorCode = -5806
	ObErrExpNeedSameDatatype                                                      ObErrorCode = -5807
	ObErrCharacterSetMismatch                                                     ObErrorCode = -5808
	ObErrRegexpNomatch                                                            ObErrorCode = -5809
	ObErrRegexpBadpat                                                             ObErrorCode = -5810
	ObErrRegexpEescape                                                            ObErrorCode = -5811
	ObErrRegexpEbrack                                                             ObErrorCode = -5812
	ObErrRegexpEparen                                                             ObErrorCode = -5813
	ObErrRegexpEsubreg                                                            ObErrorCode = -5814
	ObErrRegexpErange                                                             ObErrorCode = -5815
	ObErrRegexpEctype                                                             ObErrorCode = -5816
	ObErrRegexpEcollate                                                           ObErrorCode = -5817
	ObErrRegexpEbrace                                                             ObErrorCode = -5818
	ObErrRegexpBadbr                                                              ObErrorCode = -5819
	ObErrRegexpBadrpt                                                             ObErrorCode = -5820
	ObErrRegexpAssert                                                             ObErrorCode = -5821
	ObErrRegexpInvarg                                                             ObErrorCode = -5822
	ObErrRegexpMixed                                                              ObErrorCode = -5823
	ObErrRegexpBadopt                                                             ObErrorCode = -5824
	ObErrRegexpEtoobig                                                            ObErrorCode = -5825
	ObNotSupportedRowIdType                                                       ObErrorCode = -5826
	ObErrParallelDdlConflict                                                      ObErrorCode = -5827
	ObErrSubscriptBeyondCount                                                     ObErrorCode = -5828
	ObErrNotPartitioned                                                           ObErrorCode = -5829
	ObUnknownSubpartition                                                         ObErrorCode = -5830
	ObErrInvalidSqlRowLimiting                                                    ObErrorCode = -5831
	INCORRECT_ARGUMENTS_TO_ESCAPE                                                 ObErrorCode = -5832
	STATIC_ENG_NOT_IMPLEMENT                                                      ObErrorCode = -5833
	ObObjAlreadyExist                                                             ObErrorCode = -5834
	ObDblinkNotExistToAccess                                                      ObErrorCode = -5835
	ObDblinkNotExistToDrop                                                        ObErrorCode = -5836
	ObErrAccessIntoNull                                                           ObErrorCode = -5837
	ObErrCollecionNull                                                            ObErrorCode = -5838
	ObErrNoDataNeeded                                                             ObErrorCode = -5839
	ObErrProgramError                                                             ObErrorCode = -5840
	ObErrRowtypeMismatch                                                          ObErrorCode = -5841
	ObErrStorageError                                                             ObErrorCode = -5842
	ObErrSubscriptOutsideLimit                                                    ObErrorCode = -5843
	ObErrInvalidCursor                                                            ObErrorCode = -5844
	ObErrLoginDenied                                                              ObErrorCode = -5845
	ObErrNotLoggedOn                                                              ObErrorCode = -5846
	ObErrSelfIsNull                                                               ObErrorCode = -5847
	ObErrTimeoutOnResource                                                        ObErrorCode = -5848
	ObColumnCantChangeToNotNull                                                   ObErrorCode = -5849
	ObColumnCantChangeToNullale                                                   ObErrorCode = -5850
	ObEnableNotNullConstraintViolated                                             ObErrorCode = -5851
	ObErrArgumentShouldConstant                                                   ObErrorCode = -5852
	ObErrNotASingleGroupFunction                                                  ObErrorCode = -5853
	ObErrZeroLengthIdentifier                                                     ObErrorCode = -5854
	ObErrParamValueInvalid                                                        ObErrorCode = -5855
	ObErrDbmsSqlCursorNotExist                                                    ObErrorCode = -5856
	ObErrDbmsSqlNotAllVarBind                                                     ObErrorCode = -5857
	ObErrConflictingDeclarations                                                  ObErrorCode = -5858
	ObErrDropColReferencedMultiColsConstraint                                     ObErrorCode = -5859
	ObErrModifyColDatatyepReferencedConstraint                                    ObErrorCode = -5860
	ObErrPercentileValueInvalid                                                   ObErrorCode = -5861
	ObErrArgumentShouldNumericDateDatetimeType                                    ObErrorCode = -5862
	ObErrAlterTableRenameWithOption                                               ObErrorCode = -5863
	ObErrOnlySimpleColumnNameAllowed                                              ObErrorCode = -5864
	ObErrSafeUpdateModeNeedWhereOrLimit                                           ObErrorCode = -5865
	ObErrSpecifiyPartitionDescription                                             ObErrorCode = -5866
	ObErrSameNameSubpartition                                                     ObErrorCode = -5867
	ObErrUpdateOrderBy                                                            ObErrorCode = -5868
	ObErrUpdateLimit                                                              ObErrorCode = -5869
	ObRowIdTypeMismatch                                                           ObErrorCode = -5870
	ObRowIdNumMismatch                                                            ObErrorCode = -5871
	ObNoColumnAlias                                                               ObErrorCode = -5872
	ObErrInvalidDatatype                                                          ObErrorCode = -5874
	ObErrNotCompositePartition                                                    ObErrorCode = -5875
	ObErrSubpartitionNotExpectValuesIn                                            ObErrorCode = -5876
	ObErrSubpartitionExpectValuesIn                                               ObErrorCode = -5877
	ObErrPartitionNotExpectValuesLessThan                                         ObErrorCode = -5878
	ObErrPartitionExpectValuesLessThan                                            ObErrorCode = -5879
	ObErrProgramUnitNotExist                                                      ObErrorCode = -5880
	ObErrInvalidRestorePointName                                                  ObErrorCode = -5881
	ObErrInputTimeType                                                            ObErrorCode = -5882
	ObErrInArrayDml                                                               ObErrorCode = -5883
	ObErrTriggerCompileError                                                      ObErrorCode = -5884
	ObErrInTrimSet                                                                ObErrorCode = -5885
	ObErrMissingOrInvalidPasswordForRole                                          ObErrorCode = -5886
	ObErrMissingOrInvalidPassword                                                 ObErrorCode = -5887
	ObErrNoOptionsForAlterUser                                                    ObErrorCode = -5888
	ObErrNoMatchingUkPkForColList                                                 ObErrorCode = -5889
	ObErrDupFkInTable                                                             ObErrorCode = -5890
	ObErrDupFkExists                                                              ObErrorCode = -5891
	ObErrMissingOrInvalidPriviege                                                 ObErrorCode = -5892
	ObErrInvalidVirtualColumnType                                                 ObErrorCode = -5893
	ObErrReferencedTableHasNoPk                                                   ObErrorCode = -5894
	ObErrModifyPartColumnType                                                     ObErrorCode = -5895
	ObErrModifySubpartColumnType                                                  ObErrorCode = -5896
	ObErrDecreaseColumnLength                                                     ObErrorCode = -5897
	ObErrRemotePartIllegal                                                        ObErrorCode = -5899
	ObErrDuplicateColumnExpressionWasSpecified                                    ObErrorCode = -5900
	ObErrAViewNotAppropriateHere                                                  ObErrorCode = -5901
	ObRowIdViewNoKeyPreserved                                                     ObErrorCode = -5902
	ObRowIdViewHasDistinctEtc                                                     ObErrorCode = -5903
	ObErrAtLeastOneColumnNotVirtual                                               ObErrorCode = -5904
	ObErrOnlyPureFuncCanbeIndexed                                                 ObErrorCode = -5905
	ObErrOnlyPureFuncCanbeVirtualColumnExpression                                 ObErrorCode = -5906
	ObErrUpdateOperationOnVirtualColumns                                          ObErrorCode = -5907
	ObErrInvalidColumnExpression                                                  ObErrorCode = -5908
	ObErrIdentityColumnCountExceLimit                                             ObErrorCode = -5909
	ObErrInvalidNotNullConstraintOnIdentityColumn                                 ObErrorCode = -5910
	ObErrCannotModifyNotNullConstraintOnIdentityColumn                            ObErrorCode = -5911
	ObErrCannotDropNotNullConstraintOnIdentityColumn                              ObErrorCode = -5912
	ObErrColumnModifyToIdentityColumn                                             ObErrorCode = -5913
	ObErrIdentityColumnCannotHaveDefaultValue                                     ObErrorCode = -5914
	ObErrIdentityColumnMustBeNumericType                                          ObErrorCode = -5915
	ObErrPrebuiltTableManagedCannotBeIdentityColumn                               ObErrorCode = -5916
	ObErrCannotAlterSystemGeneratedSequence                                       ObErrorCode = -5917
	ObErrCannotDropSystemGeneratedSequence                                        ObErrorCode = -5918
	ObErrInsertIntoGeneratedAlwaysIdentityColumn                                  ObErrorCode = -5919
	ObErrUpdateGeneratedAlwaysIdentityColumn                                      ObErrorCode = -5920
	ObErrIdentityColumnSequenceMismatchAlterTableExchangePartition                ObErrorCode = -5921
	ObErrCannotRenameSystemGeneratedSequence                                      ObErrorCode = -5922
	ObErrRevokeByColumn                                                           ObErrorCode = -5923
	ObErrTypeBodyNotExist                                                         ObErrorCode = -5924
	ObErrInvalidArgumentForWidthBucket                                            ObErrorCode = -5925
	ObErrCbyNoMemory                                                              ObErrorCode = -5926
	ObErrIllegalParamForCbyPath                                                   ObErrorCode = -5927
	ObErrHostUnknown                                                              ObErrorCode = -5928
	ObErrWindowNameIsNotDefine                                                    ObErrorCode = -5929
	ObErrOpenCursorsExceeded                                                      ObErrorCode = -5930
	ObErrFetchOutSequence                                                         ObErrorCode = -5931
	ObErrUnexpectedNameStr                                                        ObErrorCode = -5932
	ObErrNoProgramUnit                                                            ObErrorCode = -5933
	ObErrArgInvalid                                                               ObErrorCode = -5934
	ObErrDbmsStatsPl                                                              ObErrorCode = -5935
	ObErrIncorrectValueForFunction                                                ObErrorCode = -5936
	ObErrUnsupportedCharacterSet                                                  ObErrorCode = -5937
	ObErrMustBeFollowedByFourHexadecimalCharactersOrAnother                       ObErrorCode = -5938
	ObErrParameterTooLong                                                         ObErrorCode = -5939
	ObErrInvalidPlsqlCcflags                                                      ObErrorCode = -5940
	ObErrRefMutuallyDep                                                           ObErrorCode = -5941
	ObErrColumnNotAllowed                                                         ObErrorCode = -5942
	ObErrCannotAccessNlsDataFilesOrInvalidEnvironmentSpecified                    ObErrorCode = -5943
	ObErrDuplicateNullSpecification                                               ObErrorCode = -5944
	ObErrNotNullConstraintViolated                                                ObErrorCode = -5945
	ObErrTableAddNotNullColumnNotEmpty                                            ObErrorCode = -5946
	ObErrColumnExpressionModificationWithOtherDdl                                 ObErrorCode = -5947
	ObErrVirtualColWithConstraintCantBeChanged                                    ObErrorCode = -5948
	ObErrInvalidNotNullConstraintOnDefaultOnNullIdentityColumn                    ObErrorCode = -5949
	ObErrInvalidDataTypeForAtTimeZone                                             ObErrorCode = -5950
	ObErrBadArg                                                                   ObErrorCode = -5951
	ObErrCannotModifyNotNullConstraintOnDefaultOnNullColumn                       ObErrorCode = -5952
	ObErrCannotDropNotNullConstraintOnDefaultOnNullColumn                         ObErrorCode = -5953
	ObErrInvalidPath                                                              ObErrorCode = -5954
	ObErrInvalidParamEncountered                                                  ObErrorCode = -5955
	ObErrIncorrectMethodUsage                                                     ObErrorCode = -5956
	ObErrTypeMismatch                                                             ObErrorCode = -5957
	ObErrFetchColumnNull                                                          ObErrorCode = -5958
	ObErrInvalidSizeSpecified                                                     ObErrorCode = -5959
	ObErrSourceEmpty                                                              ObErrorCode = -5960
	ObErrBadValueForObjectType                                                    ObErrorCode = -5961
	ObErrUnableGetSource                                                          ObErrorCode = -5962
	ObErrMissingIdentifier                                                        ObErrorCode = -5963
	ObErrDupCompileParam                                                          ObErrorCode = -5964
	ObErrDataNotWellFormat                                                        ObErrorCode = -5965
	ObErrMustCompositType                                                         ObErrorCode = -5966
	ObErrUserExceedResource                                                       ObErrorCode = -5967
	ObErrUtlEncodeArgumentInvalid                                                 ObErrorCode = -5968
	ObErrUtlEncodeCharsetInvalid                                                  ObErrorCode = -5969
	ObErrUtlEncodeMimeHeadTag                                                     ObErrorCode = -5970
	ObErrCheckOptionViolated                                                      ObErrorCode = -5971
	ObErrCheckOptionOnNonupdatableView                                            ObErrorCode = -5972
	ObErrNoDescForPos                                                             ObErrorCode = -5973
	ObErrIllObjFlag                                                               ObErrorCode = -5974
	ObErrPartitionExtendedOnView                                                  ObErrorCode = -5977
	ObErrNotAllVariableBind                                                       ObErrorCode = -5978
	ObErrBindVariableNotExist                                                     ObErrorCode = -5979
	ObErrNotValidRoutineName                                                      ObErrorCode = -5980
	ObErrDdlInIllegalContext                                                      ObErrorCode = -5981
	ObErrCteNeedQueryBlocks                                                       ObErrorCode = -5982
	ObErrWindowRowsIntervalUse                                                    ObErrorCode = -5983
	ObErrWindowRangeFrameOrderType                                                ObErrorCode = -5984
	ObErrWindowIllegalOrderBy                                                     ObErrorCode = -5985
	ObErrMultipleConstraintsWithSameName                                          ObErrorCode = -5986
	ObErrNonBooleanExprForCheckConstraint                                         ObErrorCode = -5987
	ObErrCheckConstraintNotFound                                                  ObErrorCode = -5988
	ObErrAlterConstraintEnforcementNotSupported                                   ObErrorCode = -5989
	ObErrCheckConstraintRefersAutoIncrementColumn                                 ObErrorCode = -5990
	ObErrCheckConstraintNamedFunctionIsNotAllowed                                 ObErrorCode = -5991
	ObErrCheckConstraintFunctionIsNotAllowed                                      ObErrorCode = -5992
	ObErrCheckConstraintVariables                                                 ObErrorCode = -5993
	ObErrCheckConstraintRefersUnknownColumn                                       ObErrorCode = -5994
	ObErrUseUdfInPart                                                             ObErrorCode = -5995
	ObErrUseUdfNotDetermin                                                        ObErrorCode = -5996
	ObErrIntervalClauseHasMoreThanOneColumn                                       ObErrorCode = -5997
	ObErrInvalidDataTypeIntervalTable                                             ObErrorCode = -5998
	ObErrIntervalExprNotCorrectType                                               ObErrorCode = -5999
	ObErrTableIsAlreadyARangePartitionedTable                                     ObErrorCode = -6000
	ObTransactionSetViolation                                                     ObErrorCode = -6001
	ObTransRollbacked                                                             ObErrorCode = -6002
	ObErrExclusiveLockConflict                                                    ObErrorCode = -6003
	ObErrSharedLockConflict                                                       ObErrorCode = -6004
	ObTryLockRowConflict                                                          ObErrorCode = -6005
	ObErrExclusiveLockConflictNowait                                              ObErrorCode = -6006
	ObClockOutOfOrder                                                             ObErrorCode = -6201
	ObTransHasDecided                                                             ObErrorCode = -6204
	ObTransInvalidState                                                           ObErrorCode = -6205
	ObTransStateNotChange                                                         ObErrorCode = -6206
	ObTransProtocolError                                                          ObErrorCode = -6207
	ObTransInvalidMessage                                                         ObErrorCode = -6208
	ObTransInvalidMessageType                                                     ObErrorCode = -6209
	ObPartitionIsFrozen                                                           ObErrorCode = -6214
	ObPartitionIsNotFrozen                                                        ObErrorCode = -6215
	ObTransInvalidLogType                                                         ObErrorCode = -6219
	ObTransSqlSequenceIllegal                                                     ObErrorCode = -6220
	ObTransCannotBeKilled                                                         ObErrorCode = -6221
	ObTransStateUnknown                                                           ObErrorCode = -6222
	ObTransIsExiting                                                              ObErrorCode = -6223
	ObTransNeedRollback                                                           ObErrorCode = -6224
	ObPartitionIsNotStopped                                                       ObErrorCode = -6227
	ObPartitionIsStopped                                                          ObErrorCode = -6228
	ObPartitionIsBlocked                                                          ObErrorCode = -6229
	ObTransRpcTimeout                                                             ObErrorCode = -6230
	ObReplicaNotReadable                                                          ObErrorCode = -6231
	ObPartitionIsSplitting                                                        ObErrorCode = -6232
	ObTransCommited                                                               ObErrorCode = -6233
	ObTransCtxCountReachLimit                                                     ObErrorCode = -6234
	ObTransCannotSerialize                                                        ObErrorCode = -6235
	ObTransWeakReadVersionNotReady                                                ObErrorCode = -6236
	ObGtsStandbyIsInvalid                                                         ObErrorCode = -6237
	ObGtsUpdateFailed                                                             ObErrorCode = -6238
	ObGtsIsNotServing                                                             ObErrorCode = -6239
	ObPgPartitionNotExist                                                         ObErrorCode = -6240
	ObTransStmtNeedRetry                                                          ObErrorCode = -6241
	ObSavepointNotExist                                                           ObErrorCode = -6242
	ObTransWaitSchemaRefresh                                                      ObErrorCode = -6243
	ObTransOutOfThreshold                                                         ObErrorCode = -6244
	ObTransXaNota                                                                 ObErrorCode = -6245
	ObTransXaRmfail                                                               ObErrorCode = -6246
	ObTransXaDupid                                                                ObErrorCode = -6247
	ObTransXaOutside                                                              ObErrorCode = -6248
	ObTransXaInval                                                                ObErrorCode = -6249
	ObTransXaRmerr                                                                ObErrorCode = -6250
	ObTransXaProto                                                                ObErrorCode = -6251
	ObTransXaRbrollback                                                           ObErrorCode = -6252
	ObTransXaRbtimeout                                                            ObErrorCode = -6253
	ObTransXaRdonly                                                               ObErrorCode = -6254
	ObTransXaRetry                                                                ObErrorCode = -6255
	ObErrRowNotLocked                                                             ObErrorCode = -6256
	ObEmptyPg                                                                     ObErrorCode = -6257
	ObTransXaErrCommit                                                            ObErrorCode = -6258
	ObErrRestorePointExist                                                        ObErrorCode = -6259
	ObErrRestorePointNotExist                                                     ObErrorCode = -6260
	ObErrBackupPointExist                                                         ObErrorCode = -6261
	ObErrBackupPointNotExist                                                      ObErrorCode = -6262
	ObErrRestorePointTooMany                                                      ObErrorCode = -6263
	ObTransXaBranchFail                                                           ObErrorCode = -6264
	ObObjLockNotExist                                                             ObErrorCode = -6265
	ObObjLockExist                                                                ObErrorCode = -6266
	ObTryLockObjConflict                                                          ObErrorCode = -6267
	ObTxNologcb                                                                   ObErrorCode = -6268
	ObErrAddPartitionOnInterval                                                   ObErrorCode = -6269
	ObErrMaxvaluePartitionWithInterval                                            ObErrorCode = -6270
	ObErrInvalidIntervalHighBounds                                                ObErrorCode = -6271
	ObNoPartitionForIntervalPart                                                  ObErrorCode = -6272
	ObErrIntervalCannotBeZero                                                     ObErrorCode = -6273
	ObErrPartitioningKeyMapsToAPartitionOutsideMaximumPermittedNumberOfPartitions ObErrorCode = -6274
	ObObjLockNotCompleted                                                         ObErrorCode = -6275
	ObObjUnlockConflict                                                           ObErrorCode = -6276
	ObScnOutOfBound                                                               ObErrorCode = -6277
	ObTransIdleTimeout                                                            ObErrorCode = -6278
	ObTransFreeRouteNotSupported                                                  ObErrorCode = -6279
	ObTransLiveTooMuchTime                                                        ObErrorCode = -6280
	ObTransCommitTooMuchTime                                                      ObErrorCode = -6281
	ObTransTooManyParticipants                                                    ObErrorCode = -6282
	ObLogAlreadySplit                                                             ObErrorCode = -6283
	ObLogIdNotFound                                                               ObErrorCode = -6301
	ObLsrThreadStopped                                                            ObErrorCode = -6302
	ObNoLog                                                                       ObErrorCode = -6303
	ObLogIdRangeError                                                             ObErrorCode = -6304
	ObLogIterEnough                                                               ObErrorCode = -6305
	ObClogInvalidAck                                                              ObErrorCode = -6306
	ObClogCacheInvalid                                                            ObErrorCode = -6307
	ObExtHandleUnfinish                                                           ObErrorCode = -6308
	ObCursorNotExist                                                              ObErrorCode = -6309
	ObStreamNotExist                                                              ObErrorCode = -6310
	ObStreamBusy                                                                  ObErrorCode = -6311
	ObFileRecycled                                                                ObErrorCode = -6312
	ObReplayEagainTooMuchTime                                                     ObErrorCode = -6313
	ObMemberChangeFailed                                                          ObErrorCode = -6314
	ObNoNeedBatchCtx                                                              ObErrorCode = -6315
	ObTooLargeLogId                                                               ObErrorCode = -6316
	ObAllocLogIdNeedRetry                                                         ObErrorCode = -6317
	ObTransOnePcNotAllowed                                                        ObErrorCode = -6318
	ObLogNeedRebuild                                                              ObErrorCode = -6319
	ObTooManyLogTask                                                              ObErrorCode = -6320
	ObInvalidBatchSize                                                            ObErrorCode = -6321
	ObClogSlideTimeout                                                            ObErrorCode = -6322
	ObLogReplayError                                                              ObErrorCode = -6323
	ObTryLockConfigChangeConflict                                                 ObErrorCode = -6324
	ObElectionWarnLogbufFull                                                      ObErrorCode = -7000
	ObElectionWarnLogbufEmpty                                                     ObErrorCode = -7001
	ObElectionWarnNotRunning                                                      ObErrorCode = -7002
	ObElectionWarnIsRunning                                                       ObErrorCode = -7003
	ObElectionWarnNotReachMajority                                                ObErrorCode = -7004
	ObElectionWarnInvalidServer                                                   ObErrorCode = -7005
	ObElectionWarnInvalidLeader                                                   ObErrorCode = -7006
	ObElectionWarnLeaderLeaseExpired                                              ObErrorCode = -7007
	ObElectionWarnInvalidMessage                                                  ObErrorCode = -7010
	ObElectionWarnMessageNotIntime                                                ObErrorCode = -7011
	ObElectionWarnNotCandidate                                                    ObErrorCode = -7012
	ObElectionWarnNotCandidateOrVoter                                             ObErrorCode = -7013
	ObElectionWarnProtocolError                                                   ObErrorCode = -7014
	ObElectionWarnRuntimeOutOfRange                                               ObErrorCode = -7015
	ObElectionWarnLastOperationNotDone                                            ObErrorCode = -7021
	ObElectionWarnCurrentServerNotLeader                                          ObErrorCode = -7022
	ObElectionWarnNoPrepareMessage                                                ObErrorCode = -7024
	ObElectionErrorMultiPrepareMessage                                            ObErrorCode = -7025
	ObElectionNotExist                                                            ObErrorCode = -7026
	ObElectionMgrIsRunning                                                        ObErrorCode = -7027
	ObElectionWarnNoMajorityPrepareMessage                                        ObErrorCode = -7029
	ObElectionAsyncLogWarnInit                                                    ObErrorCode = -7030
	ObElectionWaitLeaderMessage                                                   ObErrorCode = -7031
	ObElectionGroupNotExist                                                       ObErrorCode = -7032
	ObUnexpectEgVersion                                                           ObErrorCode = -7033
	ObElectionGroupMgrIsRunning                                                   ObErrorCode = -7034
	ObElectionMgrNotRunning                                                       ObErrorCode = -7035
	ObElectionErrorVoteMsgConflict                                                ObErrorCode = -7036
	ObElectionErrorDuplicatedMsg                                                  ObErrorCode = -7037
	ObElectionWarnT1NotMatch                                                      ObErrorCode = -7038
	ObElectionBelowMajority                                                       ObErrorCode = -7039
	ObElectionOverMajority                                                        ObErrorCode = -7040
	ObElectionDuringUpgrading                                                     ObErrorCode = -7041
	ObTransferTaskCompleted                                                       ObErrorCode = -7100
	ObTooManyTransferTask                                                         ObErrorCode = -7101
	ObTransferTaskExist                                                           ObErrorCode = -7102
	ObTransferTaskNotExist                                                        ObErrorCode = -7103
	ObNotAllowToRemove                                                            ObErrorCode = -7104
	ObRgNotMatch                                                                  ObErrorCode = -7105
	ObTransferTaskAborted                                                         ObErrorCode = -7106
	ObTransferInvalidMessage                                                      ObErrorCode = -7107
	ObTransferCtxTsNotMatch                                                       ObErrorCode = -7108
	ObTransferSysError                                                            ObErrorCode = -7109
	ObTransferMemberListNotSame                                                   ObErrorCode = -7110
	ObErrUnexpectedLockOwner                                                      ObErrorCode = -7111
	ObLsTransferScnTooSmall                                                       ObErrorCode = -7112
	ObTabletTransferSeqNotMatch                                                   ObErrorCode = -7113
	ObTransferDetectActiveTrans                                                   ObErrorCode = -7114
	ObTransferSrcLsNotExist                                                       ObErrorCode = -7115
	ObTransferSrcTabletNotExist                                                   ObErrorCode = -7116
	ObLsNeedRebuild                                                               ObErrorCode = -7117
	ObObsoleteClogNeedSkip                                                        ObErrorCode = -7118
	ObTransferWaitTransactionEndTimeout                                           ObErrorCode = -7119
	ObTabletGcLockConflict                                                        ObErrorCode = -7120
	ObSequenceNotMatch                                                            ObErrorCode = -7121
	ObSequenceTooSmall                                                            ObErrorCode = -7122
	ObErrInvalidXmlDatatype                                                       ObErrorCode = -7402
	ObErrXmlMissingComma                                                          ObErrorCode = -7403
	ObErrInvalidXpathExpression                                                   ObErrorCode = -7404
	ObErrExtractvalueMultiNodes                                                   ObErrorCode = -7405
	ObErrXmlFramentConvert                                                        ObErrorCode = -7406
	ObInvalidPrintOption                                                          ObErrorCode = -7407
	ObXmlCharLenTooSmall                                                          ObErrorCode = -7408
	ObXpathExpressionUnsupported                                                  ObErrorCode = -7409
	ObExtractvalueNotLeafNode                                                     ObErrorCode = -7410
	ObXmlInsertFragment                                                           ObErrorCode = -7411
	ObErrNoOrderMapSql                                                            ObErrorCode = -7412
	ObErrXmlelementAliased                                                        ObErrorCode = -7413
	ObInvalidAlterationgDatatype                                                  ObErrorCode = -7414
	ObInvalidModificationOfColumns                                                ObErrorCode = -7415
	ObErrNullForXmlConstructor                                                    ObErrorCode = -7416
	ObErrXmlIndex                                                                 ObErrorCode = -7417
	ObErrUpdateXmlWithInvalidNode                                                 ObErrorCode = -7418
	ObLobValueNotExist                                                            ObErrorCode = -7419
	ObErrJsonFunUnsupportedType                                                   ObErrorCode = -7420
	ObServerIsInit                                                                ObErrorCode = -8001
	ObServerIsStopping                                                            ObErrorCode = -8002
	ObPacketChecksumError                                                         ObErrorCode = -8003
	ObNotReadAllData                                                              ObErrorCode = -9008
	ObBuildMd5Error                                                               ObErrorCode = -9009
	ObMd5NotMatch                                                                 ObErrorCode = -9010
	ObOssDataVersionNotMatched                                                    ObErrorCode = -9012
	ObOssWriteError                                                               ObErrorCode = -9013
	ObRestoreInProgress                                                           ObErrorCode = -9014
	ObAgentInitingBackupCountError                                                ObErrorCode = -9015
	ObClusterNameNotEqual                                                         ObErrorCode = -9016
	ObRsListInvaild                                                               ObErrorCode = -9017
	ObAgentHasFailedTask                                                          ObErrorCode = -9018
	ObRestorePartitionIsComplete                                                  ObErrorCode = -9019
	ObRestorePartitionTwice                                                       ObErrorCode = -9020
	ObStopDropSchema                                                              ObErrorCode = -9022
	ObCannotStartLogArchiveBackup                                                 ObErrorCode = -9023
	ObAlreadyNoLogArchiveBackup                                                   ObErrorCode = -9024
	ObLogArchiveBackupInfoNotExist                                                ObErrorCode = -9025
	ObLogArchiveInterrupted                                                       ObErrorCode = -9027
	ObLogArchiveStatNotMatch                                                      ObErrorCode = -9028
	ObLogArchiveNotRunning                                                        ObErrorCode = -9029
	ObLogArchiveInvalidRound                                                      ObErrorCode = -9030
	ObReplicaCannotBackup                                                         ObErrorCode = -9031
	ObBackupInfoNotExist                                                          ObErrorCode = -9032
	ObBackupInfoNotMatch                                                          ObErrorCode = -9033
	ObLogArchiveAlreadyStopped                                                    ObErrorCode = -9034
	ObRestoreIndexFailed                                                          ObErrorCode = -9035
	ObBackupInProgress                                                            ObErrorCode = -9036
	ObInvalidLogArchiveStatus                                                     ObErrorCode = -9037
	ObCannotAddReplicaDuringSetMemberList                                         ObErrorCode = -9038
	ObLogArchiveLeaderChanged                                                     ObErrorCode = -9039
	ObBackupCanNotStart                                                           ObErrorCode = -9040
	ObCancelBackupNotAllowed                                                      ObErrorCode = -9041
	ObBackupDataVersionGapOverLimit                                               ObErrorCode = -9042
	ObPgLogArchiveStatusNotInit                                                   ObErrorCode = -9043
	ObBackupDeleteDataInProgress                                                  ObErrorCode = -9044
	ObBackupDeleteBackupSetNotAllowed                                             ObErrorCode = -9045
	ObInvalidBackupSetId                                                          ObErrorCode = -9046
	ObBackupInvalidPassword                                                       ObErrorCode = -9047
	ObIsolatedBackupSet                                                           ObErrorCode = -9048
	ObCannotCancelStoppedBackup                                                   ObErrorCode = -9049
	ObBackupBackupCanNotStart                                                     ObErrorCode = -9050
	ObBackupMountFileNotValid                                                     ObErrorCode = -9051
	ObBackupCleanInfoNotMatch                                                     ObErrorCode = -9052
	ObCancelDeleteBackupNotAllowed                                                ObErrorCode = -9053
	ObBackupCleanInfoNotExist                                                     ObErrorCode = -9054
	ObCannotSetBackupRegion                                                       ObErrorCode = -9057
	ObCannotSetBackupZone                                                         ObErrorCode = -9058
	ObBackupBackupReachMaxBackupTimes                                             ObErrorCode = -9059
	ObArchiveLogNotContinuesWithData                                              ObErrorCode = -9064
	ObAgentHasSuspended                                                           ObErrorCode = -9065
	ObBackupConflictValue                                                         ObErrorCode = -9066
	ObBackupDeleteBackupPieceNotAllowed                                           ObErrorCode = -9069
	ObBackupDestNotConnect                                                        ObErrorCode = -9070
	ObEsiSessionConflicts                                                         ObErrorCode = -9072
	ObBackupValidateTaskSkipped                                                   ObErrorCode = -9074
	ObEsiIoError                                                                  ObErrorCode = -9075
	ObArchiveRoundNotContinuous                                                   ObErrorCode = -9077
	ObArchiveLogToEnd                                                             ObErrorCode = -9078
	ObArchiveLogRecycled                                                          ObErrorCode = -9079
	ObBackupFormatFileNotExist                                                    ObErrorCode = -9080
	ObBackupFormatFileNotMatch                                                    ObErrorCode = -9081
	ObBackupMajorNotCoverMinor                                                    ObErrorCode = -9085
	ObBackupAdvanceCheckpointTimeout                                              ObErrorCode = -9086
	ObClogRecycleBeforeArchive                                                    ObErrorCode = -9087
	ObSourceTenantStateNotMatch                                                   ObErrorCode = -9088
	ObSourceLsStateNotMatch                                                       ObErrorCode = -9089
	ObEsiSessionNotExist                                                          ObErrorCode = -9090
	ObAlreadyInArchiveMode                                                        ObErrorCode = -9091
	ObAlreadyInNoarchiveMode                                                      ObErrorCode = -9092
	ObRestoreLogToEnd                                                             ObErrorCode = -9093
	ObLsRestoreFailed                                                             ObErrorCode = -9094
	ObNoTabletNeedBackup                                                          ObErrorCode = -9095
	ObErrRestoreStandbyVersionLag                                                 ObErrorCode = -9096
	ObErrRestorePrimaryTenantDropped                                              ObErrorCode = -9097
	ObNoSuchFileOrDirectory                                                       ObErrorCode = -9100
	ObFileOrDirectoryExist                                                        ObErrorCode = -9101
	ObFileOrDirectoryPermissionDenied                                             ObErrorCode = -9102
	ObTooManyOpenFiles                                                            ObErrorCode = -9103
	ObDirectLoadCommitError                                                       ObErrorCode = -9104
	ObErrResizeFileToSmaller                                                      ObErrorCode = -9200
	ObMarkBlockInfoTimeout                                                        ObErrorCode = -9201
	ObNotReadyToExtendFile                                                        ObErrorCode = -9202
	ObErrDuplicateHavingClauseInTableExpression                                   ObErrorCode = -9501
	ObErrInoutParamPlacementNotProperly                                           ObErrorCode = -9502
	ObErrObjectNotFound                                                           ObErrorCode = -9503
	ObErrInvalidInputValue                                                        ObErrorCode = -9504
	ObErrGotoBranchIllegal                                                        ObErrorCode = -9505
	ObErrOnlySchemaLevelAllow                                                     ObErrorCode = -9506
	ObErrDeclMoreThanOnce                                                         ObErrorCode = -9507
	ObErrDuplicateFiled                                                           ObErrorCode = -9508
	ObErrPragmaIllegal                                                            ObErrorCode = -9509
	ObErrExitContinueIllegal                                                      ObErrorCode = -9510
	ObErrLabelIllegal                                                             ObErrorCode = -9512
	ObErrCursorLeftAssign                                                         ObErrorCode = -9513
	ObErrInitNotnullIllegal                                                       ObErrorCode = -9514
	ObErrInitConstIllegal                                                         ObErrorCode = -9515
	ObErrCursorVarInPkg                                                           ObErrorCode = -9516
	ObErrLimitClause                                                              ObErrorCode = -9518
	ObErrExpressionWrongType                                                      ObErrorCode = -9519
	ObErrSpecNotExist                                                             ObErrorCode = -9520
	ObErrTypeSpecNoRoutine                                                        ObErrorCode = -9521
	ObErrTypeBodyNoRoutine                                                        ObErrorCode = -9522
	ObErrBothOrderMap                                                             ObErrorCode = -9523
	ObErrNoOrderMap                                                               ObErrorCode = -9524
	ObErrOrderMapNeedBeFunc                                                       ObErrorCode = -9525
	ObErrIdentifierTooLong                                                        ObErrorCode = -9526
	ObErrInvokeStaticByInstance                                                   ObErrorCode = -9527
	ObErrConsNameIllegal                                                          ObErrorCode = -9528
	ObErrAttrFuncConflict                                                         ObErrorCode = -9529
	ObErrSelfParamNotOut                                                          ObErrorCode = -9530
	ObErrMapRetScalarType                                                         ObErrorCode = -9531
	ObErrMapMoreThanSelfParam                                                     ObErrorCode = -9532
	ObErrOrderRetIntType                                                          ObErrorCode = -9533
	ObErrOrderParamType                                                           ObErrorCode = -9534
	ObErrObjCmpSql                                                                ObErrorCode = -9535
	ObErrMapOrderPragma                                                           ObErrorCode = -9536
	ObErrOrderParamMustInMode                                                     ObErrorCode = -9537
	ObErrOrderParamNotTwo                                                         ObErrorCode = -9538
	ObErrTypeRefRefcursive                                                        ObErrorCode = -9539
	ObErrDirectiveError                                                           ObErrorCode = -9540
	ObErrConsHasRetNode                                                           ObErrorCode = -9541
	ObErrCallWrongArg                                                             ObErrorCode = -9542
	ObErrFuncNameSameWithCons                                                     ObErrorCode = -9543
	ObErrFuncDup                                                                  ObErrorCode = -9544
	ObErrWhenClause                                                               ObErrorCode = -9545
	ObErrNewOldReferences                                                         ObErrorCode = -9546
	ObErrTypeDeclIllegal                                                          ObErrorCode = -9547
	ObErrObjectInvalid                                                            ObErrorCode = -9548
	ObErrExpNotAssignable                                                         ObErrorCode = -9550
	ObErrCursorContainBothRegularAndArray                                         ObErrorCode = -9551
	ObErrStaticBoolExpr                                                           ObErrorCode = -9552
	ObErrDirectiveContext                                                         ObErrorCode = -9553
	ObUtlFileInvalidPath                                                          ObErrorCode = -9554
	ObUtlFileInvalidMode                                                          ObErrorCode = -9555
	ObUtlFileInvalidFilehandle                                                    ObErrorCode = -9556
	ObUtlFileInvalidOperation                                                     ObErrorCode = -9557
	ObUtlFileReadError                                                            ObErrorCode = -9558
	ObUtlFileWriteError                                                           ObErrorCode = -9559
	ObUtlFileInternalError                                                        ObErrorCode = -9560
	ObUtlFileCharsetmismatch                                                      ObErrorCode = -9561
	ObUtlFileInvalidMaxlinesize                                                   ObErrorCode = -9562
	ObUtlFileInvalidFilename                                                      ObErrorCode = -9563
	ObUtlFileAccessDenied                                                         ObErrorCode = -9564
	ObUtlFileInvalidOffset                                                        ObErrorCode = -9565
	ObUtlFileDeleteFailed                                                         ObErrorCode = -9566
	ObUtlFileRenameFailed                                                         ObErrorCode = -9567
	ObErrBindTypeNotMatchColumn                                                   ObErrorCode = -9568
	ObErrNestedTableInTri                                                         ObErrorCode = -9569
	ObErrColListInTri                                                             ObErrorCode = -9570
	ObErrWhenClauseInTri                                                          ObErrorCode = -9571
	ObErrInsteadTriOnTable                                                        ObErrorCode = -9572
	ObErrReturningClause                                                          ObErrorCode = -9573
	ObErrNoReturnInFunction                                                       ObErrorCode = -9575
	ObErrStmtNotAllowInMysqlFuncTrigger                                           ObErrorCode = -9576
	ObErrTooLongStringType                                                        ObErrorCode = -9577
	ObErrWidthOutOfRange                                                          ObErrorCode = -9578
	ObErrRedefineLabel                                                            ObErrorCode = -9579
	ObErrStmtNotAllowInMysqlProcedrue                                             ObErrorCode = -9580
	ObErrTriggerNotSupport                                                        ObErrorCode = -9581
	ObErrTriggerInWrongSchema                                                     ObErrorCode = -9582
	ObErrUnknownException                                                         ObErrorCode = -9583
	ObErrTriggerCantChangeRow                                                     ObErrorCode = -9584
	ObErrItemNotInBody                                                            ObErrorCode = -9585
	ObErrWrongRowtype                                                             ObErrorCode = -9586
	ObErrRoutineNotDefine                                                         ObErrorCode = -9587
	ObErrDupNameInCursor                                                          ObErrorCode = -9588
	ObErrLocalCollInSql                                                           ObErrorCode = -9589
	ObErrTypeMismatchInFetch                                                      ObErrorCode = -9590
	ObErrOthersMustLast                                                           ObErrorCode = -9591
	ObErrRaiseNotInHandler                                                        ObErrorCode = -9592
	ObErrInvalidCursorReturnType                                                  ObErrorCode = -9593
	ObErrInCursorOpend                                                            ObErrorCode = -9594
	ObErrCursorNoReturnType                                                       ObErrorCode = -9595
	ObErrNoChoices                                                                ObErrorCode = -9596
	ObErrTypeDeclMalformed                                                        ObErrorCode = -9597
	ObErrInFormalNotDenotable                                                     ObErrorCode = -9598
	ObErrOutParamHasDefault                                                       ObErrorCode = -9599
	ObErrOnlyFuncCanPipelined                                                     ObErrorCode = -9600
	ObErrPipeReturnNotColl                                                        ObErrorCode = -9601
	ObErrMismatchSubprogram                                                       ObErrorCode = -9602
	ObErrParamInPackageSpec                                                       ObErrorCode = -9603
	ObErrNumericLiteralRequired                                                   ObErrorCode = -9604
	ObErrNonIntLiteral                                                            ObErrorCode = -9605
	ObErrImproperConstraintForm                                                   ObErrorCode = -9606
	ObErrTypeCantConstrained                                                      ObErrorCode = -9607
	ObErrAnyCsNotAllowed                                                          ObErrorCode = -9608
	ObErrSchemaTypeIllegal                                                        ObErrorCode = -9609
	ObErrUnsupportedTableIndexType                                                ObErrorCode = -9610
	ObErrArrayMustHavePositiveLimit                                               ObErrorCode = -9611
	ObErrForallIterNotAllowed                                                     ObErrorCode = -9612
	ObErrBulkInBind                                                               ObErrorCode = -9613
	ObErrForallBulkTogether                                                       ObErrorCode = -9614
	ObErrForallDmlWithoutBulk                                                     ObErrorCode = -9615
	ObErrShouldCollectionType                                                     ObErrorCode = -9616
	ObErrAssocElemType                                                            ObErrorCode = -9617
	ObErrIntoClauseExpected                                                       ObErrorCode = -9618
	ObErrSubprogramViolatesPragma                                                 ObErrorCode = -9619
	ObErrExprSqlType                                                              ObErrorCode = -9620
	ObErrPragmaDeclTwice                                                          ObErrorCode = -9621
	ObErrPragmaFollowDecl                                                         ObErrorCode = -9622
	ObErrPipeStmtInNonPipelinedFunc                                               ObErrorCode = -9623
	ObErrImplRestriction                                                          ObErrorCode = -9624
	ObErrInsufficientPrivilege                                                    ObErrorCode = -9625
	ObErrIllegalOption                                                            ObErrorCode = -9626
	ObErrNoFunctionExist                                                          ObErrorCode = -9627
	ObErrOutOfScope                                                               ObErrorCode = -9628
	ObErrIllegalErrorNum                                                          ObErrorCode = -9629
	ObErrDefaultNotMatch                                                          ObErrorCode = -9630
	ObErrTableSingleIndex                                                         ObErrorCode = -9631
	ObErrPragmaDecl                                                               ObErrorCode = -9632
	ObErrIncorrectArguments                                                       ObErrorCode = -9633
	ObErrReturnValueRequired                                                      ObErrorCode = -9634
	ObErrReturnExprIllegal                                                        ObErrorCode = -9635
	ObErrLimitIllegal                                                             ObErrorCode = -9636
	ObErrIntoExprIllegal                                                          ObErrorCode = -9637
	ObErrBulkSqlRestriction                                                       ObErrorCode = -9638
	ObErrMixSingleMulti                                                           ObErrorCode = -9639
	ObErrTriggerNoSuchRow                                                         ObErrorCode = -9640
	ObErrSetUsage                                                                 ObErrorCode = -9641
	ObErrModifierConflicts                                                        ObErrorCode = -9642
	ObErrDuplicateModifier                                                        ObErrorCode = -9643
	ObErrStrLiteralTooLong                                                        ObErrorCode = -9644
	ObErrSelfParamNotInout                                                        ObErrorCode = -9645
	ObErrConstructMustReturnSelf                                                  ObErrorCode = -9646
	ObErrFirstParamMustNotNull                                                    ObErrorCode = -9647
	ObErrCoalesceAtLeastOneNotNull                                                ObErrorCode = -9648
	ObErrStaticMethodHasSelf                                                      ObErrorCode = -9649
	ObErrNoAttrFound                                                              ObErrorCode = -9650
	ObErrIllegalTypeForObject                                                     ObErrorCode = -9651
	ObErrUnsupportedType                                                          ObErrorCode = -9652
	ObErrPositionalFollowName                                                     ObErrorCode = -9653
	ObErrNeedALabel                                                               ObErrorCode = -9654
	ObErrReferSamePackage                                                         ObErrorCode = -9655
	ObErrPlCommon                                                                 ObErrorCode = -9656
	ObErrIdentEmpty                                                               ObErrorCode = -9657
	ObErrPragmaStrUnsupport                                                       ObErrorCode = -9658
	ObErrEndLabelNotMatch                                                         ObErrorCode = -9659
	ObErrWrongFetchIntoNum                                                        ObErrorCode = -9660
	ObErrPragmaFirstArg                                                           ObErrorCode = -9661
	ObErrTriggerCantChangeOldRow                                                  ObErrorCode = -9662
	ObErrTriggerCantCrtOnRoView                                                   ObErrorCode = -9663
	ObErrTriggerInvalidRefName                                                    ObErrorCode = -9664
	ObErrExpNotIntoTarget                                                         ObErrorCode = -9665
	ObErrCaseNull                                                                 ObErrorCode = -9666
	ObErrInvalidGoto                                                              ObErrorCode = -9667
	ObErrPrivateUdfUseInSql                                                       ObErrorCode = -9668
	ObErrFieldNotDenotable                                                        ObErrorCode = -9669
	ObNumericPrecisionNotInteger                                                  ObErrorCode = -9670
	ObErrRequireInteger                                                           ObErrorCode = -9671
	ObErrIndexTableOfCursor                                                       ObErrorCode = -9672
	ObNullCheckError                                                              ObErrorCode = -9673
	ObErrExNameArg                                                                ObErrorCode = -9674
	ObErrExArgNum                                                                 ObErrorCode = -9675
	ObErrExSecondArg                                                              ObErrorCode = -9676
	ObObenCursorNumberIsZero                                                      ObErrorCode = -9677
	ObNoStmtParse                                                                 ObErrorCode = -9678
	ObArrayCntIsIllegal                                                           ObErrorCode = -9679
	ObErrWrongSchemaRef                                                           ObErrorCode = -9680
	ObErrComponentUndeclared                                                      ObErrorCode = -9681
	ObErrFuncOnlyInSql                                                            ObErrorCode = -9682
	ObErrUndefined                                                                ObErrorCode = -9683
	ObErrSubtypeNotnullMismatch                                                   ObErrorCode = -9684
	ObErrBindVarNotExist                                                          ObErrorCode = -9685
	ObErrCursorInOpenDynamicSql                                                   ObErrorCode = -9686
	ObErrInvalidInputArgument                                                     ObErrorCode = -9687
	ObErrClientIdentifierTooLong                                                  ObErrorCode = -9688
	ObErrInvalidNamespaceValue                                                    ObErrorCode = -9689
	ObErrInvalidNamespaceBeg                                                      ObErrorCode = -9690
	ObErrSessionContextExceeded                                                   ObErrorCode = -9691
	ObErrNotCursorNameInCurrentOf                                                 ObErrorCode = -9692
	ObErrNotForUpdateCursorInCurrentOf                                            ObErrorCode = -9693
	ObErrDupSignalSet                                                             ObErrorCode = -9694
	ObErrSignalNotFound                                                           ObErrorCode = -9695
	ObErrInvalidConditionNumber                                                   ObErrorCode = -9696
	ObErrRecursiveSqlLevelsExceeded                                               ObErrorCode = -9697
	ObErrInvalidSection                                                           ObErrorCode = -9698
	ObErrDuplicateTriggerSection                                                  ObErrorCode = -9699
	ObErrParsePlsql                                                               ObErrorCode = -9700
	ObErrSignalWarn                                                               ObErrorCode = -9701
	ObErrResignalWithoutActiveHandler                                             ObErrorCode = -9702
	ObErrCannotUpdateVirtualColInTrg                                              ObErrorCode = -9703
	ObErrTrgOrder                                                                 ObErrorCode = -9704
	ObErrRefAnotherTableInTrg                                                     ObErrorCode = -9705
	ObErrRefTypeInTrg                                                             ObErrorCode = -9706
	ObErrRefCyclicInTrg                                                           ObErrorCode = -9707
	ObErrCannotSpecifyPrecedesInTrg                                               ObErrorCode = -9708
	ObErrCannotPerformDmlInsideQuery                                              ObErrorCode = -9709
	ObErrCannotPerformDdlCommitOrRollbackInsideQueryOrDmlTips                     ObErrorCode = -9710
	ObErrStatementStringInExecuteImmediateIsNullOrZeroLength                      ObErrorCode = -9711
	ObErrMissingIntoKeyword                                                       ObErrorCode = -9712
	ObErrClauseReturnIllegal                                                      ObErrorCode = -9713
	ObErrNameHasTooManyParts                                                      ObErrorCode = -9714
	ObErrLobSpanTransaction                                                       ObErrorCode = -9715
	ObErrInvalidMultiset                                                          ObErrorCode = -9716
	ObErrInvalidCastUdt                                                           ObErrorCode = -9717
	ObErrPolicyExist                                                              ObErrorCode = -9718
	ObErrPolicyNotExist                                                           ObErrorCode = -9719
	ObErrAddPolicyToSysObject                                                     ObErrorCode = -9720
	ObErrInvalidInputString                                                       ObErrorCode = -9721
	ObErrSecColumnOnView                                                          ObErrorCode = -9722
	ObErrInvalidInputForArgument                                                  ObErrorCode = -9723
	ObErrPolicyDisabled                                                           ObErrorCode = -9724
	ObErrCircularPolicies                                                         ObErrorCode = -9725
	ObErrTooManyPolicies                                                          ObErrorCode = -9726
	ObErrPolicyFunction                                                           ObErrorCode = -9727
	ObErrNoPrivEvalPredicate                                                      ObErrorCode = -9728
	ObErrExecutePolicyFunction                                                    ObErrorCode = -9729
	ObErrPolicyPredicate                                                          ObErrorCode = -9730
	ObErrNoPrivDirectPathAccess                                                   ObErrorCode = -9731
	ObErrIntegrityConstraintViolated                                              ObErrorCode = -9732
	ObErrPolicyGroupExist                                                         ObErrorCode = -9733
	ObErrPolicyGroupNotExist                                                      ObErrorCode = -9734
	ObErrDrivingContextExist                                                      ObErrorCode = -9735
	ObErrDrivingContextNotExist                                                   ObErrorCode = -9736
	ObErrUpdateDefaultGroup                                                       ObErrorCode = -9737
	ObErrContextContainInvalidGroup                                               ObErrorCode = -9738
	ObErrInvalidSecColumnType                                                     ObErrorCode = -9739
	ObErrUnprotectedVirtualColumn                                                 ObErrorCode = -9740
	ObErrAttributeAssociation                                                     ObErrorCode = -9741
	ObErrMergeIntoWithPolicy                                                      ObErrorCode = -9742
	ObErrSpNoDropSp                                                               ObErrorCode = -9743
	ObErrRecompilationObject                                                      ObErrorCode = -9744
	ObErrVariableNotInSelectList                                                  ObErrorCode = -9745
	ObErrMultiRecord                                                              ObErrorCode = -9746
	ObErrMalformedPsPacket                                                        ObErrorCode = -9747
	ObErrViewSelectContainQuestionmark                                            ObErrorCode = -9748
	ObErrObjectNotExist                                                           ObErrorCode = -9749
	ObErrTableOutOfRange                                                          ObErrorCode = -9750
	ObErrWrongUsage                                                               ObErrorCode = -9751
	ObErrForallOnRemoteTable                                                      ObErrorCode = -9752
	ObErrSequenceNotDefine                                                        ObErrorCode = -9753
	ObErrDebugIdNotExist                                                          ObErrorCode = -9754
	ObTTLNotEnable                                                                ObErrorCode = -10501
	ObTTLColumnNotExist                                                           ObErrorCode = -10502
	ObTTLColumnTypeNotSupported                                                   ObErrorCode = -10503
	ObTTLCmdNotAllowed                                                            ObErrorCode = -10504
	ObTTLNoTaskRunning                                                            ObErrorCode = -10505
	ObTTLTenantIsRestore                                                          ObErrorCode = -10506
	ObTTLInvalidHbaseTtl                                                          ObErrorCode = -10507
	ObTTLInvalidHbaseMaxVersions                                                  ObErrorCode = -10508
	ObKvCredentialNotMatch                                                        ObErrorCode = -10509
	ObKvRowkeyCountNotMatch                                                       ObErrorCode = -10510
	ObKvColumnTypeNotMatch                                                        ObErrorCode = -10511
	ObKvCollationMismatch                                                         ObErrorCode = -10512
	ObKvScanRangeMissing                                                          ObErrorCode = -10513
	ObKvRedisParseError                                                           ObErrorCode = -10515
	ObErrValuesClauseNeedHaveColumn                                               ObErrorCode = -11000
	ObErrValuesClauseCannotUseDefaultValues                                       ObErrorCode = -11001
	ObWrongPartitionName                                                          ObErrorCode = -11002
	ObErrPluginIsNotLoaded                                                        ObErrorCode = -11003
	ObErrArgumentShouldConstantOrGroupExpr                                        ObErrorCode = -11010
	ObSpRaiseApplicationError                                                     ObErrorCode = -20000
	ObSpRaiseApplicationErrorNum                                                  ObErrorCode = -21000
	ObClobOnlySupportWithMultibyteFun                                             ObErrorCode = -22998
	ObErrUpdateTwice                                                              ObErrorCode = -30926
	ObErrFlashbackQueryWithUpdate                                                 ObErrorCode = -32491
	ObErrUpdateOnExpr                                                             ObErrorCode = -38104
	ObErrSpecifiedRowNoLongerExists                                               ObErrorCode = -38105
)

var ObErrorNames = map[ObErrorCode]string{
	ObSuccess:                                        "ObSuccess",
	ObError:                                          "ObError",
	ObObjTypeError:                                   "ObObjTypeError",
	ObInvalidArgument:                                "ObInvalidArgument",
	ObArrayOutOfRange:                                "ObArrayOutOfRange",
	ObServerListenError:                              "ObServerListenError",
	ObInitTwice:                                      "ObInitTwice",
	ObNotInit:                                        "ObNotInit",
	ObNotSupported:                                   "ObNotSupported",
	ObIterEnd:                                        "ObIterEnd",
	ObIoError:                                        "ObIoError",
	ObErrorFuncVersion:                               "ObErrorFuncVersion",
	ObTimeout:                                        "ObTimeout",
	ObAllocateMemoryFailed:                           "ObAllocateMemoryFailed",
	ObInnerStatError:                                 "ObInnerStatError",
	ObErrSys:                                         "ObErrSys",
	ObErrUnexpected:                                  "ObErrUnexpected",
	ObEntryExist:                                     "ObEntryExist",
	ObEntryNotExist:                                  "ObEntryNotExist",
	ObSizeOverflow:                                   "ObSizeOverflow",
	ObRefNumNotZero:                                  "ObRefNumNotZero",
	ObConflictValue:                                  "ObConflictValue",
	ObItemNotSetted:                                  "ObItemNotSetted",
	ObEagain:                                         "ObEagain",
	ObBufNotEnough:                                   "ObBufNotEnough",
	ObReadNothing:                                    "ObReadNothing",
	ObFileNotExist:                                   "ObFileNotExist",
	ObDiscontinuousLog:                               "ObDiscontinuousLog",
	ObSerializeError:                                 "ObSerializeError",
	ObDeserializeError:                               "ObDeserializeError",
	ObAioTimeout:                                     "ObAioTimeout",
	ObNeedRetry:                                      "ObNeedRetry",
	ObNotMaster:                                      "ObNotMaster",
	ObDecryptFailed:                                  "ObDecryptFailed",
	ObNotTheObject:                                   "ObNotTheObject",
	ObLastLogRuinned:                                 "ObLastLogRuinned",
	ObInvalidError:                                   "ObInvalidError",
	ObDecimalOverflowWarn:                            "ObDecimalOverflowWarn",
	ObEmptyRange:                                     "ObEmptyRange",
	ObDirNotExist:                                    "ObDirNotExist",
	ObInvalidData:                                    "ObInvalidData",
	ObCanceled:                                       "ObCanceled",
	ObLogNotAlign:                                    "ObLogNotAlign",
	ObNotImplement:                                   "ObNotImplement",
	ObDivisionByZero:                                 "ObDivisionByZero",
	ObExceedMemLimit:                                 "ObExceedMemLimit",
	ObQueueOverflow:                                  "ObQueueOverflow",
	ObStartLogCursorInvalid:                          "ObStartLogCursorInvalid",
	ObLockNotMatch:                                   "ObLockNotMatch",
	ObDeadLock:                                       "ObDeadLock",
	ObChecksumError:                                  "ObChecksumError",
	ObInitFail:                                       "ObInitFail",
	ObRowkeyOrderError:                               "ObRowkeyOrderError",
	ObPhysicChecksumError:                            "ObPhysicChecksumError",
	ObStateNotMatch:                                  "ObStateNotMatch",
	ObInStopState:                                    "ObInStopState",
	ObLogNotClear:                                    "ObLogNotClear",
	ObFileAlreadyExist:                               "ObFileAlreadyExist",
	ObUnknownPacket:                                  "ObUnknownPacket",
	ObRpcPacketTooLong:                               "ObRpcPacketTooLong",
	ObLogTooLarge:                                    "ObLogTooLarge",
	ObRpcSendError:                                   "ObRpcSendError",
	ObRpcPostError:                                   "ObRpcPostError",
	ObLibeasyError:                                   "ObLibeasyError",
	ObConnectError:                                   "ObConnectError",
	ObRpcPacketInvalid:                               "ObRpcPacketInvalid",
	ObBadAddress:                                     "ObBadAddress",
	ObErrMinValue:                                    "ObErrMinValue",
	ObErrMaxValue:                                    "ObErrMaxValue",
	ObErrNullValue:                                   "ObErrNullValue",
	ObResourceOut:                                    "ObResourceOut",
	ObErrSqlClient:                                   "ObErrSqlClient",
	ObOperateOverflow:                                "ObOperateOverflow",
	ObInvalidDateFormat:                              "ObInvalidDateFormat",
	ObInvalidArgumentNum:                             "ObInvalidArgumentNum",
	ObEmptyResult:                                    "ObEmptyResult",
	ObLogInvalidModId:                                "ObLogInvalidModId",
	ObLogModuleUnknown:                               "ObLogModuleUnknown",
	ObLogLevelInvalid:                                "ObLogLevelInvalid",
	ObLogParserSyntaxErr:                             "ObLogParserSyntaxErr",
	ObUnknownConnection:                              "ObUnknownConnection",
	ObErrorOutOfRange:                                "ObErrorOutOfRange",
	ObOpNotAllow:                                     "ObOpNotAllow",
	ObErrAlreadyExists:                               "ObErrAlreadyExists",
	ObSearchNotFound:                                 "ObSearchNotFound",
	ObItemNotMatch:                                   "ObItemNotMatch",
	ObInvalidDateFormatEnd:                           "ObInvalidDateFormatEnd",
	ObDdlTaskExecuteTooMuchTime:                      "ObDdlTaskExecuteTooMuchTime",
	ObHashExist:                                      "ObHashExist",
	ObHashNotExist:                                   "ObHashNotExist",
	ObHashGetTimeout:                                 "ObHashGetTimeout",
	ObHashPlacementRetry:                             "ObHashPlacementRetry",
	ObHashFull:                                       "ObHashFull",
	ObWaitNextTimeout:                                "ObWaitNextTimeout",
	ObMajorFreezeNotFinished:                         "ObMajorFreezeNotFinished",
	ObInvalidDateValue:                               "ObInvalidDateValue",
	ObInactiveSqlClient:                              "ObInactiveSqlClient",
	ObInactiveRpcProxy:                               "ObInactiveRpcProxy",
	ObIntervalWithMonth:                              "ObIntervalWithMonth",
	ObTooManyDatetimeParts:                           "ObTooManyDatetimeParts",
	ObDataOutOfRange:                                 "ObDataOutOfRange",
	ObErrTruncatedWrongValueForField:                 "ObErrTruncatedWrongValueForField",
	ObErrOutOfLowerBound:                             "ObErrOutOfLowerBound",
	ObErrOutOfUpperBound:                             "ObErrOutOfUpperBound",
	ObBadNullError:                                   "ObBadNullError",
	ObFileNotOpened:                                  "ObFileNotOpened",
	ObErrDataTruncated:                               "ObErrDataTruncated",
	ObNotRunning:                                     "ObNotRunning",
	ObErrCompressDecompressData:                      "ObErrCompressDecompressData",
	ObErrIncorrectStringValue:                        "ObErrIncorrectStringValue",
	ObDatetimeFunctionOverflow:                       "ObDatetimeFunctionOverflow",
	ObErrDoubleTruncated:                             "ObErrDoubleTruncated",
	ObCacheFreeBlockNotEnough:                        "ObCacheFreeBlockNotEnough",
	ObLastLogNotComplete:                             "ObLastLogNotComplete",
	ObUnexpectInternalError:                          "ObUnexpectInternalError",
	ObErrTooMuchTime:                                 "ObErrTooMuchTime",
	ObErrThreadPanic:                                 "ObErrThreadPanic",
	ObErrIntervalPartitionExist:                      "ObErrIntervalPartitionExist",
	ObErrIntervalPartitionError:                      "ObErrIntervalPartitionError",
	ObFrozenInfoAlreadyExist:                         "ObFrozenInfoAlreadyExist",
	ObCreateStandbyTenantFailed:                      "ObCreateStandbyTenantFailed",
	ObLsWaitingSafeDestroy:                           "ObLsWaitingSafeDestroy",
	ObLsLockConflict:                                 "ObLsLockConflict",
	ObInvalidRootKey:                                 "ObInvalidRootKey",
	ObErrParserSyntax:                                "ObErrParserSyntax",
	ObErrColumnNotFound:                              "ObErrColumnNotFound",
	ObErrSysVariableUnknown:                          "ObErrSysVariableUnknown",
	ObErrReadOnly:                                    "ObErrReadOnly",
	ObIntegerPrecisionOverflow:                       "ObIntegerPrecisionOverflow",
	ObDecimalPrecisionOverflow:                       "ObDecimalPrecisionOverflow",
	ObNumericOverflow:                                "ObNumericOverflow",
	ObErrSysConfigUnknown:                            "ObErrSysConfigUnknown",
	ObInvalidArgumentForExtract:                      "ObInvalidArgumentForExtract",
	ObInvalidArgumentForIs:                           "ObInvalidArgumentForIs",
	ObInvalidArgumentForLength:                       "ObInvalidArgumentForLength",
	ObInvalidArgumentForSubstr:                       "ObInvalidArgumentForSubstr",
	ObInvalidArgumentForTimeToUsec:                   "ObInvalidArgumentForTimeToUsec",
	ObInvalidArgumentForUsecToTime:                   "ObInvalidArgumentForUsecToTime",
	ObInvalidNumeric:                                 "ObInvalidNumeric",
	ObErrRegexpError:                                 "ObErrRegexpError",
	ObErrUnknownCharset:                              "ObErrUnknownCharset",
	ObErrUnknownCollation:                            "ObErrUnknownCollation",
	ObErrCollationMismatch:                           "ObErrCollationMismatch",
	ObErrWrongValueForVar:                            "ObErrWrongValueForVar",
	ObTenantNotInServer:                              "ObTenantNotInServer",
	ObTenantNotExist:                                 "ObTenantNotExist",
	ObErrDataTooLong:                                 "ObErrDataTooLong",
	ObErrWrongValueCountOnRow:                        "ObErrWrongValueCountOnRow",
	ObCantAggregate_2collations:                      "ObCantAggregate_2collations",
	ObErrUnknownTimeZone:                             "ObErrUnknownTimeZone",
	ObErrTooBigPrecision:                             "ObErrTooBigPrecision",
	ObErrMBiggerThanD:                                "ObErrMBiggerThanD",
	ObErrTruncatedWrongValue:                         "ObErrTruncatedWrongValue",
	ObErrWrongValue:                                  "ObErrWrongValue",
	ObErrUnexpectedTzTransition:                      "ObErrUnexpectedTzTransition",
	ObErrInvalidTimezoneRegionId:                     "ObErrInvalidTimezoneRegionId",
	ObErrInvalidHexNumber:                            "ObErrInvalidHexNumber",
	ObErrFieldNotFoundInDatetimeOrInterval:           "ObErrFieldNotFoundInDatetimeOrInterval",
	ObErrInvalidJsonText:                             "ObErrInvalidJsonText",
	ObErrInvalidJsonTextInParam:                      "ObErrInvalidJsonTextInParam",
	ObErrInvalidJsonBinaryData:                       "ObErrInvalidJsonBinaryData",
	ObErrInvalidJsonPath:                             "ObErrInvalidJsonPath",
	ObErrInvalidJsonCharset:                          "ObErrInvalidJsonCharset",
	ObErrInvalidJsonCharsetInFunction:                "ObErrInvalidJsonCharsetInFunction",
	ObErrInvalidTypeForJson:                          "ObErrInvalidTypeForJson",
	ObErrInvalidCastToJson:                           "ObErrInvalidCastToJson",
	ObErrInvalidJsonPathCharset:                      "ObErrInvalidJsonPathCharset",
	ObErrInvalidJsonPathWildcard:                     "ObErrInvalidJsonPathWildcard",
	ObErrJsonValueTooBig:                             "ObErrJsonValueTooBig",
	ObErrJsonKeyTooBig:                               "ObErrJsonKeyTooBig",
	ObErrJsonUsedAsKey:                               "ObErrJsonUsedAsKey",
	ObErrJsonVacuousPath:                             "ObErrJsonVacuousPath",
	ObErrJsonBadOneOrAllArg:                          "ObErrJsonBadOneOrAllArg",
	ObErrNumericJsonValueOutOfRange:                  "ObErrNumericJsonValueOutOfRange",
	ObErrInvalidJsonValueForCast:                     "ObErrInvalidJsonValueForCast",
	ObErrJsonOutOfDepth:                              "ObErrJsonOutOfDepth",
	ObErrJsonDocumentNullKey:                         "ObErrJsonDocumentNullKey",
	ObErrBlobCantHaveDefault:                         "ObErrBlobCantHaveDefault",
	ObErrInvalidJsonPathArrayCell:                    "ObErrInvalidJsonPathArrayCell",
	ObErrMissingJsonValue:                            "ObErrMissingJsonValue",
	ObErrMultipleJsonValues:                          "ObErrMultipleJsonValues",
	ObInvalidArgumentForTimestampToScn:               "ObInvalidArgumentForTimestampToScn",
	ObInvalidArgumentForScnToTimestamp:               "ObInvalidArgumentForScnToTimestamp",
	ObErrInvalidJsonType:                             "ObErrInvalidJsonType",
	ObErrJsonPathSyntaxError:                         "ObErrJsonPathSyntaxError",
	ObErrJsonValueNoScalar:                           "ObErrJsonValueNoScalar",
	ObErrDuplicateKey:                                "ObErrDuplicateKey",
	ObErrJsonPathExpressionSyntaxError:               "ObErrJsonPathExpressionSyntaxError",
	ObErrNotIso8601Format:                            "ObErrNotIso8601Format",
	ObErrValueExceededMax:                            "ObErrValueExceededMax",
	ObErrBoolNotConvertNumber:                        "ObErrBoolNotConvertNumber",
	ObErrJsonKeyNotFound:                             "ObErrJsonKeyNotFound",
	ObErrYearConflictsWithJulianDate:                 "ObErrYearConflictsWithJulianDate",
	ObErrDayOfYearConflictsWithJulianDate:            "ObErrDayOfYearConflictsWithJulianDate",
	ObErrMonthConflictsWithJulianDate:                "ObErrMonthConflictsWithJulianDate",
	ObErrDayOfMonthConflictsWithJulianDate:           "ObErrDayOfMonthConflictsWithJulianDate",
	ObErrDayOfWeekConflictsWithJulianDate:            "ObErrDayOfWeekConflictsWithJulianDate",
	ObErrHourConflictsWithSecondsInDay:               "ObErrHourConflictsWithSecondsInDay",
	ObErrMinutesOfHourConflictsWithSecondsInDay:      "ObErrMinutesOfHourConflictsWithSecondsInDay",
	ObErrSecondsOfMinuteConflictsWithSecondsInDay:    "ObErrSecondsOfMinuteConflictsWithSecondsInDay",
	ObErrDateNotValidForMonthSpecified:               "ObErrDateNotValidForMonthSpecified",
	ObErrInputValueNotLongEnough:                     "ObErrInputValueNotLongEnough",
	ObErrInvalidYearValue:                            "ObErrInvalidYearValue",
	ObErrInvalidQuarterValue:                         "ObErrInvalidQuarterValue",
	ObErrInvalidMonth:                                "ObErrInvalidMonth",
	ObErrInvalidDayOfTheWeek:                         "ObErrInvalidDayOfTheWeek",
	ObErrInvalidDayOfYearValue:                       "ObErrInvalidDayOfYearValue",
	ObErrInvalidHour12Value:                          "ObErrInvalidHour12Value",
	ObErrInvalidHour24Value:                          "ObErrInvalidHour24Value",
	ObErrInvalidMinutesValue:                         "ObErrInvalidMinutesValue",
	ObErrInvalidSecondsValue:                         "ObErrInvalidSecondsValue",
	ObErrInvalidSecondsInDayValue:                    "ObErrInvalidSecondsInDayValue",
	ObErrInvalidJulianDateValue:                      "ObErrInvalidJulianDateValue",
	ObErrAmOrPmRequired:                              "ObErrAmOrPmRequired",
	ObErrBcOrAdRequired:                              "ObErrBcOrAdRequired",
	ObErrFormatCodeAppearsTwice:                      "ObErrFormatCodeAppearsTwice",
	ObErrDayOfWeekSpecifiedMoreThanOnce:              "ObErrDayOfWeekSpecifiedMoreThanOnce",
	ObErrSignedYearPrecludesUseOfBcAd:                "ObErrSignedYearPrecludesUseOfBcAd",
	ObErrJulianDatePrecludesUseOfDayOfYear:           "ObErrJulianDatePrecludesUseOfDayOfYear",
	ObErrYearMayOnlyBeSpecifiedOnce:                  "ObErrYearMayOnlyBeSpecifiedOnce",
	ObErrHourMayOnlyBeSpecifiedOnce:                  "ObErrHourMayOnlyBeSpecifiedOnce",
	ObErrAmPmConflictsWithUseOfAmDotPmDot:            "ObErrAmPmConflictsWithUseOfAmDotPmDot",
	ObErrBcAdConflictWithUseOfBcDotAdDot:             "ObErrBcAdConflictWithUseOfBcDotAdDot",
	ObErrMonthMayOnlyBeSpecifiedOnce:                 "ObErrMonthMayOnlyBeSpecifiedOnce",
	ObErrDayOfWeekMayOnlyBeSpecifiedOnce:             "ObErrDayOfWeekMayOnlyBeSpecifiedOnce",
	ObErrFormatCodeCannotAppear:                      "ObErrFormatCodeCannotAppear",
	ObErrNonNumericCharacterValue:                    "ObErrNonNumericCharacterValue",
	ObInvalidMeridianIndicatorUse:                    "ObInvalidMeridianIndicatorUse",
	ObErrDayOfMonthRange:                             "ObErrDayOfMonthRange",
	ObErrArgumentOutOfRange:                          "ObErrArgumentOutOfRange",
	ObErrIntervalInvalid:                             "ObErrIntervalInvalid",
	ObErrTheLeadingPrecisionOfTheIntervalIsTooSmall:  "ObErrTheLeadingPrecisionOfTheIntervalIsTooSmall",
	ObErrInvalidTimeZoneHour:                         "ObErrInvalidTimeZoneHour",
	ObErrInvalidTimeZoneMinute:                       "ObErrInvalidTimeZoneMinute",
	ObErrNotAValidTimeZone:                           "ObErrNotAValidTimeZone",
	ObErrDateFormatIsTooLongForInternalBuffer:        "ObErrDateFormatIsTooLongForInternalBuffer",
	ObErrOperatorCannotBeUsedWithList:                "ObErrOperatorCannotBeUsedWithList",
	ObInvalidRowId:                                   "ObInvalidRowId",
	ObErrNumericNotMatchFormatLength:                 "ObErrNumericNotMatchFormatLength",
	ObErrDatetimeIntervalInternalError:               "ObErrDatetimeIntervalInternalError",
	ObErrDblinkRemoteEcode:                           "ObErrDblinkRemoteEcode",
	ObErrDblinkNoLib:                                 "ObErrDblinkNoLib",
	ObSwitchingToFollowerGracefully:                  "ObSwitchingToFollowerGracefully",
	ObMaskSetNoNode:                                  "ObMaskSetNoNode",
	ObTransTimeout:                                   "ObTransTimeout",
	ObTransKilled:                                    "ObTransKilled",
	ObTransStmtTimeout:                               "ObTransStmtTimeout",
	ObTransCtxNotExist:                               "ObTransCtxNotExist",
	ObTransUnknown:                                   "ObTransUnknown",
	ObErrReadOnlyTransaction:                         "ObErrReadOnlyTransaction",
	ObErrGisDifferentSrids:                           "ObErrGisDifferentSrids",
	ObErrGisUnsupportedArgument:                      "ObErrGisUnsupportedArgument",
	ObErrGisUnknownError:                             "ObErrGisUnknownError",
	ObErrGisUnknownException:                         "ObErrGisUnknownException",
	ObErrGisInvalidData:                              "ObErrGisInvalidData",
	ObErrBoostGeometryEmptyInputException:            "ObErrBoostGeometryEmptyInputException",
	ObErrBoostGeometryCentroidException:              "ObErrBoostGeometryCentroidException",
	ObErrBoostGeometryOverlayInvalidInputException:   "ObErrBoostGeometryOverlayInvalidInputException",
	ObErrBoostGeometryTurnInfoException:              "ObErrBoostGeometryTurnInfoException",
	ObErrBoostGeometrySelfIntersectionPointException: "ObErrBoostGeometrySelfIntersectionPointException",
	ObErrBoostGeometryUnknownException:               "ObErrBoostGeometryUnknownException",
	ObErrGisDataWrongEndianess:                       "ObErrGisDataWrongEndianess",
	ObErrAlterOperationNotSupportedReasonGis:         "ObErrAlterOperationNotSupportedReasonGis",
	ObErrBoostGeometryInconsistentTurnsException:     "ObErrBoostGeometryInconsistentTurnsException",
	ObErrGisMaxPointsInGeometryOverflowed:            "ObErrGisMaxPointsInGeometryOverflowed",
	ObErrUnexpectedGeometryType:                      "ObErrUnexpectedGeometryType",
	ObErrSrsParseError:                               "ObErrSrsParseError",
	ObErrSrsProjParameterMissing:                     "ObErrSrsProjParameterMissing",
	ObErrWarnSrsNotFound:                             "ObErrWarnSrsNotFound",
	ObErrSrsNotCartesian:                             "ObErrSrsNotCartesian",
	ObErrSrsNotCartesianUndefined:                    "ObErrSrsNotCartesianUndefined",
	ObErrSrsNotFound:                                 "ObErrSrsNotFound",
	ObErrWarnSrsNotFoundAxisOrder:                    "ObErrWarnSrsNotFoundAxisOrder",
	ObErrNotImplementedForGeographicSrs:              "ObErrNotImplementedForGeographicSrs",
	ObErrWrongSridForColumn:                          "ObErrWrongSridForColumn",
	ObErrCannotAlterSridDueToIndex:                   "ObErrCannotAlterSridDueToIndex",
	ObErrWarnUselessSpatialIndex:                     "ObErrWarnUselessSpatialIndex",
	ObErrOnlyImplementedForSrid_0And_4326:            "ObErrOnlyImplementedForSrid_0And_4326",
	ObErrNotImplementedForCartesianSrs:               "ObErrNotImplementedForCartesianSrs",
	ObErrNotImplementedForProjectedSrs:               "ObErrNotImplementedForProjectedSrs",
	ObErrSrsMissingMandatoryAttribute:                "ObErrSrsMissingMandatoryAttribute",
	ObErrSrsMultipleAttributeDefinitions:             "ObErrSrsMultipleAttributeDefinitions",
	ObErrSrsNameCantBeEmptyOrWhitespace:              "ObErrSrsNameCantBeEmptyOrWhitespace",
	ObErrSrsOrganizationCantBeEmptyOrWhitespace:      "ObErrSrsOrganizationCantBeEmptyOrWhitespace",
	ObErrSrsIdAlreadyExists:                          "ObErrSrsIdAlreadyExists",
	ObErrWarnSrsIdAlreadyExists:                      "ObErrWarnSrsIdAlreadyExists",
	ObErrCantModifySrid_0:                            "ObErrCantModifySrid_0",
	ObErrWarnReservedSridRange:                       "ObErrWarnReservedSridRange",
	ObErrCantModifySrsUsedByColumn:                   "ObErrCantModifySrsUsedByColumn",
	ObErrSrsInvalidCharacterInAttribute:              "ObErrSrsInvalidCharacterInAttribute",
	ObErrSrsAttributeStringTooLong:                   "ObErrSrsAttributeStringTooLong",
	ObErrSrsNotGeographic:                            "ObErrSrsNotGeographic",
	ObErrPolygonTooLarge:                             "ObErrPolygonTooLarge",
	ObErrSpatialUniqueIndex:                          "ObErrSpatialUniqueIndex",
	ObErrGeometryParamLongitudeOutOfRange:            "ObErrGeometryParamLongitudeOutOfRange",
	ObErrGeometryParamLatitudeOutOfRange:             "ObErrGeometryParamLatitudeOutOfRange",
	ObErrSrsGeogcsInvalidAxes:                        "ObErrSrsGeogcsInvalidAxes",
	ObErrSrsInvalidSemiMajorAxis:                     "ObErrSrsInvalidSemiMajorAxis",
	ObErrSrsInvalidInverseFlattening:                 "ObErrSrsInvalidInverseFlattening",
	ObErrSrsInvalidAngularUnit:                       "ObErrSrsInvalidAngularUnit",
	ObErrSrsInvalidPrimeMeridian:                     "ObErrSrsInvalidPrimeMeridian",
	ObErrTransformSourceSrsNotSupported:              "ObErrTransformSourceSrsNotSupported",
	ObErrTransformTargetSrsNotSupported:              "ObErrTransformTargetSrsNotSupported",
	ObErrTransformSourceSrsMissingTowgs84:            "ObErrTransformSourceSrsMissingTowgs84",
	ObErrTransformTargetSrsMissingTowgs84:            "ObErrTransformTargetSrsMissingTowgs84",
	ObErrFunctionalIndexOnJsonOrGeometryFunction:     "ObErrFunctionalIndexOnJsonOrGeometryFunction",
	ObErrSpatialFunctionalIndex:                      "ObErrSpatialFunctionalIndex",
	ObErrGeometryInUnknownLengthUnit:                 "ObErrGeometryInUnknownLengthUnit",
	ObErrInvalidCastToGeometry:                       "ObErrInvalidCastToGeometry",
	ObErrInvalidCastPolygonRingDirection:             "ObErrInvalidCastPolygonRingDirection",
	ObErrGisDifferentSridsAggregation:                "ObErrGisDifferentSridsAggregation",
	ObErrLongitudeOutOfRange:                         "ObErrLongitudeOutOfRange",
	ObErrLatitudeOutOfRange:                          "ObErrLatitudeOutOfRange",
	ObErrStdBadAllocError:                            "ObErrStdBadAllocError",
	ObErrStdDomainError:                              "ObErrStdDomainError",
	ObErrStdLengthError:                              "ObErrStdLengthError",
	ObErrStdInvalidArgument:                          "ObErrStdInvalidArgument",
	ObErrStdOutOfRangeError:                          "ObErrStdOutOfRangeError",
	ObErrStdOverflowError:                            "ObErrStdOverflowError",
	ObErrStdRangeError:                               "ObErrStdRangeError",
	ObErrStdUnderflowError:                           "ObErrStdUnderflowError",
	ObErrStdLogicError:                               "ObErrStdLogicError",
	ObErrStdRuntimeError:                             "ObErrStdRuntimeError",
	ObErrStdUnknownException:                         "ObErrStdUnknownException",
	ObErrCantCreateGeometryObject:                    "ObErrCantCreateGeometryObject",
	ObErrSridWrongUsage:                              "ObErrSridWrongUsage",
	ObErrIndexOrderWrongUsage:                        "ObErrIndexOrderWrongUsage",
	ObErrSpatialMustHaveGeomCol:                      "ObErrSpatialMustHaveGeomCol",
	ObErrSpatialCantHaveNull:                         "ObErrSpatialCantHaveNull",
	ObErrIndexTypeNotSupportedForSpatialIndex:        "ObErrIndexTypeNotSupportedForSpatialIndex",
	ObErrUnitNotFound:                                "ObErrUnitNotFound",
	ObErrInvalidOptionKeyValuePair:                   "ObErrInvalidOptionKeyValuePair",
	ObErrNonpositiveRadius:                           "ObErrNonpositiveRadius",
	ObErrSrsEmpty:                                    "ObErrSrsEmpty",
	ObErrInvalidOptionKey:                            "ObErrInvalidOptionKey",
	ObErrInvalidOptionValue:                          "ObErrInvalidOptionValue",
	ObErrInvalidGeometryType:                         "ObErrInvalidGeometryType",
	ObPacketClusterIdNotMatch:                        "ObPacketClusterIdNotMatch",
	ObTenantIdNotMatch:                               "ObTenantIdNotMatch",
	ObUriError:                                       "ObUriError",
	ObFinalMd5Error:                                  "ObFinalMd5Error",
	ObOssError:                                       "ObOssError",
	ObInitMd5Error:                                   "ObInitMd5Error",
	ObOutOfElement:                                   "ObOutOfElement",
	ObUpdateMd5Error:                                 "ObUpdateMd5Error",
	ObFileLengthInvalid:                              "ObFileLengthInvalid",
	ObBackupFileNotExist:                             "ObBackupFileNotExist",
	ObInvalidBackupDest:                              "ObInvalidBackupDest",
	ObCosError:                                       "ObCosError",
	ObIoLimit:                                        "ObIoLimit",
	ObBackupBackupReachCopyLimit:                     "ObBackupBackupReachCopyLimit",
	ObBackupIoProhibited:                             "ObBackupIoProhibited",
	ObBackupPermissionDenied:                         "ObBackupPermissionDenied",
	ObEsiObsError:                                    "ObEsiObsError",
	ObBackupMetaIndexNotExist:                        "ObBackupMetaIndexNotExist",
	ObBackupDeviceOutOfSpace:                         "ObBackupDeviceOutOfSpace",
	ObBackupPwriteOffsetNotMatch:                     "ObBackupPwriteOffsetNotMatch",
	ObBackupPwriteContentNotMatch:                    "ObBackupPwriteContentNotMatch",
	ObCloudObjectNotAppendable:                       "ObCloudObjectNotAppendable",
	ObRestoreTenantFailed:                            "ObRestoreTenantFailed",
	ObErrXmlParse:                                    "ObErrXmlParse",
	ObErrXsltParse:                                   "ObErrXsltParse",
	ObPacketNotSent:                                  "ObPacketNotSent",
	ObPartialFailed:                                  "ObPartialFailed",
	ObSchemaError:                                    "ObSchemaError",
	ObTenantOutOfMem:                                 "ObTenantOutOfMem",
	ObUnknownObj:                                     "ObUnknownObj",
	ObNoMonitorData:                                  "ObNoMonitorData",
	ObTooManySstable:                                 "ObTooManySstable",
	ObKilledByThrottling:                             "ObKilledByThrottling",
	ObUserNotExist:                                   "ObUserNotExist",
	ObPasswordWrong:                                  "ObPasswordWrong",
	ObSkeyVersionWrong:                               "ObSkeyVersionWrong",
	ObPushdownStatusChanged:                          "ObPushdownStatusChanged",
	ObStorageSchemaInvalid:                           "ObStorageSchemaInvalid",
	ObMediumCompactionInfoInvalid:                    "ObMediumCompactionInfoInvalid",
	ObNotRegistered:                                  "ObNotRegistered",
	ObWaitqueueTimeout:                               "ObWaitqueueTimeout",
	ObAlreadyRegistered:                              "ObAlreadyRegistered",
	ObNoCsSelected:                                   "ObNoCsSelected",
	ObNoTabletsCreated:                               "ObNoTabletsCreated",
	ObDecimalUnlegalError:                            "ObDecimalUnlegalError",
	ObObjDivideError:                                 "ObObjDivideError",
	ObNotADecimal:                                    "ObNotADecimal",
	ObDecimalPrecisionNotEqual:                       "ObDecimalPrecisionNotEqual",
	ObSessionKilled:                                  "ObSessionKilled",
	ObLogNotSync:                                     "ObLogNotSync",
	ObSessionNotFound:                                "ObSessionNotFound",
	ObInvalidLog:                                     "ObInvalidLog",
	ObAlreadyDone:                                    "ObAlreadyDone",
	ObLogSrcChanged:                                  "ObLogSrcChanged",
	ObLogMissing:                                     "ObLogMissing",
	ObNeedWait:                                       "ObNeedWait",
	ObResultUnknown:                                  "ObResultUnknown",
	ObNoResult:                                       "ObNoResult",
	ObLogIdRangeNotContinuous:                        "ObLogIdRangeNotContinuous",
	ObTermLagged:                                     "ObTermLagged",
	ObTermNotMatch:                                   "ObTermNotMatch",
	ObPartialLog:                                     "ObPartialLog",
	ObNotEnoughStore:                                 "ObNotEnoughStore",
	ObBlockSwitched:                                  "ObBlockSwitched",
	ObReadZeroLog:                                    "ObReadZeroLog",
	ObBlockNeedFreeze:                                "ObBlockNeedFreeze",
	ObBlockFrozen:                                    "ObBlockFrozen",
	ObInFatalState:                                   "ObInFatalState",
	ObUpsMasterExists:                                "ObUpsMasterExists",
	ObNotFree:                                        "ObNotFree",
	ObInitSqlContextError:                            "ObInitSqlContextError",
	ObSkipInvalidRow:                                 "ObSkipInvalidRow",
	ObNoTablet:                                       "ObNoTablet",
	ObSnapshotDiscarded:                              "ObSnapshotDiscarded",
	ObDataNotUptodate:                                "ObDataNotUptodate",
	ObRowModified:                                    "ObRowModified",
	ObVersionNotMatch:                                "ObVersionNotMatch",
	ObEnqueueFailed:                                  "ObEnqueueFailed",
	ObInvalidConfig:                                  "ObInvalidConfig",
	ObStmtExpired:                                    "ObStmtExpired",
	ObMetaTableWithoutUseTable:                       "ObMetaTableWithoutUseTable",
	ObDiscardPacket:                                  "ObDiscardPacket",
	ObPoolRegisteredFailed:                           "ObPoolRegisteredFailed",
	ObPoolUnregisteredFailed:                         "ObPoolUnregisteredFailed",
	ObLeaseNotEnough:                                 "ObLeaseNotEnough",
	ObLeaseNotMatch:                                  "ObLeaseNotMatch",
	ObUpsSwitchNotHappen:                             "ObUpsSwitchNotHappen",
	ObCacheNotHit:                                    "ObCacheNotHit",
	ObNestedLoopNotSupport:                           "ObNestedLoopNotSupport",
	ObIndexOutOfRange:                                "ObIndexOutOfRange",
	ObIntUnderflow:                                   "ObIntUnderflow",
	ObCacheShrinkFailed:                              "ObCacheShrinkFailed",
	ObOldSchemaVersion:                               "ObOldSchemaVersion",
	ObReleaseSchemaError:                             "ObReleaseSchemaError",
	ObNoEmptyEntry:                                   "ObNoEmptyEntry",
	ObBeyondTheRange:                                 "ObBeyondTheRange",
	ObServerOutofDiskSpace:                           "ObServerOutofDiskSpace",
	ObColumnGroupNotFound:                            "ObColumnGroupNotFound",
	ObCsCompressLibError:                             "ObCsCompressLibError",
	ObSchedulerTaskCntMismatch:                       "ObSchedulerTaskCntMismatch",
	ObInvalidMacroBlockType:                          "ObInvalidMacroBlockType",
	ObPgIsRemoved:                                    "ObPgIsRemoved",
	ObPacketProcessed:                                "ObPacketProcessed",
	ObLeaderNotExist:                                 "ObLeaderNotExist",
	ObPrepareMajorFreezeFailed:                       "ObPrepareMajorFreezeFailed",
	ObCommitMajorFreezeFailed:                        "ObCommitMajorFreezeFailed",
	ObAbortMajorFreezeFailed:                         "ObAbortMajorFreezeFailed",
	ObPartitionNotLeader:                             "ObPartitionNotLeader",
	ObWaitMajorFreezeResponseTimeout:                 "ObWaitMajorFreezeResponseTimeout",
	ObCurlError:                                      "ObCurlError",
	ObMajorFreezeNotAllow:                            "ObMajorFreezeNotAllow",
	ObPrepareFreezeFailed:                            "ObPrepareFreezeFailed",
	ObPartitionNotExist:                              "ObPartitionNotExist",
	ObErrNoDefaultForField:                           "ObErrNoDefaultForField",
	ObErrFieldSpecifiedTwice:                         "ObErrFieldSpecifiedTwice",
	ObErrTooLongTableComment:                         "ObErrTooLongTableComment",
	ObErrTooLongFieldComment:                         "ObErrTooLongFieldComment",
	ObErrTooLongIndexComment:                         "ObErrTooLongIndexComment",
	ObNotFollower:                                    "ObNotFollower",
	ObObconfigReturnError:                            "ObObconfigReturnError",
	ObObconfigAppnameMismatch:                        "ObObconfigAppnameMismatch",
	ObErrViewSelectDerived:                           "ObErrViewSelectDerived",
	ObCantMjPath:                                     "ObCantMjPath",
	ObErrNoJoinOrderGenerated:                        "ObErrNoJoinOrderGenerated",
	ObErrNoPathGenerated:                             "ObErrNoPathGenerated",
	ObErrWaitRemoteSchemaRefresh:                     "ObErrWaitRemoteSchemaRefresh",
	ObTimerTaskHasScheduled:                          "ObTimerTaskHasScheduled",
	ObTimerTaskHasNotScheduled:                       "ObTimerTaskHasNotScheduled",
	ObParseDebugSyncError:                            "ObParseDebugSyncError",
	ObUnknownDebugSyncPoint:                          "ObUnknownDebugSyncPoint",
	ObErrInterrupted:                                 "ObErrInterrupted",
	ObInvalidPartition:                               "ObInvalidPartition",
	ObErrTimeoutTruncated:                            "ObErrTimeoutTruncated",
	ObErrTooLongTenantComment:                        "ObErrTooLongTenantComment",
	ObErrNetPacketTooLarge:                           "ObErrNetPacketTooLarge",
	ObTraceDescNotExist:                              "ObTraceDescNotExist",
	ObErrNoDefault:                                   "ObErrNoDefault",
	ObIsChangingLeader:                               "ObIsChangingLeader",
	ObMinorFreezeNotAllow:                            "ObMinorFreezeNotAllow",
	ObLogOutofDiskSpace:                              "ObLogOutofDiskSpace",
	ObRpcConnectError:                                "ObRpcConnectError",
	ObMinorMergeNotAllow:                             "ObMinorMergeNotAllow",
	ObCacheInvalid:                                   "ObCacheInvalid",
	ObReachServerDataCopyInConcurrencyLimit:          "ObReachServerDataCopyInConcurrencyLimit",
	ObWorkingPartitionExist:                          "ObWorkingPartitionExist",
	ObWorkingPartitionNotExist:                       "ObWorkingPartitionNotExist",
	ObLibeasyReachMemLimit:                           "ObLibeasyReachMemLimit",
	ObSyncWashMbTimeout:                              "ObSyncWashMbTimeout",
	ObNotAllowMigrateIn:                              "ObNotAllowMigrateIn",
	ObSchedulerTaskCntMistach:                        "ObSchedulerTaskCntMistach",
	ObMissArgument:                                   "ObMissArgument",
	ObTableIsDeleted:                                 "ObTableIsDeleted",
	ObVersionRangeNotContinues:                       "ObVersionRangeNotContinues",
	ObInvalidIoBuffer:                                "ObInvalidIoBuffer",
	ObPartitionIsRemoved:                             "ObPartitionIsRemoved",
	ObGtsNotReady:                                    "ObGtsNotReady",
	ObMajorSstableNotExist:                           "ObMajorSstableNotExist",
	ObVersionRangeDiscarded:                          "ObVersionRangeDiscarded",
	ObMajorSstableHasMerged:                          "ObMajorSstableHasMerged",
	ObMinorSstableRangeCross:                         "ObMinorSstableRangeCross",
	ObMemtableCannotMinorMerge:                       "ObMemtableCannotMinorMerge",
	ObTaskExist:                                      "ObTaskExist",
	ObAllocateDiskSpaceFailed:                        "ObAllocateDiskSpaceFailed",
	ObCantFindUdf:                                    "ObCantFindUdf",
	ObCantInitializeUdf:                              "ObCantInitializeUdf",
	ObUdfNoPaths:                                     "ObUdfNoPaths",
	ObUdfExists:                                      "ObUdfExists",
	ObCantOpenLibrary:                                "ObCantOpenLibrary",
	ObCantFindDlEntry:                                "ObCantFindDlEntry",
	ObObjectNameExist:                                "ObObjectNameExist",
	ObObjectNameNotExist:                             "ObObjectNameNotExist",
	ObErrDupArgument:                                 "ObErrDupArgument",
	ObErrInvalidSequenceName:                         "ObErrInvalidSequenceName",
	ObErrDupMaxvalueSpec:                             "ObErrDupMaxvalueSpec",
	ObErrDupMinvalueSpec:                             "ObErrDupMinvalueSpec",
	ObErrDupCycleSpec:                                "ObErrDupCycleSpec",
	ObErrDupCacheSpec:                                "ObErrDupCacheSpec",
	ObErrDupOrderSpec:                                "ObErrDupOrderSpec",
	ObErrConflMaxvalueSpec:                           "ObErrConflMaxvalueSpec",
	ObErrConflMinvalueSpec:                           "ObErrConflMinvalueSpec",
	ObErrConflCycleSpec:                              "ObErrConflCycleSpec",
	ObErrConflCacheSpec:                              "ObErrConflCacheSpec",
	ObErrConflOrderSpec:                              "ObErrConflOrderSpec",
	ObErrAlterStartSeqNumberNotAllowed:               "ObErrAlterStartSeqNumberNotAllowed",
	ObErrDupIncrementBySpec:                          "ObErrDupIncrementBySpec",
	ObErrDupStartWithSpec:                            "ObErrDupStartWithSpec",
	ObErrRequireAlterSeqOption:                       "ObErrRequireAlterSeqOption",
	ObErrSeqNotAllowedHere:                           "ObErrSeqNotAllowedHere",
	ObErrSeqNotExist:                                 "ObErrSeqNotExist",
	ObErrSeqOptionMustBeInteger:                      "ObErrSeqOptionMustBeInteger",
	ObErrSeqIncrementCanNotBeZero:                    "ObErrSeqIncrementCanNotBeZero",
	ObErrSeqOptionExceedRange:                        "ObErrSeqOptionExceedRange",
	ObErrMinvalueLargerThanMaxvalue:                  "ObErrMinvalueLargerThanMaxvalue",
	ObErrSeqIncrementTooLarge:                        "ObErrSeqIncrementTooLarge",
	ObErrStartWithLessThanMinvalue:                   "ObErrStartWithLessThanMinvalue",
	ObErrMinvalueExceedCurrval:                       "ObErrMinvalueExceedCurrval",
	ObErrStartWithExceedMaxvalue:                     "ObErrStartWithExceedMaxvalue",
	ObErrMaxvalueExceedCurrval:                       "ObErrMaxvalueExceedCurrval",
	ObErrSeqCacheTooSmall:                            "ObErrSeqCacheTooSmall",
	ObErrSeqOptionOutOfRange:                         "ObErrSeqOptionOutOfRange",
	ObErrSeqCacheTooLarge:                            "ObErrSeqCacheTooLarge",
	ObErrSeqRequireMinvalue:                          "ObErrSeqRequireMinvalue",
	ObErrSeqRequireMaxvalue:                          "ObErrSeqRequireMaxvalue",
	ObErrSeqNoLongerExist:                            "ObErrSeqNoLongerExist",
	ObErrSeqValueExceedLimit:                         "ObErrSeqValueExceedLimit",
	ObErrDivisorIsZero:                               "ObErrDivisorIsZero",
	ObErrAesDecrypt:                                  "ObErrAesDecrypt",
	ObErrAesEncrypt:                                  "ObErrAesEncrypt",
	ObErrAesIvLength:                                 "ObErrAesIvLength",
	ObStoreDirError:                                  "ObStoreDirError",
	ObOpenTwice:                                      "ObOpenTwice",
	ObRaidSuperBlockNotMacth:                         "ObRaidSuperBlockNotMacth",
	ObNotOpen:                                        "ObNotOpen",
	ObNotInService:                                   "ObNotInService",
	ObRaidDiskNotNormal:                              "ObRaidDiskNotNormal",
	ObTenantSchemaNotFull:                            "ObTenantSchemaNotFull",
	ObInvalidQueryTimestamp:                          "ObInvalidQueryTimestamp",
	ObDirNotEmpty:                                    "ObDirNotEmpty",
	ObSchemaNotUptodate:                              "ObSchemaNotUptodate",
	ObRoleNotExist:                                   "ObRoleNotExist",
	ObRoleExist:                                      "ObRoleExist",
	ObPrivDup:                                        "ObPrivDup",
	ObKeystoreExist:                                  "ObKeystoreExist",
	ObKeystoreNotExist:                               "ObKeystoreNotExist",
	ObKeystoreWrongPassword:                          "ObKeystoreWrongPassword",
	ObTablespaceExist:                                "ObTablespaceExist",
	ObTablespaceNotExist:                             "ObTablespaceNotExist",
	ObTablespaceDeleteNotEmpty:                       "ObTablespaceDeleteNotEmpty",
	ObFloatPrecisionOutRange:                         "ObFloatPrecisionOutRange",
	ObNumericPrecisionOutRange:                       "ObNumericPrecisionOutRange",
	ObNumericScaleOutRange:                           "ObNumericScaleOutRange",
	ObKeystoreNotOpen:                                "ObKeystoreNotOpen",
	ObKeystoreOpenNoMasterKey:                        "ObKeystoreOpenNoMasterKey",
	ObSlogReachMaxConcurrency:                        "ObSlogReachMaxConcurrency",
	ObErrByAccessOrSessionClauseNotAllowedForNoaudit: "ObErrByAccessOrSessionClauseNotAllowedForNoaudit",
	ObErrAuditingTheObjectIsNotSupported:             "ObErrAuditingTheObjectIsNotSupported",
	ObErrDdlStatementCannotBeAuditedWithBySessionSpecified: "ObErrDdlStatementCannotBeAuditedWithBySessionSpecified",
	ObErrNotValidPassword:                              "ObErrNotValidPassword",
	ObErrMustChangePassword:                            "ObErrMustChangePassword",
	ObOversizeNeedRetry:                                "ObOversizeNeedRetry",
	ObObconfigClusterNotExist:                          "ObObconfigClusterNotExist",
	ObErrGetMasterKey:                                  "ObErrGetMasterKey",
	ObErrTdeMethod:                                     "ObErrTdeMethod",
	ObKmsServerConnectError:                            "ObKmsServerConnectError",
	ObKmsServerIsBusy:                                  "ObKmsServerIsBusy",
	ObKmsServerUpdateKeyConflict:                       "ObKmsServerUpdateKeyConflict",
	ObErrValueLargerThanAllowed:                        "ObErrValueLargerThanAllowed",
	ObDiskError:                                        "ObDiskError",
	ObUnimplementedFeature:                             "ObUnimplementedFeature",
	ObErrDefensiveCheck:                                "ObErrDefensiveCheck",
	ObClusterNameHashConflict:                          "ObClusterNameHashConflict",
	ObHeapTableExausted:                                "ObHeapTableExausted",
	ObErrIndexKeyNotFound:                              "ObErrIndexKeyNotFound",
	ObUnsupportedDeprecatedFeature:                     "ObUnsupportedDeprecatedFeature",
	ObErrDupRestartSpec:                                "ObErrDupRestartSpec",
	ObGtiNotReady:                                      "ObGtiNotReady",
	ObStackOverflow:                                    "ObStackOverflow",
	ObNotAllowRemovingLeader:                           "ObNotAllowRemovingLeader",
	ObNeedSwitchConsumerGroup:                          "ObNeedSwitchConsumerGroup",
	ObErrRemoteSchemaNotFull:                           "ObErrRemoteSchemaNotFull",
	ObDdlSstableRangeCross:                             "ObDdlSstableRangeCross",
	ObDiskHung:                                         "ObDiskHung",
	ObErrObserverStart:                                 "ObErrObserverStart",
	ObErrObserverStop:                                  "ObErrObserverStop",
	ObErrObserviceStart:                                "ObErrObserviceStart",
	ObEncodingEstSizeOverflow:                          "ObEncodingEstSizeOverflow",
	ObInvalidSubPartitionType:                          "ObInvalidSubPartitionType",
	ObErrUnexpectedUnitStatus:                          "ObErrUnexpectedUnitStatus",
	ObAutoincCacheNotEqual:                             "ObAutoincCacheNotEqual",
	ObImportNotInServer:                                "ObImportNotInServer",
	ObConvertError:                                     "ObConvertError",
	ObBypassTimeout:                                    "ObBypassTimeout",
	ObRsStateNotAllow:                                  "ObRsStateNotAllow",
	ObNoReplicaValid:                                   "ObNoReplicaValid",
	ObNoNeedUpdate:                                     "ObNoNeedUpdate",
	ObCacheTimeout:                                     "ObCacheTimeout",
	ObIterStop:                                         "ObIterStop",
	ObZoneAlreadyMaster:                                "ObZoneAlreadyMaster",
	ObIpPortIsNotSlaveZone:                             "ObIpPortIsNotSlaveZone",
	ObZoneIsNotSlave:                                   "ObZoneIsNotSlave",
	ObZoneIsNotMaster:                                  "ObZoneIsNotMaster",
	ObConfigNotSync:                                    "ObConfigNotSync",
	ObIpPortIsNotZone:                                  "ObIpPortIsNotZone",
	ObMasterZoneNotExist:                               "ObMasterZoneNotExist",
	ObZoneInfoNotExist:                                 "ObZoneInfoNotExist",
	ObGetZoneMasterUpsFailed:                           "ObGetZoneMasterUpsFailed",
	ObMultipleMasterZonesExist:                         "ObMultipleMasterZonesExist",
	ObIndexingZoneInvalid:                              "ObIndexingZoneInvalid",
	ObRootTableRangeNotExist:                           "ObRootTableRangeNotExist",
	ObRootMigrateConcurrencyFull:                       "ObRootMigrateConcurrencyFull",
	ObRootMigrateInfoNotFound:                          "ObRootMigrateInfoNotFound",
	ObNotDataLoadTable:                                 "ObNotDataLoadTable",
	ObDataLoadTableDuplicated:                          "ObDataLoadTableDuplicated",
	ObRootTableIdExist:                                 "ObRootTableIdExist",
	ObIndexTimeout:                                     "ObIndexTimeout",
	ObRootNotIntegrated:                                "ObRootNotIntegrated",
	ObIndexIneligible:                                  "ObIndexIneligible",
	ObRebalanceExecTimeout:                             "ObRebalanceExecTimeout",
	ObMergeNotStarted:                                  "ObMergeNotStarted",
	ObMergeAlreadyStarted:                              "ObMergeAlreadyStarted",
	ObRootserviceExist:                                 "ObRootserviceExist",
	ObRsShutdown:                                       "ObRsShutdown",
	ObServerMigrateInDenied:                            "ObServerMigrateInDenied",
	ObRebalanceTaskCantExec:                            "ObRebalanceTaskCantExec",
	ObPartitionCntReachRootserverLimit:                 "ObPartitionCntReachRootserverLimit",
	ObRebalanceTaskNotInProgress:                       "ObRebalanceTaskNotInProgress",
	ObDataSourceNotExist:                               "ObDataSourceNotExist",
	ObDataSourceTableNotExist:                          "ObDataSourceTableNotExist",
	ObDataSourceRangeNotExist:                          "ObDataSourceRangeNotExist",
	ObDataSourceDataNotExist:                           "ObDataSourceDataNotExist",
	ObDataSourceSysError:                               "ObDataSourceSysError",
	ObDataSourceTimeout:                                "ObDataSourceTimeout",
	ObDataSourceConcurrencyFull:                        "ObDataSourceConcurrencyFull",
	ObDataSourceWrongUriFormat:                         "ObDataSourceWrongUriFormat",
	ObSstableVersionUnequal:                            "ObSstableVersionUnequal",
	ObUpsRenewLeaseNotAllowed:                          "ObUpsRenewLeaseNotAllowed",
	ObUpsCountOverLimit:                                "ObUpsCountOverLimit",
	ObNoUpsMajority:                                    "ObNoUpsMajority",
	ObIndexCountReachTheLimit:                          "ObIndexCountReachTheLimit",
	ObTaskExpired:                                      "ObTaskExpired",
	ObTablegroupNotEmpty:                               "ObTablegroupNotEmpty",
	ObInvalidServerStatus:                              "ObInvalidServerStatus",
	ObWaitElecLeaderTimeout:                            "ObWaitElecLeaderTimeout",
	ObWaitAllRsOnlineTimeout:                           "ObWaitAllRsOnlineTimeout",
	ObAllReplicasOnMergeZone:                           "ObAllReplicasOnMergeZone",
	ObMachineResourceNotEnough:                         "ObMachineResourceNotEnough",
	ObNotServerCanHoldSoftly:                           "ObNotServerCanHoldSoftly",
	ObResourcePoolAlreadyGranted:                       "ObResourcePoolAlreadyGranted",
	ObServerAlreadyDeleted:                             "ObServerAlreadyDeleted",
	ObServerNotDeleting:                                "ObServerNotDeleting",
	ObServerNotInWhiteList:                             "ObServerNotInWhiteList",
	ObServerZoneNotMatch:                               "ObServerZoneNotMatch",
	ObOverZoneNumLimit:                                 "ObOverZoneNumLimit",
	ObZoneStatusNotMatch:                               "ObZoneStatusNotMatch",
	ObResourceUnitIsReferenced:                         "ObResourceUnitIsReferenced",
	ObDifferentPrimaryZone:                             "ObDifferentPrimaryZone",
	ObServerNotActive:                                  "ObServerNotActive",
	ObRsNotMaster:                                      "ObRsNotMaster",
	ObCandidateListError:                               "ObCandidateListError",
	ObPartitionZoneDuplicated:                          "ObPartitionZoneDuplicated",
	ObZoneDuplicated:                                   "ObZoneDuplicated",
	ObNotAllZoneActive:                                 "ObNotAllZoneActive",
	ObPrimaryZoneNotInZoneList:                         "ObPrimaryZoneNotInZoneList",
	ObReplicaNumNotMatch:                               "ObReplicaNumNotMatch",
	ObZoneListPoolListNotMatch:                         "ObZoneListPoolListNotMatch",
	ObInvalidTenantName:                                "ObInvalidTenantName",
	ObEmptyResourcePoolList:                            "ObEmptyResourcePoolList",
	ObResourceUnitNotExist:                             "ObResourceUnitNotExist",
	ObResourceUnitExist:                                "ObResourceUnitExist",
	ObResourcePoolNotExist:                             "ObResourcePoolNotExist",
	ObResourcePoolExist:                                "ObResourcePoolExist",
	ObWaitLeaderSwitchTimeout:                          "ObWaitLeaderSwitchTimeout",
	ObLocationNotExist:                                 "ObLocationNotExist",
	ObLocationLeaderNotExist:                           "ObLocationLeaderNotExist",
	ObZoneNotActive:                                    "ObZoneNotActive",
	ObUnitNumOverServerCount:                           "ObUnitNumOverServerCount",
	ObPoolServerIntersect:                              "ObPoolServerIntersect",
	ObNotSingleResourcePool:                            "ObNotSingleResourcePool",
	ObResourceUnitValueBelowLimit:                      "ObResourceUnitValueBelowLimit",
	ObStopServerInMultipleZones:                        "ObStopServerInMultipleZones",
	ObSessionEntryExist:                                "ObSessionEntryExist",
	ObGotSignalAborting:                                "ObGotSignalAborting",
	ObServerNotAlive:                                   "ObServerNotAlive",
	ObGetLocationTimeOut:                               "ObGetLocationTimeOut",
	ObUnitIsMigrating:                                  "ObUnitIsMigrating",
	ObClusterNoMatch:                                   "ObClusterNoMatch",
	ObCheckZoneMergeOrder:                              "ObCheckZoneMergeOrder",
	ObErrZoneNotEmpty:                                  "ObErrZoneNotEmpty",
	ObDifferentLocality:                                "ObDifferentLocality",
	ObEmptyLocality:                                    "ObEmptyLocality",
	ObFullReplicaNumNotEnough:                          "ObFullReplicaNumNotEnough",
	ObReplicaNumNotEnough:                              "ObReplicaNumNotEnough",
	ObDataSourceNotValid:                               "ObDataSourceNotValid",
	ObRunJobNotSuccess:                                 "ObRunJobNotSuccess",
	ObNoNeedRebuild:                                    "ObNoNeedRebuild",
	ObNeedRemoveUnneedTable:                            "ObNeedRemoveUnneedTable",
	ObNoNeedMerge:                                      "ObNoNeedMerge",
	ObConflictOption:                                   "ObConflictOption",
	ObDuplicateOption:                                  "ObDuplicateOption",
	ObInvalidOption:                                    "ObInvalidOption",
	ObRpcNeedReconnect:                                 "ObRpcNeedReconnect",
	ObCannotCopyMajorSstable:                           "ObCannotCopyMajorSstable",
	ObSrcDoNotAllowedMigrate:                           "ObSrcDoNotAllowedMigrate",
	ObTooManyTenantPartitionsError:                     "ObTooManyTenantPartitionsError",
	ObActiveMemtbaleNotExsit:                           "ObActiveMemtbaleNotExsit",
	ObUseDupFollowAfterDml:                             "ObUseDupFollowAfterDml",
	ObNoDiskNeedRebuild:                                "ObNoDiskNeedRebuild",
	ObStandbyReadOnly:                                  "ObStandbyReadOnly",
	ObInvaldWebServiceContent:                          "ObInvaldWebServiceContent",
	ObPrimaryClusterExist:                              "ObPrimaryClusterExist",
	ObArrayBindingSwitchIterator:                       "ObArrayBindingSwitchIterator",
	ObErrStandbyClusterNotEmpty:                        "ObErrStandbyClusterNotEmpty",
	ObNotPrimaryCluster:                                "ObNotPrimaryCluster",
	ObErrCheckDropColumnFailed:                         "ObErrCheckDropColumnFailed",
	ObNotStandbyCluster:                                "ObNotStandbyCluster",
	ObClusterVersionNotCompatible:                      "ObClusterVersionNotCompatible",
	ObWaitTransTableMergeTimeout:                       "ObWaitTransTableMergeTimeout",
	ObSkipRenewLocationByRpc:                           "ObSkipRenewLocationByRpc",
	ObRenewLocationByRpcFailed:                         "ObRenewLocationByRpcFailed",
	ObClusterIdNoMatch:                                 "ObClusterIdNoMatch",
	ObErrParamInvalid:                                  "ObErrParamInvalid",
	ObErrResObjAlreadyExist:                            "ObErrResObjAlreadyExist",
	ObErrResPlanNotExist:                               "ObErrResPlanNotExist",
	ObErrPercentageOutOfRange:                          "ObErrPercentageOutOfRange",
	ObErrPlanDirectiveNotExist:                         "ObErrPlanDirectiveNotExist",
	ObErrPlanDirectiveAlreadyExist:                     "ObErrPlanDirectiveAlreadyExist",
	ObErrInvalidPlanDirectiveName:                      "ObErrInvalidPlanDirectiveName",
	ObFailoverNotAllow:                                 "ObFailoverNotAllow",
	ObAddClusterNotAllowed:                             "ObAddClusterNotAllowed",
	ObErrConsumerGroupNotExist:                         "ObErrConsumerGroupNotExist",
	ObClusterNotAccessible:                             "ObClusterNotAccessible",
	ObTenantResourceUnitExist:                          "ObTenantResourceUnitExist",
	ObErrDropTruncatePartitionRebuildIndex:             "ObErrDropTruncatePartitionRebuildIndex",
	ObErrAtlerTableIllegalFk:                           "ObErrAtlerTableIllegalFk",
	ObErrNoResourceManagerPrivilege:                    "ObErrNoResourceManagerPrivilege",
	ObLeaderCoordinatorNeedRetry:                       "ObLeaderCoordinatorNeedRetry",
	ObRebalanceTaskNeedRetry:                           "ObRebalanceTaskNeedRetry",
	ObErrResMgrPlanNotExist:                            "ObErrResMgrPlanNotExist",
	ObLsNotExist:                                       "ObLsNotExist",
	ObTooManyTenantLs:                                  "ObTooManyTenantLs",
	ObLsLocationNotExist:                               "ObLsLocationNotExist",
	ObLsLocationLeaderNotExist:                         "ObLsLocationLeaderNotExist",
	ObMappingBetweenTabletAndLsNotExist:                "ObMappingBetweenTabletAndLsNotExist",
	ObTabletExist:                                      "ObTabletExist",
	ObTabletNotExist:                                   "ObTabletNotExist",
	ObErrStandbyStatus:                                 "ObErrStandbyStatus",
	ObLsNeedRevoke:                                     "ObLsNeedRevoke",
	ObErrLastPartitionInTheRangeSectionCannotBeDropped: "ObErrLastPartitionInTheRangeSectionCannotBeDropped",
	ObErrSetIntervalIsNotLegalOnThisTable:              "ObErrSetIntervalIsNotLegalOnThisTable",
	ObCheckClusterStatus:                               "ObCheckClusterStatus",
	ObZoneResourceNotEnough:                            "ObZoneResourceNotEnough",
	ObZoneServerNotEnough:                              "ObZoneServerNotEnough",
	ObSstableNotExist:                                  "ObSstableNotExist",
	ObResourceUnitValueInvalid:                         "ObResourceUnitValueInvalid",
	ObLsExist:                                          "ObLsExist",
	ObDeviceExist:                                      "ObDeviceExist",
	ObDeviceNotExist:                                   "ObDeviceNotExist",
	ObLsReplicaTaskResultUncertain:                     "ObLsReplicaTaskResultUncertain",
	ObWaitReplayTimeout:                                "ObWaitReplayTimeout",
	ObWaitTabletReadyTimeout:                           "ObWaitTabletReadyTimeout",
	ObFreezeServiceEpochMismatch:                       "ObFreezeServiceEpochMismatch",
	ObDeleteServerNotAllowed:                           "ObDeleteServerNotAllowed",
	ObPacketStatusUnknown:                              "ObPacketStatusUnknown",
	ObArbitrationServiceNotExist:                       "ObArbitrationServiceNotExist",
	ObArbitrationServiceAlreadyExist:                   "ObArbitrationServiceAlreadyExist",
	ObUnexpectedTabletStatus:                           "ObUnexpectedTabletStatus",
	ObInvalidTableStore:                                "ObInvalidTableStore",
	ObWaitDegrationTimeout:                             "ObWaitDegrationTimeout",
	ObErrRootserviceStart:                              "ObErrRootserviceStart",
	ObErrRootserviceStop:                               "ObErrRootserviceStop",
	ObErrRootInspection:                                "ObErrRootInspection",
	ObErrRootserviceThreadHung:                         "ObErrRootserviceThreadHung",
	ObMigrateNotCompatible:                             "ObMigrateNotCompatible",
	ObClusterInfoMaybeRemained:                         "ObClusterInfoMaybeRemained",
	ObArbitrationInfoQueryFailed:                       "ObArbitrationInfoQueryFailed",
	ObIgnoreErrAccessVirtualTable:                      "ObIgnoreErrAccessVirtualTable",
	ObLsOffline:                                        "ObLsOffline",
	ObLsIsDeleted:                                      "ObLsIsDeleted",
	ObSkipCheckingLsStatus:                             "ObSkipCheckingLsStatus",
	ObErrUseRowIdForUpdate:                             "ObErrUseRowIdForUpdate",
	ObErrUnknownSetOption:                              "ObErrUnknownSetOption",
	ObLsNotLeader:                                      "ObLsNotLeader",
	ObErrParserInit:                                    "ObErrParserInit",
	ObErrParseSql:                                      "ObErrParseSql",
	ObErrResolveSql:                                    "ObErrResolveSql",
	ObErrGenPlan:                                       "ObErrGenPlan",
	ObErrColumnSize:                                    "ObErrColumnSize",
	ObErrColumnDuplicate:                               "ObErrColumnDuplicate",
	ObErrOperatorUnknown:                               "ObErrOperatorUnknown",
	ObErrStarDuplicate:                                 "ObErrStarDuplicate",
	ObErrIllegalId:                                     "ObErrIllegalId",
	ObErrIllegalValue:                                  "ObErrIllegalValue",
	ObErrColumnAmbiguous:                               "ObErrColumnAmbiguous",
	ObErrLogicalPlanFaild:                              "ObErrLogicalPlanFaild",
	ObErrSchemaUnset:                                   "ObErrSchemaUnset",
	ObErrIllegalName:                                   "ObErrIllegalName",
	ObTableNotExist:                                    "ObTableNotExist",
	ObErrTableExist:                                    "ObErrTableExist",
	ObErrExprUnknown:                                   "ObErrExprUnknown",
	ObErrIllegalType:                                   "ObErrIllegalType",
	ObErrPrimaryKeyDuplicate:                           "ObErrPrimaryKeyDuplicate",
	ObErrKeyNameDuplicate:                              "ObErrKeyNameDuplicate",
	ObErrCreatetimeDuplicate:                           "ObErrCreatetimeDuplicate",
	ObErrModifytimeDuplicate:                           "ObErrModifytimeDuplicate",
	ObErrIllegalIndex:                                  "ObErrIllegalIndex",
	ObErrInvalidSchema:                                 "ObErrInvalidSchema",
	ObErrInsertNullRowkey:                              "ObErrInsertNullRowkey",
	ObErrDeleteNullRowkey:                              "ObErrDeleteNullRowkey",
	ObErrUserEmpty:                                     "ObErrUserEmpty",
	ObErrUserNotExist:                                  "ObErrUserNotExist",
	ObErrNoPrivilege:                                   "ObErrNoPrivilege",
	ObErrNoAvailablePrivilegeEntry:                     "ObErrNoAvailablePrivilegeEntry",
	ObErrWrongPassword:                                 "ObErrWrongPassword",
	ObErrUserIsLocked:                                  "ObErrUserIsLocked",
	ObErrUpdateRowkeyColumn:                            "ObErrUpdateRowkeyColumn",
	ObErrUpdateJoinColumn:                              "ObErrUpdateJoinColumn",
	ObErrInvalidColumnNum:                              "ObErrInvalidColumnNum",
	ObErrPrepareStmtNotFound:                           "ObErrPrepareStmtNotFound",
	ObErrOlderPrivilegeVersion:                         "ObErrOlderPrivilegeVersion",
	ObErrLackOfRowkeyCol:                               "ObErrLackOfRowkeyCol",
	ObErrUserExist:                                     "ObErrUserExist",
	ObErrPasswordEmpty:                                 "ObErrPasswordEmpty",
	ObErrGrantPrivilegesToCreateTable:                  "ObErrGrantPrivilegesToCreateTable",
	ObErrWrongDynamicParam:                             "ObErrWrongDynamicParam",
	ObErrParamSize:                                     "ObErrParamSize",
	ObErrFunctionUnknown:                               "ObErrFunctionUnknown",
	ObErrCreatModifyTimeColumn:                         "ObErrCreatModifyTimeColumn",
	ObErrModifyPrimaryKey:                              "ObErrModifyPrimaryKey",
	ObErrParamDuplicate:                                "ObErrParamDuplicate",
	ObErrTooManySessions:                               "ObErrTooManySessions",
	ObErrTooManyPs:                                     "ObErrTooManyPs",
	ObErrHintUnknown:                                   "ObErrHintUnknown",
	ObErrWhenUnsatisfied:                               "ObErrWhenUnsatisfied",
	ObErrQueryInterrupted:                              "ObErrQueryInterrupted",
	ObErrSessionInterrupted:                            "ObErrSessionInterrupted",
	ObErrUnknownSessionId:                              "ObErrUnknownSessionId",
	ObErrProtocolNotRecognize:                          "ObErrProtocolNotRecognize",
	ObErrWriteAuthError:                                "ObErrWriteAuthError",
	ObErrParseJoinInfo:                                 "ObErrParseJoinInfo",
	ObErrAlterIndexColumn:                              "ObErrAlterIndexColumn",
	ObErrModifyIndexTable:                              "ObErrModifyIndexTable",
	ObErrIndexUnavailable:                              "ObErrIndexUnavailable",
	ObErrNopValue:                                      "ObErrNopValue",
	ObErrPsTooManyParam:                                "ObErrPsTooManyParam",
	ObErrInvalidTypeForOp:                              "ObErrInvalidTypeForOp",
	ObErrCastVarcharToBool:                             "ObErrCastVarcharToBool",
	ObErrCastVarcharToNumber:                           "ObErrCastVarcharToNumber",
	ObErrCastVarcharToTime:                             "ObErrCastVarcharToTime",
	ObErrCastNumberOverflow:                            "ObErrCastNumberOverflow",
	ObSchemaNumberPrecisionOverflow:                    "ObSchemaNumberPrecisionOverflow",
	ObSchemaNumberScaleOverflow:                        "ObSchemaNumberScaleOverflow",
	ObErrIndexUnknown:                                  "ObErrIndexUnknown",
	ObErrTooManyJoinTables:                             "ObErrTooManyJoinTables",
	ObErrDdlOnRemoteDatabase:                           "ObErrDdlOnRemoteDatabase",
	ObErrMissingKeyword:                                "ObErrMissingKeyword",
	ObErrDatabaseLinkExpected:                          "ObErrDatabaseLinkExpected",
	ObErrVarcharTooLong:                                "ObErrVarcharTooLong",
	ObErrLocalVariable:                                 "ObErrLocalVariable",
	ObErrGlobalVariable:                                "ObErrGlobalVariable",
	ObErrVariableIsReadonly:                            "ObErrVariableIsReadonly",
	ObErrIncorrectGlobalLocalVar:                       "ObErrIncorrectGlobalLocalVar",
	ObErrExpireInfoTooLong:                             "ObErrExpireInfoTooLong",
	ObErrExpireCondTooLong:                             "ObErrExpireCondTooLong",
	ObErrUserVariableUnknown:                           "ObErrUserVariableUnknown",
	ObIllegalUsageOfMergingFrozenTime:                  "ObIllegalUsageOfMergingFrozenTime",
	ObSqlLogOpSetchildOverflow:                         "ObSqlLogOpSetchildOverflow",
	ObSqlExplainFailed:                                 "ObSqlExplainFailed",
	ObSqlOptCopyOpFailed:                               "ObSqlOptCopyOpFailed",
	ObSqlOptGenPlanFalied:                              "ObSqlOptGenPlanFalied",
	ObSqlOptCreateRawexprFailed:                        "ObSqlOptCreateRawexprFailed",
	ObSqlOptJoinOrderFailed:                            "ObSqlOptJoinOrderFailed",
	ObSqlOptError:                                      "ObSqlOptError",
	ObErrOciInitTimezone:                               "ObErrOciInitTimezone",
	ObErrZlibData:                                      "ObErrZlibData",
	ObErrDblinkSessionKilled:                           "ObErrDblinkSessionKilled",
	ObSqlResolverNoMemory:                              "ObSqlResolverNoMemory",
	ObSqlDmlOnly:                                       "ObSqlDmlOnly",
	ObErrNoGrant:                                       "ObErrNoGrant",
	ObErrNoDbSelected:                                  "ObErrNoDbSelected",
	ObSqlPcOverflow:                                    "ObSqlPcOverflow",
	ObSqlPcPlanDuplicate:                               "ObSqlPcPlanDuplicate",
	ObSqlPcPlanExpire:                                  "ObSqlPcPlanExpire",
	ObSqlPcNotExist:                                    "ObSqlPcNotExist",
	ObSqlParamsLimit:                                   "ObSqlParamsLimit",
	ObSqlPcPlanSizeLimit:                               "ObSqlPcPlanSizeLimit",
	ObUnknownPartition:                                 "ObUnknownPartition",
	ObPartitionNotMatch:                                "ObPartitionNotMatch",
	ObErPasswdLength:                                   "ObErPasswdLength",
	ObErrInsertInnerJoinColumn:                         "ObErrInsertInnerJoinColumn",
	ObTablegroupNotExist:                               "ObTablegroupNotExist",
	ObSubQueryTooManyRow:                               "ObSubQueryTooManyRow",
	ObErrBadDatabase:                                   "ObErrBadDatabase",
	ObCannotUser:                                       "ObCannotUser",
	ObTenantExist:                                      "ObTenantExist",
	ObDatabaseExist:                                    "ObDatabaseExist",
	ObTablegroupExist:                                  "ObTablegroupExist",
	ObErrInvalidTenantName:                             "ObErrInvalidTenantName",
	ObEmptyTenant:                                      "ObEmptyTenant",
	ObWrongDbName:                                      "ObWrongDbName",
	ObWrongTableName:                                   "ObWrongTableName",
	ObWrongColumnName:                                  "ObWrongColumnName",
	ObErrColumnSpec:                                    "ObErrColumnSpec",
	ObErrDbDropExists:                                  "ObErrDbDropExists",
	ObErrCreateUserWithGrant:                           "ObErrCreateUserWithGrant",
	ObErrNoDbPrivilege:                                 "ObErrNoDbPrivilege",
	ObErrNoTablePrivilege:                              "ObErrNoTablePrivilege",
	ObInvalidOnUpdate:                                  "ObInvalidOnUpdate",
	ObInvalidDefault:                                   "ObInvalidDefault",
	ObErrUpdateTableUsed:                               "ObErrUpdateTableUsed",
	ObErrCoulumnValueNotMatch:                          "ObErrCoulumnValueNotMatch",
	ObErrInvalidGroupFuncUse:                           "ObErrInvalidGroupFuncUse",
	ObErrFieldTypeNotAllowedAsPartitionField:           "ObErrFieldTypeNotAllowedAsPartitionField",
	ObErrTooLongIdent:                                  "ObErrTooLongIdent",
	ObErrWrongTypeForVar:                               "ObErrWrongTypeForVar",
	ObWrongUserNameLength:                              "ObWrongUserNameLength",
	ObErrPrivUsage:                                     "ObErrPrivUsage",
	ObIllegalGrantForTable:                             "ObIllegalGrantForTable",
	ObErrReachAutoincMax:                               "ObErrReachAutoincMax",
	ObErrNoTablesUsed:                                  "ObErrNoTablesUsed",
	ObCantRemoveAllFields:                              "ObCantRemoveAllFields",
	ObTooManyPartitionsError:                           "ObTooManyPartitionsError",
	ObNoPartsError:                                     "ObNoPartsError",
	ObWrongSubKey:                                      "ObWrongSubKey",
	ObKeyPart_0:                                        "ObKeyPart_0",
	ObErrWrongAutoKey:                                  "ObErrWrongAutoKey",
	ObErrTooManyKeys:                                   "ObErrTooManyKeys",
	ObErrTooManyRowkeyColumns:                          "ObErrTooManyRowkeyColumns",
	ObErrTooLongKeyLength:                              "ObErrTooLongKeyLength",
	ObErrTooManyColumns:                                "ObErrTooManyColumns",
	ObErrTooLongColumnLength:                           "ObErrTooLongColumnLength",
	ObErrTooBigRowsize:                                 "ObErrTooBigRowsize",
	ObErrUnknownTable:                                  "ObErrUnknownTable",
	ObErrBadTable:                                      "ObErrBadTable",
	ObErrTooBigScale:                                   "ObErrTooBigScale",
	ObErrTooBigDisplaywidth:                            "ObErrTooBigDisplaywidth",
	ObWrongGroupField:                                  "ObWrongGroupField",
	ObNonUniqError:                                     "ObNonUniqError",
	ObErrNonuniqTable:                                  "ObErrNonuniqTable",
	ObErrCantDropFieldOrKey:                            "ObErrCantDropFieldOrKey",
	ObErrMultiplePriKey:                                "ObErrMultiplePriKey",
	ObErrKeyColumnDoesNotExits:                         "ObErrKeyColumnDoesNotExits",
	ObErrAutoPartitionKey:                              "ObErrAutoPartitionKey",
	ObErrCantUseOptionHere:                             "ObErrCantUseOptionHere",
	ObErrWrongObject:                                   "ObErrWrongObject",
	ObErrOnRename:                                      "ObErrOnRename",
	ObErrWrongKeyColumn:                                "ObErrWrongKeyColumn",
	ObErrBadFieldError:                                 "ObErrBadFieldError",
	ObErrWrongFieldWithGroup:                           "ObErrWrongFieldWithGroup",
	ObErrCantChangeTxCharacteristics:                   "ObErrCantChangeTxCharacteristics",
	ObErrCantExecuteInReadOnlyTransaction:              "ObErrCantExecuteInReadOnlyTransaction",
	ObErrMixOfGroupFuncAndFields:                       "ObErrMixOfGroupFuncAndFields",
	ObErrWrongIdentName:                                "ObErrWrongIdentName",
	ObWrongNameForIndex:                                "ObWrongNameForIndex",
	ObIllegalReference:                                 "ObIllegalReference",
	ObReachMemoryLimit:                                 "ObReachMemoryLimit",
	ObErrPasswordFormat:                                "ObErrPasswordFormat",
	ObErrNonUpdatableTable:                             "ObErrNonUpdatableTable",
	ObErrWarnDataOutOfRange:                            "ObErrWarnDataOutOfRange",
	ObErrWrongExprInPartitionFuncError:                 "ObErrWrongExprInPartitionFuncError",
	ObErrViewInvalid:                                   "ObErrViewInvalid",
	ObErrOptionPreventsStatement:                       "ObErrOptionPreventsStatement",
	ObErrDbReadOnly:                                    "ObErrDbReadOnly",
	ObErrTableReadOnly:                                 "ObErrTableReadOnly",
	ObErrLockOrActiveTransaction:                       "ObErrLockOrActiveTransaction",
	ObErrSameNamePartitionField:                        "ObErrSameNamePartitionField",
	ObErrTablenameNotAllowedHere:                       "ObErrTablenameNotAllowedHere",
	ObErrViewRecursive:                                 "ObErrViewRecursive",
	ObErrQualifier:                                     "ObErrQualifier",
	ObErrViewWrongList:                                 "ObErrViewWrongList",
	ObSysVarsMaybeDiffVersion:                          "ObSysVarsMaybeDiffVersion",
	ObErrAutoIncrementConflict:                         "ObErrAutoIncrementConflict",
	ObErrTaskSkipped:                                   "ObErrTaskSkipped",
	ObErrNameBecomesEmpty:                              "ObErrNameBecomesEmpty",
	ObErrRemovedSpaces:                                 "ObErrRemovedSpaces",
	ObWarnAddAutoincrementColumn:                       "ObWarnAddAutoincrementColumn",
	ObWarnChamgeNullAttribute:                          "ObWarnChamgeNullAttribute",
	ObErrInvalidCharacterString:                        "ObErrInvalidCharacterString",
	ObErrKillDenied:                                    "ObErrKillDenied",
	ObErrColumnDefinitionAmbiguous:                     "ObErrColumnDefinitionAmbiguous",
	ObErrEmptyQuery:                                    "ObErrEmptyQuery",
	ObErrCutValueGroupConcat:                           "ObErrCutValueGroupConcat",
	ObErrFieldNotFoundPart:                             "ObErrFieldNotFoundPart",
	ObErrPrimaryCantHaveNull:                           "ObErrPrimaryCantHaveNull",
	ObErrPartitionFuncNotAllowedError:                  "ObErrPartitionFuncNotAllowedError",
	ObErrInvalidBlockSize:                              "ObErrInvalidBlockSize",
	ObErrUnknownStorageEngine:                          "ObErrUnknownStorageEngine",
	ObErrTenantIsLocked:                                "ObErrTenantIsLocked",
	ObEerUniqueKeyNeedAllFieldsInPf:                    "ObEerUniqueKeyNeedAllFieldsInPf",
	ObErrPartitionFunctionIsNotAllowed:                 "ObErrPartitionFunctionIsNotAllowed",
	ObErrAggregateOrderForUnion:                        "ObErrAggregateOrderForUnion",
	ObErrOutlineExist:                                  "ObErrOutlineExist",
	ObOutlineNotExist:                                  "ObOutlineNotExist",
	ObWarnOptionBelowLimit:                             "ObWarnOptionBelowLimit",
	ObInvalidOutline:                                   "ObInvalidOutline",
	ObReachMaxConcurrentNum:                            "ObReachMaxConcurrentNum",
	ObErrOperationOnRecycleObject:                      "ObErrOperationOnRecycleObject",
	ObErrObjectNotInRecyclebin:                         "ObErrObjectNotInRecyclebin",
	ObErrConCountError:                                 "ObErrConCountError",
	ObErrOutlineContentExist:                           "ObErrOutlineContentExist",
	ObErrOutlineMaxConcurrentExist:                     "ObErrOutlineMaxConcurrentExist",
	ObErrValuesIsNotIntTypeError:                       "ObErrValuesIsNotIntTypeError",
	ObErrWrongTypeColumnValueError:                     "ObErrWrongTypeColumnValueError",
	ObErrPartitionColumnListError:                      "ObErrPartitionColumnListError",
	ObErrTooManyValuesError:                            "ObErrTooManyValuesError",
	ObErrPartitionValueError:                           "ObErrPartitionValueError",
	ObErrPartitionIntervalError:                        "ObErrPartitionIntervalError",
	ObErrSameNamePartition:                             "ObErrSameNamePartition",
	ObErrRangeNotIncreasingError:                       "ObErrRangeNotIncreasingError",
	ObErrParsePartitionRange:                           "ObErrParsePartitionRange",
	ObErrUniqueKeyNeedAllFieldsInPf:                    "ObErrUniqueKeyNeedAllFieldsInPf",
	ObNoPartitionForGivenValue:                         "ObNoPartitionForGivenValue",
	ObEerNullInValuesLessThan:                          "ObEerNullInValuesLessThan",
	ObErrPartitionConstDomainError:                     "ObErrPartitionConstDomainError",
	ObErrTooManyPartitionFuncFields:                    "ObErrTooManyPartitionFuncFields",
	ObErrBadFtColumn:                                   "ObErrBadFtColumn",
	ObErrKeyDoesNotExists:                              "ObErrKeyDoesNotExists",
	ObNonDefaultValueForGeneratedColumn:                "ObNonDefaultValueForGeneratedColumn",
	ObErrBadCtxcatColumn:                               "ObErrBadCtxcatColumn",
	ObErrUnsupportedActionOnGeneratedColumn:            "ObErrUnsupportedActionOnGeneratedColumn",
	ObErrDependentByGeneratedColumn:                    "ObErrDependentByGeneratedColumn",
	ObErrTooManyRows:                                   "ObErrTooManyRows",
	ObWrongFieldTerminators:                            "ObWrongFieldTerminators",
	ObNoReadableReplica:                                "ObNoReadableReplica",
	ObErrSynonymExist:                                  "ObErrSynonymExist",
	ObSynonymNotExist:                                  "ObSynonymNotExist",
	ObErrMissOrderByExpr:                               "ObErrMissOrderByExpr",
	ObErrNotConstExpr:                                  "ObErrNotConstExpr",
	ObErrPartitionMgmtOnNonpartitioned:                 "ObErrPartitionMgmtOnNonpartitioned",
	ObErrDropPartitionNonExistent:                      "ObErrDropPartitionNonExistent",
	ObErrPartitionMgmtOnTwopartTable:                   "ObErrPartitionMgmtOnTwopartTable",
	ObErrOnlyOnRangeListPartition:                      "ObErrOnlyOnRangeListPartition",
	ObErrDropLastPartition:                             "ObErrDropLastPartition",
	ObErrParallelServersTargetNotEnough:                "ObErrParallelServersTargetNotEnough",
	ObErrIgnoreUserHostName:                            "ObErrIgnoreUserHostName",
	ObIgnoreSqlInRestore:                               "ObIgnoreSqlInRestore",
	ObErrTemporaryTableWithPartition:                   "ObErrTemporaryTableWithPartition",
	ObErrInvalidColumnId:                               "ObErrInvalidColumnId",
	ObSyncDdlDuplicate:                                 "ObSyncDdlDuplicate",
	ObSyncDdlError:                                     "ObSyncDdlError",
	ObErrRowIsReferenced:                               "ObErrRowIsReferenced",
	ObErrNoReferencedRow:                               "ObErrNoReferencedRow",
	ObErrFuncResultTooLarge:                            "ObErrFuncResultTooLarge",
	ObErrCannotAddForeign:                              "ObErrCannotAddForeign",
	ObErrWrongFkDef:                                    "ObErrWrongFkDef",
	ObErrInvalidChildColumnLengthFk:                    "ObErrInvalidChildColumnLengthFk",
	ObErrAlterColumnFk:                                 "ObErrAlterColumnFk",
	ObErrConnectByRequired:                             "ObErrConnectByRequired",
	ObErrInvalidPseudoColumnPlace:                      "ObErrInvalidPseudoColumnPlace",
	ObErrNocycleRequired:                               "ObErrNocycleRequired",
	ObErrConnectByLoop:                                 "ObErrConnectByLoop",
	ObErrInvalidSiblings:                               "ObErrInvalidSiblings",
	ObErrInvalidSeparator:                              "ObErrInvalidSeparator",
	ObErrInvalidSynonymName:                            "ObErrInvalidSynonymName",
	ObErrLoopOfSynonym:                                 "ObErrLoopOfSynonym",
	ObErrSynonymSameAsObject:                           "ObErrSynonymSameAsObject",
	ObErrSynonymTranslationInvalid:                     "ObErrSynonymTranslationInvalid",
	ObErrExistObject:                                   "ObErrExistObject",
	ObErrIllegalValueForType:                           "ObErrIllegalValueForType",
	ObErTooLongSetEnumValue:                            "ObErTooLongSetEnumValue",
	ObErDuplicatedValueInType:                          "ObErDuplicatedValueInType",
	ObErTooBigEnum:                                     "ObErTooBigEnum",
	ObErrTooBigSet:                                     "ObErrTooBigSet",
	ObErrWrongRowId:                                    "ObErrWrongRowId",
	ObErrInvalidWindowFunctionPlace:                    "ObErrInvalidWindowFunctionPlace",
	ObErrParsePartitionList:                            "ObErrParsePartitionList",
	ObErrMultipleDefConstInListPart:                    "ObErrMultipleDefConstInListPart",
	ObErrWrongFuncArgumentsType:                        "ObErrWrongFuncArgumentsType",
	ObErrMultiUpdateKeyConflict:                        "ObErrMultiUpdateKeyConflict",
	ObErrInsufficientPxWorker:                          "ObErrInsufficientPxWorker",
	ObErrForUpdateExprNotAllowed:                       "ObErrForUpdateExprNotAllowed",
	ObErrWinFuncArgNotInPartitionBy:                    "ObErrWinFuncArgNotInPartitionBy",
	ObErrTooLongStringInConcat:                         "ObErrTooLongStringInConcat",
	ObErrWrongTimestampLtzColumnValueError:             "ObErrWrongTimestampLtzColumnValueError",
	ObErrUpdCausePartChange:                            "ObErrUpdCausePartChange",
	ObErrInvalidTypeForArgument:                        "ObErrInvalidTypeForArgument",
	ObErrAddPartBounNotInc:                             "ObErrAddPartBounNotInc",
	ObErrDataTooLongInPartCheck:                        "ObErrDataTooLongInPartCheck",
	ObErrWrongTypeColumnValueV2Error:                   "ObErrWrongTypeColumnValueV2Error",
	ObCantAggregate_3collations:                        "ObCantAggregate_3collations",
	ObCantAggregateNcollations:                         "ObCantAggregateNcollations",
	ObErrDuplicatedUniqueKey:                           "ObErrDuplicatedUniqueKey",
	ObDoubleOverflow:                                   "ObDoubleOverflow",
	ObErrNoSysPrivilege:                                "ObErrNoSysPrivilege",
	ObErrNoLoginPrivilege:                              "ObErrNoLoginPrivilege",
	ObErrCannotRevokePrivilegesYouDidNotGrant:          "ObErrCannotRevokePrivilegesYouDidNotGrant",
	ObErrSystemPrivilegesNotGrantedTo:                  "ObErrSystemPrivilegesNotGrantedTo",
	ObErrOnlySelectAndAlterPrivilegesAreValidForSequences:          "ObErrOnlySelectAndAlterPrivilegesAreValidForSequences",
	ObErrExecutePrivilegeNotAllowedForTables:                       "ObErrExecutePrivilegeNotAllowedForTables",
	ObErrOnlyExecuteAndDebugPrivilegesAreValidForProcedures:        "ObErrOnlyExecuteAndDebugPrivilegesAreValidForProcedures",
	ObErrOnlyExecuteDebugAndUnderPrivilegesAreValidForTypes:        "ObErrOnlyExecuteDebugAndUnderPrivilegesAreValidForTypes",
	ObErrAdminOptionNotGrantedForRole:                              "ObErrAdminOptionNotGrantedForRole",
	ObErrUserOrRoleDoesNotExist:                                    "ObErrUserOrRoleDoesNotExist",
	ObErrMissingOnKeyword:                                          "ObErrMissingOnKeyword",
	ObErrNoGrantOption:                                             "ObErrNoGrantOption",
	ObErrAlterIndexAndExecuteNotAllowedForViews:                    "ObErrAlterIndexAndExecuteNotAllowedForViews",
	ObErrCircularRoleGrantDetected:                                 "ObErrCircularRoleGrantDetected",
	ObErrInvalidPrivilegeOnDirectories:                             "ObErrInvalidPrivilegeOnDirectories",
	ObErrDirectoryAccessDenied:                                     "ObErrDirectoryAccessDenied",
	ObErrMissingOrInvalidRoleName:                                  "ObErrMissingOrInvalidRoleName",
	ObErrRoleNotGrantedOrDoesNotExist:                              "ObErrRoleNotGrantedOrDoesNotExist",
	ObErrDefaultRoleNotGrantedToUser:                               "ObErrDefaultRoleNotGrantedToUser",
	ObErrRoleNotGrantedTo:                                          "ObErrRoleNotGrantedTo",
	ObErrCannotGrantToARoleWithGrantOption:                         "ObErrCannotGrantToARoleWithGrantOption",
	ObErrDuplicateUsernameInList:                                   "ObErrDuplicateUsernameInList",
	ObErrCannotGrantStringToARole:                                  "ObErrCannotGrantStringToARole",
	ObErrCascadeConstraintsMustBeSpecifiedToPerformThisRevoke:      "ObErrCascadeConstraintsMustBeSpecifiedToPerformThisRevoke",
	ObErrYouMayNotRevokePrivilegesFromYourself:                     "ObErrYouMayNotRevokePrivilegesFromYourself",
	ObErrMissErrLogMandatoryColumn:                                 "ObErrMissErrLogMandatoryColumn",
	ObTableDefinitionChanged:                                       "ObTableDefinitionChanged",
	ObErrObjectStringDoesNotExist:                                  "ObErrObjectStringDoesNotExist",
	ObErrResultantDataTypeOfVirtualColumnIsNotSupported:            "ObErrResultantDataTypeOfVirtualColumnIsNotSupported",
	ObErrGetStackedDiagnostics:                                     "ObErrGetStackedDiagnostics",
	ObDdlSchemaVersionNotMatch:                                     "ObDdlSchemaVersionNotMatch",
	ObErrColumnGroupDuplicate:                                      "ObErrColumnGroupDuplicate",
	ObErrReservedSyntax:                                            "ObErrReservedSyntax",
	ObErrInvalidParamToProcedure:                                   "ObErrInvalidParamToProcedure",
	ObErrWrongParametersToNativeFct:                                "ObErrWrongParametersToNativeFct",
	ObErrCteMaxRecursionDepth:                                      "ObErrCteMaxRecursionDepth",
	ObDuplicateObjectNameExist:                                     "ObDuplicateObjectNameExist",
	ObErrRefreshSchemaTooLong:                                      "ObErrRefreshSchemaTooLong",
	ObSqlRetrySpm:                                                  "ObSqlRetrySpm",
	ObOutlineNotReproducible:                                       "ObOutlineNotReproducible",
	ObEerWindowNoChildPartitioning:                                 "ObEerWindowNoChildPartitioning",
	ObEerWindowNoInheritFrame:                                      "ObEerWindowNoInheritFrame",
	ObEerWindowNoRedefineOrderBy:                                   "ObEerWindowNoRedefineOrderBy",
	ObErrInvalidDataTypeReturning:                                  "ObErrInvalidDataTypeReturning",
	ObErrJsonValueNoValue:                                          "ObErrJsonValueNoValue",
	ObErrDefaultValueNotLiteral:                                    "ObErrDefaultValueNotLiteral",
	ObErrJsonSyntaxError:                                           "ObErrJsonSyntaxError",
	ObErrJsonEqualOutsidePredicate:                                 "ObErrJsonEqualOutsidePredicate",
	ObErrWithoutArrWrapper:                                         "ObErrWithoutArrWrapper",
	ObErrJsonPatchInvalid:                                          "ObErrJsonPatchInvalid",
	ObErrOrderSiblingsByNotAllowed:                                 "ObErrOrderSiblingsByNotAllowed",
	ObErrLobTypeNotSorting:                                         "ObErrLobTypeNotSorting",
	ObErrJsonIllegalZeroLengthIdentifierError:                      "ObErrJsonIllegalZeroLengthIdentifierError",
	ObErrNoValueInPassing:                                          "ObErrNoValueInPassing",
	ObErrInvalidColumnSpe:                                          "ObErrInvalidColumnSpe",
	ObErrInputJsonNotBeNull:                                        "ObErrInputJsonNotBeNull",
	ObErrInvalidDataType:                                           "ObErrInvalidDataType",
	ObErrInvalidClause:                                             "ObErrInvalidClause",
	ObErrInvalidCmpOp:                                              "ObErrInvalidCmpOp",
	ObErrInvalidInput:                                              "ObErrInvalidInput",
	ObErrEmptyInputToJsonOperator:                                  "ObErrEmptyInputToJsonOperator",
	ObErrAdditionalIsJson:                                          "ObErrAdditionalIsJson",
	ObErrFunctionInvalidState:                                      "ObErrFunctionInvalidState",
	ObErrMissValue:                                                 "ObErrMissValue",
	ObErrDifferentTypeSelected:                                     "ObErrDifferentTypeSelected",
	ObErrNoValueSelected:                                           "ObErrNoValueSelected",
	ObErrNonTextRetNotsupport:                                      "ObErrNonTextRetNotsupport",
	ObErrPlJsontypeUsage:                                           "ObErrPlJsontypeUsage",
	ObErrNullInput:                                                 "ObErrNullInput",
	ObErrDefaultValueNotMatch:                                      "ObErrDefaultValueNotMatch",
	ObErrConversionFail:                                            "ObErrConversionFail",
	ObErrNotObjRef:                                                 "ObErrNotObjRef",
	ObErrUnsupportTruncateType:                                     "ObErrUnsupportTruncateType",
	ObErrUnimplementJsonFeature:                                    "ObErrUnimplementJsonFeature",
	ObErrUsageKeyword:                                              "ObErrUsageKeyword",
	ObErrInputJsonTable:                                            "ObErrInputJsonTable",
	ObErrBoolCastNumber:                                            "ObErrBoolCastNumber",
	ObErrNestedPathDisjunct:                                        "ObErrNestedPathDisjunct",
	ObErrInvalidVariableInJsonPath:                                 "ObErrInvalidVariableInJsonPath",
	ObErrInvalidDefaultValueProvided:                               "ObErrInvalidDefaultValueProvided",
	ObErrPathExpressionNotLiteral:                                  "ObErrPathExpressionNotLiteral",
	ObErrInvalidArgumentForJsonCall:                                "ObErrInvalidArgumentForJsonCall",
	ObErrSchemaHistoryEmpty:                                        "ObErrSchemaHistoryEmpty",
	ObErrTableNameNotInList:                                        "ObErrTableNameNotInList",
	ObErrDefaultNotAtLastInListPart:                                "ObErrDefaultNotAtLastInListPart",
	ObErrMysqlCharacterSetMismatch:                                 "ObErrMysqlCharacterSetMismatch",
	ObErrRenamePartitionNameDuplicate:                              "ObErrRenamePartitionNameDuplicate",
	ObErrRenameSubpartitionNameDuplicate:                           "ObErrRenameSubpartitionNameDuplicate",
	ObErrInvalidWaitInterval:                                       "ObErrInvalidWaitInterval",
	ObErrFunctionalIndexRefAutoIncrement:                           "ObErrFunctionalIndexRefAutoIncrement",
	ObErrDependentByFunctionalIndex:                                "ObErrDependentByFunctionalIndex",
	ObErrFunctionalIndexOnLob:                                      "ObErrFunctionalIndexOnLob",
	ObErrFunctionalIndexOnField:                                    "ObErrFunctionalIndexOnField",
	ObErrGencolLegitCheckFailed:                                    "ObErrGencolLegitCheckFailed",
	ObErrGroupingFuncWithoutGroupBy:                                "ObErrGroupingFuncWithoutGroupBy",
	ObErrDependentByPartitionFunc:                                  "ObErrDependentByPartitionFunc",
	ObErrViewSelectContainInto:                                     "ObErrViewSelectContainInto",
	ObErrDefaultNotAllowed:                                         "ObErrDefaultNotAllowed",
	ObErrModifyRealcolToGencol:                                     "ObErrModifyRealcolToGencol",
	ObErrModifyTypeOfGencol:                                        "ObErrModifyTypeOfGencol",
	ObErrWindowFrameIllegal:                                        "ObErrWindowFrameIllegal",
	ObErrWindowRangeFrameTemporalType:                              "ObErrWindowRangeFrameTemporalType",
	ObErrWindowRangeFrameNumericType:                               "ObErrWindowRangeFrameNumericType",
	ObErrWindowRangeBoundNotConstant:                               "ObErrWindowRangeBoundNotConstant",
	ObErrDefaultForModifyingViews:                                  "ObErrDefaultForModifyingViews",
	ObErrFkColumnNotNull:                                           "ObErrFkColumnNotNull",
	ObErrUnsupportedFkSetNullOnGeneratedColumn:                     "ObErrUnsupportedFkSetNullOnGeneratedColumn",
	ObJsonProcessingError:                                          "ObJsonProcessingError",
	ObErrTableWithoutAlias:                                         "ObErrTableWithoutAlias",
	ObErrDeprecatedSyntax:                                          "ObErrDeprecatedSyntax",
	ObErrSpAlreadyExists:                                           "ObErrSpAlreadyExists",
	ObErrSpDoesNotExist:                                            "ObErrSpDoesNotExist",
	ObErrSpUndeclaredVar:                                           "ObErrSpUndeclaredVar",
	ObErrSpUndeclaredType:                                          "ObErrSpUndeclaredType",
	ObErrSpCondMismatch:                                            "ObErrSpCondMismatch",
	ObErrSpLilabelMismatch:                                         "ObErrSpLilabelMismatch",
	ObErrSpCursorMismatch:                                          "ObErrSpCursorMismatch",
	ObErrSpDupParam:                                                "ObErrSpDupParam",
	ObErrSpDupVar:                                                  "ObErrSpDupVar",
	ObErrSpDupType:                                                 "ObErrSpDupType",
	ObErrSpDupCondition:                                            "ObErrSpDupCondition",
	ObErrSpDupLabel:                                                "ObErrSpDupLabel",
	ObErrSpDupCursor:                                               "ObErrSpDupCursor",
	ObErrSpInvalidFetchArg:                                         "ObErrSpInvalidFetchArg",
	ObErrSpWrongArgNum:                                             "ObErrSpWrongArgNum",
	ObErrSpUnhandledException:                                      "ObErrSpUnhandledException",
	ObErrSpBadConditionType:                                        "ObErrSpBadConditionType",
	ObErrPackageAlreadyExists:                                      "ObErrPackageAlreadyExists",
	ObErrPackageDoseNotExist:                                       "ObErrPackageDoseNotExist",
	ObEerUnknownStmtHandler:                                        "ObEerUnknownStmtHandler",
	ObErrInvalidWindowFuncUse:                                      "ObErrInvalidWindowFuncUse",
	ObErrConstraintDuplicate:                                       "ObErrConstraintDuplicate",
	ObErrContraintNotFound:                                         "ObErrContraintNotFound",
	ObErrAlterTableAlterDuplicatedIndex:                            "ObErrAlterTableAlterDuplicatedIndex",
	ObEerInvalidArgumentForLogarithm:                               "ObEerInvalidArgumentForLogarithm",
	ObErrReorganizeOutsideRange:                                    "ObErrReorganizeOutsideRange",
	ObErSpRecursionLimit:                                           "ObErSpRecursionLimit",
	ObErUnsupportedPs:                                              "ObErUnsupportedPs",
	ObErStmtNotAllowedInSfOrTrg:                                    "ObErStmtNotAllowedInSfOrTrg",
	ObErSpNoRecursion:                                              "ObErSpNoRecursion",
	ObErSpCaseNotFound:                                             "ObErSpCaseNotFound",
	ObErrInvalidSplitCount:                                         "ObErrInvalidSplitCount",
	ObErrInvalidSplitGrammar:                                       "ObErrInvalidSplitGrammar",
	ObErrMissValues:                                                "ObErrMissValues",
	ObErrMissAtValues:                                              "ObErrMissAtValues",
	ObErCommitNotAllowedInSfOrTrg:                                  "ObErCommitNotAllowedInSfOrTrg",
	ObPcGetLocationError:                                           "ObPcGetLocationError",
	ObPcLockConflict:                                               "ObPcLockConflict",
	ObErSpNoRetset:                                                 "ObErSpNoRetset",
	ObErSpNoreturnend:                                              "ObErSpNoreturnend",
	ObErrSpDupHandler:                                              "ObErrSpDupHandler",
	ObErSpNoRecursiveCreate:                                        "ObErSpNoRecursiveCreate",
	ObErSpBadreturn:                                                "ObErSpBadreturn",
	ObErSpBadCursorSelect:                                          "ObErSpBadCursorSelect",
	ObErSpBadSqlstate:                                              "ObErSpBadSqlstate",
	ObErSpVarcondAfterCurshndlr:                                    "ObErSpVarcondAfterCurshndlr",
	ObErSpCursorAfterHandler:                                       "ObErSpCursorAfterHandler",
	ObErSpWrongName:                                                "ObErSpWrongName",
	ObErSpCursorAlreadyOpen:                                        "ObErSpCursorAlreadyOpen",
	ObErSpCursorNotOpen:                                            "ObErSpCursorNotOpen",
	ObErSpCantSetAutocommit:                                        "ObErSpCantSetAutocommit",
	ObErSpNotVarArg:                                                "ObErSpNotVarArg",
	ObErSpLilabelMismatch:                                          "ObErSpLilabelMismatch",
	ObErrTruncateIllegalFk:                                         "ObErrTruncateIllegalFk",
	ObErrDupKey:                                                    "ObErrDupKey",
	ObErInvalidUseOfNull:                                           "ObErInvalidUseOfNull",
	ObErrSplitListLessValue:                                        "ObErrSplitListLessValue",
	ObErrAddPartitionToDefaultList:                                 "ObErrAddPartitionToDefaultList",
	ObErrSplitIntoOnePartition:                                     "ObErrSplitIntoOnePartition",
	ObErrNoTenantPrivilege:                                         "ObErrNoTenantPrivilege",
	ObErrInvalidPercentage:                                         "ObErrInvalidPercentage",
	ObErrCollectHistogram:                                          "ObErrCollectHistogram",
	ObErTempTableInUse:                                             "ObErTempTableInUse",
	ObErrInvalidNlsParameterString:                                 "ObErrInvalidNlsParameterString",
	ObErrDatetimeIntervalPrecisionOutOfRange:                       "ObErrDatetimeIntervalPrecisionOutOfRange",
	ObErrCmdNotProperlyEnded:                                       "ObErrCmdNotProperlyEnded",
	ObErrInvalidNumberFormatModel:                                  "ObErrInvalidNumberFormatModel",
	ObWarnNonAsciiSeparatorNotImplemented:                          "ObWarnNonAsciiSeparatorNotImplemented",
	ObWarnAmbiguousFieldTerm:                                       "ObWarnAmbiguousFieldTerm",
	ObWarnTooFewRecords:                                            "ObWarnTooFewRecords",
	ObWarnTooManyRecords:                                           "ObWarnTooManyRecords",
	ObErrTooManyValues:                                             "ObErrTooManyValues",
	ObErrNotEnoughValues:                                           "ObErrNotEnoughValues",
	ObErrMoreThanOneRow:                                            "ObErrMoreThanOneRow",
	ObErrNotSubQuery:                                               "ObErrNotSubQuery",
	ObInappropriateInto:                                            "ObInappropriateInto",
	ObErrTableIsReferenced:                                         "ObErrTableIsReferenced",
	ObErrQualifierExistsForUsingColumn:                             "ObErrQualifierExistsForUsingColumn",
	ObErrOuterJoinNested:                                           "ObErrOuterJoinNested",
	ObErrMultiOuterJoinTable:                                       "ObErrMultiOuterJoinTable",
	ObErrOuterJoinOnCorrelationColumn:                              "ObErrOuterJoinOnCorrelationColumn",
	ObErrOuterJoinAmbiguous:                                        "ObErrOuterJoinAmbiguous",
	ObErrOuterJoinWithSubQuery:                                     "ObErrOuterJoinWithSubQuery",
	ObErrOuterJoinWithAnsiJoin:                                     "ObErrOuterJoinWithAnsiJoin",
	ObErrOuterJoinNotAllowed:                                       "ObErrOuterJoinNotAllowed",
	ObSchemaEagain:                                                 "ObSchemaEagain",
	ObErrZeroLenCol:                                                "ObErrZeroLenCol",
	ObErrInvalidCharFollowingEscapeChar:                            "ObErrInvalidCharFollowingEscapeChar",
	ObErrInvalidEscapeCharLength:                                   "ObErrInvalidEscapeCharLength",
	ObErrNotSelectedExpr:                                           "ObErrNotSelectedExpr",
	ObErrUkPkDuplicate:                                             "ObErrUkPkDuplicate",
	ObErrColumnListAlreadyIndexed:                                  "ObErrColumnListAlreadyIndexed",
	ObErrBushyTreeNotSupported:                                     "ObErrBushyTreeNotSupported",
	ObErrOrderByItemNotInSelectList:                                "ObErrOrderByItemNotInSelectList",
	ObErrNumericOrValueError:                                       "ObErrNumericOrValueError",
	ObErrConstraintNameDuplicate:                                   "ObErrConstraintNameDuplicate",
	ObErrOnlyHaveInvisibleColInTable:                               "ObErrOnlyHaveInvisibleColInTable",
	ObErrInvisibleColOnUnsupportedTableType:                        "ObErrInvisibleColOnUnsupportedTableType",
	ObErrModifyColVisibilityCombinedWithOtherOption:                "ObErrModifyColVisibilityCombinedWithOtherOption",
	ObErrModifyColVisibilityBySysUser:                              "ObErrModifyColVisibilityBySysUser",
	ObErrTooManyArgsForFun:                                         "ObErrTooManyArgsForFun",
	ObPxSqlNeedRetry:                                               "ObPxSqlNeedRetry",
	ObTenantHasBeenDropped:                                         "ObTenantHasBeenDropped",
	ObErrExtractFieldInvalid:                                       "ObErrExtractFieldInvalid",
	ObErrPackageCompileError:                                       "ObErrPackageCompileError",
	ObErrSpEmptyBlock:                                              "ObErrSpEmptyBlock",
	ObArrayBindingRollback:                                         "ObArrayBindingRollback",
	ObErrInvalidSubQueryUse:                                        "ObErrInvalidSubQueryUse",
	ObErrDateOrSysVarCannotInCheckCst:                              "ObErrDateOrSysVarCannotInCheckCst",
	ObErrNonexistentConstraint:                                     "ObErrNonexistentConstraint",
	ObErrCheckConstraintViolated:                                   "ObErrCheckConstraintViolated",
	ObErrGroupFuncNotAllowed:                                       "ObErrGroupFuncNotAllowed",
	ObErrPolicyStringNotFound:                                      "ObErrPolicyStringNotFound",
	ObErrInvalidLabelString:                                        "ObErrInvalidLabelString",
	ObErrUndefinedCompartmentStringForPolicyString:                 "ObErrUndefinedCompartmentStringForPolicyString",
	ObErrUndefinedLevelStringForPolicyString:                       "ObErrUndefinedLevelStringForPolicyString",
	ObErrUndefinedGroupStringForPolicyString:                       "ObErrUndefinedGroupStringForPolicyString",
	ObErrLbacError:                                                 "ObErrLbacError",
	ObErrPolicyRoleAlreadyExistsForPolicyString:                    "ObErrPolicyRoleAlreadyExistsForPolicyString",
	ObErrNullOrInvalidUserLabel:                                    "ObErrNullOrInvalidUserLabel",
	ObErrAddIndex:                                                  "ObErrAddIndex",
	ObErrProfileStringDoesNotExist:                                 "ObErrProfileStringDoesNotExist",
	ObErrInvalidResourceLimit:                                      "ObErrInvalidResourceLimit",
	ObErrProfileStringAlreadyExists:                                "ObErrProfileStringAlreadyExists",
	ObErrProfileStringHasUsersAssigned:                             "ObErrProfileStringHasUsersAssigned",
	ObErrAddCheckConstraintViolated:                                "ObErrAddCheckConstraintViolated",
	ObErrIllegalViewUpdate:                                         "ObErrIllegalViewUpdate",
	ObErrVirtualColNotAllowed:                                      "ObErrVirtualColNotAllowed",
	ObErrOViewMultiupdate:                                          "ObErrOViewMultiupdate",
	ObErrNonInsertableTable:                                        "ObErrNonInsertableTable",
	ObErrViewMultiupdate:                                           "ObErrViewMultiupdate",
	ObErrNonupdateableColumn:                                       "ObErrNonupdateableColumn",
	ObErrViewDeleteMergeView:                                       "ObErrViewDeleteMergeView",
	ObErrODeleteViewNonKeyPreserved:                                "ObErrODeleteViewNonKeyPreserved",
	ObErrOUpdateViewNonKeyPreserved:                                "ObErrOUpdateViewNonKeyPreserved",
	ObErrModifyReadOnlyView:                                        "ObErrModifyReadOnlyView",
	ObErrInvalidInitransValue:                                      "ObErrInvalidInitransValue",
	ObErrInvalidMaxtransValue:                                      "ObErrInvalidMaxtransValue",
	ObErrInvalidPctfreeOrPctusedValue:                              "ObErrInvalidPctfreeOrPctusedValue",
	ObErrProxyReroute:                                              "ObErrProxyReroute",
	ObErrIllegalArgumentForFunction:                                "ObErrIllegalArgumentForFunction",
	ObErrInvalidSamplingRange:                                      "ObErrInvalidSamplingRange",
	ObErrSpecifyDatabaseNotAllowed:                                 "ObErrSpecifyDatabaseNotAllowed",
	ObErrStmtTriggerWithWhenClause:                                 "ObErrStmtTriggerWithWhenClause",
	ObErrTriggerNotExist:                                           "ObErrTriggerNotExist",
	ObErrTriggerAlreadyExist:                                       "ObErrTriggerAlreadyExist",
	ObErrTriggerExistOnOtherTable:                                  "ObErrTriggerExistOnOtherTable",
	ObErrSignaledInParallelQueryServer:                             "ObErrSignaledInParallelQueryServer",
	ObErrCteIllegalQueryName:                                       "ObErrCteIllegalQueryName",
	ObErrCteUnsupportedColumnAliasing:                              "ObErrCteUnsupportedColumnAliasing",
	ObErrUnsupportedUseOfCte:                                       "ObErrUnsupportedUseOfCte",
	ObErrCteColumnNumberNotMatch:                                   "ObErrCteColumnNumberNotMatch",
	ObErrNeedColumnAliasListInRecursiveCte:                         "ObErrNeedColumnAliasListInRecursiveCte",
	ObErrNeedUnionAllInRecursiveCte:                                "ObErrNeedUnionAllInRecursiveCte",
	ObErrNeedOnlyTwoBranchInRecursiveCte:                           "ObErrNeedOnlyTwoBranchInRecursiveCte",
	ObErrNeedReferenceItselfDirectlyInRecursiveCte:                 "ObErrNeedReferenceItselfDirectlyInRecursiveCte",
	ObErrNeedInitBranchInRecursiveCte:                              "ObErrNeedInitBranchInRecursiveCte",
	ObErrCycleFoundInRecursiveCte:                                  "ObErrCycleFoundInRecursiveCte",
	ObErrCteReachMaxLevelRecursion:                                 "ObErrCteReachMaxLevelRecursion",
	ObErrCteIllegalSearchPseudoName:                                "ObErrCteIllegalSearchPseudoName",
	ObErrCteIllegalCycleNonCycleValue:                              "ObErrCteIllegalCycleNonCycleValue",
	ObErrCteIllegalCyclePseudoName:                                 "ObErrCteIllegalCyclePseudoName",
	ObErrCteColumnAliasDuplicate:                                   "ObErrCteColumnAliasDuplicate",
	ObErrCteIllegalSearchCycleClause:                               "ObErrCteIllegalSearchCycleClause",
	ObErrCteDuplicateCycleNonCycleValue:                            "ObErrCteDuplicateCycleNonCycleValue",
	ObErrCteDuplicateSeqNameCycleColumn:                            "ObErrCteDuplicateSeqNameCycleColumn",
	ObErrCteDuplicateNameInSearchClause:                            "ObErrCteDuplicateNameInSearchClause",
	ObErrCteDuplicateNameInCycleClause:                             "ObErrCteDuplicateNameInCycleClause",
	ObErrCteIllegalColumnInCycleClause:                             "ObErrCteIllegalColumnInCycleClause",
	ObErrCteIllegalRecursiveBranch:                                 "ObErrCteIllegalRecursiveBranch",
	ObErrIllegalJoinInRecursiveCte:                                 "ObErrIllegalJoinInRecursiveCte",
	ObErrCteNeedColumnAliasList:                                    "ObErrCteNeedColumnAliasList",
	ObErrCteIllegalColumnInSerachCaluse:                            "ObErrCteIllegalColumnInSerachCaluse",
	ObErrCteRecursiveQueryNameReferencedMoreThanOnce:               "ObErrCteRecursiveQueryNameReferencedMoreThanOnce",
	ObErrCbyPseudoColumnNotAllowed:                                 "ObErrCbyPseudoColumnNotAllowed",
	ObErrCbyLoop:                                                   "ObErrCbyLoop",
	ObErrCbyJoinNotAllowed:                                         "ObErrCbyJoinNotAllowed",
	ObErrCbyConnectByRequired:                                      "ObErrCbyConnectByRequired",
	ObErrCbyConnectByPathNotAllowed:                                "ObErrCbyConnectByPathNotAllowed",
	ObErrCbyConnectByPathIllegalParam:                              "ObErrCbyConnectByPathIllegalParam",
	ObErrCbyConnectByPathInvalidSeparator:                          "ObErrCbyConnectByPathInvalidSeparator",
	ObErrCbyConnectByRootIllegalUsed:                               "ObErrCbyConnectByRootIllegalUsed",
	ObErrCbyOrederSiblingsByNotAllowed:                             "ObErrCbyOrederSiblingsByNotAllowed",
	ObErrCbyNocycleRequired:                                        "ObErrCbyNocycleRequired",
	ObErrNotEnoughArgsForFun:                                       "ObErrNotEnoughArgsForFun",
	ObErrPrepareStmtChecksum:                                       "ObErrPrepareStmtChecksum",
	ObErrEnableNonexistentConstraint:                               "ObErrEnableNonexistentConstraint",
	ObErrDisableNonexistentConstraint:                              "ObErrDisableNonexistentConstraint",
	ObErrDowngradeDop:                                              "ObErrDowngradeDop",
	ObErrDowngradeParallelMaxServers:                               "ObErrDowngradeParallelMaxServers",
	ObErrOrphanedChildRecordExists:                                 "ObErrOrphanedChildRecordExists",
	ObErrColCheckCstReferAnotherCol:                                "ObErrColCheckCstReferAnotherCol",
	ObBatchedMultiStmtRollback:                                     "ObBatchedMultiStmtRollback",
	ObErrForUpdateSelectViewCannot:                                 "ObErrForUpdateSelectViewCannot",
	ObErrPolicyWithCheckOptionViolation:                            "ObErrPolicyWithCheckOptionViolation",
	ObErrPolicyAlreadyAppliedToTable:                               "ObErrPolicyAlreadyAppliedToTable",
	ObErrMutatingTableOperation:                                    "ObErrMutatingTableOperation",
	ObErrModifyOrDropMultiColumnConstraint:                         "ObErrModifyOrDropMultiColumnConstraint",
	ObErrDropParentKeyColumn:                                       "ObErrDropParentKeyColumn",
	ObAutoincServiceBusy:                                           "ObAutoincServiceBusy",
	ObErrConstraintConstraintDisableValidate:                       "ObErrConstraintConstraintDisableValidate",
	ObErrAutonomousTransactionRollback:                             "ObErrAutonomousTransactionRollback",
	ObOrderbyClauseNotAllowed:                                      "ObOrderbyClauseNotAllowed",
	ObDistinctNotAllowed:                                           "ObDistinctNotAllowed",
	ObErrAssignUserVariableNotAllowed:                              "ObErrAssignUserVariableNotAllowed",
	ObErrModifyNonexistentConstraint:                               "ObErrModifyNonexistentConstraint",
	ObErrSpExceptionHandleIllegal:                                  "ObErrSpExceptionHandleIllegal",
	ObErrInvalidInsertColumn:                                       "ObErrInvalidInsertColumn",
	ObIncorrectUseOfOperator:                                       "ObIncorrectUseOfOperator",
	ObErrNonConstExprIsNotAllowedForPivotUnpivotValues:             "ObErrNonConstExprIsNotAllowedForPivotUnpivotValues",
	ObErrExpectAggregateFunctionInsidePivotOperation:               "ObErrExpectAggregateFunctionInsidePivotOperation",
	ObErrExpNeedSameDatatype:                                       "ObErrExpNeedSameDatatype",
	ObErrCharacterSetMismatch:                                      "ObErrCharacterSetMismatch",
	ObErrRegexpNomatch:                                             "ObErrRegexpNomatch",
	ObErrRegexpBadpat:                                              "ObErrRegexpBadpat",
	ObErrRegexpEescape:                                             "ObErrRegexpEescape",
	ObErrRegexpEbrack:                                              "ObErrRegexpEbrack",
	ObErrRegexpEparen:                                              "ObErrRegexpEparen",
	ObErrRegexpEsubreg:                                             "ObErrRegexpEsubreg",
	ObErrRegexpErange:                                              "ObErrRegexpErange",
	ObErrRegexpEctype:                                              "ObErrRegexpEctype",
	ObErrRegexpEcollate:                                            "ObErrRegexpEcollate",
	ObErrRegexpEbrace:                                              "ObErrRegexpEbrace",
	ObErrRegexpBadbr:                                               "ObErrRegexpBadbr",
	ObErrRegexpBadrpt:                                              "ObErrRegexpBadrpt",
	ObErrRegexpAssert:                                              "ObErrRegexpAssert",
	ObErrRegexpInvarg:                                              "ObErrRegexpInvarg",
	ObErrRegexpMixed:                                               "ObErrRegexpMixed",
	ObErrRegexpBadopt:                                              "ObErrRegexpBadopt",
	ObErrRegexpEtoobig:                                             "ObErrRegexpEtoobig",
	ObNotSupportedRowIdType:                                        "ObNotSupportedRowIdType",
	ObErrParallelDdlConflict:                                       "ObErrParallelDdlConflict",
	ObErrSubscriptBeyondCount:                                      "ObErrSubscriptBeyondCount",
	ObErrNotPartitioned:                                            "ObErrNotPartitioned",
	ObUnknownSubpartition:                                          "ObUnknownSubpartition",
	ObErrInvalidSqlRowLimiting:                                     "ObErrInvalidSqlRowLimiting",
	INCORRECT_ARGUMENTS_TO_ESCAPE:                                  " INCORRECT_ARGUMENTS_TO_ESCAPE",
	STATIC_ENG_NOT_IMPLEMENT:                                       " STATIC_ENG_NOT_IMPLEMENT",
	ObObjAlreadyExist:                                              "ObObjAlreadyExist",
	ObDblinkNotExistToAccess:                                       "ObDblinkNotExistToAccess",
	ObDblinkNotExistToDrop:                                         "ObDblinkNotExistToDrop",
	ObErrAccessIntoNull:                                            "ObErrAccessIntoNull",
	ObErrCollecionNull:                                             "ObErrCollecionNull",
	ObErrNoDataNeeded:                                              "ObErrNoDataNeeded",
	ObErrProgramError:                                              "ObErrProgramError",
	ObErrRowtypeMismatch:                                           "ObErrRowtypeMismatch",
	ObErrStorageError:                                              "ObErrStorageError",
	ObErrSubscriptOutsideLimit:                                     "ObErrSubscriptOutsideLimit",
	ObErrInvalidCursor:                                             "ObErrInvalidCursor",
	ObErrLoginDenied:                                               "ObErrLoginDenied",
	ObErrNotLoggedOn:                                               "ObErrNotLoggedOn",
	ObErrSelfIsNull:                                                "ObErrSelfIsNull",
	ObErrTimeoutOnResource:                                         "ObErrTimeoutOnResource",
	ObColumnCantChangeToNotNull:                                    "ObColumnCantChangeToNotNull",
	ObColumnCantChangeToNullale:                                    "ObColumnCantChangeToNullale",
	ObEnableNotNullConstraintViolated:                              "ObEnableNotNullConstraintViolated",
	ObErrArgumentShouldConstant:                                    "ObErrArgumentShouldConstant",
	ObErrNotASingleGroupFunction:                                   "ObErrNotASingleGroupFunction",
	ObErrZeroLengthIdentifier:                                      "ObErrZeroLengthIdentifier",
	ObErrParamValueInvalid:                                         "ObErrParamValueInvalid",
	ObErrDbmsSqlCursorNotExist:                                     "ObErrDbmsSqlCursorNotExist",
	ObErrDbmsSqlNotAllVarBind:                                      "ObErrDbmsSqlNotAllVarBind",
	ObErrConflictingDeclarations:                                   "ObErrConflictingDeclarations",
	ObErrDropColReferencedMultiColsConstraint:                      "ObErrDropColReferencedMultiColsConstraint",
	ObErrModifyColDatatyepReferencedConstraint:                     "ObErrModifyColDatatyepReferencedConstraint",
	ObErrPercentileValueInvalid:                                    "ObErrPercentileValueInvalid",
	ObErrArgumentShouldNumericDateDatetimeType:                     "ObErrArgumentShouldNumericDateDatetimeType",
	ObErrAlterTableRenameWithOption:                                "ObErrAlterTableRenameWithOption",
	ObErrOnlySimpleColumnNameAllowed:                               "ObErrOnlySimpleColumnNameAllowed",
	ObErrSafeUpdateModeNeedWhereOrLimit:                            "ObErrSafeUpdateModeNeedWhereOrLimit",
	ObErrSpecifiyPartitionDescription:                              "ObErrSpecifiyPartitionDescription",
	ObErrSameNameSubpartition:                                      "ObErrSameNameSubpartition",
	ObErrUpdateOrderBy:                                             "ObErrUpdateOrderBy",
	ObErrUpdateLimit:                                               "ObErrUpdateLimit",
	ObRowIdTypeMismatch:                                            "ObRowIdTypeMismatch",
	ObRowIdNumMismatch:                                             "ObRowIdNumMismatch",
	ObNoColumnAlias:                                                "ObNoColumnAlias",
	ObErrInvalidDatatype:                                           "ObErrInvalidDatatype",
	ObErrNotCompositePartition:                                     "ObErrNotCompositePartition",
	ObErrSubpartitionNotExpectValuesIn:                             "ObErrSubpartitionNotExpectValuesIn",
	ObErrSubpartitionExpectValuesIn:                                "ObErrSubpartitionExpectValuesIn",
	ObErrPartitionNotExpectValuesLessThan:                          "ObErrPartitionNotExpectValuesLessThan",
	ObErrPartitionExpectValuesLessThan:                             "ObErrPartitionExpectValuesLessThan",
	ObErrProgramUnitNotExist:                                       "ObErrProgramUnitNotExist",
	ObErrInvalidRestorePointName:                                   "ObErrInvalidRestorePointName",
	ObErrInputTimeType:                                             "ObErrInputTimeType",
	ObErrInArrayDml:                                                "ObErrInArrayDml",
	ObErrTriggerCompileError:                                       "ObErrTriggerCompileError",
	ObErrInTrimSet:                                                 "ObErrInTrimSet",
	ObErrMissingOrInvalidPasswordForRole:                           "ObErrMissingOrInvalidPasswordForRole",
	ObErrMissingOrInvalidPassword:                                  "ObErrMissingOrInvalidPassword",
	ObErrNoOptionsForAlterUser:                                     "ObErrNoOptionsForAlterUser",
	ObErrNoMatchingUkPkForColList:                                  "ObErrNoMatchingUkPkForColList",
	ObErrDupFkInTable:                                              "ObErrDupFkInTable",
	ObErrDupFkExists:                                               "ObErrDupFkExists",
	ObErrMissingOrInvalidPriviege:                                  "ObErrMissingOrInvalidPriviege",
	ObErrInvalidVirtualColumnType:                                  "ObErrInvalidVirtualColumnType",
	ObErrReferencedTableHasNoPk:                                    "ObErrReferencedTableHasNoPk",
	ObErrModifyPartColumnType:                                      "ObErrModifyPartColumnType",
	ObErrModifySubpartColumnType:                                   "ObErrModifySubpartColumnType",
	ObErrDecreaseColumnLength:                                      "ObErrDecreaseColumnLength",
	ObErrRemotePartIllegal:                                         "ObErrRemotePartIllegal",
	ObErrDuplicateColumnExpressionWasSpecified:                     "ObErrDuplicateColumnExpressionWasSpecified",
	ObErrAViewNotAppropriateHere:                                   "ObErrAViewNotAppropriateHere",
	ObRowIdViewNoKeyPreserved:                                      "ObRowIdViewNoKeyPreserved",
	ObRowIdViewHasDistinctEtc:                                      "ObRowIdViewHasDistinctEtc",
	ObErrAtLeastOneColumnNotVirtual:                                "ObErrAtLeastOneColumnNotVirtual",
	ObErrOnlyPureFuncCanbeIndexed:                                  "ObErrOnlyPureFuncCanbeIndexed",
	ObErrOnlyPureFuncCanbeVirtualColumnExpression:                  "ObErrOnlyPureFuncCanbeVirtualColumnExpression",
	ObErrUpdateOperationOnVirtualColumns:                           "ObErrUpdateOperationOnVirtualColumns",
	ObErrInvalidColumnExpression:                                   "ObErrInvalidColumnExpression",
	ObErrIdentityColumnCountExceLimit:                              "ObErrIdentityColumnCountExceLimit",
	ObErrInvalidNotNullConstraintOnIdentityColumn:                  "ObErrInvalidNotNullConstraintOnIdentityColumn",
	ObErrCannotModifyNotNullConstraintOnIdentityColumn:             "ObErrCannotModifyNotNullConstraintOnIdentityColumn",
	ObErrCannotDropNotNullConstraintOnIdentityColumn:               "ObErrCannotDropNotNullConstraintOnIdentityColumn",
	ObErrColumnModifyToIdentityColumn:                              "ObErrColumnModifyToIdentityColumn",
	ObErrIdentityColumnCannotHaveDefaultValue:                      "ObErrIdentityColumnCannotHaveDefaultValue",
	ObErrIdentityColumnMustBeNumericType:                           "ObErrIdentityColumnMustBeNumericType",
	ObErrPrebuiltTableManagedCannotBeIdentityColumn:                "ObErrPrebuiltTableManagedCannotBeIdentityColumn",
	ObErrCannotAlterSystemGeneratedSequence:                        "ObErrCannotAlterSystemGeneratedSequence",
	ObErrCannotDropSystemGeneratedSequence:                         "ObErrCannotDropSystemGeneratedSequence",
	ObErrInsertIntoGeneratedAlwaysIdentityColumn:                   "ObErrInsertIntoGeneratedAlwaysIdentityColumn",
	ObErrUpdateGeneratedAlwaysIdentityColumn:                       "ObErrUpdateGeneratedAlwaysIdentityColumn",
	ObErrIdentityColumnSequenceMismatchAlterTableExchangePartition: "ObErrIdentityColumnSequenceMismatchAlterTableExchangePartition",
	ObErrCannotRenameSystemGeneratedSequence:                       "ObErrCannotRenameSystemGeneratedSequence",
	ObErrRevokeByColumn:                                            "ObErrRevokeByColumn",
	ObErrTypeBodyNotExist:                                          "ObErrTypeBodyNotExist",
	ObErrInvalidArgumentForWidthBucket:                             "ObErrInvalidArgumentForWidthBucket",
	ObErrCbyNoMemory:                                               "ObErrCbyNoMemory",
	ObErrIllegalParamForCbyPath:                                    "ObErrIllegalParamForCbyPath",
	ObErrHostUnknown:                                               "ObErrHostUnknown",
	ObErrWindowNameIsNotDefine:                                     "ObErrWindowNameIsNotDefine",
	ObErrOpenCursorsExceeded:                                       "ObErrOpenCursorsExceeded",
	ObErrFetchOutSequence:                                          "ObErrFetchOutSequence",
	ObErrUnexpectedNameStr:                                         "ObErrUnexpectedNameStr",
	ObErrNoProgramUnit:                                             "ObErrNoProgramUnit",
	ObErrArgInvalid:                                                "ObErrArgInvalid",
	ObErrDbmsStatsPl:                                               "ObErrDbmsStatsPl",
	ObErrIncorrectValueForFunction:                                 "ObErrIncorrectValueForFunction",
	ObErrUnsupportedCharacterSet:                                   "ObErrUnsupportedCharacterSet",
	ObErrMustBeFollowedByFourHexadecimalCharactersOrAnother:        "ObErrMustBeFollowedByFourHexadecimalCharactersOrAnother",
	ObErrParameterTooLong:                                          "ObErrParameterTooLong",
	ObErrInvalidPlsqlCcflags:                                       "ObErrInvalidPlsqlCcflags",
	ObErrRefMutuallyDep:                                            "ObErrRefMutuallyDep",
	ObErrColumnNotAllowed:                                          "ObErrColumnNotAllowed",
	ObErrCannotAccessNlsDataFilesOrInvalidEnvironmentSpecified:     "ObErrCannotAccessNlsDataFilesOrInvalidEnvironmentSpecified",
	ObErrDuplicateNullSpecification:                                "ObErrDuplicateNullSpecification",
	ObErrNotNullConstraintViolated:                                 "ObErrNotNullConstraintViolated",
	ObErrTableAddNotNullColumnNotEmpty:                             "ObErrTableAddNotNullColumnNotEmpty",
	ObErrColumnExpressionModificationWithOtherDdl:                  "ObErrColumnExpressionModificationWithOtherDdl",
	ObErrVirtualColWithConstraintCantBeChanged:                     "ObErrVirtualColWithConstraintCantBeChanged",
	ObErrInvalidNotNullConstraintOnDefaultOnNullIdentityColumn:     "ObErrInvalidNotNullConstraintOnDefaultOnNullIdentityColumn",
	ObErrInvalidDataTypeForAtTimeZone:                              "ObErrInvalidDataTypeForAtTimeZone",
	ObErrBadArg:                                                    "ObErrBadArg",
	ObErrCannotModifyNotNullConstraintOnDefaultOnNullColumn:        "ObErrCannotModifyNotNullConstraintOnDefaultOnNullColumn",
	ObErrCannotDropNotNullConstraintOnDefaultOnNullColumn:          "ObErrCannotDropNotNullConstraintOnDefaultOnNullColumn",
	ObErrInvalidPath:                                               "ObErrInvalidPath",
	ObErrInvalidParamEncountered:                                   "ObErrInvalidParamEncountered",
	ObErrIncorrectMethodUsage:                                      "ObErrIncorrectMethodUsage",
	ObErrTypeMismatch:                                              "ObErrTypeMismatch",
	ObErrFetchColumnNull:                                           "ObErrFetchColumnNull",
	ObErrInvalidSizeSpecified:                                      "ObErrInvalidSizeSpecified",
	ObErrSourceEmpty:                                               "ObErrSourceEmpty",
	ObErrBadValueForObjectType:                                     "ObErrBadValueForObjectType",
	ObErrUnableGetSource:                                           "ObErrUnableGetSource",
	ObErrMissingIdentifier:                                         "ObErrMissingIdentifier",
	ObErrDupCompileParam:                                           "ObErrDupCompileParam",
	ObErrDataNotWellFormat:                                         "ObErrDataNotWellFormat",
	ObErrMustCompositType:                                          "ObErrMustCompositType",
	ObErrUserExceedResource:                                        "ObErrUserExceedResource",
	ObErrUtlEncodeArgumentInvalid:                                  "ObErrUtlEncodeArgumentInvalid",
	ObErrUtlEncodeCharsetInvalid:                                   "ObErrUtlEncodeCharsetInvalid",
	ObErrUtlEncodeMimeHeadTag:                                      "ObErrUtlEncodeMimeHeadTag",
	ObErrCheckOptionViolated:                                       "ObErrCheckOptionViolated",
	ObErrCheckOptionOnNonupdatableView:                             "ObErrCheckOptionOnNonupdatableView",
	ObErrNoDescForPos:                                              "ObErrNoDescForPos",
	ObErrIllObjFlag:                                                "ObErrIllObjFlag",
	ObErrPartitionExtendedOnView:                                   "ObErrPartitionExtendedOnView",
	ObErrNotAllVariableBind:                                        "ObErrNotAllVariableBind",
	ObErrBindVariableNotExist:                                      "ObErrBindVariableNotExist",
	ObErrNotValidRoutineName:                                       "ObErrNotValidRoutineName",
	ObErrDdlInIllegalContext:                                       "ObErrDdlInIllegalContext",
	ObErrCteNeedQueryBlocks:                                        "ObErrCteNeedQueryBlocks",
	ObErrWindowRowsIntervalUse:                                     "ObErrWindowRowsIntervalUse",
	ObErrWindowRangeFrameOrderType:                                 "ObErrWindowRangeFrameOrderType",
	ObErrWindowIllegalOrderBy:                                      "ObErrWindowIllegalOrderBy",
	ObErrMultipleConstraintsWithSameName:                           "ObErrMultipleConstraintsWithSameName",
	ObErrNonBooleanExprForCheckConstraint:                          "ObErrNonBooleanExprForCheckConstraint",
	ObErrCheckConstraintNotFound:                                   "ObErrCheckConstraintNotFound",
	ObErrAlterConstraintEnforcementNotSupported:                    "ObErrAlterConstraintEnforcementNotSupported",
	ObErrCheckConstraintRefersAutoIncrementColumn:                  "ObErrCheckConstraintRefersAutoIncrementColumn",
	ObErrCheckConstraintNamedFunctionIsNotAllowed:                  "ObErrCheckConstraintNamedFunctionIsNotAllowed",
	ObErrCheckConstraintFunctionIsNotAllowed:                       "ObErrCheckConstraintFunctionIsNotAllowed",
	ObErrCheckConstraintVariables:                                  "ObErrCheckConstraintVariables",
	ObErrCheckConstraintRefersUnknownColumn:                        "ObErrCheckConstraintRefersUnknownColumn",
	ObErrUseUdfInPart:                                              "ObErrUseUdfInPart",
	ObErrUseUdfNotDetermin:                                         "ObErrUseUdfNotDetermin",
	ObErrIntervalClauseHasMoreThanOneColumn:                        "ObErrIntervalClauseHasMoreThanOneColumn",
	ObErrInvalidDataTypeIntervalTable:                              "ObErrInvalidDataTypeIntervalTable",
	ObErrIntervalExprNotCorrectType:                                "ObErrIntervalExprNotCorrectType",
	ObErrTableIsAlreadyARangePartitionedTable:                      "ObErrTableIsAlreadyARangePartitionedTable",
	ObTransactionSetViolation:                                      "ObTransactionSetViolation",
	ObTransRollbacked:                                              "ObTransRollbacked",
	ObErrExclusiveLockConflict:                                     "ObErrExclusiveLockConflict",
	ObErrSharedLockConflict:                                        "ObErrSharedLockConflict",
	ObTryLockRowConflict:                                           "ObTryLockRowConflict",
	ObErrExclusiveLockConflictNowait:                               "ObErrExclusiveLockConflictNowait",
	ObClockOutOfOrder:                                              "ObClockOutOfOrder",
	ObTransHasDecided:                                              "ObTransHasDecided",
	ObTransInvalidState:                                            "ObTransInvalidState",
	ObTransStateNotChange:                                          "ObTransStateNotChange",
	ObTransProtocolError:                                           "ObTransProtocolError",
	ObTransInvalidMessage:                                          "ObTransInvalidMessage",
	ObTransInvalidMessageType:                                      "ObTransInvalidMessageType",
	ObPartitionIsFrozen:                                            "ObPartitionIsFrozen",
	ObPartitionIsNotFrozen:                                         "ObPartitionIsNotFrozen",
	ObTransInvalidLogType:                                          "ObTransInvalidLogType",
	ObTransSqlSequenceIllegal:                                      "ObTransSqlSequenceIllegal",
	ObTransCannotBeKilled:                                          "ObTransCannotBeKilled",
	ObTransStateUnknown:                                            "ObTransStateUnknown",
	ObTransIsExiting:                                               "ObTransIsExiting",
	ObTransNeedRollback:                                            "ObTransNeedRollback",
	ObPartitionIsNotStopped:                                        "ObPartitionIsNotStopped",
	ObPartitionIsStopped:                                           "ObPartitionIsStopped",
	ObPartitionIsBlocked:                                           "ObPartitionIsBlocked",
	ObTransRpcTimeout:                                              "ObTransRpcTimeout",
	ObReplicaNotReadable:                                           "ObReplicaNotReadable",
	ObPartitionIsSplitting:                                         "ObPartitionIsSplitting",
	ObTransCommited:                                                "ObTransCommited",
	ObTransCtxCountReachLimit:                                      "ObTransCtxCountReachLimit",
	ObTransCannotSerialize:                                         "ObTransCannotSerialize",
	ObTransWeakReadVersionNotReady:                                 "ObTransWeakReadVersionNotReady",
	ObGtsStandbyIsInvalid:                                          "ObGtsStandbyIsInvalid",
	ObGtsUpdateFailed:                                              "ObGtsUpdateFailed",
	ObGtsIsNotServing:                                              "ObGtsIsNotServing",
	ObPgPartitionNotExist:                                          "ObPgPartitionNotExist",
	ObTransStmtNeedRetry:                                           "ObTransStmtNeedRetry",
	ObSavepointNotExist:                                            "ObSavepointNotExist",
	ObTransWaitSchemaRefresh:                                       "ObTransWaitSchemaRefresh",
	ObTransOutOfThreshold:                                          "ObTransOutOfThreshold",
	ObTransXaNota:                                                  "ObTransXaNota",
	ObTransXaRmfail:                                                "ObTransXaRmfail",
	ObTransXaDupid:                                                 "ObTransXaDupid",
	ObTransXaOutside:                                               "ObTransXaOutside",
	ObTransXaInval:                                                 "ObTransXaInval",
	ObTransXaRmerr:                                                 "ObTransXaRmerr",
	ObTransXaProto:                                                 "ObTransXaProto",
	ObTransXaRbrollback:                                            "ObTransXaRbrollback",
	ObTransXaRbtimeout:                                             "ObTransXaRbtimeout",
	ObTransXaRdonly:                                                "ObTransXaRdonly",
	ObTransXaRetry:                                                 "ObTransXaRetry",
	ObErrRowNotLocked:                                              "ObErrRowNotLocked",
	ObEmptyPg:                                                      "ObEmptyPg",
	ObTransXaErrCommit:                                             "ObTransXaErrCommit",
	ObErrRestorePointExist:                                         "ObErrRestorePointExist",
	ObErrRestorePointNotExist:                                      "ObErrRestorePointNotExist",
	ObErrBackupPointExist:                                          "ObErrBackupPointExist",
	ObErrBackupPointNotExist:                                       "ObErrBackupPointNotExist",
	ObErrRestorePointTooMany:                                       "ObErrRestorePointTooMany",
	ObTransXaBranchFail:                                            "ObTransXaBranchFail",
	ObObjLockNotExist:                                              "ObObjLockNotExist",
	ObObjLockExist:                                                 "ObObjLockExist",
	ObTryLockObjConflict:                                           "ObTryLockObjConflict",
	ObTxNologcb:                                                    "ObTxNologcb",
	ObErrAddPartitionOnInterval:                                    "ObErrAddPartitionOnInterval",
	ObErrMaxvaluePartitionWithInterval:                             "ObErrMaxvaluePartitionWithInterval",
	ObErrInvalidIntervalHighBounds:                                 "ObErrInvalidIntervalHighBounds",
	ObNoPartitionForIntervalPart:                                   "ObNoPartitionForIntervalPart",
	ObErrIntervalCannotBeZero:                                      "ObErrIntervalCannotBeZero",
	ObErrPartitioningKeyMapsToAPartitionOutsideMaximumPermittedNumberOfPartitions: "ObErrPartitioningKeyMapsToAPartitionOutsideMaximumPermittedNumberOfPartitions",
	ObObjLockNotCompleted:                       "ObObjLockNotCompleted",
	ObObjUnlockConflict:                         "ObObjUnlockConflict",
	ObScnOutOfBound:                             "ObScnOutOfBound",
	ObTransIdleTimeout:                          "ObTransIdleTimeout",
	ObTransFreeRouteNotSupported:                "ObTransFreeRouteNotSupported",
	ObTransLiveTooMuchTime:                      "ObTransLiveTooMuchTime",
	ObTransCommitTooMuchTime:                    "ObTransCommitTooMuchTime",
	ObTransTooManyParticipants:                  "ObTransTooManyParticipants",
	ObLogAlreadySplit:                           "ObLogAlreadySplit",
	ObLogIdNotFound:                             "ObLogIdNotFound",
	ObLsrThreadStopped:                          "ObLsrThreadStopped",
	ObNoLog:                                     "ObNoLog",
	ObLogIdRangeError:                           "ObLogIdRangeError",
	ObLogIterEnough:                             "ObLogIterEnough",
	ObClogInvalidAck:                            "ObClogInvalidAck",
	ObClogCacheInvalid:                          "ObClogCacheInvalid",
	ObExtHandleUnfinish:                         "ObExtHandleUnfinish",
	ObCursorNotExist:                            "ObCursorNotExist",
	ObStreamNotExist:                            "ObStreamNotExist",
	ObStreamBusy:                                "ObStreamBusy",
	ObFileRecycled:                              "ObFileRecycled",
	ObReplayEagainTooMuchTime:                   "ObReplayEagainTooMuchTime",
	ObMemberChangeFailed:                        "ObMemberChangeFailed",
	ObNoNeedBatchCtx:                            "ObNoNeedBatchCtx",
	ObTooLargeLogId:                             "ObTooLargeLogId",
	ObAllocLogIdNeedRetry:                       "ObAllocLogIdNeedRetry",
	ObTransOnePcNotAllowed:                      "ObTransOnePcNotAllowed",
	ObLogNeedRebuild:                            "ObLogNeedRebuild",
	ObTooManyLogTask:                            "ObTooManyLogTask",
	ObInvalidBatchSize:                          "ObInvalidBatchSize",
	ObClogSlideTimeout:                          "ObClogSlideTimeout",
	ObLogReplayError:                            "ObLogReplayError",
	ObTryLockConfigChangeConflict:               "ObTryLockConfigChangeConflict",
	ObElectionWarnLogbufFull:                    "ObElectionWarnLogbufFull",
	ObElectionWarnLogbufEmpty:                   "ObElectionWarnLogbufEmpty",
	ObElectionWarnNotRunning:                    "ObElectionWarnNotRunning",
	ObElectionWarnIsRunning:                     "ObElectionWarnIsRunning",
	ObElectionWarnNotReachMajority:              "ObElectionWarnNotReachMajority",
	ObElectionWarnInvalidServer:                 "ObElectionWarnInvalidServer",
	ObElectionWarnInvalidLeader:                 "ObElectionWarnInvalidLeader",
	ObElectionWarnLeaderLeaseExpired:            "ObElectionWarnLeaderLeaseExpired",
	ObElectionWarnInvalidMessage:                "ObElectionWarnInvalidMessage",
	ObElectionWarnMessageNotIntime:              "ObElectionWarnMessageNotIntime",
	ObElectionWarnNotCandidate:                  "ObElectionWarnNotCandidate",
	ObElectionWarnNotCandidateOrVoter:           "ObElectionWarnNotCandidateOrVoter",
	ObElectionWarnProtocolError:                 "ObElectionWarnProtocolError",
	ObElectionWarnRuntimeOutOfRange:             "ObElectionWarnRuntimeOutOfRange",
	ObElectionWarnLastOperationNotDone:          "ObElectionWarnLastOperationNotDone",
	ObElectionWarnCurrentServerNotLeader:        "ObElectionWarnCurrentServerNotLeader",
	ObElectionWarnNoPrepareMessage:              "ObElectionWarnNoPrepareMessage",
	ObElectionErrorMultiPrepareMessage:          "ObElectionErrorMultiPrepareMessage",
	ObElectionNotExist:                          "ObElectionNotExist",
	ObElectionMgrIsRunning:                      "ObElectionMgrIsRunning",
	ObElectionWarnNoMajorityPrepareMessage:      "ObElectionWarnNoMajorityPrepareMessage",
	ObElectionAsyncLogWarnInit:                  "ObElectionAsyncLogWarnInit",
	ObElectionWaitLeaderMessage:                 "ObElectionWaitLeaderMessage",
	ObElectionGroupNotExist:                     "ObElectionGroupNotExist",
	ObUnexpectEgVersion:                         "ObUnexpectEgVersion",
	ObElectionGroupMgrIsRunning:                 "ObElectionGroupMgrIsRunning",
	ObElectionMgrNotRunning:                     "ObElectionMgrNotRunning",
	ObElectionErrorVoteMsgConflict:              "ObElectionErrorVoteMsgConflict",
	ObElectionErrorDuplicatedMsg:                "ObElectionErrorDuplicatedMsg",
	ObElectionWarnT1NotMatch:                    "ObElectionWarnT1NotMatch",
	ObElectionBelowMajority:                     "ObElectionBelowMajority",
	ObElectionOverMajority:                      "ObElectionOverMajority",
	ObElectionDuringUpgrading:                   "ObElectionDuringUpgrading",
	ObTransferTaskCompleted:                     "ObTransferTaskCompleted",
	ObTooManyTransferTask:                       "ObTooManyTransferTask",
	ObTransferTaskExist:                         "ObTransferTaskExist",
	ObTransferTaskNotExist:                      "ObTransferTaskNotExist",
	ObNotAllowToRemove:                          "ObNotAllowToRemove",
	ObRgNotMatch:                                "ObRgNotMatch",
	ObTransferTaskAborted:                       "ObTransferTaskAborted",
	ObTransferInvalidMessage:                    "ObTransferInvalidMessage",
	ObTransferCtxTsNotMatch:                     "ObTransferCtxTsNotMatch",
	ObTransferSysError:                          "ObTransferSysError",
	ObTransferMemberListNotSame:                 "ObTransferMemberListNotSame",
	ObErrUnexpectedLockOwner:                    "ObErrUnexpectedLockOwner",
	ObLsTransferScnTooSmall:                     "ObLsTransferScnTooSmall",
	ObTabletTransferSeqNotMatch:                 "ObTabletTransferSeqNotMatch",
	ObTransferDetectActiveTrans:                 "ObTransferDetectActiveTrans",
	ObTransferSrcLsNotExist:                     "ObTransferSrcLsNotExist",
	ObTransferSrcTabletNotExist:                 "ObTransferSrcTabletNotExist",
	ObLsNeedRebuild:                             "ObLsNeedRebuild",
	ObObsoleteClogNeedSkip:                      "ObObsoleteClogNeedSkip",
	ObTransferWaitTransactionEndTimeout:         "ObTransferWaitTransactionEndTimeout",
	ObTabletGcLockConflict:                      "ObTabletGcLockConflict",
	ObSequenceNotMatch:                          "ObSequenceNotMatch",
	ObSequenceTooSmall:                          "ObSequenceTooSmall",
	ObErrInvalidXmlDatatype:                     "ObErrInvalidXmlDatatype",
	ObErrXmlMissingComma:                        "ObErrXmlMissingComma",
	ObErrInvalidXpathExpression:                 "ObErrInvalidXpathExpression",
	ObErrExtractvalueMultiNodes:                 "ObErrExtractvalueMultiNodes",
	ObErrXmlFramentConvert:                      "ObErrXmlFramentConvert",
	ObInvalidPrintOption:                        "ObInvalidPrintOption",
	ObXmlCharLenTooSmall:                        "ObXmlCharLenTooSmall",
	ObXpathExpressionUnsupported:                "ObXpathExpressionUnsupported",
	ObExtractvalueNotLeafNode:                   "ObExtractvalueNotLeafNode",
	ObXmlInsertFragment:                         "ObXmlInsertFragment",
	ObErrNoOrderMapSql:                          "ObErrNoOrderMapSql",
	ObErrXmlelementAliased:                      "ObErrXmlelementAliased",
	ObInvalidAlterationgDatatype:                "ObInvalidAlterationgDatatype",
	ObInvalidModificationOfColumns:              "ObInvalidModificationOfColumns",
	ObErrNullForXmlConstructor:                  "ObErrNullForXmlConstructor",
	ObErrXmlIndex:                               "ObErrXmlIndex",
	ObErrUpdateXmlWithInvalidNode:               "ObErrUpdateXmlWithInvalidNode",
	ObLobValueNotExist:                          "ObLobValueNotExist",
	ObErrJsonFunUnsupportedType:                 "ObErrJsonFunUnsupportedType",
	ObServerIsInit:                              "ObServerIsInit",
	ObServerIsStopping:                          "ObServerIsStopping",
	ObPacketChecksumError:                       "ObPacketChecksumError",
	ObNotReadAllData:                            "ObNotReadAllData",
	ObBuildMd5Error:                             "ObBuildMd5Error",
	ObMd5NotMatch:                               "ObMd5NotMatch",
	ObOssDataVersionNotMatched:                  "ObOssDataVersionNotMatched",
	ObOssWriteError:                             "ObOssWriteError",
	ObRestoreInProgress:                         "ObRestoreInProgress",
	ObAgentInitingBackupCountError:              "ObAgentInitingBackupCountError",
	ObClusterNameNotEqual:                       "ObClusterNameNotEqual",
	ObRsListInvaild:                             "ObRsListInvaild",
	ObAgentHasFailedTask:                        "ObAgentHasFailedTask",
	ObRestorePartitionIsComplete:                "ObRestorePartitionIsComplete",
	ObRestorePartitionTwice:                     "ObRestorePartitionTwice",
	ObStopDropSchema:                            "ObStopDropSchema",
	ObCannotStartLogArchiveBackup:               "ObCannotStartLogArchiveBackup",
	ObAlreadyNoLogArchiveBackup:                 "ObAlreadyNoLogArchiveBackup",
	ObLogArchiveBackupInfoNotExist:              "ObLogArchiveBackupInfoNotExist",
	ObLogArchiveInterrupted:                     "ObLogArchiveInterrupted",
	ObLogArchiveStatNotMatch:                    "ObLogArchiveStatNotMatch",
	ObLogArchiveNotRunning:                      "ObLogArchiveNotRunning",
	ObLogArchiveInvalidRound:                    "ObLogArchiveInvalidRound",
	ObReplicaCannotBackup:                       "ObReplicaCannotBackup",
	ObBackupInfoNotExist:                        "ObBackupInfoNotExist",
	ObBackupInfoNotMatch:                        "ObBackupInfoNotMatch",
	ObLogArchiveAlreadyStopped:                  "ObLogArchiveAlreadyStopped",
	ObRestoreIndexFailed:                        "ObRestoreIndexFailed",
	ObBackupInProgress:                          "ObBackupInProgress",
	ObInvalidLogArchiveStatus:                   "ObInvalidLogArchiveStatus",
	ObCannotAddReplicaDuringSetMemberList:       "ObCannotAddReplicaDuringSetMemberList",
	ObLogArchiveLeaderChanged:                   "ObLogArchiveLeaderChanged",
	ObBackupCanNotStart:                         "ObBackupCanNotStart",
	ObCancelBackupNotAllowed:                    "ObCancelBackupNotAllowed",
	ObBackupDataVersionGapOverLimit:             "ObBackupDataVersionGapOverLimit",
	ObPgLogArchiveStatusNotInit:                 "ObPgLogArchiveStatusNotInit",
	ObBackupDeleteDataInProgress:                "ObBackupDeleteDataInProgress",
	ObBackupDeleteBackupSetNotAllowed:           "ObBackupDeleteBackupSetNotAllowed",
	ObInvalidBackupSetId:                        "ObInvalidBackupSetId",
	ObBackupInvalidPassword:                     "ObBackupInvalidPassword",
	ObIsolatedBackupSet:                         "ObIsolatedBackupSet",
	ObCannotCancelStoppedBackup:                 "ObCannotCancelStoppedBackup",
	ObBackupBackupCanNotStart:                   "ObBackupBackupCanNotStart",
	ObBackupMountFileNotValid:                   "ObBackupMountFileNotValid",
	ObBackupCleanInfoNotMatch:                   "ObBackupCleanInfoNotMatch",
	ObCancelDeleteBackupNotAllowed:              "ObCancelDeleteBackupNotAllowed",
	ObBackupCleanInfoNotExist:                   "ObBackupCleanInfoNotExist",
	ObCannotSetBackupRegion:                     "ObCannotSetBackupRegion",
	ObCannotSetBackupZone:                       "ObCannotSetBackupZone",
	ObBackupBackupReachMaxBackupTimes:           "ObBackupBackupReachMaxBackupTimes",
	ObArchiveLogNotContinuesWithData:            "ObArchiveLogNotContinuesWithData",
	ObAgentHasSuspended:                         "ObAgentHasSuspended",
	ObBackupConflictValue:                       "ObBackupConflictValue",
	ObBackupDeleteBackupPieceNotAllowed:         "ObBackupDeleteBackupPieceNotAllowed",
	ObBackupDestNotConnect:                      "ObBackupDestNotConnect",
	ObEsiSessionConflicts:                       "ObEsiSessionConflicts",
	ObBackupValidateTaskSkipped:                 "ObBackupValidateTaskSkipped",
	ObEsiIoError:                                "ObEsiIoError",
	ObArchiveRoundNotContinuous:                 "ObArchiveRoundNotContinuous",
	ObArchiveLogToEnd:                           "ObArchiveLogToEnd",
	ObArchiveLogRecycled:                        "ObArchiveLogRecycled",
	ObBackupFormatFileNotExist:                  "ObBackupFormatFileNotExist",
	ObBackupFormatFileNotMatch:                  "ObBackupFormatFileNotMatch",
	ObBackupMajorNotCoverMinor:                  "ObBackupMajorNotCoverMinor",
	ObBackupAdvanceCheckpointTimeout:            "ObBackupAdvanceCheckpointTimeout",
	ObClogRecycleBeforeArchive:                  "ObClogRecycleBeforeArchive",
	ObSourceTenantStateNotMatch:                 "ObSourceTenantStateNotMatch",
	ObSourceLsStateNotMatch:                     "ObSourceLsStateNotMatch",
	ObEsiSessionNotExist:                        "ObEsiSessionNotExist",
	ObAlreadyInArchiveMode:                      "ObAlreadyInArchiveMode",
	ObAlreadyInNoarchiveMode:                    "ObAlreadyInNoarchiveMode",
	ObRestoreLogToEnd:                           "ObRestoreLogToEnd",
	ObLsRestoreFailed:                           "ObLsRestoreFailed",
	ObNoTabletNeedBackup:                        "ObNoTabletNeedBackup",
	ObErrRestoreStandbyVersionLag:               "ObErrRestoreStandbyVersionLag",
	ObErrRestorePrimaryTenantDropped:            "ObErrRestorePrimaryTenantDropped",
	ObNoSuchFileOrDirectory:                     "ObNoSuchFileOrDirectory",
	ObFileOrDirectoryExist:                      "ObFileOrDirectoryExist",
	ObFileOrDirectoryPermissionDenied:           "ObFileOrDirectoryPermissionDenied",
	ObTooManyOpenFiles:                          "ObTooManyOpenFiles",
	ObDirectLoadCommitError:                     "ObDirectLoadCommitError",
	ObErrResizeFileToSmaller:                    "ObErrResizeFileToSmaller",
	ObMarkBlockInfoTimeout:                      "ObMarkBlockInfoTimeout",
	ObNotReadyToExtendFile:                      "ObNotReadyToExtendFile",
	ObErrDuplicateHavingClauseInTableExpression: "ObErrDuplicateHavingClauseInTableExpression",
	ObErrInoutParamPlacementNotProperly:         "ObErrInoutParamPlacementNotProperly",
	ObErrObjectNotFound:                         "ObErrObjectNotFound",
	ObErrInvalidInputValue:                      "ObErrInvalidInputValue",
	ObErrGotoBranchIllegal:                      "ObErrGotoBranchIllegal",
	ObErrOnlySchemaLevelAllow:                   "ObErrOnlySchemaLevelAllow",
	ObErrDeclMoreThanOnce:                       "ObErrDeclMoreThanOnce",
	ObErrDuplicateFiled:                         "ObErrDuplicateFiled",
	ObErrPragmaIllegal:                          "ObErrPragmaIllegal",
	ObErrExitContinueIllegal:                    "ObErrExitContinueIllegal",
	ObErrLabelIllegal:                           "ObErrLabelIllegal",
	ObErrCursorLeftAssign:                       "ObErrCursorLeftAssign",
	ObErrInitNotnullIllegal:                     "ObErrInitNotnullIllegal",
	ObErrInitConstIllegal:                       "ObErrInitConstIllegal",
	ObErrCursorVarInPkg:                         "ObErrCursorVarInPkg",
	ObErrLimitClause:                            "ObErrLimitClause",
	ObErrExpressionWrongType:                    "ObErrExpressionWrongType",
	ObErrSpecNotExist:                           "ObErrSpecNotExist",
	ObErrTypeSpecNoRoutine:                      "ObErrTypeSpecNoRoutine",
	ObErrTypeBodyNoRoutine:                      "ObErrTypeBodyNoRoutine",
	ObErrBothOrderMap:                           "ObErrBothOrderMap",
	ObErrNoOrderMap:                             "ObErrNoOrderMap",
	ObErrOrderMapNeedBeFunc:                     "ObErrOrderMapNeedBeFunc",
	ObErrIdentifierTooLong:                      "ObErrIdentifierTooLong",
	ObErrInvokeStaticByInstance:                 "ObErrInvokeStaticByInstance",
	ObErrConsNameIllegal:                        "ObErrConsNameIllegal",
	ObErrAttrFuncConflict:                       "ObErrAttrFuncConflict",
	ObErrSelfParamNotOut:                        "ObErrSelfParamNotOut",
	ObErrMapRetScalarType:                       "ObErrMapRetScalarType",
	ObErrMapMoreThanSelfParam:                   "ObErrMapMoreThanSelfParam",
	ObErrOrderRetIntType:                        "ObErrOrderRetIntType",
	ObErrOrderParamType:                         "ObErrOrderParamType",
	ObErrObjCmpSql:                              "ObErrObjCmpSql",
	ObErrMapOrderPragma:                         "ObErrMapOrderPragma",
	ObErrOrderParamMustInMode:                   "ObErrOrderParamMustInMode",
	ObErrOrderParamNotTwo:                       "ObErrOrderParamNotTwo",
	ObErrTypeRefRefcursive:                      "ObErrTypeRefRefcursive",
	ObErrDirectiveError:                         "ObErrDirectiveError",
	ObErrConsHasRetNode:                         "ObErrConsHasRetNode",
	ObErrCallWrongArg:                           "ObErrCallWrongArg",
	ObErrFuncNameSameWithCons:                   "ObErrFuncNameSameWithCons",
	ObErrFuncDup:                                "ObErrFuncDup",
	ObErrWhenClause:                             "ObErrWhenClause",
	ObErrNewOldReferences:                       "ObErrNewOldReferences",
	ObErrTypeDeclIllegal:                        "ObErrTypeDeclIllegal",
	ObErrObjectInvalid:                          "ObErrObjectInvalid",
	ObErrExpNotAssignable:                       "ObErrExpNotAssignable",
	ObErrCursorContainBothRegularAndArray:       "ObErrCursorContainBothRegularAndArray",
	ObErrStaticBoolExpr:                         "ObErrStaticBoolExpr",
	ObErrDirectiveContext:                       "ObErrDirectiveContext",
	ObUtlFileInvalidPath:                        "ObUtlFileInvalidPath",
	ObUtlFileInvalidMode:                        "ObUtlFileInvalidMode",
	ObUtlFileInvalidFilehandle:                  "ObUtlFileInvalidFilehandle",
	ObUtlFileInvalidOperation:                   "ObUtlFileInvalidOperation",
	ObUtlFileReadError:                          "ObUtlFileReadError",
	ObUtlFileWriteError:                         "ObUtlFileWriteError",
	ObUtlFileInternalError:                      "ObUtlFileInternalError",
	ObUtlFileCharsetmismatch:                    "ObUtlFileCharsetmismatch",
	ObUtlFileInvalidMaxlinesize:                 "ObUtlFileInvalidMaxlinesize",
	ObUtlFileInvalidFilename:                    "ObUtlFileInvalidFilename",
	ObUtlFileAccessDenied:                       "ObUtlFileAccessDenied",
	ObUtlFileInvalidOffset:                      "ObUtlFileInvalidOffset",
	ObUtlFileDeleteFailed:                       "ObUtlFileDeleteFailed",
	ObUtlFileRenameFailed:                       "ObUtlFileRenameFailed",
	ObErrBindTypeNotMatchColumn:                 "ObErrBindTypeNotMatchColumn",
	ObErrNestedTableInTri:                       "ObErrNestedTableInTri",
	ObErrColListInTri:                           "ObErrColListInTri",
	ObErrWhenClauseInTri:                        "ObErrWhenClauseInTri",
	ObErrInsteadTriOnTable:                      "ObErrInsteadTriOnTable",
	ObErrReturningClause:                        "ObErrReturningClause",
	ObErrNoReturnInFunction:                     "ObErrNoReturnInFunction",
	ObErrStmtNotAllowInMysqlFuncTrigger:         "ObErrStmtNotAllowInMysqlFuncTrigger",
	ObErrTooLongStringType:                      "ObErrTooLongStringType",
	ObErrWidthOutOfRange:                        "ObErrWidthOutOfRange",
	ObErrRedefineLabel:                          "ObErrRedefineLabel",
	ObErrStmtNotAllowInMysqlProcedrue:           "ObErrStmtNotAllowInMysqlProcedrue",
	ObErrTriggerNotSupport:                      "ObErrTriggerNotSupport",
	ObErrTriggerInWrongSchema:                   "ObErrTriggerInWrongSchema",
	ObErrUnknownException:                       "ObErrUnknownException",
	ObErrTriggerCantChangeRow:                   "ObErrTriggerCantChangeRow",
	ObErrItemNotInBody:                          "ObErrItemNotInBody",
	ObErrWrongRowtype:                           "ObErrWrongRowtype",
	ObErrRoutineNotDefine:                       "ObErrRoutineNotDefine",
	ObErrDupNameInCursor:                        "ObErrDupNameInCursor",
	ObErrLocalCollInSql:                         "ObErrLocalCollInSql",
	ObErrTypeMismatchInFetch:                    "ObErrTypeMismatchInFetch",
	ObErrOthersMustLast:                         "ObErrOthersMustLast",
	ObErrRaiseNotInHandler:                      "ObErrRaiseNotInHandler",
	ObErrInvalidCursorReturnType:                "ObErrInvalidCursorReturnType",
	ObErrInCursorOpend:                          "ObErrInCursorOpend",
	ObErrCursorNoReturnType:                     "ObErrCursorNoReturnType",
	ObErrNoChoices:                              "ObErrNoChoices",
	ObErrTypeDeclMalformed:                      "ObErrTypeDeclMalformed",
	ObErrInFormalNotDenotable:                   "ObErrInFormalNotDenotable",
	ObErrOutParamHasDefault:                     "ObErrOutParamHasDefault",
	ObErrOnlyFuncCanPipelined:                   "ObErrOnlyFuncCanPipelined",
	ObErrPipeReturnNotColl:                      "ObErrPipeReturnNotColl",
	ObErrMismatchSubprogram:                     "ObErrMismatchSubprogram",
	ObErrParamInPackageSpec:                     "ObErrParamInPackageSpec",
	ObErrNumericLiteralRequired:                 "ObErrNumericLiteralRequired",
	ObErrNonIntLiteral:                          "ObErrNonIntLiteral",
	ObErrImproperConstraintForm:                 "ObErrImproperConstraintForm",
	ObErrTypeCantConstrained:                    "ObErrTypeCantConstrained",
	ObErrAnyCsNotAllowed:                        "ObErrAnyCsNotAllowed",
	ObErrSchemaTypeIllegal:                      "ObErrSchemaTypeIllegal",
	ObErrUnsupportedTableIndexType:              "ObErrUnsupportedTableIndexType",
	ObErrArrayMustHavePositiveLimit:             "ObErrArrayMustHavePositiveLimit",
	ObErrForallIterNotAllowed:                   "ObErrForallIterNotAllowed",
	ObErrBulkInBind:                             "ObErrBulkInBind",
	ObErrForallBulkTogether:                     "ObErrForallBulkTogether",
	ObErrForallDmlWithoutBulk:                   "ObErrForallDmlWithoutBulk",
	ObErrShouldCollectionType:                   "ObErrShouldCollectionType",
	ObErrAssocElemType:                          "ObErrAssocElemType",
	ObErrIntoClauseExpected:                     "ObErrIntoClauseExpected",
	ObErrSubprogramViolatesPragma:               "ObErrSubprogramViolatesPragma",
	ObErrExprSqlType:                            "ObErrExprSqlType",
	ObErrPragmaDeclTwice:                        "ObErrPragmaDeclTwice",
	ObErrPragmaFollowDecl:                       "ObErrPragmaFollowDecl",
	ObErrPipeStmtInNonPipelinedFunc:             "ObErrPipeStmtInNonPipelinedFunc",
	ObErrImplRestriction:                        "ObErrImplRestriction",
	ObErrInsufficientPrivilege:                  "ObErrInsufficientPrivilege",
	ObErrIllegalOption:                          "ObErrIllegalOption",
	ObErrNoFunctionExist:                        "ObErrNoFunctionExist",
	ObErrOutOfScope:                             "ObErrOutOfScope",
	ObErrIllegalErrorNum:                        "ObErrIllegalErrorNum",
	ObErrDefaultNotMatch:                        "ObErrDefaultNotMatch",
	ObErrTableSingleIndex:                       "ObErrTableSingleIndex",
	ObErrPragmaDecl:                             "ObErrPragmaDecl",
	ObErrIncorrectArguments:                     "ObErrIncorrectArguments",
	ObErrReturnValueRequired:                    "ObErrReturnValueRequired",
	ObErrReturnExprIllegal:                      "ObErrReturnExprIllegal",
	ObErrLimitIllegal:                           "ObErrLimitIllegal",
	ObErrIntoExprIllegal:                        "ObErrIntoExprIllegal",
	ObErrBulkSqlRestriction:                     "ObErrBulkSqlRestriction",
	ObErrMixSingleMulti:                         "ObErrMixSingleMulti",
	ObErrTriggerNoSuchRow:                       "ObErrTriggerNoSuchRow",
	ObErrSetUsage:                               "ObErrSetUsage",
	ObErrModifierConflicts:                      "ObErrModifierConflicts",
	ObErrDuplicateModifier:                      "ObErrDuplicateModifier",
	ObErrStrLiteralTooLong:                      "ObErrStrLiteralTooLong",
	ObErrSelfParamNotInout:                      "ObErrSelfParamNotInout",
	ObErrConstructMustReturnSelf:                "ObErrConstructMustReturnSelf",
	ObErrFirstParamMustNotNull:                  "ObErrFirstParamMustNotNull",
	ObErrCoalesceAtLeastOneNotNull:              "ObErrCoalesceAtLeastOneNotNull",
	ObErrStaticMethodHasSelf:                    "ObErrStaticMethodHasSelf",
	ObErrNoAttrFound:                            "ObErrNoAttrFound",
	ObErrIllegalTypeForObject:                   "ObErrIllegalTypeForObject",
	ObErrUnsupportedType:                        "ObErrUnsupportedType",
	ObErrPositionalFollowName:                   "ObErrPositionalFollowName",
	ObErrNeedALabel:                             "ObErrNeedALabel",
	ObErrReferSamePackage:                       "ObErrReferSamePackage",
	ObErrPlCommon:                               "ObErrPlCommon",
	ObErrIdentEmpty:                             "ObErrIdentEmpty",
	ObErrPragmaStrUnsupport:                     "ObErrPragmaStrUnsupport",
	ObErrEndLabelNotMatch:                       "ObErrEndLabelNotMatch",
	ObErrWrongFetchIntoNum:                      "ObErrWrongFetchIntoNum",
	ObErrPragmaFirstArg:                         "ObErrPragmaFirstArg",
	ObErrTriggerCantChangeOldRow:                "ObErrTriggerCantChangeOldRow",
	ObErrTriggerCantCrtOnRoView:                 "ObErrTriggerCantCrtOnRoView",
	ObErrTriggerInvalidRefName:                  "ObErrTriggerInvalidRefName",
	ObErrExpNotIntoTarget:                       "ObErrExpNotIntoTarget",
	ObErrCaseNull:                               "ObErrCaseNull",
	ObErrInvalidGoto:                            "ObErrInvalidGoto",
	ObErrPrivateUdfUseInSql:                     "ObErrPrivateUdfUseInSql",
	ObErrFieldNotDenotable:                      "ObErrFieldNotDenotable",
	ObNumericPrecisionNotInteger:                "ObNumericPrecisionNotInteger",
	ObErrRequireInteger:                         "ObErrRequireInteger",
	ObErrIndexTableOfCursor:                     "ObErrIndexTableOfCursor",
	ObNullCheckError:                            "ObNullCheckError",
	ObErrExNameArg:                              "ObErrExNameArg",
	ObErrExArgNum:                               "ObErrExArgNum",
	ObErrExSecondArg:                            "ObErrExSecondArg",
	ObObenCursorNumberIsZero:                    "ObObenCursorNumberIsZero",
	ObNoStmtParse:                               "ObNoStmtParse",
	ObArrayCntIsIllegal:                         "ObArrayCntIsIllegal",
	ObErrWrongSchemaRef:                         "ObErrWrongSchemaRef",
	ObErrComponentUndeclared:                    "ObErrComponentUndeclared",
	ObErrFuncOnlyInSql:                          "ObErrFuncOnlyInSql",
	ObErrUndefined:                              "ObErrUndefined",
	ObErrSubtypeNotnullMismatch:                 "ObErrSubtypeNotnullMismatch",
	ObErrBindVarNotExist:                        "ObErrBindVarNotExist",
	ObErrCursorInOpenDynamicSql:                 "ObErrCursorInOpenDynamicSql",
	ObErrInvalidInputArgument:                   "ObErrInvalidInputArgument",
	ObErrClientIdentifierTooLong:                "ObErrClientIdentifierTooLong",
	ObErrInvalidNamespaceValue:                  "ObErrInvalidNamespaceValue",
	ObErrInvalidNamespaceBeg:                    "ObErrInvalidNamespaceBeg",
	ObErrSessionContextExceeded:                 "ObErrSessionContextExceeded",
	ObErrNotCursorNameInCurrentOf:               "ObErrNotCursorNameInCurrentOf",
	ObErrNotForUpdateCursorInCurrentOf:          "ObErrNotForUpdateCursorInCurrentOf",
	ObErrDupSignalSet:                           "ObErrDupSignalSet",
	ObErrSignalNotFound:                         "ObErrSignalNotFound",
	ObErrInvalidConditionNumber:                 "ObErrInvalidConditionNumber",
	ObErrRecursiveSqlLevelsExceeded:             "ObErrRecursiveSqlLevelsExceeded",
	ObErrInvalidSection:                         "ObErrInvalidSection",
	ObErrDuplicateTriggerSection:                "ObErrDuplicateTriggerSection",
	ObErrParsePlsql:                             "ObErrParsePlsql",
	ObErrSignalWarn:                             "ObErrSignalWarn",
	ObErrResignalWithoutActiveHandler:           "ObErrResignalWithoutActiveHandler",
	ObErrCannotUpdateVirtualColInTrg:            "ObErrCannotUpdateVirtualColInTrg",
	ObErrTrgOrder:                               "ObErrTrgOrder",
	ObErrRefAnotherTableInTrg:                   "ObErrRefAnotherTableInTrg",
	ObErrRefTypeInTrg:                           "ObErrRefTypeInTrg",
	ObErrRefCyclicInTrg:                         "ObErrRefCyclicInTrg",
	ObErrCannotSpecifyPrecedesInTrg:             "ObErrCannotSpecifyPrecedesInTrg",
	ObErrCannotPerformDmlInsideQuery:            "ObErrCannotPerformDmlInsideQuery",
	ObErrCannotPerformDdlCommitOrRollbackInsideQueryOrDmlTips: "ObErrCannotPerformDdlCommitOrRollbackInsideQueryOrDmlTips",
	ObErrStatementStringInExecuteImmediateIsNullOrZeroLength:  "ObErrStatementStringInExecuteImmediateIsNullOrZeroLength",
	ObErrMissingIntoKeyword:                                   "ObErrMissingIntoKeyword",
	ObErrClauseReturnIllegal:                                  "ObErrClauseReturnIllegal",
	ObErrNameHasTooManyParts:                                  "ObErrNameHasTooManyParts",
	ObErrLobSpanTransaction:                                   "ObErrLobSpanTransaction",
	ObErrInvalidMultiset:                                      "ObErrInvalidMultiset",
	ObErrInvalidCastUdt:                                       "ObErrInvalidCastUdt",
	ObErrPolicyExist:                                          "ObErrPolicyExist",
	ObErrPolicyNotExist:                                       "ObErrPolicyNotExist",
	ObErrAddPolicyToSysObject:                                 "ObErrAddPolicyToSysObject",
	ObErrInvalidInputString:                                   "ObErrInvalidInputString",
	ObErrSecColumnOnView:                                      "ObErrSecColumnOnView",
	ObErrInvalidInputForArgument:                              "ObErrInvalidInputForArgument",
	ObErrPolicyDisabled:                                       "ObErrPolicyDisabled",
	ObErrCircularPolicies:                                     "ObErrCircularPolicies",
	ObErrTooManyPolicies:                                      "ObErrTooManyPolicies",
	ObErrPolicyFunction:                                       "ObErrPolicyFunction",
	ObErrNoPrivEvalPredicate:                                  "ObErrNoPrivEvalPredicate",
	ObErrExecutePolicyFunction:                                "ObErrExecutePolicyFunction",
	ObErrPolicyPredicate:                                      "ObErrPolicyPredicate",
	ObErrNoPrivDirectPathAccess:                               "ObErrNoPrivDirectPathAccess",
	ObErrIntegrityConstraintViolated:                          "ObErrIntegrityConstraintViolated",
	ObErrPolicyGroupExist:                                     "ObErrPolicyGroupExist",
	ObErrPolicyGroupNotExist:                                  "ObErrPolicyGroupNotExist",
	ObErrDrivingContextExist:                                  "ObErrDrivingContextExist",
	ObErrDrivingContextNotExist:                               "ObErrDrivingContextNotExist",
	ObErrUpdateDefaultGroup:                                   "ObErrUpdateDefaultGroup",
	ObErrContextContainInvalidGroup:                           "ObErrContextContainInvalidGroup",
	ObErrInvalidSecColumnType:                                 "ObErrInvalidSecColumnType",
	ObErrUnprotectedVirtualColumn:                             "ObErrUnprotectedVirtualColumn",
	ObErrAttributeAssociation:                                 "ObErrAttributeAssociation",
	ObErrMergeIntoWithPolicy:                                  "ObErrMergeIntoWithPolicy",
	ObErrSpNoDropSp:                                           "ObErrSpNoDropSp",
	ObErrRecompilationObject:                                  "ObErrRecompilationObject",
	ObErrVariableNotInSelectList:                              "ObErrVariableNotInSelectList",
	ObErrMultiRecord:                                          "ObErrMultiRecord",
	ObErrMalformedPsPacket:                                    "ObErrMalformedPsPacket",
	ObErrViewSelectContainQuestionmark:                        "ObErrViewSelectContainQuestionmark",
	ObErrObjectNotExist:                                       "ObErrObjectNotExist",
	ObErrTableOutOfRange:                                      "ObErrTableOutOfRange",
	ObErrWrongUsage:                                           "ObErrWrongUsage",
	ObErrForallOnRemoteTable:                                  "ObErrForallOnRemoteTable",
	ObErrSequenceNotDefine:                                    "ObErrSequenceNotDefine",
	ObErrDebugIdNotExist:                                      "ObErrDebugIdNotExist",
	ObTTLNotEnable:                                            "ObTTLNotEnable",
	ObTTLColumnNotExist:                                       "ObTTLColumnNotExist",
	ObTTLColumnTypeNotSupported:                               "ObTTLColumnTypeNotSupported",
	ObTTLCmdNotAllowed:                                        "ObTTLCmdNotAllowed",
	ObTTLNoTaskRunning:                                        "ObTTLNoTaskRunning",
	ObTTLTenantIsRestore:                                      "ObTTLTenantIsRestore",
	ObTTLInvalidHbaseTtl:                                      "ObTTLInvalidHbaseTtl",
	ObTTLInvalidHbaseMaxVersions:                              "ObTTLInvalidHbaseMaxVersions",
	ObKvCredentialNotMatch:                                    "ObKvCredentialNotMatch",
	ObKvRowkeyCountNotMatch:                                   "ObKvRowkeyCountNotMatch",
	ObKvColumnTypeNotMatch:                                    "ObKvColumnTypeNotMatch",
	ObKvCollationMismatch:                                     "ObKvCollationMismatch",
	ObKvScanRangeMissing:                                      "ObKvScanRangeMissing",
	ObKvRedisParseError:                                       "ObKvRedisParseError",
	ObErrValuesClauseNeedHaveColumn:                           "ObErrValuesClauseNeedHaveColumn",
	ObErrValuesClauseCannotUseDefaultValues:                   "ObErrValuesClauseCannotUseDefaultValues",
	ObWrongPartitionName:                                      "ObWrongPartitionName",
	ObErrPluginIsNotLoaded:                                    "ObErrPluginIsNotLoaded",
	ObErrArgumentShouldConstantOrGroupExpr:                    "ObErrArgumentShouldConstantOrGroupExpr",
	ObSpRaiseApplicationError:                                 "ObSpRaiseApplicationError",
	ObSpRaiseApplicationErrorNum:                              "ObSpRaiseApplicationErrorNum",
	ObClobOnlySupportWithMultibyteFun:                         "ObClobOnlySupportWithMultibyteFun",
	ObErrUpdateTwice:                                          "ObErrUpdateTwice",
	ObErrFlashbackQueryWithUpdate:                             "ObErrFlashbackQueryWithUpdate",
	ObErrUpdateOnExpr:                                         "ObErrUpdateOnExpr",
	ObErrSpecifiedRowNoLongerExists:                           "ObErrSpecifiedRowNoLongerExists",
}

func (c ObErrorCode) IsRefreshTableErrorCode() bool {
	return c == ObSchemaError ||
		c == ObGtsNotReady ||
		c == ObSchemaEagain
}

func (c ObErrorCode) GetErrorCodeName() string {
	if name, ok := ObErrorNames[c]; ok {
		return name
	}
	return "UnknownErrorCode"
}
