package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/hazeliscoding/wx-cli/api"
	"github.com/hazeliscoding/wx-cli/config"
)

func DisplayWeather(weather *api.WeatherResponse, ipInfo config.IPInfoResponse) {
	today := weather.Forecast.ForecastDay[0]
	todayDate, _ := time.Parse("2006-01-02", today.Date)

	fmt.Println("┌────────────────────────────────────────────────────────────┐")
	fmt.Printf("│ 🌍 Location: %s, %s, %s\n", ipInfo.City, ipInfo.Region, ipInfo.Country)
	fmt.Printf("│ 📅 Date: %s | %s\n", todayDate.Format("01/02"), todayDate.Format("Monday"))
	fmt.Println("├────────────────────────────────────────────────────────────┤")
	fmt.Printf("│ 🌡️ Current: %s, %.0f°F (Feels like %.0f°F)\n", strings.ToLower(weather.Current.Condition.Text),
		today.Day.AvgTempF, weather.Current.FeelsLikeF)
	fmt.Printf("│ 🔽 Min: %.0f°F | 🔼 Max: %.0f°F\n", today.Day.MinTempF, today.Day.MaxTempF)
	fmt.Printf("│ 💧 Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Println("└────────────────────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("📅 6-Day Forecast:")
	fmt.Println("┌───────────┬───────────┬─────────────┬──────────────────────┬──────────┬──────────┐")
	fmt.Println("│   Date    │ Temp (°F) │ Rain Chance │      Condition       │ Sunrise  │  Sunset  │")
	fmt.Println("├───────────┼───────────┼─────────────┼──────────────────────┼──────────┼──────────┤")

	for i := range weather.Forecast.ForecastDay {
		day := weather.Forecast.ForecastDay[i]
		date, _ := time.Parse("2006-01-02", day.Date)
		condition := strings.ToLower(day.Day.Condition.Text)

		if len(condition) > 20 {
			condition = condition[:20] + "..."
		}

		dateStr := fmt.Sprintf("%s %s", date.Format("01/02"), date.Format("Mon"))
		rainChance := fmt.Sprintf("%.0f%%", day.Day.TotalPrecipIn)

		fmt.Printf(
			"│ %-9s │ %-9.0f │ %-11s │ %-20s │ %-8s │ %-8s │\n",
			dateStr,
			day.Day.AvgTempF,
			rainChance,
			condition,
			day.Astro.Sunrise,
			day.Astro.Sunset,
		)
	}

	fmt.Println("└───────────┴───────────┴─────────────┴──────────────────────┴──────────┴──────────┘")
}
