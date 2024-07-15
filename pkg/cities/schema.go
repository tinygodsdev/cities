package cities

import (
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
)

type City struct {
	Name                           string    `json:"name"`
	Timestamp                      time.Time `json:"timestamp"`
	Temperature                    string    `json:"temperature"`
	Humidity                       string    `json:"humidity"`
	Pressure                       string    `json:"pressure"`
	Description                    string    `json:"description"`
	Co                             string    `json:"co"`
	No2                            string    `json:"no2"`
	O3                             string    `json:"o3"`
	Pm10                           string    `json:"pm10"`
	Pm25                           string    `json:"pm25"`
	So2                            string    `json:"so2"`
	CPI                            string    `json:"cpi"`
	GDPPerCapita                   string    `json:"gdp_per_capita"`
	Exports                        string    `json:"exports"`
	Imports                        string    `json:"imports"`
	Unemployment                   string    `json:"unemployment"`
	IndividualsUsingInternet       string    `json:"internet_users"`
	TaxRevenue                     string    `json:"tax_revenue"`
	LifeExpectancy                 string    `json:"life_expectancy"`
	MortalityRateUnder5            string    `json:"infant_mortality"`
	GovtExpenditureEducation       string    `json:"spending_on_education"`
	CO2Emissions                   string    `json:"co2_emissions"`
	LiteracyRate                   string    `json:"literacy_rate"`
	CurrentHealthExpenditure       string    `json:"spending_on_health"`
	PovertyHeadcount               string    `json:"poverty"`
	HealthExpenditurePerCapita     string    `json:"health_spending_per_capita"`
	PairOfJeans                    string    `json:"pair_of_jeans"`
	Apartment1BedroomOutsideCentre string    `json:"rent_1_bedroom_apartment_non_central"`
	Apartment1BedroomCityCentre    string    `json:"rent_1_bedroom_apartment_central"`
	Apples                         string    `json:"1kg_of_apples"`
	MonthlyNetSalary               string    `json:"monthly_net_salary"`
	Banana                         string    `json:"1kg_of_bananas"`
	BasicUtilities                 string    `json:"basic_utilities_for_85m2_apartment"`
	BeefRound                      string    `json:"1kg_of_beef_round"`
	BottleOfWine                   string    `json:"bottle_of_wine"`
	Cappuccino                     string    `json:"cappuccino"`
	ChickenFillets                 string    `json:"1kg_of_chicken_fillets"`
	Cigarettes                     string    `json:"pack_of_cigarettes"`
	DomesticBeerBottle             string    `json:"bottle_of_beer"`
	Eggs                           string    `json:"12_eggs"`
	Gasoline                       string    `json:"1l_of_gasoline"`
	InternationalPrimarySchool     string    `json:"primary_school_yearly"`
	Internet                       string    `json:"internet_60_mbps_or_more"`
	Bread                          string    `json:"1kg_of_bread"`
	LocalCheese                    string    `json:"1kg_of_cheese"`
	McMeal                         string    `json:"meal_at_mcdonalds_or_similar"`
	MealFor2                       string    `json:"meal_for_2_in_mid_range_restaurant"`
	Milk                           string    `json:"1l_of_milk"`
	MobilePlan                     string    `json:"mobile_plan_monthly"`
	MonthlyPass                    string    `json:"monthly_pass"`
	MortgageInterestRate           string    `json:"mortgage_interest_rate"`
	OneWayTicket                   string    `json:"local_one_way_ticket"`
	Oranges                        string    `json:"1kg_of_oranges"`
	Potato                         string    `json:"1kg_of_potatoes"`
	Preschool                      string    `json:"kindergarten_monthly"`
	Taxi1Mile                      string    `json:"taxi_1_km"`
	Tomato                         string    `json:"1kg_of_tomatoes"`
	WaterBottle                    string    `json:"1.5l_water_bottle"`
	Motto                          string    `json:"motto"`
	PopulationTotal                string    `json:"population"`
	AreaTotal                      string    `json:"area"`
	Elevation                      string    `json:"elevation"`
	Timezone                       string    `json:"timezone"`
	PopulationDensity              string    `json:"density"`
}

func NewCity(name string, attributes map[string]data.Attribute, timestamp time.Time) City {
	city := City{
		Name:                           name,
		Timestamp:                      timestamp,
		Temperature:                    getAttributeValue(attributes, AttributeTemperature),
		Humidity:                       getAttributeValue(attributes, AttributeHumidity),
		Pressure:                       getAttributeValue(attributes, AttributePressure),
		Description:                    getAttributeValue(attributes, AttributeDescription),
		Co:                             getAttributeValue(attributes, AttributeCo),
		No2:                            getAttributeValue(attributes, AttributeNo2),
		O3:                             getAttributeValue(attributes, AttributeO3),
		Pm10:                           getAttributeValue(attributes, AttributePm10),
		Pm25:                           getAttributeValue(attributes, AttributePm25),
		So2:                            getAttributeValue(attributes, AttributeSo2),
		CPI:                            getAttributeValue(attributes, AttributeCPI),
		GDPPerCapita:                   getAttributeValue(attributes, AttributeGDPPerCapita),
		Exports:                        getAttributeValue(attributes, AttributeExports),
		Imports:                        getAttributeValue(attributes, AttributeImports),
		Unemployment:                   getAttributeValue(attributes, AttributeUnemployment),
		IndividualsUsingInternet:       getAttributeValue(attributes, AttributeIndividualsUsingInternet),
		TaxRevenue:                     getAttributeValue(attributes, AttributeTaxRevenue),
		LifeExpectancy:                 getAttributeValue(attributes, AttributeLifeExpectancy),
		MortalityRateUnder5:            getAttributeValue(attributes, AttributeMortalityRateUnder5),
		GovtExpenditureEducation:       getAttributeValue(attributes, AttributeGovtExpenditureEducation),
		CO2Emissions:                   getAttributeValue(attributes, AttributeCO2Emissions),
		LiteracyRate:                   getAttributeValue(attributes, AttributeLiteracyRate),
		CurrentHealthExpenditure:       getAttributeValue(attributes, AttributeCurrentHealthExpenditure),
		PovertyHeadcount:               getAttributeValue(attributes, AttributePovertyHeadcount),
		HealthExpenditurePerCapita:     getAttributeValue(attributes, AttributeHealthExpenditurePerCapita),
		PairOfJeans:                    getAttributeValue(attributes, AttributePairOfJeans),
		Apartment1BedroomOutsideCentre: getAttributeValue(attributes, AttributeApartment1BedroomOutsideCentre),
		Apartment1BedroomCityCentre:    getAttributeValue(attributes, AttributeApartment1BedroomCityCentre),
		Apples:                         getAttributeValue(attributes, AttributeApples),
		MonthlyNetSalary:               getAttributeValue(attributes, AttributeMonthlyNetSalary),
		Banana:                         getAttributeValue(attributes, AttributeBanana),
		BasicUtilities:                 getAttributeValue(attributes, AttributeBasicUtilities),
		BeefRound:                      getAttributeValue(attributes, AttributeBeefRound),
		BottleOfWine:                   getAttributeValue(attributes, AttributeBottleOfWine),
		Cappuccino:                     getAttributeValue(attributes, AttributeCappuccino),
		ChickenFillets:                 getAttributeValue(attributes, AttributeChickenFillets),
		Cigarettes:                     getAttributeValue(attributes, AttributeCigarettes),
		DomesticBeerBottle:             getAttributeValue(attributes, AttributeDomesticBeerBottle),
		Eggs:                           getAttributeValue(attributes, AttributeEggs),
		Gasoline:                       getAttributeValue(attributes, AttributeGasoline),
		InternationalPrimarySchool:     getAttributeValue(attributes, AttributeInternationalPrimarySchool),
		Internet:                       getAttributeValue(attributes, AttributeInternet),
		Bread:                          getAttributeValue(attributes, AttributeBread),
		LocalCheese:                    getAttributeValue(attributes, AttributeLocalCheese),
		McMeal:                         getAttributeValue(attributes, AttributeMcMeal),
		MealFor2:                       getAttributeValue(attributes, AttributeMealFor2),
		Milk:                           getAttributeValue(attributes, AttributeMilk),
		MobilePlan:                     getAttributeValue(attributes, AttributeMobilePlan),
		MonthlyPass:                    getAttributeValue(attributes, AttributeMonthlyPass),
		MortgageInterestRate:           getAttributeValue(attributes, AttributeMortgageInterestRate),
		OneWayTicket:                   getAttributeValue(attributes, AttributeOneWayTicket),
		Oranges:                        getAttributeValue(attributes, AttributeOranges),
		Potato:                         getAttributeValue(attributes, AttributePotato),
		Preschool:                      getAttributeValue(attributes, AttributePreschool),
		Taxi1Mile:                      getAttributeValue(attributes, AttributeTaxi1Mile),
		Tomato:                         getAttributeValue(attributes, AttributeTomato),
		WaterBottle:                    getAttributeValue(attributes, AttributeWaterBottle),
		Motto:                          getAttributeValue(attributes, AttributeMotto),
		PopulationTotal:                getAttributeValue(attributes, AttributePopulationTotal),
		AreaTotal:                      getAttributeValue(attributes, AttributeAreaTotal),
		Elevation:                      getAttributeValue(attributes, AttributeElevation),
		Timezone:                       getAttributeValue(attributes, AttributeTimezone),
		PopulationDensity:              getAttributeValue(attributes, AttributePopulationDensity),
	}
	return city
}

func getAttributeValue(attributes map[string]data.Attribute, label string) string {
	if attr, exists := attributes[label]; exists && len(attr.Values) > 0 {
		return attr.Values[0]
	}
	return ""
}
