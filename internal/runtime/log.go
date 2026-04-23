package runtime

import "github.com/sirupsen/logrus"

var (
	log    *logrus.Logger
	logger *logrus.Entry
)

// NewLogger initializes the package-level logger with the given level.
// If level is empty or invalid, it falls back to InfoLevel.
// Accepted values: trace, debug, info, warn, warning, error, fatal, panic.
func NewLogger(level string) *logrus.Entry {
	if log == nil {
		log = logrus.New()
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	log.SetLevel(lvl)

	logger = logrus.NewEntry(log)
	return logger
}

func Logger() *logrus.Entry {
	if logger == nil {
		NewLogger("info")
	}

	return logger
}

// SetLevel updates the current logger level at runtime.
func SetLevel(level string) error {
	if log == nil {
		NewLogger(level)
		return nil
	}
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
