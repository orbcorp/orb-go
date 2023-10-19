// File generated from our OpenAPI spec by Stainless.

package shared

import "time"

type UnionTime time.Time

func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddStartDate()  {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsAddEndDate()    {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditEndDate()   {}
func (UnionTime) ImplementsSubscriptionPriceIntervalsParamsEditStartDate() {}
