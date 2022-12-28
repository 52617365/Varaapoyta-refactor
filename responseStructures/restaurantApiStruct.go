package responseStructures

type RestaurantApiResponse struct {
	Data struct {
		ListRestaurantsByLocation struct {
			TotalCount int     `json:"totalCount"`
			Edges      []Edges `json:"edges"`
		} `json:"listRestaurantsByLocation"`
	} `json:"data"`
}

type Edges struct {
	ID                string `json:"id"`
	ReservationPageID string `json:"-"`
	Name              struct {
		FiFI string `json:"fi_FI"`
	} `json:"name"`
	URLPath struct {
		FiFI string `json:"fi_FI"`
	} `json:"urlPath"`
	Address struct {
		Municipality struct {
			FiFI string `json:"fi_FI"`
		} `json:"municipality"`
		Street struct {
			FiFI string `json:"fi_FI"`
		} `json:"street"`
		ZipCode string `json:"zipCode"`
	} `json:"address"`
	Features struct {
		Accessible bool `json:"accessible"`
	} `json:"features"`
	OpeningTime struct {
		RestaurantTime struct {
			Ranges []Ranges `json:"ranges"`
		} `json:"restaurantTime"`
		KitchenTime struct {
			Ranges []Ranges `json:"ranges"`
		} `json:"kitchenTime"`
	} `json:"openingTime"`
	Links struct {
		TableReservationLocalized struct {
			FiFI string `json:"fi_FI"`
		} `json:"tableReservationLocalized"`
		HomepageLocalized struct {
			FiFI string `json:"fi_FI"`
		} `json:"homepageLocalized"`
	}
}

type Ranges struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	EndNextDay bool   `json:"endNextDay"`
}
