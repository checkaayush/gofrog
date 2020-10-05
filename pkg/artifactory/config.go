package artifactory

// Config represents Artifactory API configuration
type Config struct {
	BaseURL  string `toml:"base_url" validate:"required,url"`
	User     string `toml:"user" validate:"required"`
	Password string `toml:"password" validate:"required"`
}
