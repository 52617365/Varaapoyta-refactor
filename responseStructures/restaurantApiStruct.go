package responseStructures

type RestaurantApiResponse struct {
	Data struct {
		ListRestaurantsByLocation struct {
			TotalCount int `json:"totalCount"`
			Edges      []struct {
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
						Ranges interface{} `json:"ranges"`
					} `json:"restaurantTime"`
					KitchenTime struct {
						Ranges interface{} `json:"ranges"`
					} `json:"kitchenTime"`
				} `json:"openingTime"`
				Links struct {
					TableReservationLocalized struct {
						FiFI string `json:"fi_FI"`
					} `json:"tableReservationLocalized"`
					HomepageLocalized struct {
						FiFI string `json:"fi_FI"`
					} `json:"homepageLocalized"`
				} `json:"links"`
			} `json:"edges"`
		} `json:"listRestaurantsByLocation"`
	} `json:"data"`
}
