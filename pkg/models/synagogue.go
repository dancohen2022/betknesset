package betknesset

import (
	"fmt"
)

type Synagogues struct {
	ListSynagogues []string
}

var synagogues Synagogues

func init() {
	//Note the following customizable parameters and their meanings:
	params := []string{"https://www.hebcal.com/hebcal?"}
	params = append(params, "v=1")       // version. Required.
	params = append(params, "&cfg=json") //output JSON instead of HTML
	//cfg=json for JSON format (described in more detail below)
	//cfg=fc for fullcalendar.io integration
	//cfg=rss for RSS (Atom 1.0)
	//cfg=ics for iCalendar RFC 5545
	//cfg=csv for Comma Separated Values
	params = append(params, "&year=now") // – “now” for current year, or 4-digit YYYY such as 2003
	params = append(params, "&month=x")  // – “x” for entire Gregorian year, or use a numeric month (1=January, 6=June, etc.)
	params = append(params, "&maj=on")   //– Major holidays
	params = append(params, "&min=on")   // – Minor holidays (Tu BiShvat, Lag B’Omer, …)
	params = append(params, "&nx=on")    // – Rosh Chodesh
	params = append(params, "&mf=on")    // – Minor fasts (Ta’anit Esther, Tzom Gedaliah, …)
	params = append(params, "&ss=on")    // – Special Shabbatot (Shabbat Shekalim, Zachor, …)
	params = append(params, "&mod=on")   // – Modern holidays (Yom HaShoah, Yom HaAtzma’ut, …)
	params = append(params, "&s=on")     // – Parashat ha-Shavuah on Saturday
	params = append(params, "&c=on")     // – Candle lighting times. See also candle-lighting options below.
	// b=18 – Candle-lighting time minutes before sunset (the default is 18). For Jerusalem, the default is b=40
	// M=on – Havdalah at nightfall – tzeit hakochavim, the point when 3 small stars are observable in the night sky with the naked eye (sun 8.5° below the horizon). This option is an excellent default for most places on the planet)
	// m=50 – Havdalah 50 minutes after sundown. This option is available for those whose minhag is to end Shabbat a fixed number of minutes after sundown. Typically one would enter 42 min for three medium-sized stars, 50 min for three small stars, 72 min for Rabbeinu Tam, or 0 to suppress Havdalah times. Set to m=0 (zero) to disable Havdalah times
	params = append(params, "&D=on") // – Hebrew date for dates with some event
	params = append(params, "&d=on") // – Hebrew date for entire date range
	params = append(params, "&o=on") // – Days of the Omer

	//As an alternative to specifying year=2021, you may specify a range of dates using both start and end:

	//start=2021-12-29 – Gregorian start date in YYYY-MM-DD format
	//end=2022-01-04 – Gregorian end date in YYYY-MM-DD format
	//Mutually exclusive options for Diaspora/Israel holidays and Torah Readings:

	//i=off – Diaspora holidays and Torah readings (default if unspecified)
	params = append(params, "&i=on") // – Israel holidays and Torah readings
	//Mutually exclusive location for candle-lighting times:

	//geo=none – no candle-lighting location (default if unspecified)
	//geo=geoname – location specified by GeoNames.org numeric ID
	//requires additional parameter geonameid=3448439
	//Hebcal.com supports approximately 47,000 different GeoNames IDs. These are cities with a population of 5000+. See cities5000.zip from https://download.geonames.org/export/dump/.
	//geo=zip – location specified by United States ZIP code
	//requires additional parameter zip=90210
	//geo=city – location specified by one of the Hebcal.com legacy city identifiers
	//requires additional parameter city=GB-London
	//geo=pos – location specified by latitude, longitude, and timezone. Requires additional 3 parameters:
	//latitude=[-90.0 to 90.0] – latitude in decimal format (e.g. 31.76904 or -23.5475)
	//longitude=[-180.0 to 180.0] – longitude decimal format (e.g. 35.21633 or -46.63611)
	//tzid=TimezoneIdentifier. See List of tz database time zones. Be sure to use the “TZ database name” such as America/New_York or Europe/Paris, not a UTC offset
	params = append(params, "&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel") // for Raanana
	//url = "https://www.hebcal.com/hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&month=6&ss=on&mf=on&c=on&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel&M=on&s=on"

	var url string
	for _, item := range params {
		url = url + item
	}
	fmt.Println(url)
	//synagogues =
	//{"title":"Candle lighting: 7:25pm","date":"2022-06-03T19:25:00+03:00","category":"candles","title_orig":"Candle lighting","hebrew":"הדלקת נרות","memo":"Parashat Bamidbar"}
}

func GetSynagogues() *Synagogues {
	return &synagogues
}
