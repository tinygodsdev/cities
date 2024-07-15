package cities

const (
	// weather labels
	AttributeTemperature = "temperature"
	AttributeHumidity    = "humidity"
	AttributePressure    = "pressure"
	AttributeDescription = "description"

	// air quality labels
	AttributeCo   = "co"
	AttributeNo2  = "no2"
	AttributeO3   = "o3"
	AttributePm10 = "pm10"
	AttributePm25 = "pm25"
	AttributeSo2  = "so2"

	// world bank labels
	AttributeCPI                             = "Consumer price index (2010 = 100)"
	AttributeCPIShort                        = "Consumer price index (2010=100)"
	AttributeGDPPerCapita                    = "GDP per capita (current US$)"
	AttributeGDPPerCapitaShort               = "GDP per capita"
	AttributeExports                         = "Merchandise exports (current US$)"
	AttributeExportsShort                    = "Exports"
	AttributeImports                         = "Merchandise imports (current US$)"
	AttributeImportsShort                    = "Imports"
	AttributeUnemployment                    = "Unemployment, total (% of total labor force) (modeled ILO estimate)"
	AttributeUnemploymentShort               = "Unemployment"
	AttributeIndividualsUsingInternet        = "Individuals using the Internet (% of population)"
	AttributeIndividualsUsingInternetShort   = "Internet users"
	AttributeTaxRevenue                      = "Tax revenue (% of GDP)"
	AttributeTaxRevenueShort                 = "Tax revenue"
	AttributeLifeExpectancy                  = "Life expectancy at birth, total (years)"
	AttributeLifeExpectancyShort             = "Life expectancy"
	AttributeMortalityRateUnder5             = "Mortality rate, under-5 (per 1,000 live births)"
	AttributeMortalityRateUnder5Short        = "Infant mortality"
	AttributeGovtExpenditureEducation        = "Government expenditure on education, total (% of GDP)"
	AttributeGovtExpenditureEducationShort   = "Spending on education"
	AttributeCO2Emissions                    = "CO2 emissions (metric tons per capita)"
	AttributeCO2EmissionsShort               = "CO2 emissions"
	AttributeLiteracyRate                    = "Literacy rate, adult total (% of people ages 15 and above)"
	AttributeLiteracyRateShort               = "Literacy rate"
	AttributeCurrentHealthExpenditure        = "Current health expenditure (% of GDP)"
	AttributeCurrentHealthExpenditureShort   = "Spending on health"
	AttributePovertyHeadcount                = "Poverty headcount ratio at $2.15 a day (2017 PPP) (% of population)"
	AttributePovertyHeadcountShort           = "Poverty"
	AttributeHealthExpenditurePerCapita      = "Current health expenditure per capita, PPP (current international $)"
	AttributeHealthExpenditurePerCapitaShort = "Health spending per capita"

	// prices
	AttributePairOfJeans      = "1 Pair of Jeans (Levis 501 Or Similar)"
	AttributePairOfJeansShort = "Pair of Jeans"

	AttributeApartment1BedroomOutsideCentre      = "Apartment (1 bedroom) Outside of Centre"
	AttributeApartment1BedroomOutsideCentreShort = "Rent 1-bedroom apartment (non-central)"

	AttributeApartment1BedroomCityCentre      = "Apartment (1 bedroom) in City Centre"
	AttributeApartment1BedroomCityCentreShort = "Rent 1-bedroom apartment (central)"

	AttributeApples      = "Apples (1 lb)"
	AttributeApplesShort = "1kg of apples"

	AttributeMonthlyNetSalary      = "Average Monthly Net Salary (After Tax)"
	AttributeMonthlyNetSalaryShort = "Monthly net salary"

	AttributeBanana      = "Banana (1 lb)"
	AttributeBananaShort = "1kg of bananas"

	AttributeBasicUtilities      = "Basic (Electricity, Heating, Cooling, Water, Garbage) for 915 sq ft Apartment"
	AttributeBasicUtilitiesShort = "Basic utilities for 85m² apartment"

	AttributeBeefRound      = "Beef Round (1 lb) (or Equivalent Back Leg Red Meat)"
	AttributeBeefRoundShort = "1kg of beef round"

	AttributeBottleOfWine      = "Bottle of Wine (Mid-Range)"
	AttributeBottleOfWineShort = "Bottle of wine"

	AttributeCappuccino      = "Cappuccino (regular)"
	AttributeCappuccinoShort = "Cappuccino"

	AttributeChickenFillets      = "Chicken Fillets (1 lb)"
	AttributeChickenFilletsShort = "1kg of chicken fillets"

	AttributeCigarettes      = "Cigarettes 20 Pack (Marlboro)"
	AttributeCigarettesShort = "Pack of cigarettes"

	AttributeDomesticBeerBottle      = "Domestic Beer (0.5 liter bottle)"
	AttributeDomesticBeerBottleShort = "Bottle of beer"

	AttributeEggs      = "Eggs (regular) (12)"
	AttributeEggsShort = "12 eggs"

	AttributeGasoline      = "Gasoline (1 gallon)"
	AttributeGasolineShort = "1L of gasoline"

	AttributeInternationalPrimarySchool      = "International Primary School, Yearly for 1 Child"
	AttributeInternationalPrimarySchoolShort = "Primary school (yearly)"

	AttributeInternet      = "Internet (60 Mbps or More, Unlimited Data, Cable/ADSL)"
	AttributeInternetShort = "Internet (60 Mbps or more)"

	AttributeBread      = "Loaf of Fresh White Bread (1 lb)"
	AttributeBreadShort = "1kg of bread"

	AttributeLocalCheese      = "Local Cheese (1 lb)"
	AttributeLocalCheeseShort = "1kg of cheese"

	AttributeMcMeal      = "McMeal at McDonalds (or Equivalent Combo Meal)"
	AttributeMcMealShort = "Meal at McDonalds or similar"

	AttributeMealFor2      = "Meal for 2 People, Mid-range Restaurant, Three-course"
	AttributeMealFor2Short = "Meal for 2 in mid-range restaurant"

	AttributeMilk      = "Milk (regular), (1 gallon)"
	AttributeMilkShort = "1L of milk"

	AttributeMobilePlan      = "Mobile Phone Monthly Plan with Calls and 10GB+ Data"
	AttributeMobilePlanShort = "Mobile plan (monthly)"

	AttributeMonthlyPass      = "Monthly Pass (Regular Price)"
	AttributeMonthlyPassShort = "Monthly pass"

	AttributeMortgageInterestRate      = "Mortgage Interest Rate in Percentages (%), Yearly, for 20 Years Fixed-Rate"
	AttributeMortgageInterestRateShort = "Mortgage interest rate"

	AttributeOneWayTicket      = "One-way Ticket (Local Transport)"
	AttributeOneWayTicketShort = "Local one-way ticket"

	AttributeOranges      = "Oranges (1 lb)"
	AttributeOrangesShort = "1kg of oranges"

	AttributePotato      = "Potato (1 lb)"
	AttributePotatoShort = "1kg of potatoes"

	AttributePreschool      = "Preschool (or Kindergarten), Full Day, Private, Monthly for 1 Child"
	AttributePreschoolShort = "Kindergarten (monthly)"

	AttributeTaxi1Mile      = "Taxi 1 mile (Normal Tariff)"
	AttributeTaxi1MileShort = "Taxi 1 km"

	AttributeTomato      = "Tomato (1 lb)"
	AttributeTomatoShort = "1kg of tomatoes"

	AttributeWaterBottle      = "Water (1.5 liter bottle)"
	AttributeWaterBottleShort = "1.5L water bottle"

	// ignored for now
	AttributePricePerSqFtOutsideCentre       = "Price per Square Feet to Buy Apartment Outside of Centre"
	AttributePricePerSqFtCityCentre          = "Price per Square Feet to Buy Apartment in City Centre"
	AttributeMealInexpensiveRestaurant       = "Meal, Inexpensive Restaurant"
	AttributeDomesticBeerPint                = "Domestic Beer (1 pint draught)"
	AttributeRice                            = "Rice (white), (1 lb)"
	AttributeTaxi1HourWaiting                = "Taxi 1hour Waiting (Normal Tariff)"
	AttributeFitnessClubFee                  = "Fitness Club, Monthly Fee for 1 Adult"
	AttributeTaxiStart                       = "Taxi Start (Normal Tariff)"
	AttributeTennisCourtRent                 = "Tennis Court Rent (1 Hour on Weekend)"
	AttributeToyotaCorolla                   = "Toyota Corolla Sedan 1.6l 97kW Comfort (Or Equivalent New Car)"
	AttributeApartment3BedroomsOutsideCentre = "Apartment (3 bedrooms) Outside of Centre"
	AttributeApartment3BedroomsCityCentre    = "Apartment (3 bedrooms) in City Centre"
	AttributeMenLeatherBusinessShoes         = "1 Pair of Men Leather Business Shoes"
	AttributeNikeRunningShoes                = "1 Pair of Nike Running Shoes (Mid-Range)"
	AttributeSummerDressChainStore           = "1 Summer Dress in a Chain Store (Zara, H&M, ...)"
	AttributeVolkswagenGolf                  = "Volkswagen Golf 1.4 90 KW Trendline (Or Equivalent New Car)"
	AttributeWaterSmallBottle                = "Water (12 oz small bottle)"
	AttributeOnion                           = "Onion (1 lb)"
	AttributeImportedBeer                    = "Imported Beer (12 oz small bottle)"
	AttributeLettuce                         = "Lettuce (1 head)"
	AttributeCinemaSeat                      = "Cinema, International Release, 1 Seat"
	AttributeCokePepsi                       = "Coke/Pepsi (12 oz small bottle)"

	// general info
	AttributeMotto                  = "motto"
	AttributePopulationTotal        = "population_total"
	AttributePopulationTotalShort   = "Population"
	AttributeAreaTotal              = "area_total_km2"
	AttributeAreaTotalShort         = "Area"
	AttributeElevation              = "elevation_m"
	AttributeElevationShort         = "Elevation"
	AttributeTimezone               = "timezone"
	AttributeTimezoneShort          = "Timezone"
	AttributePopulationDensity      = "population_density_km2"
	AttributePopulationDensityShort = "Density"

	// categories
	CategoryWeather    = "weather"
	CategoryAirQuality = "air_quality"
	CategoryWorldBank  = "world_bank"
	CategoryPrices     = "prices"
	CategoryInfo       = "info"

	// tags
	TagCategory = "category"
	TagCity     = "city"
)
