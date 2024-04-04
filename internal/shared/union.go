// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddStartDateUnion()  {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddEndDateUnion()    {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditEndDateUnion()   {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditStartDateUnion() {}

type UnionString string

func (UnionString) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}

type UnionBool bool

func (UnionBool) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsEvaluatePriceGroupGroupingValuesUnion() {}
