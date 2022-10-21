package config

import "github.com/kelseyhightower/envconfig"

type Config struct {

    // TCP port
    Port            string  `envconfig:"PORT"`

    // Timeout (seconds) for incoming calls
    Timeout         int     `envconfig:"TIMEOUT"`

    // Twilio API key
    TwilioApiKey    string  `envconfig:"TWILIO_API_KEY"`
    TwilioApiSecret string  `envconfig:"TWILIO_API_KEY_SECRET"`
}

func NewDefaultConfig() (*Config, error) {
    cfg := new(Config)

    err := envconfig.Process("", cfg)
    if err != nil {
        return nil, err
    }

    return cfg, nil
}
