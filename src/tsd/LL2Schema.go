package tsd

import "time"

type LL2StandardResponse struct {
	Count int     `json:"count" bson:"count"`
	Next  *string `json:"next" bson:"next"`
	Prev  *string `json:"previous" bson:"previous"`
}

type LL2LaunchesResponse struct {
	Results []LL2Launch `json:"results" bson:"results,inline"`
	LL2StandardResponse
}

type LL2LaunchStatus struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Abbrev      string `json:"abbrev" bson:"abbrev"`
	Description string `json:"description" bson:"description"`
}

type LL2LaunchUpdate struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Abbrev      string `json:"abbrev" bson:"abbrev"`
	Description string `json:"description" bson:"description"`
}

type LL2LaunchNetPrecision struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Abbrev      string `json:"abbrev" bson:"abbrev"`
	Description string `json:"description" bson:"description"`
}

type LL2GenericTypeStruct struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type LL2LSP struct {
	ID                            int                  `json:"id" bson:"id"`
	URL                           string               `json:"url" bson:"url"`
	Name                          string               `json:"name" bson:"name"`
	Featured                      bool                 `json:"featured" bson:"featured"`
	Type                          LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	CountryCode                   string               `json:"country_code" bson:"country_code"`
	Abbrev                        string               `json:"abbrev" bson:"abbrev"`
	Description                   string               `json:"description" bson:"description"`
	Administrator                 string               `json:"administrator" bson:"administrator"`
	FoundingYear                  int                  `json:"founding_year" bson:"founding_year"`
	Launchers                     string               `json:"launchers" bson:"launchers"`
	Spacecraft                    string               `json:"spacecraft" bson:"spacecraft"`
	LaunchLibraryURL              string               `json:"launch_library_url" bson:"launch_library_url"`
	TotalLaunchCount              int                  `json:"total_launch_count" bson:"total_launch_count"`
	ConsecutiveSuccessfulLaunches int                  `json:"consecutive_successful_launches" bson:"consecutive_successful_launches"`
	SuccessfulLaunches            int                  `json:"successful_launches" bson:"successful_launches"`
	FailedLaunches                int                  `json:"failed_launches" bson:"failed_launches"`
	PendingLaunches               int                  `json:"pending_launches" bson:"pending_launches"`
	ConsecutiveSuccessfulLandings int                  `json:"consecutive_successful_landings" bson:"consecutive_successful_landings"`
	SuccessfulLandings            int                  `json:"successful_landings" bson:"successful_landings"`
	FailedLandings                int                  `json:"failed_landings" bson:"failed_landings"`
	AttemptedLandings             int                  `json:"attempted_landings" bson:"attempted_landings"`
	InfoURL                       string               `json:"info_url" bson:"info_url"`
	WikiURL                       string               `json:"wiki_url" bson:"wiki_url"`
	LogoURL                       string               `json:"logo_url" bson:"logo_url"`
	ImageURL                      string               `json:"image_url" bson:"image_url"`
	NationURL                     string               `json:"nation_url" bson:"nation_url"`
}

type LL2LaunchImage struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	ImageUrl     string `json:"image_url" bson:"image_url"`
	ThumbnailUrl string `json:"thumbnail_url" bson:"thumbnail_url"`
	Credit       string `json:"credit" bson:"credit"`
	License      struct {
		ID       int    `json:"id" bson:"id"`
		Name     string `json:"name" bson:"name"`
		Priority int    `json:"priority" bson:"priority"`
		Link     string `json:"link" bson:"link"`
	} `json:"license" bson:"license,inline"`
	SingleUse bool          `json:"single_use" bson:"single_use"`
	Variants  []interface{} `json:"variants" bson:"variants"`
}

type LL2Rocket struct {
	ID              int                    `json:"id" bson:"id"`
	Configuration   LL2RocketConfiguration `json:"configuration" bson:"configuration,inline"`
	LauncherStage   []LL2LauncherStage     `json:"launcher_stage" bson:"launcher_stage,inline"`
	SpacecraftStage []LL2SpacecraftStage   `json:"spacecraft_stage" bson:"spacecraft_stage,inline"`
}

type LL2SpacecraftStage struct {
	ID            int               `json:"id" bson:"id"`
	URL           string            `json:"url" bson:"url"`
	MissionEnd    time.Time         `json:"mission_end" bson:"mission_end"`
	Destination   string            `json:"destination" bson:"destination"`
	LaunchCrew    []LL2Crew         `json:"launch_crew" bson:"launch_crew,inline"`
	OnboardCrew   []LL2Crew         `json:"onboard_crew" bson:"onboard_crew,inline"`
	LandingCrew   []LL2Crew         `json:"landing_crew" bson:"landing_crew,inline"`
	Spacecraft    LL2Spacecraft     `json:"spacecraft" bson:"spacecraft,inline"`
	Landing       LL2Landing        `json:"landing" bson:"landing,inline"`
	DockingEvents []LL2DockingEvent `json:"docking_events" bson:"docking_events,inline"`
}

type LL2DockingEvent struct {
	ID              int                `json:"id" bson:"id"`
	Spacestation    LL2Spacestation    `json:"spacestation" bson:"spacestation,inline"`
	Docking         time.Time          `json:"docking" bson:"docking"`
	Departure       time.Time          `json:"departure" bson:"departure"`
	DockingLocation LL2DockingLocation `json:"docking_location" bson:"docking_location,inline"`
}

type LL2Spacestation struct {
	ID          int                  `json:"id" bson:"id"`
	URL         string               `json:"url" bson:"url"`
	Name        string               `json:"name" bson:"name"`
	Status      LL2GenericTypeStruct `json:"status" bson:"status,inline"`
	Founded     string               `json:"founded" bson:"founded"`
	Description string               `json:"description" bson:"description"`
	Orbit       string               `json:"orbit" bson:"orbit"`
	ImageURL    string               `json:"image_url" bson:"image_url"`
}

type LL2DockingLocation struct {
	ID           int                    `json:"id" bson:"id"`
	Name         string                 `json:"name" bson:"name"`
	Spacestation LL2PartialSpaceStation `json:"spacestation" bson:"spacestation,inline"`
}

type LL2PartialSpaceStation struct {
	ID   int    `json:"id" bson:"id"`
	URL  string `json:"url" bson:"url"`
	Name string `json:"name" bson:"name"`
}

type LL2Spacecraft struct {
	ID               int                  `json:"id" bson:"id"`
	URL              string               `json:"url" bson:"url"`
	Name             string               `json:"name" bson:"name"`
	SerialNumber     string               `json:"serial_number" bson:"serial_number"`
	IsPlaceholder    bool                 `json:"is_placeholder" bson:"is_placeholder"`
	InSpace          bool                 `json:"in_space" bson:"in_space"`
	TimeInSpace      string               `json:"time_in_space" bson:"time_in_space"`
	TimeDocked       string               `json:"time_docked" bson:"time_docked"`
	FlightsCount     int                  `json:"flights_count" bson:"flights_count"`
	MissionEndsCount int                  `json:"mission_ends_count" bson:"mission_ends_count"`
	Status           LL2GenericTypeStruct `json:"status" bson:"status,inline"`
	Description      string               `json:"description" bson:"description"`
	SpacecraftConfig LL2SpacecraftConfig  `json:"spacecraft_config" bson:"spacecraft_config,inline"`
}

type LL2SpacecraftConfig struct {
	ID                    int                  `json:"id" bson:"id"`
	URL                   string               `json:"url" bson:"url"`
	Name                  string               `json:"name" bson:"name"`
	Type                  LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	Agency                LL2Agency            `json:"agency" bson:"agency,inline"`
	InUse                 bool                 `json:"in_use" bson:"in_use"`
	Capability            string               `json:"capability" bson:"capability"`
	History               string               `json:"history" bson:"history"`
	Details               string               `json:"details" bson:"details"`
	MaidenFlight          string               `json:"maiden_flight" bson:"maiden_flight"`
	Height                float64              `json:"height" bson:"height"`
	Diameter              float64              `json:"diameter" bson:"diameter"`
	HumanRated            bool                 `json:"human_rated" bson:"human_rated"`
	CrewCapacity          int                  `json:"crew_capacity" bson:"crew_capacity"`
	PayloadCapacity       int                  `json:"payload_capacity" bson:"payload_capacity"`
	PayloadReturnCapacity int                  `json:"payload_return_capacity" bson:"payload_return_capacity"`
	FlightLife            string               `json:"flight_life" bson:"flight_life"`
	ImageURL              string               `json:"image_url" bson:"image_url"`
	NationURL             string               `json:"nation_url" bson:"nation_url"`
	WikiLink              string               `json:"wiki_link" bson:"wiki_link"`
	InfoLink              string               `json:"info_link" bson:"info_link"`
}

type LL2Crew struct {
	ID        int          `json:"id" bson:"id"`
	Role      LL2Role      `json:"role" bson:"role,inline"`
	Astronaut LL2Astronaut `json:"astronaut" bson:"astronaut,inline"`
}

type LL2Role struct {
	ID       int    `json:"id" bson:"id"`
	Role     string `json:"role" bson:"role"`
	Priority int    `json:"priority" bson:"priority"`
}

type LL2Astronaut struct {
	ID           int                  `json:"id" bson:"id"`
	URL          string               `json:"url" bson:"url"`
	Name         string               `json:"name" bson:"name"`
	Type         LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	InSpace      bool                 `json:"in_space" bson:"in_space"`
	TimeInSpace  string               `json:"time_in_space" bson:"time_in_space"`
	Status       LL2GenericTypeStruct `json:"status" bson:"status,inline"`
	Agency       LL2Agency            `json:"agency" bson:"agency,inline"`
	DateOfBirth  string               `json:"date_of_birth" bson:"date_of_birth"`
	DateOfDeath  string               `json:"date_of_death" bson:"date_of_death"`
	Nationality  string               `json:"nationality" bson:"nationality"`
	Twitter      string               `json:"twitter" bson:"twitter"`
	Instagram    string               `json:"instagram" bson:"instagram"`
	Bio          string               `json:"bio" bson:"bio"`
	ProfileImage string               `json:"profile_image" bson:"profile_image"`
	Wiki         string               `json:"wiki" bson:"wiki"`
	LastFlight   time.Time            `json:"last_flight" bson:"last_flight"`
	FirstFlight  time.Time            `json:"first_flight" bson:"first_flight"`
}

type LL2LauncherStage struct {
	ID                   int               `json:"id" bson:"id"`
	Type                 string            `json:"type" bson:"type"`
	Reused               bool              `json:"reused" bson:"reused"`
	LauncherFlightNumber int               `json:"launcher_flight_number" bson:"launcher_flight_number"`
	Launcher             LL2Launcher       `json:"launcher" bson:"launcher,inline"`
	Landing              LL2Landing        `json:"landing" bson:"landing,inline"`
	PreviousFlightDate   time.Time         `json:"previous_flight_date" bson:"previous_flight_date"`
	TurnAroundTimeDays   int               `json:"turn_around_time_days" bson:"turn_around_time_days"`
	PreviousFlight       LL2PreviousFlight `json:"previous_flight" bson:"previous_flight,inline"`
}

type LL2PreviousFlight struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type LL2Launcher struct {
	ID                 int                  `json:"id" bson:"id"`
	URL                string               `json:"url" bson:"url"`
	Details            string               `json:"details" bson:"details"`
	FlightProven       bool                 `json:"flight_proven" bson:"flight_proven"`
	SerialNumber       string               `json:"serial_number" bson:"serial_number"`
	Status             LL2GenericTypeStruct `json:"status" bson:"status,inline"`
	ImageURL           string               `json:"image_url" bson:"image_url"`
	SuccessfulLandings int                  `json:"successful_landings" bson:"successful_landings"`
	AttemptedLandings  int                  `json:"attempted_landings" bson:"attempted_landings"`
	Flights            int                  `json:"flights" bson:"flights"`
	LastLaunchDate     time.Time            `json:"last_launch_date" bson:"last_launch_date"`
	FirstLaunchDate    time.Time            `json:"first_launch_date" bson:"first_launch_date"`
}

type LL2Landing struct {
	ID                int                  `json:"id" bson:"id"`
	Attempt           bool                 `json:"attempt" bson:"attempt"`
	Success           bool                 `json:"success" bson:"success"`
	Description       string               `json:"description" bson:"description"`
	DownrangeDistance float64              `json:"downrange_distance" bson:"downrange_distance"`
	Location          LL2Location          `json:"location" bson:"location,inline"`
	Type              LL2GenericTypeStruct `json:"type" bson:"type,inline"`
}

type LL2Location struct {
	ID                 int                    `json:"id"`
	Name               string                 `json:"name"`
	Abbrev             string                 `json:"abbrev"`
	Description        string                 `json:"description"`
	Location           LL2LocationSubLocation `json:"location" bson:"location,inline"`
	SuccessfulLandings string                 `json:"successful_landings"`
}

type LL2LocationSubLocation struct {
	ID                int    `json:"id"`
	URL               string `json:"url"`
	Name              string `json:"name"`
	CountryCode       string `json:"country_code"`
	Description       string `json:"description"`
	MapImage          string `json:"map_image"`
	TimezoneName      string `json:"timezone_name"`
	TotalLaunchCount  int    `json:"total_launch_count"`
	TotalLandingCount int    `json:"total_landing_count"`
}

type LL2RocketConfiguration struct {
	ID                            int             `json:"id"`
	URL                           string          `json:"url"`
	Name                          string          `json:"name"`
	Active                        bool            `json:"active"`
	Reusable                      bool            `json:"reusable"`
	Description                   string          `json:"description"`
	Family                        string          `json:"family"`
	FullName                      string          `json:"full_name"`
	Manufacturer                  LL2Manufacturer `json:"manufacturer" bson:"manufacturer,inline"`
	Program                       []LL2Program    `json:"program" bson:"program,inline"`
	Variant                       string          `json:"variant"`
	Alias                         string          `json:"alias"`
	MinStage                      int             `json:"min_stage"`
	MaxStage                      int             `json:"max_stage"`
	Length                        float64         `json:"length"`
	Diameter                      float64         `json:"diameter"`
	MaidenFlight                  string          `json:"maiden_flight"`
	LaunchCost                    int             `json:"launch_cost"`
	LaunchMass                    float64         `json:"launch_mass"`
	LeoCapacity                   float64         `json:"leo_capacity"`
	GtoCapacity                   float64         `json:"gto_capacity"`
	ToThrust                      float64         `json:"to_thrust"`
	Apogee                        float64         `json:"apogee"`
	VehicleRange                  float64         `json:"vehicle_range"`
	ImageURL                      string          `json:"image_url"`
	InfoURL                       string          `json:"info_url"`
	WikiURL                       string          `json:"wiki_url"`
	TotalLaunchCount              int             `json:"total_launch_count"`
	ConsecutiveSuccessfulLaunches int             `json:"consecutive_successful_launches"`
	SuccessfulLaunches            int             `json:"successful_launches"`
	FailedLaunches                int             `json:"failed_launches"`
	PendingLaunches               int             `json:"pending_launches"`
	AttemptedLandings             int             `json:"attempted_landings"`
	SuccessfulLandings            int             `json:"successful_landings"`
	FailedLandings                int             `json:"failed_landings"`
	ConsecutiveSuccessfulLandings int             `json:"consecutive_successful_landings"`
}

type LL2Manufacturer struct {
	ID                            int                  `json:"id" bson:"id" `
	URL                           string               `json:"url" bson:"url"`
	Name                          string               `json:"name" bson:"name"`
	Featured                      bool                 `json:"featured" bson:"featured"`
	Type                          LL2GenericTypeStruct `json:"type" bson:"type, inline"`
	CountryCode                   string               `json:"country_code" bson:"country_code"`
	Abbrev                        string               `json:"abbrev" bson:"abbrev"`
	Description                   string               `json:"description" bson:"description"`
	Administrator                 string               `json:"administrator" bson:"administrator"`
	FoundingYear                  int                  `json:"founding_year" bson:"founding_year"`
	Launchers                     string               `json:"launchers" bson:"launchers"`
	Spacecraft                    string               `json:"spacecraft" bson:"spacecraft"`
	LaunchLibraryURL              string               `json:"launch_library_url" bson:"launch_library_url"`
	TotalLaunchCount              int                  `json:"total_launch_count" bson:"total_launch_count"`
	ConsecutiveSuccessfulLaunches int                  `json:"consecutive_successful_launches" bson:"consecutive_successful_launches"`
	SuccessfulLaunches            int                  `json:"successful_launches" bson:"successful_launches"`
	FailedLaunches                int                  `json:"failed_launches" bson:"failed_launches"`
	PendingLaunches               int                  `json:"pending_launches" bson:"pending_launches"`
	ConsecutiveSuccessfulLandings int                  `json:"consecutive_successful_landings" bson:"consecutive_successful_landings"`
	SuccessfulLandings            int                  `json:"successful_landings" bson:"successful_landings"`
	FailedLandings                int                  `json:"failed_landings" bson:"failed_landings"`
	AttemptedLandings             int                  `json:"attempted_landings" bson:"attempted_landings"`
	InfoURL                       string               `json:"info_url" bson:"info_url"`
	WikiURL                       string               `json:"wiki_url" bson:"wiki_url"`
	LogoURL                       string               `json:"logo_url" bson:"logo_url"`
	ImageURL                      string               `json:"image_url" bson:"image_url"`
	NationURL                     string               `json:"nation_url" bson:"nation_url"`
}

type LL2MissionPatch struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Priority int       `json:"priority"`
	ImageURL string    `json:"image_url"`
	Agency   LL2Agency `json:"agency" bson:"agency,inline"`
}

type LL2Agency struct {
	ID   int                  `json:"id" bson:"id"`
	URL  string               `json:"url" bson:"url"`
	Name string               `json:"name" bson:"name"`
	Type LL2GenericTypeStruct `json:"type" bson:"type,inline"`
}

type LL2InfoURL struct {
	Priority     int                  `json:"priority" bson:"priority"`
	Source       string               `json:"source" bson:"source"`
	Title        string               `json:"title" bson:"title"`
	Description  string               `json:"description" bson:"description"`
	FeatureImage string               `json:"feature_image" bson:"feature_image"`
	URL          string               `json:"url" bson:"url"`
	Type         LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	Language     LL2Language          `json:"language" bson:"language,inline"`
}

type LL2VidURL struct {
	Priority     int                  `json:"priority" bson:"priority"`
	Source       string               `json:"source" bson:"source"`
	Publisher    string               `json:"publisher" bson:"publisher"`
	Title        string               `json:"title" bson:"title"`
	Description  string               `json:"description" bson:"description"`
	FeatureImage string               `json:"feature_image" bson:"feature_image"`
	URL          string               `json:"url" bson:"url"`
	Type         LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	Language     LL2Language          `json:"language" bson:"language,inline"`
	StartTime    time.Time            `json:"start_time" bson:"start_time"`
	EndTime      time.Time            `json:"end_time" bson:"end_time"`
}

type LL2Language struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

type LL2Mission struct {
	ID               int                   `json:"id" bson:"id"`
	Name             string                `json:"name" bson:"name"`
	Description      string                `json:"description" bson:"description"`
	LaunchDesignator string                `json:"launch_designator" bson:"launch_designator"`
	Type             string                `json:"type" bson:"type"`
	Orbit            LL2LaunchNetPrecision `json:"orbit" bson:"orbit,inline"`
	Agencies         []LL2Agency           `json:"agencies" bson:"agencies,inline"`
	InfoURLs         []LL2InfoURL          `json:"info_urls" bson:"info_urls,inline"`
	VidURLs          []LL2VidURL           `json:"vid_urls" bson:"vid_urls,inline"`
}

type LL2Pad struct {
	ID               int         `json:"id" bson:"id"`
	URL              string      `json:"url" bson:"url"`
	AgencyID         int         `json:"agency_id" bson:"agency_id"`
	Name             string      `json:"name" bson:"name"`
	Description      string      `json:"description" bson:"description"`
	InfoURL          string      `json:"info_url" bson:"info_url"`
	WikiURL          string      `json:"wiki_url" bson:"wiki_url"`
	MapURL           string      `json:"map_url" bson:"map_url"`
	Latitude         float64     `json:"latitude" bson:"latitude"`
	Longitude        float64     `json:"longitude" bson:"longitude"`
	Location         LL2Location `json:"location" bson:"location"`
	CountryCode      string      `json:"country_code" bson:"country_code"`
	MapImage         string      `json:"map_image" bson:"map_image"`
	Timezone         string      `json:"timezone" bson:"timezone"`
	TotalLaunchCount int         `json:"total_launch_count" bson:"total_launch_count"`
}

type LL2URLResource struct {
	Priority     int                  `json:"priority" bson:"priority"`
	Source       string               `json:"source" bson:"source"`
	Title        string               `json:"title" bson:"title"`
	Description  string               `json:"description" bson:"description"`
	FeatureImage string               `json:"feature_image" bson:"feature_image"`
	URL          string               `json:"url" bson:"url"`
	Type         LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	Language     LL2Language          `json:"language" bson:"language,inline"`
}

type LL2MissionTimeline struct {
	Type         LL2GenericTypeStruct `json:"type" bson:"type,inline"`
	RelativeTime string               `json:"relative_time" bson:"relative_time"`
}

type LL2VideoResource struct {
	*LL2URLResource
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
}

type LL2Program struct {
	ID             int               `json:"id" bson:"id"`
	URL            string            `json:"url" bson:"url"`
	Name           string            `json:"name" bson:"name"`
	Description    string            `json:"description" bson:"description"`
	Agencies       []LL2Agency       `json:"agencies" bson:"agencies,inline"`
	ImageURL       string            `json:"image_url" bson:"image_url"`
	StartDate      time.Time         `json:"start_date" bson:"start_date"`
	EndDate        time.Time         `json:"end_date" bson:"end_date"`
	InfoURL        string            `json:"info_url" bson:"info_url"`
	WikiURL        string            `json:"wiki_url" bson:"wiki_url"`
	MissionPatches []LL2MissionPatch `json:"mission_patches" bson:"mission_patches,inline"`
}

type LL2Launch struct {
	ID                             string                `json:"id"`
	URL                            string                `json:"url"`
	Slug                           string                `json:"slug"`
	FlightClubURL                  string                `json:"flightclub_url"`
	RSpacexAPIID                   string                `json:"r_spacex_api_id"`
	Name                           string                `json:"name"`
	Status                         LL2LaunchStatus       `json:"status" bson:"status,inline"`
	LastUpdated                    time.Time             `json:"last_updated"`
	Updates                        []LL2LaunchUpdate     `json:"updates" bson:"updates,inline"`
	Net                            time.Time             `json:"net"`
	NetPrecision                   LL2LaunchNetPrecision `json:"net_precision" bson:"net_precision,inline"`
	WindowEnd                      time.Time             `json:"window_end"`
	WindowStart                    time.Time             `json:"window_start"`
	Probability                    int                   `json:"probability"`
	WeatherConcerns                string                `json:"weather_concerns"`
	HoldReason                     string                `json:"holdreason"`
	FailReason                     string                `json:"failreason"`
	Hashtag                        string                `json:"hashtag"`
	LaunchServiceProvider          LL2LSP                `json:"launch_service_provider" bson:"launch_service_provider"`
	Rocket                         LL2Rocket             `json:"rocket" bson:"rocket,inline"`
	Mission                        LL2Mission            `json:"mission" bson:"mission,inline"`
	Pad                            LL2Pad                `json:"pad" bson:"pad,inline"`
	InfoURLs                       []LL2URLResource      `json:"info_urls" bson:"info_urls,inline"`
	VidURLs                        []LL2VideoResource    `json:"vid_urls" bson:"vid_urls,inline"`
	WebcastLive                    bool                  `json:"webcast_live" bson:"webcast_live"`
	Timeline                       []LL2MissionTimeline  `json:"timeline" bson:"timeline,inline"`
	Image                          LL2LaunchImage        `json:"image" bson:"image,inline"`
	Infographic                    interface{}           `json:"infographic" bson:"infographic"`
	Program                        []LL2Program          `json:"program" bson:"program,inline"`
	OrbitalLaunchAttemptCount      int                   `json:"orbital_launch_attempt_count"`
	LocationLaunchAttemptCount     int                   `json:"location_launch_attempt_count"`
	PadLaunchAttemptCount          int                   `json:"pad_launch_attempt_count"`
	AgencyLaunchAttemptCount       int                   `json:"agency_launch_attempt_count"`
	OrbitalLaunchAttemptCountYear  int                   `json:"orbital_launch_attempt_count_year"`
	LocationLaunchAttemptCountYear int                   `json:"location_launch_attempt_count_year"`
	PadLaunchAttemptCountYear      int                   `json:"pad_launch_attempt_count_year"`
	AgencyLaunchAttemptCountYear   int                   `json:"agency_launch_attempt_count_year"`
	PadTurnaround                  string                `json:"pad_turnaround"`
	MissionPatches                 []LL2MissionPatch     `json:"mission_patches" bson:"mission_patches,inline"`
	Type                           LL2GenericTypeStruct  `json:"type"`
}

type LL2ThrottleResponse struct {
	RequestLimit       int    `json:"your_request_limit" bson:"your_request_limit"`
	LimitFrequencySecs int    `json:"limit_frequency_secs" bson:"limit_frequency_secs"`
	CurrentUse         int    `json:"current_use" bson:"current_use"`
	NextUseSecs        int    `json:"next_use_secs" bson:"next_use_secs"`
	Identifier         string `json:"ident" bson:"ident"`
}
