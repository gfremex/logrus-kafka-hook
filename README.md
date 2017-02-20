## logrus-kafka-hook


A [logrus.Hook](https://godoc.org/github.com/Sirupsen/logrus#Hook) which sends a single
log entry to multiple kafka topics simultaneously.

## How to use

### Import package

```Go
import lkh "github.com/gfremex/logrus-kafka-hook"
```

### Create a hook (KafkaHook)

```Go
NewKafkaHook(id string, levels []logrus.Level, formatter logrus.Formatter, brokers []string) (*KafkaHook, error)
```

- id: Hook Id
- levels: [logrus.Levels](https://godoc.org/github.com/Sirupsen/logrus#Level) supported by the hook
- formatter: [logrus.Formatter](https://godoc.org/github.com/Sirupsen/logrus#Formatter) used by the hook
- brokers: Kafka brokers

For example:

```Go
hook, err := lkh.NewKafkaHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{},
		[]string{"192.168.60.5:9092", "192.168.60.6:9092", "192.168.60.7:9092"},
	)
```

### Create a [logrus.Logger](https://godoc.org/github.com/Sirupsen/logrus#Logger)

For example:

```Go
logger := logrus.New()
```

### Add hook to logger

```Go
logger.Hooks.Add(hook)
```

### Topic

For only one topic pass the Kafka topic name as the Id.

### Send messages to logger

For example:

```Go
l.Debug("This must not be logged")

l.Info("This is an Info msg")

l.Warn("This is a Warn msg")

l.Error("This is an Error msg")
```

#### Complete examples

[https://github.com/gfremex/logrus-kafka-hook/tree/master/examples](https://github.com/gfremex/logrus-kafka-hook/tree/master/examples)
