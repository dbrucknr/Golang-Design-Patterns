package examples

import "fmt"

type WeatherObserver interface {
	Update(data WeatherData)
}

type WeatherSubject interface {
	RegisterObserver(observer WeatherObserver)
	RemoveObserver(observer WeatherObserver)
	PushNotification()
}

type WeatherData struct {
	Temperature float64
	WindSpeed   float64
	Pressure    float64
}

type WeatherStation struct {
	temperature float64
	windSpeed   float64
	pressure    float64
	observers   []WeatherObserver
}

func (ws *WeatherStation) RegisterObserver(observer WeatherObserver) {
	ws.observers = append(ws.observers, observer)
}

func (ws *WeatherStation) RemoveObserver(observer WeatherObserver) {
	for i, obs := range ws.observers {
		// If this pattern were to be adjusted so that the RemoveObserver method accepted a custom domain param object
		// - This match logic will need to be modified
		if obs == observer {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
			return
		}
	}
}

func (ws *WeatherStation) PushNotification() {
	// NOTE - all instances of ws.observers must satisfy the WeatherObserver interface constraint
	//   - Enables calling of `observer.Update(WeatherData{...`
	for _, observer := range ws.observers {
		observer.Update(WeatherData{
			Temperature: ws.temperature,
			WindSpeed:   ws.windSpeed,
			Pressure:    ws.pressure,
		})
	}
}

func (ws *WeatherStation) SetMeasurements(temp, wind, pressure float64) {
	ws.temperature = temp
	ws.windSpeed = wind
	ws.pressure = pressure
	ws.PushNotification()
}

// UI Concept / Sample
// - This could be a real GUI link, that receives updates over an SSE channel
// - There could be a special method that would forward traffic to a different entity
type UserInterface struct {
}

func (ui *UserInterface) Display(temperature, windSpeed, pressure float64) {
	fmt.Printf("UI: Temperature: %.2f, Wind Speed: %.2f, Pressure: %.2f\n", temperature, windSpeed, pressure)
}
func (ui *UserInterface) Update(data WeatherData) {
	ui.Display(data.Temperature, data.WindSpeed, data.Pressure)
}

// Logger Concept / Sample
// - This represents a system dependency / entity that could be a native system logger or a remote logging service
// - The Log method could be designed to broadcast / submit data to a remote repository
type Logger struct {
}

func (l *Logger) Log(temperature, windSpeed, pressure float64) {
	fmt.Printf("Logger: Temperature: %.2f, Wind Speed: %.2f, Pressure: %.2f\n", temperature, windSpeed, pressure)
}
func (l *Logger) Update(data WeatherData) {
	l.Log(data.Temperature, data.WindSpeed, data.Pressure)
}

// Alert Concept / Sample
// - This represents a service that could be responsible for activating internal or external remote alerts
// - The internal Alert Method could be expanded / multiplied to only be called on certain conditions (severity, weather type, etc)
type AlertSystem struct {
}

func (as *AlertSystem) Alert(temperature, windSpeed, pressure float64) {
	fmt.Printf("Alert: Temperature: %.2f, Wind Speed: %.2f, Pressure: %.2f\n", temperature, windSpeed, pressure)
}
func (as *AlertSystem) Update(data WeatherData) {
	// Set conditions for alerting
	if data.WindSpeed > 50 {
		// Severe Wind Alert
		as.Alert(data.Temperature, data.WindSpeed, data.Pressure)
	}

	if data.Temperature > 80 {
		// High Temperature Alert
		as.Alert(data.Temperature, data.WindSpeed, data.Pressure)
	}

	//...
}
