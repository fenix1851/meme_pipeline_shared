package config

import (
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Config — единая структура для всех микросервисов.
type Config struct {
	App          AppConfig           `yaml:"app"`
	Database     DatabaseConfig      `yaml:"database"`
	RabbitMQ     *RabbitMQConfig     `yaml:"rabbitmq,omitempty"`
	Storage      *StorageConfig      `yaml:"storage,omitempty"`
	MinIO        *MinIOConfig        `yaml:"minio,omitempty"`
	OpenAI       *OpenAIConfig       `yaml:"openai,omitempty"`
	LabelStudio  *LabelStudioConfig  `yaml:"label_studio,omitempty"`
	ClientReddit *ClientRedditConfig `yaml:"client_reddit,omitempty"`
	Jobs         JobsConfig          `yaml:"jobs,omitempty"`
	Instagram    *InstagramConfig    `yaml:"instagram,omitempty"`
}

// AppConfig — настройки приложения.
type AppConfig struct {
	Env         string `yaml:"env"`
	LogLevel    string `yaml:"log_level"`
	LogFilePath string `yaml:"log_file_path"`
}

type InstagramConfig struct {
	AccessToken  string `yaml:"access_token"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	UserID       string `yaml:"user_id"`
}

// DatabaseConfig — настройки подключения к БД.
type DatabaseConfig struct {
	URI string `yaml:"uri"`
}

// RabbitMQConfig — настройки RabbitMQ.
type RabbitMQConfig struct {
	URL   string `yaml:"url"`
	Queue string `yaml:"queue"`
}

// StorageConfig — настройки хранения файлов.
type StorageConfig struct {
	Path string `yaml:"path"`
}

// MinIOConfig — настройки MinIO.
type MinIOConfig struct {
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	UseSSL    bool   `yaml:"use_ssl"`
}

// OpenAIConfig — настройки для работы с OpenAI.
type OpenAIConfig struct {
	APIKey      string  `yaml:"api_key"`
	Model       string  `yaml:"model"`
	MaxTokens   int     `yaml:"max_tokens"`
	Temperature float64 `yaml:"temperature"`
}

// LabelStudioConfig — настройки Label Studio.
type LabelStudioConfig struct {
	APIURL    string `yaml:"api_url"`
	APIToken  string `yaml:"api_token"`
	ProjectID string `yaml:"project_id"`
}

// ClientRedditConfig — настройки для подключения к Reddit API.
type ClientRedditConfig struct {
	AccessToken  string         `yaml:"access_token"`
	RefreshToken string         `yaml:"refresh_token"`
	AuthData     AuthDataConfig `yaml:"auth_data"`
	URL          string         `yaml:"url"`
}

// AuthDataConfig — данные для авторизации в Reddit.
type AuthDataConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	UserAgent    string `yaml:"user_agent"`
}

// JobSettings — общие настройки для задачи.
type JobSettings struct {
	Enabled  bool   `yaml:"enabled"`
	Interval string `yaml:"interval"`
}

// JobParsePostsConfig — настройки задачи по парсингу постов.
type JobParsePostsConfig struct {
	Interval      string `yaml:"interval"`
	PeriodOfParse string `yaml:"period_of_parse"`
	LimitOfPosts  string `yaml:"limit_of_posts"`
}

// JobsConfig — объединение конфигураций задач.
type JobsConfig struct {
	// Например, задача генерации мемов (раньше называлась meme_crossing_sheduler)
	MemeCrosser *JobSettings `yaml:"meme_crossing_scheduler,omitempty"`
	// Дополнительные задачи (например, для сбора/разметки мемов)
	MemeFetching  *JobSettings         `yaml:"meme_fetching,omitempty"`
	MemeLabeling  *JobSettings         `yaml:"meme_labeling,omitempty"`
	ParsePosts    *JobParsePostsConfig `yaml:"job_parse_posts,omitempty"`
	TopicAnalysis *JobSettings         `yaml:"topic_analysis,omitempty"`
}

// interpolateEnv выполняет замену шаблонов ${VAR_NAME} на значения переменных окружения.
func interpolateEnv(input string) string {
	re := regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		groups := re.FindStringSubmatch(match)
		if len(groups) == 2 {
			if val, ok := os.LookupEnv(groups[1]); ok {
				return val
			}
		}
		// Если переменная не найдена, оставляем исходное значение
		return match
	})
}

// interpolateConfig выполняет интерполяцию переменных окружения во всём YAML-содержимом.
func interpolateConfig(data []byte) []byte {
	return []byte(interpolateEnv(string(data)))
}

// LoadConfig загружает конфигурацию из YAML-файла с учётом подстановки переменных окружения.
func LoadConfig(configPath string) (*Config, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	interpolated := interpolateConfig(raw)

	var cfg Config
	if err := yaml.Unmarshal(interpolated, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
