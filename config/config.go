package config

import (
	"os"
	"strings"

	// Viper: ใช้โหลดค่าคอนฟิกจากไฟล์และตัวแปรแวดล้อม
	"github.com/spf13/viper"
	// Zap: ตัว logger ประสิทธิภาพสูง
	"go.uber.org/zap"
)

// สร้าง logger สำหรับใช้ภายในแพ็กเกจนี้
var logger, _ = zap.NewProduction()

// LoadConfig: เตรียมค่า Viper ให้อ่านคอนฟิกจากไฟล์และตัวแปรแวดล้อม
func LoadConfig() {
	// กำหนดชื่อ/ชนิดไฟล์คอนฟิก และที่อยู่ของไฟล์ (รูทโปรเจกต์ และโฟลเดอร์ config/)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// อ่านไฟล์คอนฟิกถ้ามี (ถ้าไม่มีจะไม่ถือว่าเป็นความผิดพลาดร้ายแรง)
	if err := viper.ReadInConfig(); err != nil {
		// ล็อกเตือนเฉพาะกรณีที่ไม่ใช่ข้อผิดพลาด "ไม่พบไฟล์"
		if !strings.Contains(strings.ToLower(err.Error()), "not found") {
			logger.Warn("Could not read config file", zap.Error(err))
		}
	}

	// เปิดให้อ่านค่าจากตัวแปรแวดล้อม โดยแปลง key ที่มีจุดเป็นขีดล่าง
	// เช่น server.port -> SERVER_PORT
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// GetServerPort: คืนค่าพอร์ตของเซิร์ฟเวอร์ โดยลำดับความสำคัญคือ
// 1) ENV: PORT  2) config: server.port  3) ค่าเริ่มต้น 8080
func GetServerPort() string {
	LoadConfig()
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	if v := viper.GetString("server.port"); v != "" {
		return v
	}
	return "8080"
}
