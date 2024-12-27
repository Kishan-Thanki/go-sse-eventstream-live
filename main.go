package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", ":8080", "The address to listen on")

func main() {
	flag.Parse()

	// Create a server instance
	server := &http.Server{
		Addr:    *addr,
		Handler: http.DefaultServeMux,
	}

	// Listen for shutdown signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		slog.Info("Server started")
		http.HandleFunc("/", home)
		http.HandleFunc("/events", events)

		// Start the server and listen for errors
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// Wait for a shutdown signal
	sig := <-done
	slog.Info("Received signal, shutting down", slog.String("signal", sig.String()))

	// Graceful shutdown with timeout
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error during server shutdown", slog.String("error", err.Error()))
	} else {
		slog.Info("Server shut down gracefully")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "typer.html")
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Alert message
	alert := []string{"Alert:", "Someone", "just", "forgot", "a", "semicolon", "in", "production!"}

	// Stream each word of the alert message
	for _, word := range alert {
		content := fmt.Sprintf("data: %s\n\n", word)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 500)
	}

	// Reaction message
	reaction := []string{"I", "just", "deployed", "a", "week's", "work.", "Wait!", "Whattttt???", "Nooooo!"}

	// Stream each word of the reaction message
	for _, word := range reaction {
		content := fmt.Sprintf("data: %s\n\n", word)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 500)
	}

	// Panic message
	panicMsg := []string{"ERROR:", "The", "server", "is", "now", "in", "panic", "mode...", "Please", "consider", "rebooting", "your", "life...", "and", "the", "server 🖥️⚡💔!"}

	// Stream each word of the panic message
	for _, word := range panicMsg {
		content := fmt.Sprintf("data: %s\n\n", word)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 500)
	}

	// Log the panic on the server side
	fmt.Println("Connection lost due to panic!")
}
