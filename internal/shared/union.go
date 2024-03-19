// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddStartDate()  {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddEndDate()    {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditEndDate()   {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditStartDate() {}

type UnionString string

func (UnionString) ImplementsEvaluatePriceGroupGroupingValue() {}

type UnionBool bool

func (UnionBool) ImplementsEvaluatePriceGroupGroupingValue() {}

type UnionFloat float64

func (UnionFloat) ImplementsEvaluatePriceGroupGroupingValue() {}
