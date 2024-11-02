package tsd

import "time"

type LL2StandardResponse struct {
	Count int     `json:"count"`
	Next  *string `json:"next"`
	Prev  *string `json:"previous"`
}

type LL2LaunchesResponse struct {
	Results []LL2Launch `json:"results"`
	LL2StandardResponse
}

type LL2LaunchStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Abbrev      string `json:"abbrev"`
	Description string `json:"description"`
}

type LL2LaunchUpdate struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Abbrev      string `json:"abbrev"`
	Description string `json:"description"`
}

type LL2LaunchNetPrecision struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Abbrev      string `json:"abbrev"`
	Description string `json:"description"`
}

type LL2GenericTypeStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LL2LSP struct {
	ID                            int                  `json:"id"`
	URL                           string               `json:"url"`
	Name                          string               `json:"name"`
	Featured                      bool                 `json:"featured"`
	Type                          LL2GenericTypeStruct `json:"type"`
	CountryCode                   string               `json:"country_code"`
	Abbrev                        string               `json:"abbrev"`
	Description                   string               `json:"description"`
	Administrator                 string               `json:"administrator"`
	FoundingYear                  int                  `json:"founding_year"`
	Launchers                     string               `json:"launchers"`
	Spacecraft                    string               `json:"spacecraft"`
	LaunchLibraryURL              string               `json:"launch_library_url"`
	TotalLaunchCount              int                  `json:"total_launch_count"`
	ConsecutiveSuccessfulLaunches int                  `json:"consecutive_successful_launches"`
	SuccessfulLaunches            int                  `json:"successful_launches"`
	FailedLaunches                int                  `json:"failed_launches"`
	PendingLaunches               int                  `json:"pending_launches"`
	ConsecutiveSuccessfulLandings int                  `json:"consecutive_successful_landings"`
	SuccessfulLandings            int                  `json:"successful_landings"`
	FailedLandings                int                  `json:"failed_landings"`
	AttemptedLandings             int                  `json:"attempted_landings"`
	InfoURL                       string               `json:"info_url"`
	WikiURL                       string               `json:"wiki_url"`
	LogoURL                       string               `json:"logo_url"`
	ImageURL                      string               `json:"image_url"`
	NationURL                     string               `json:"nation_url"`
}

type LL2LaunchImage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ImageUrl     string `json:"image_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Credit       string `json:"credit"`
	License      struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Priority int    `json:"priority"`
		Link     string `json:"link"`
	} `json:"license"`
	SingleUse bool          `json:"single_use"`
	Variants  []interface{} `json:"variants"`
}

type LL2Launch struct {
	ID                    string                `json:"id"`
	URL                   string                `json:"url"`
	Slug                  string                `json:"slug"`
	FlightClubURL         string                `json:"flightclub_url"`
	RSpacexAPIID          string                `json:"r_spacex_api_id"`
	Name                  string                `json:"name"`
	Status                LL2LaunchStatus       `json:"status"`
	LastUpdated           time.Time             `json:"last_updated"`
	Updates               []LL2LaunchUpdate     `json:"updates"`
	Net                   time.Time             `json:"net"`
	NetPrecision          LL2LaunchNetPrecision `json:"net_precision"`
	WindowEnd             time.Time             `json:"window_end"`
	WindowStart           time.Time             `json:"window_start"`
	Probability           int                   `json:"probability"`
	WeatherConcerns       string                `json:"weather_concerns"`
	HoldReason            string                `json:"holdreason"`
	FailReason            string                `json:"failreason"`
	Hashtag               string                `json:"hashtag"`
	LaunchServiceProvider LL2LSP                `json:"launch_service_provider"`
	Rocket                struct {
		ID            int `json:"id"`
		Configuration struct {
			ID           int    `json:"id"`
			URL          string `json:"url"`
			Name         string `json:"name"`
			Active       bool   `json:"active"`
			Reusable     bool   `json:"reusable"`
			Description  string `json:"description"`
			Family       string `json:"family"`
			FullName     string `json:"full_name"`
			Manufacturer struct {
				ID                            int                  `json:"id"`
				URL                           string               `json:"url"`
				Name                          string               `json:"name"`
				Featured                      bool                 `json:"featured"`
				Type                          LL2GenericTypeStruct `json:"type"`
				CountryCode                   string               `json:"country_code"`
				Abbrev                        string               `json:"abbrev"`
				Description                   string               `json:"description"`
				Administrator                 string               `json:"administrator"`
				FoundingYear                  int                  `json:"founding_year"`
				Launchers                     string               `json:"launchers"`
				Spacecraft                    string               `json:"spacecraft"`
				LaunchLibraryURL              string               `json:"launch_library_url"`
				TotalLaunchCount              int                  `json:"total_launch_count"`
				ConsecutiveSuccessfulLaunches int                  `json:"consecutive_successful_launches"`
				SuccessfulLaunches            int                  `json:"successful_launches"`
				FailedLaunches                int                  `json:"failed_launches"`
				PendingLaunches               int                  `json:"pending_launches"`
				ConsecutiveSuccessfulLandings int                  `json:"consecutive_successful_landings"`
				SuccessfulLandings            int                  `json:"successful_landings"`
				FailedLandings                int                  `json:"failed_landings"`
				AttemptedLandings             int                  `json:"attempted_landings"`
				InfoURL                       string               `json:"info_url"`
				WikiURL                       string               `json:"wiki_url"`
				LogoURL                       string               `json:"logo_url"`
				ImageURL                      string               `json:"image_url"`
				NationURL                     string               `json:"nation_url"`
			} `json:"manufacturer"`
			Program []struct {
				ID          int    `json:"id"`
				URL         string `json:"url"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Agencies    []struct {
					ID   int                  `json:"id"`
					URL  string               `json:"url"`
					Name string               `json:"name"`
					Type LL2GenericTypeStruct `json:"type"`
				} `json:"agencies"`
				ImageURL       string    `json:"image_url"`
				StartDate      time.Time `json:"start_date"`
				EndDate        time.Time `json:"end_date"`
				InfoURL        string    `json:"info_url"`
				WikiURL        string    `json:"wiki_url"`
				MissionPatches []struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Priority int    `json:"priority"`
					ImageURL string `json:"image_url"`
					Agency   struct {
						ID   int                  `json:"id"`
						URL  string               `json:"url"`
						Name string               `json:"name"`
						Type LL2GenericTypeStruct `json:"type"`
					} `json:"agency"`
				} `json:"mission_patches"`
				Type LL2GenericTypeStruct `json:"type"`
			} `json:"program"`
			Variant                       string  `json:"variant"`
			Alias                         string  `json:"alias"`
			MinStage                      int     `json:"min_stage"`
			MaxStage                      int     `json:"max_stage"`
			Length                        float64 `json:"length"`
			Diameter                      float64 `json:"diameter"`
			MaidenFlight                  string  `json:"maiden_flight"`
			LaunchCost                    int     `json:"launch_cost"`
			LaunchMass                    float64 `json:"launch_mass"`
			LeoCapacity                   float64 `json:"leo_capacity"`
			GtoCapacity                   float64 `json:"gto_capacity"`
			ToThrust                      float64 `json:"to_thrust"`
			Apogee                        float64 `json:"apogee"`
			VehicleRange                  float64 `json:"vehicle_range"`
			ImageURL                      string  `json:"image_url"`
			InfoURL                       string  `json:"info_url"`
			WikiURL                       string  `json:"wiki_url"`
			TotalLaunchCount              int     `json:"total_launch_count"`
			ConsecutiveSuccessfulLaunches int     `json:"consecutive_successful_launches"`
			SuccessfulLaunches            int     `json:"successful_launches"`
			FailedLaunches                int     `json:"failed_launches"`
			PendingLaunches               int     `json:"pending_launches"`
			AttemptedLandings             int     `json:"attempted_landings"`
			SuccessfulLandings            int     `json:"successful_landings"`
			FailedLandings                int     `json:"failed_landings"`
			ConsecutiveSuccessfulLandings int     `json:"consecutive_successful_landings"`
		} `json:"configuration"`
		LauncherStage []struct {
			ID                   int    `json:"id"`
			Type                 string `json:"type"`
			Reused               bool   `json:"reused"`
			LauncherFlightNumber int    `json:"launcher_flight_number"`
			Launcher             struct {
				ID                 int                  `json:"id"`
				URL                string               `json:"url"`
				Details            string               `json:"details"`
				FlightProven       bool                 `json:"flight_proven"`
				SerialNumber       string               `json:"serial_number"`
				Status             LL2GenericTypeStruct `json:"status"`
				ImageURL           string               `json:"image_url"`
				SuccessfulLandings int                  `json:"successful_landings"`
				AttemptedLandings  int                  `json:"attempted_landings"`
				Flights            int                  `json:"flights"`
				LastLaunchDate     time.Time            `json:"last_launch_date"`
				FirstLaunchDate    time.Time            `json:"first_launch_date"`
			} `json:"launcher"`
			Landing struct {
				ID                int     `json:"id"`
				Attempt           bool    `json:"attempt"`
				Success           bool    `json:"success"`
				Description       string  `json:"description"`
				DownrangeDistance float64 `json:"downrange_distance"`
				Location          struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Abbrev      string `json:"abbrev"`
					Description string `json:"description"`
					Location    struct {
						ID                int    `json:"id"`
						URL               string `json:"url"`
						Name              string `json:"name"`
						CountryCode       string `json:"country_code"`
						Description       string `json:"description"`
						MapImage          string `json:"map_image"`
						TimezoneName      string `json:"timezone_name"`
						TotalLaunchCount  int    `json:"total_launch_count"`
						TotalLandingCount int    `json:"total_landing_count"`
					} `json:"location"`
					SuccessfulLandings string `json:"successful_landings"`
				} `json:"location"`
				Type struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Abbrev      string `json:"abbrev"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"landing"`
			PreviousFlightDate time.Time `json:"previous_flight_date"`
			TurnAroundTimeDays int       `json:"turn_around_time_days"`
			PreviousFlight     struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"previous_flight"`
		} `json:"launcher_stage"`
		SpacecraftStage []struct {
			ID          int       `json:"id"`
			URL         string    `json:"url"`
			MissionEnd  time.Time `json:"mission_end"`
			Destination string    `json:"destination"`
			LaunchCrew  []struct {
				ID   int `json:"id"`
				Role struct {
					ID       int    `json:"id"`
					Role     string `json:"role"`
					Priority int    `json:"priority"`
				} `json:"role"`
				Astronaut struct {
					ID          int                  `json:"id"`
					URL         string               `json:"url"`
					Name        string               `json:"name"`
					Type        LL2GenericTypeStruct `json:"type"`
					InSpace     bool                 `json:"in_space"`
					TimeInSpace string               `json:"time_in_space"`
					Status      struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"status"`
					Agency struct {
						ID   int                  `json:"id"`
						URL  string               `json:"url"`
						Name string               `json:"name"`
						Type LL2GenericTypeStruct `json:"type"`
					} `json:"agency"`
					DateOfBirth  string    `json:"date_of_birth"`
					DateOfDeath  string    `json:"date_of_death"`
					Nationality  string    `json:"nationality"`
					Twitter      string    `json:"twitter"`
					Instagram    string    `json:"instagram"`
					Bio          string    `json:"bio"`
					ProfileImage string    `json:"profile_image"`
					Wiki         string    `json:"wiki"`
					LastFlight   time.Time `json:"last_flight"`
					FirstFlight  time.Time `json:"first_flight"`
				} `json:"astronaut"`
			} `json:"launch_crew"`
			OnboardCrew []struct {
				ID   int `json:"id"`
				Role struct {
					ID       int    `json:"id"`
					Role     string `json:"role"`
					Priority int    `json:"priority"`
				} `json:"role"`
				Astronaut struct {
					ID          int                  `json:"id"`
					URL         string               `json:"url"`
					Name        string               `json:"name"`
					Type        LL2GenericTypeStruct `json:"type"`
					InSpace     bool                 `json:"in_space"`
					TimeInSpace string               `json:"time_in_space"`
					Status      struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"status"`
					Agency struct {
						ID   int                  `json:"id"`
						URL  string               `json:"url"`
						Name string               `json:"name"`
						Type LL2GenericTypeStruct `json:"type"`
					} `json:"agency"`
					DateOfBirth  string    `json:"date_of_birth"`
					DateOfDeath  string    `json:"date_of_death"`
					Nationality  string    `json:"nationality"`
					Twitter      string    `json:"twitter"`
					Instagram    string    `json:"instagram"`
					Bio          string    `json:"bio"`
					ProfileImage string    `json:"profile_image"`
					Wiki         string    `json:"wiki"`
					LastFlight   time.Time `json:"last_flight"`
					FirstFlight  time.Time `json:"first_flight"`
				} `json:"astronaut"`
			} `json:"onboard_crew"`
			LandingCrew []struct {
				ID   int `json:"id"`
				Role struct {
					ID       int    `json:"id"`
					Role     string `json:"role"`
					Priority int    `json:"priority"`
				} `json:"role"`
				Astronaut struct {
					ID          int                  `json:"id"`
					URL         string               `json:"url"`
					Name        string               `json:"name"`
					Type        LL2GenericTypeStruct `json:"type"`
					InSpace     bool                 `json:"in_space"`
					TimeInSpace string               `json:"time_in_space"`
					Status      struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"status"`
					Agency struct {
						ID   int                  `json:"id"`
						URL  string               `json:"url"`
						Name string               `json:"name"`
						Type LL2GenericTypeStruct `json:"type"`
					} `json:"agency"`
					DateOfBirth  string    `json:"date_of_birth"`
					DateOfDeath  string    `json:"date_of_death"`
					Nationality  string    `json:"nationality"`
					Twitter      string    `json:"twitter"`
					Instagram    string    `json:"instagram"`
					Bio          string    `json:"bio"`
					ProfileImage string    `json:"profile_image"`
					Wiki         string    `json:"wiki"`
					LastFlight   time.Time `json:"last_flight"`
					FirstFlight  time.Time `json:"first_flight"`
				} `json:"astronaut"`
			} `json:"landing_crew"`
			Spacecraft struct {
				ID               int    `json:"id"`
				URL              string `json:"url"`
				Name             string `json:"name"`
				SerialNumber     string `json:"serial_number"`
				IsPlaceholder    bool   `json:"is_placeholder"`
				InSpace          bool   `json:"in_space"`
				TimeInSpace      string `json:"time_in_space"`
				TimeDocked       string `json:"time_docked"`
				FlightsCount     int    `json:"flights_count"`
				MissionEndsCount int    `json:"mission_ends_count"`
				Status           struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"status"`
				Description      string `json:"description"`
				SpacecraftConfig struct {
					ID     int                  `json:"id"`
					URL    string               `json:"url"`
					Name   string               `json:"name"`
					Type   LL2GenericTypeStruct `json:"type"`
					Agency struct {
						ID            int                  `json:"id"`
						URL           string               `json:"url"`
						Name          string               `json:"name"`
						Featured      bool                 `json:"featured"`
						Type          LL2GenericTypeStruct `json:"type"`
						CountryCode   string               `json:"country_code"`
						Abbrev        string               `json:"abbrev"`
						Description   string               `json:"description"`
						Administrator string               `json:"administrator"`
						FoundingYear  int                  `json:"founding_year"`
						Launchers     interface{}          `json:"launchers"`
						Spacecraft    interface{}          `json:"spacecraft"`
						Parent        string               `json:"parent"`
						ImageURL      string               `json:"image_url"`
						LogoURL       string               `json:"logo_url"`
					} `json:"agency"`
					InUse                 bool    `json:"in_use"`
					Capability            string  `json:"capability"`
					History               string  `json:"history"`
					Details               string  `json:"details"`
					MaidenFlight          string  `json:"maiden_flight"`
					Height                float64 `json:"height"`
					Diameter              float64 `json:"diameter"`
					HumanRated            bool    `json:"human_rated"`
					CrewCapacity          int     `json:"crew_capacity"`
					PayloadCapacity       int     `json:"payload_capacity"`
					PayloadReturnCapacity int     `json:"payload_return_capacity"`
					FlightLife            string  `json:"flight_life"`
					ImageURL              string  `json:"image_url"`
					NationURL             string  `json:"nation_url"`
					WikiLink              string  `json:"wiki_link"`
					InfoLink              string  `json:"info_link"`
				} `json:"spacecraft_config"`
			} `json:"spacecraft"`
			Landing struct {
				ID                int    `json:"id"`
				Attempt           bool   `json:"attempt"`
				Success           bool   `json:"success"`
				Description       string `json:"description"`
				DownrangeDistance int    `json:"downrange_distance"`
				Location          struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Abbrev      string `json:"abbrev"`
					Description string `json:"description"`
					Location    struct {
						ID                int    `json:"id"`
						URL               string `json:"url"`
						Name              string `json:"name"`
						CountryCode       string `json:"country_code"`
						Description       string `json:"description"`
						MapImage          string `json:"map_image"`
						TimezoneName      string `json:"timezone_name"`
						TotalLaunchCount  int    `json:"total_launch_count"`
						TotalLandingCount int    `json:"total_landing_count"`
					} `json:"location"`
					SuccessfulLandings string `json:"successful_landings"`
				} `json:"location"`
				Type struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Abbrev      string `json:"abbrev"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"landing"`
			DockingEvents []struct {
				ID           int `json:"id"`
				Spacestation struct {
					ID     int    `json:"id"`
					URL    string `json:"url"`
					Name   string `json:"name"`
					Status struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"status"`
					Founded     string `json:"founded"`
					Description string `json:"description"`
					Orbit       string `json:"orbit"`
					ImageURL    string `json:"image_url"`
				} `json:"spacestation"`
				Docking         time.Time `json:"docking"`
				Departure       time.Time `json:"departure"`
				DockingLocation struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					Spacestation struct {
						ID   int    `json:"id"`
						URL  string `json:"url"`
						Name string `json:"name"`
					} `json:"spacestation"`
				} `json:"docking_location"`
			} `json:"docking_events"`
		} `json:"spacecraft_stage"`
	} `json:"rocket"`
	Mission struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		LaunchDesignator string `json:"launch_designator"`
		Type             string `json:"type"`
		Orbit            struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Abbrev string `json:"abbrev"`
		} `json:"orbit"`
		Agencies []struct {
			ID                            int                  `json:"id"`
			URL                           string               `json:"url"`
			Name                          string               `json:"name"`
			Featured                      bool                 `json:"featured"`
			Type                          LL2GenericTypeStruct `json:"type"`
			CountryCode                   string               `json:"country_code"`
			Abbrev                        string               `json:"abbrev"`
			Description                   string               `json:"description"`
			Administrator                 string               `json:"administrator"`
			FoundingYear                  int                  `json:"founding_year"`
			Launchers                     interface{}          `json:"launchers"`
			Spacecraft                    interface{}          `json:"spacecraft"`
			LaunchLibraryURL              string               `json:"launch_library_url"`
			TotalLaunchCount              int                  `json:"total_launch_count"`
			ConsecutiveSuccessfulLaunches int                  `json:"consecutive_successful_launches"`
			SuccessfulLaunches            int                  `json:"successful_launches"`
			FailedLaunches                int                  `json:"failed_launches"`
			PendingLaunches               int                  `json:"pending_launches"`
			ConsecutiveSuccessfulLandings int                  `json:"consecutive_successful_landings"`
			SuccessfulLandings            int                  `json:"successful_landings"`
			FailedLandings                int                  `json:"failed_landings"`
			AttemptedLandings             int                  `json:"attempted_landings"`
			InfoURL                       string               `json:"info_url"`
			WikiURL                       string               `json:"wiki_url"`
			LogoURL                       string               `json:"logo_url"`
			ImageURL                      string               `json:"image_url"`
			NationURL                     string               `json:"nation_url"`
		} `json:"agencies"`
		InfoUrls []struct {
			Priority     int                  `json:"priority"`
			Source       string               `json:"source"`
			Title        string               `json:"title"`
			Description  string               `json:"description"`
			FeatureImage string               `json:"feature_image"`
			URL          string               `json:"url"`
			Type         LL2GenericTypeStruct `json:"type"`
			Language     struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"language"`
		} `json:"info_urls"`
		VidUrls []struct {
			Priority     int                  `json:"priority"`
			Source       string               `json:"source"`
			Publisher    string               `json:"publisher"`
			Title        string               `json:"title"`
			Description  string               `json:"description"`
			FeatureImage string               `json:"feature_image"`
			URL          string               `json:"url"`
			Type         LL2GenericTypeStruct `json:"type"`
			Language     struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"language"`
			StartTime time.Time `json:"start_time"`
			EndTime   time.Time `json:"end_time"`
		} `json:"vid_urls"`
	} `json:"mission"`
	Pad struct {
		ID          int     `json:"id"`
		URL         string  `json:"url"`
		AgencyID    int     `json:"agency_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		InfoURL     string  `json:"info_url"`
		WikiURL     string  `json:"wiki_url"`
		MapURL      string  `json:"map_url"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Location    struct {
			ID                int    `json:"id"`
			URL               string `json:"url"`
			Name              string `json:"name"`
			CountryCode       string `json:"country_code"`
			Description       string `json:"description"`
			MapImage          string `json:"map_image"`
			TimezoneName      string `json:"timezone_name"`
			TotalLaunchCount  int    `json:"total_launch_count"`
			TotalLandingCount int    `json:"total_landing_count"`
		} `json:"location"`
		CountryCode               string `json:"country_code"`
		MapImage                  string `json:"map_image"`
		TotalLaunchCount          int    `json:"total_launch_count"`
		OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count"`
	} `json:"pad"`
	InfoURLs []struct {
		Priority     int    `json:"priority"`
		Source       string `json:"source"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		FeatureImage string `json:"feature_image"`
		URL          string `json:"url"`
		Type         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Language struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"language"`
	} `json:"infoURLs"`
	VidURLs []struct {
		Priority     int    `json:"priority"`
		Source       string `json:"source"`
		Publisher    string `json:"publisher"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		FeatureImage string `json:"feature_image"`
		URL          string `json:"url"`
		Type         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Language struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"language"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	} `json:"vidURLs"`
	WebcastLive bool `json:"webcast_live"`
	Timeline    []struct {
		Type struct {
			ID          int    `json:"id"`
			Abbrev      string `json:"abbrev"`
			Description string `json:"description"`
		} `json:"type"`
		RelativeTime string `json:"relative_time"`
	} `json:"timeline"`
	Image       LL2LaunchImage `json:"image"`
	Infographic interface{}    `json:"infographic"`
	Program     []struct {
		ID          int    `json:"id"`
		URL         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Agencies    []struct {
			ID   int                  `json:"id"`
			URL  string               `json:"url"`
			Name string               `json:"name"`
			Type LL2GenericTypeStruct `json:"type"`
		} `json:"agencies"`
		ImageURL       string    `json:"image_url"`
		StartDate      time.Time `json:"start_date"`
		EndDate        time.Time `json:"end_date"`
		InfoURL        string    `json:"info_url"`
		WikiURL        string    `json:"wiki_url"`
		MissionPatches []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Priority int    `json:"priority"`
			ImageURL string `json:"image_url"`
			Agency   struct {
				ID   int                  `json:"id"`
				URL  string               `json:"url"`
				Name string               `json:"name"`
				Type LL2GenericTypeStruct `json:"type"`
			} `json:"agency"`
		} `json:"mission_patches"`
		Type LL2GenericTypeStruct `json:"type"`
	} `json:"program"`
	OrbitalLaunchAttemptCount      int    `json:"orbital_launch_attempt_count"`
	LocationLaunchAttemptCount     int    `json:"location_launch_attempt_count"`
	PadLaunchAttemptCount          int    `json:"pad_launch_attempt_count"`
	AgencyLaunchAttemptCount       int    `json:"agency_launch_attempt_count"`
	OrbitalLaunchAttemptCountYear  int    `json:"orbital_launch_attempt_count_year"`
	LocationLaunchAttemptCountYear int    `json:"location_launch_attempt_count_year"`
	PadLaunchAttemptCountYear      int    `json:"pad_launch_attempt_count_year"`
	AgencyLaunchAttemptCountYear   int    `json:"agency_launch_attempt_count_year"`
	PadTurnaround                  string `json:"pad_turnaround"`
	MissionPatches                 []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Priority int    `json:"priority"`
		ImageURL string `json:"image_url"`
		Agency   struct {
			ID   int                  `json:"id"`
			URL  string               `json:"url"`
			Name string               `json:"name"`
			Type LL2GenericTypeStruct `json:"type"`
		} `json:"agency"`
	} `json:"mission_patches"`
	Type LL2GenericTypeStruct `json:"type"`
}

type LL2ThrottleResponse struct {
	RequestLimit       int    `json:"your_request_limit"`
	LimitFrequencySecs int    `json:"limit_frequency_secs"`
	CurrentUse         int    `json:"current_use"`
	NextUseSecs        int    `json:"next_use_secs"`
	Identifier         string `json:"ident"`
}
