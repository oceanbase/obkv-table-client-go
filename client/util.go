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

package client

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/pkg/errors"
)

// TransferQueryRange sets the query range into tableQuery.
func TransferQueryRange(rangePair []*table.RangePair) ([]*protocol.ObNewRange, error) {
	queryRanges := make([]*protocol.ObNewRange, 0, len(rangePair))
	for _, rangePair := range rangePair {
		if len(rangePair.Start()) != len(rangePair.End()) {
			return nil, errors.New("startRange and endRange key length is not equal")
		}
		startObjs := make([]*protocol.ObObject, 0, len(rangePair.Start()))
		endObjs := make([]*protocol.ObObject, 0, len(rangePair.End()))
		for i := 0; i < len(rangePair.Start()); i++ {
			// append start obj
			objMeta, err := protocol.DefaultObjMeta(rangePair.Start()[i].Value())
			if err != nil {
				return nil, errors.WithMessage(err, "create obj meta by Range key")
			}
			startObjs = append(startObjs, protocol.NewObObjectWithParams(objMeta, rangePair.Start()[i].Value()))

			// append end obj
			objMeta, err = protocol.DefaultObjMeta(rangePair.End()[i].Value())
			if err != nil {
				return nil, errors.WithMessage(err, "create obj meta by Range key")
			}
			endObjs = append(endObjs, protocol.NewObObjectWithParams(objMeta, rangePair.End()[i].Value()))
		}
		borderFlag := protocol.NewObBorderFlag()
		if !rangePair.IncludeStart() {
			borderFlag.UnSetInclusiveStart()
		}
		if !rangePair.IncludeEnd() {
			borderFlag.UnSetInclusiveEnd()
		}
		queryRanges = append(queryRanges, protocol.NewObNewRangeWithParams(startObjs, endObjs, borderFlag))
	}
	return queryRanges, nil
}
