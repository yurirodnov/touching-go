// The program displays information about the system in which it was launched

package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
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

	fmt.Println("SYSTEM INFORMATION")
	fmt.Println("---------------------------->")

	fmt.Print("\n")

	fmt.Println("Operational system")
	fmt.Println("---------------------------->")
	
	hostInfo, _ := host.Info();
	uptime := time.Duration(hostInfo.Uptime) * time.Second	

	fmt.Printf("System uptime: %v\n", uptime)

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


	cpuUtillization, _ := cpu.Percent(time.Second, false)
	fmt.Printf("CPU utilaztion: %.2f%%\n", cpuUtillization[0])

	fmt.Println("<----------------------------")

	// Disk info block
	fmt.Println("Disk")
	fmt.Println("---------------------------->")
	diskInfo, _ := disk.Partitions(true)
	fmt.Printf("Disk %v info\n", diskInfo[0].Device) 
	fmt.Println("<----------------------------")



	
	
}