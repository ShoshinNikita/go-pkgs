package log_test

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/ShoshinNikita/go-pkgs/log"
)

const (
	tempFile = "temp.txt"
	message1 = "[DBG] Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	message2 = "[INF] Lorem ipsum dolor sit amet, consectetur adipiscing elit."
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Remove(tempFile)

	os.Exit(code)
}

func BenchmarkPrintlnSTD(b *testing.B) {
	f, _ := os.Create(tempFile)
	defer f.Close()

	stdlog.SetOutput(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stdlog.Println(message1)
		stdlog.Println(message2)
	}
}

func BenchmarkPrintlnDebugModeOn(b *testing.B) {
	f, _ := os.Create(tempFile)
	defer f.Close()

	log.GetStdLogger().SetOutput(f)

	// Turn debug mode on
	log.SetOptions(log.Debug)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println(message1)
		log.Println(message2)
	}
}

func BenchmarkPrintlnDebugModeOff(b *testing.B) {
	f, _ := os.Create(tempFile)
	defer f.Close()

	log.GetStdLogger().SetOutput(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println(message1)
		log.Println(message2)
	}
}
