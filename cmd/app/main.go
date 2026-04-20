package main

import (
	"os"

	"github.com/ladislavkosenko/lab1-tooling/internal/calc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	// 1. Налаштування красивого виводу логів для терміналу
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// 2. Ініціалізація Viper для читання конфігурації
	viper.SetConfigName("config") // ім'я файлу без розширення
	viper.SetConfigType("yaml")   // тип файлу
	viper.AddConfigPath(".")      // шукати файл у поточній директорії

	// Спроба прочитати файл
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Помилка читання файлу конфігурації")
	}

	// Читаємо значення з конфігурації
	appName := viper.GetString("app_name")
	port := viper.GetInt("port")

	// Виводимо лог інформаційного рівня
	log.Info().Str("app", appName).Int("port", port).Msg("Програма успішно запущена!")

	// Використовуємо наш калькулятор з першої лаби
	result, err := calc.Divide(10.0, 2.0)
	if err != nil {
		log.Error().Err(err).Msg("Помилка під час ділення")
		return
	}

	log.Info().Float64("result", result).Msg("Успішне обчислення")
}
