package controller

import (
	"encoding/binary"
	"fmt"
	"os"
)


func RevpiIO() {
	// Replace this path with the actual PDI path for your analog input
	pdiFilePath := "/dev/piControl0"

	// Open the PDI device file for reading
	pdiFile, err := os.OpenFile(pdiFilePath, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Error opening PDI file:", err)
		return
	}
	defer pdiFile.Close()

	// Seek to the offset where the analog input value is stored.
	// Replace 'offset' with the actual byte offset for your specific analog input.
	offset := int64(0) // Change this to the actual offset
	_, err = pdiFile.Seek(offset, 0)
	if err != nil {
		fmt.Println("Error seeking to offset:", err)
		return
	}

	// Read 2 bytes (16 bits) from the PDI device file
	buffer := make([]byte, 2)
	_, err = pdiFile.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from PDI file:", err)
		return
	}

	// Convert the 2 bytes to a 16-bit unsigned integer
	analogValue := binary.LittleEndian.Uint16(buffer)

	// Convert the 16-bit unsigned integer to a voltage between 0 and 5V
	voltage := (float64(analogValue) / 65535.0) * 5.0

	// Print the analog input value and the calculated voltage
	fmt.Printf("Analog Value: %d\n", analogValue)
	fmt.Printf("Voltage: %.2f V\n", voltage)
}