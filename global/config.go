package global

type Config struct {
	Ip            string `ini:"ip"`
	Port          int    `ini:"port"`
	Deduplication int    `ini:"deduplication"`
}
