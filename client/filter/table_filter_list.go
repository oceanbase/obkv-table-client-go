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

package filter

import (
	"strings"

	"github.com/pkg/errors"
)

type Operator uint8

const (
	OperatorAnd Operator = iota
	OperatorOr
)

type ObTableFilterList struct {
	op      Operator
	filters []ObTableFilter
}

func NewObTableFilterListWithOperator(op Operator) *ObTableFilterList {
	return &ObTableFilterList{
		op:      op,
		filters: nil,
	}
}

func NewObTableFilterListWithOperatorAndTableFilter(op Operator, filters ...ObTableFilter) *ObTableFilterList {
	return &ObTableFilterList{
		op:      op,
		filters: filters,
	}
}

func (l *ObTableFilterList) AddFilter(filters ...ObTableFilter) {
	l.filters = append(l.filters, filters...)
}

func (l *ObTableFilterList) Size() int {
	return len(l.filters)
}

func (l *ObTableFilterList) Get(pos int) (ObTableFilter, error) {
	if pos >= len(l.filters) {
		return nil, errors.Errorf("pos: %d is out of range: %d", pos, len(l.filters))
	}
	return l.filters[pos], nil
}

func (l *ObTableFilterList) String() string {
	if len(l.filters) == 0 {
		return ""
	}
	var builder strings.Builder
	var stringOperator string

	if l.op == OperatorAnd {
		stringOperator = " && "
	} else {
		stringOperator = " || "
	}

	for i, tableFilter := range l.filters {
		filterString := tableFilter.String()
		if filterString == "" {
			continue
		} else {
			if i != 0 {
				builder.WriteString(stringOperator)
			}
			if _, ok := tableFilter.(*ObTableValueFilter); ok {
				builder.WriteString(filterString)
			} else {
				builder.WriteString("(")
				builder.WriteString(filterString)
				builder.WriteString(")")
			}
		}
	}
	return builder.String()
}
