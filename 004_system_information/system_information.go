// Shows information about your system

package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)


func convertBytes(b uint64) string {
	const baseUnit uint64 = 1024
	
	if b < baseUnit{
		return fmt.Sprintf("%d B", b)
	}

	unitOfMeasure := []string{"KB", "MB", "GB", "TB", "PB", "EB"}

	div := uint64(baseUnit)
	level := 0

	for n := b / baseUnit; n >= baseUnit; n /= baseUnit {
		div *= baseUnit
		level++
	}

	return fmt.Sprintf("%.2f %s", float64(b) / float64(div), unitOfMeasure[level])

}

func main() {

	fmt.Print("\n")
	// OS info block
	fmt.Println("Operational system")
	fmt.Println("---------------------------->")

	osInfo := runtime.GOOS

	fmt.Printf("Operational system name: %s\n", osInfo)

	currentUser, err := user.Current()
	if(err != nil){
		fmt.Printf("Getting user error: %v", err)
	}
	
	fmt.Printf("User name: %v\n", currentUser.Name)
	fmt.Printf("Home directory: %v\n", currentUser.HomeDir)
	

	hostname, err := os.Hostname()
	if(err != nil){
		fmt.Printf("Getting hostname: %v\n", err)
	}

	fmt.Printf("Hostname: %v\n", hostname)
	fmt.Println("<----------------------------")
	fmt.Print("\n")

	// Memory info block

	fmt.Println("Memory")
	fmt.Println("---------------------------->")

	v, _ := mem.VirtualMemory()

	var totalMemory string = convertBytes(v.Total)
	var availableMemory string = convertBytes(v.Available)
	var usedMemory string = convertBytes(v.Used)
	var usedMemoryPercent int = int(v.UsedPercent)
	var availableMemoryPercent int = 100 - usedMemoryPercent

	fmt.Printf("Total memory: %v\n", totalMemory)
	fmt.Printf("Used memory: %v (%d%%)\n", usedMemory, usedMemoryPercent)
	fmt.Printf("Available memory: %v (%d%%)\n", availableMemory, availableMemoryPercent)
	fmt.Println("<----------------------------")
	fmt.Print("\n")

	// CPU info block
	fmt.Println("CPU")
	fmt.Println("---------------------------->")

	cpuInfo, _ := cpu.Info()
	cpuCoresLogical, _ := cpu.Counts(true)
	cpuCoresPhysical, _ := cpu.Counts(false)
		

	fmt.Printf("CPU name: %s\n", cpuInfo[0].ModelName)
	fmt.Printf("CPU cores logical: %d\n", cpuCoresLogical)
	fmt.Printf("CPU cores physical: %d\n", cpuCoresPhysical)
	fmt.Println("<----------------------------")


	
	
}