package configuration

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/panjf2000/gnet"
)

// Default configuration for the server if no options are given
// TCP_PORT="9000"
// REUSE_PORT="false"
// MULTICORE_MODE="false"
// LOAD_BALANCING="0"

type Config struct {
	Port          string             `json:"port"`
	ReusePort     bool               `json:"max_connections"`
	MultiCore     bool               `json:"multicore"`
	LoadBalancing gnet.LoadBalancing `json:"load_balancing"`
}

func GetConfiguraation() Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// This function should load the configuration from a file or environment variables.
	// For now, we return a default configuration.
	return Config{
		Port:          getPort(),                // getPort(),
		ReusePort:     getReusePortOption(),     // Default to false
		MultiCore:     getMultiCoreOption(),     // Default to false
		LoadBalancing: getLoadBalancingOption(), // Default to RoundRobin
	}
}

func getPort() string {
	// This function should return the port number from the configuration.
	// For now, we return a default value.
	port := "9000"
	envValue := os.Getenv("TCP_PORT")

	if envValue != "" {
		port = envValue
	}

	return port
}

// TCP_PORT="9000"
// REUSE_PORT="false"
// MULTICORE_MODE="false"
// LOAD_BALANCING="0"

func getReusePortOption() bool {
	// This function should return the port number from the configuration.
	// For now, we return a default value.
	reusePort := false
	envValue := os.Getenv("REUSE_PORT")

	if envValue != "" {
		parsedVal, err := strconv.ParseBool(envValue)
		if err != nil {
			return false
		}

		reusePort = parsedVal
	}

	return reusePort
}

func getMultiCoreOption() bool {
	// This function should return the port number from the configuration.
	// For now, we return a default value.
	multiCore := false
	envValue := os.Getenv("MULTICORE_MODE")

	if envValue != "" {
		parsedVal, err := strconv.ParseBool(envValue)
		if err != nil {
			return false
		}

		multiCore = parsedVal
	}

	return multiCore
}

func getLoadBalancingOption() gnet.LoadBalancing {
	loadBalancing := gnet.RoundRobin
	envValue := os.Getenv("LOAD_BALANCING")

	if envValue != "" {
		parsedVal, err := strconv.Atoi(envValue)
		if err != nil {
			return gnet.RoundRobin
		}

		switch parsedVal {
		case 0:
			loadBalancing = gnet.RoundRobin
		case 1:
			loadBalancing = gnet.LeastConnections
		case 2:
			loadBalancing = gnet.SourceAddrHash
		default:
			loadBalancing = gnet.RoundRobin
		}
	}

	return loadBalancing
}
