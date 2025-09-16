package main // โปรแกรมหลักที่รันได้ (entry point)

import (
	// โมดูลภายในโปรเจกต์: อ่านคอนฟิก (เช่น พอร์ต) และลงทะเบียนเส้นทาง (routes)
	"context"
	"jrppor/go-expense-tracker-api/config"
	"jrppor/go-expense-tracker-api/routes"
	"net/http"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	// ไลบรารีมาตรฐานสำหรับล็อกและอ่าน environment variables
	"log"
	"os"

	// เว็บเฟรมเวิร์ก Gin
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
)

func main() {
	// ข้อความบอกเริ่มบูตเซิร์ฟเวอร์
	log.Println("Booting API server…")

	// สร้าง Gin Engine และตั้งค่าโหมดจาก ENV ชื่อ GIN_MODE (ถ้าไม่ตั้งจะเป็นโหมด debug)
	if m := os.Getenv("GIN_MODE"); m != "" {
		gin.SetMode(m)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	// ไม่เชื่อถือ proxy ใดๆ โดยค่าเริ่มต้น (ป้องกันปัญหา header จาก proxy)
	_ = router.SetTrustedProxies(nil)

	// ตั้งค่าเส้นทาง API
	routes.SetupRoutes(router)

	// เชื่อมต่อฐานข้อมูล Postgres ทันทีตอนบูต (อ่านค่าจาก ENV/.env)
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	if sqlDB, e := db.DB(); e == nil {
		_ = sqlDB.Ping()
	}
	log.Println("Database connected")

	// Apply migrations (if migrations/ exists) before starting the server
	if sqlDB, e := db.DB(); e == nil {
		wd, _ := os.Getwd()
		dir := filepath.Join(wd, "migrations")
		if _, statErr := os.Stat(dir); os.IsNotExist(statErr) {
			log.Printf("No migrations directory at %s; skipping migrations", dir)
		} else {
			if err := goose.SetDialect("postgres"); err != nil {
				log.Printf("goose dialect error: %v", err)
			} else if err := goose.Up(sqlDB, dir); err != nil {
				log.Printf("goose up error: %v", err)
			} else {
				log.Println("Migrations applied")
			}
		}
	}

	// สร้าง HTTP server เพื่อรองรับการปิดแบบนุ่มนวล (graceful shutdown)
	port := config.GetServerPort()
	srv := &http.Server{Addr: ":" + port, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// รอรับสัญญาณจากระบบ (SIGINT/SIGTERM) เพื่อสั่งปิดเซิร์ฟเวอร์แบบนุ่มนวล
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// สร้าง context ที่มี timeout 5 วินาที สำหรับรอให้รีเควสต์ค้างอยู่ทำงานเสร็จ
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
