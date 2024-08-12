package service

import "github.com/souvik03-136/more-torque/pkg/client"

func DecodeVIN(vin string) (interface{}, error) {
	// Call NHTSA client and decode VIN
	return client.DecodeVIN(vin)
}

func AddVehicle(vin string, org string) (interface{}, error) {
	// Add vehicle to the system
	return nil, nil
}

func GetVehicle(vin string) (interface{}, error) {
	// Retrieve vehicle details from the system
	return nil, nil
}
