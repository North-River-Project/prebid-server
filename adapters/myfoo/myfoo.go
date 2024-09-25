package myfoo

import (
	"encoding/json"
	"github.com/prebid/openrtb/v20/openrtb2"
	"github.com/prebid/prebid-server/v2/adapters"
	"github.com/prebid/prebid-server/v2/config"
	"github.com/prebid/prebid-server/v2/openrtb_ext"
	"math/rand"
	"net/http"
)

type myFooAdapter struct {
	request *openrtb2.BidRequest
}

func Builder(_ openrtb_ext.BidderName, _ config.Adapter, _ config.Server) (adapters.Bidder, error) {
	return &myFooAdapter{}, nil
}

// MakeRequests create the object for myFoo Request.
func (o *myFooAdapter) MakeRequests(request *openrtb2.BidRequest, _ *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	o.request = request

	reqJSON, err := json.Marshal(*request)
	if err != nil {
		return nil, []error{err}
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")

	return []*adapters.RequestData{
		{
			Method:  "GET",
			Uri:     "http://127.0.0.1:8000/",
			Body:    reqJSON,
			Headers: headers,
			ImpIDs:  openrtb_ext.GetImpIDs(request.Imp),
		},
	}, nil
}

// MakeBids make the bids for the bid response.
func (o *myFooAdapter) MakeBids(_ *openrtb2.BidRequest, _ *adapters.RequestData, _ *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	bidResponse := adapters.NewBidderResponseWithBidsCapacity(1)

	bidResponse.Bids = append(bidResponse.Bids, &adapters.TypedBid{
		Bid: &openrtb2.Bid{
			ID:    o.request.ID,
			ImpID: o.request.Imp[0].ID,
			Price: rand.Float64() * 10.0,
			CrID:  "CreativeID",
		},
		BidType: openrtb_ext.BidTypeBanner,
	})
	return bidResponse, nil
}
