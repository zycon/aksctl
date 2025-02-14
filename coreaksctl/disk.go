package coreaksctl

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func CreateDisk(diskName string, diskResourcegroup string, diskLocation string, diskSize string) {

	//Create AKS Disk
	cmd := exec.Command("az", "disk", "create", "-g", diskResourcegroup, "-n",
		diskName, "-l", diskLocation, "-z", diskSize)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	color.Green("Disk Created")
}

func DeleteDisk(diskName string, diskResourceGroup string) {

	//Delete AKS Disk
	cmd := exec.Command("az", "disk", "delete", "--name", diskName,
		"--resource-group", diskResourceGroup, "--yes")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Disk Deleted")
}

func UpdateDisk(diskName string, diskResourceGroup string, diskSize string) {

	//Update AKS Disk
	cmd := exec.Command("az", "disk", "update", "--name", diskName,
		"--resource-group", diskResourceGroup, "--size-gb", diskSize)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Disk Updated")
}

func GetDisk(diskResourceGroup string) {

	//List AKS Disk
	cmd := exec.Command("az", "disk", "list",
		"--resource-group", diskResourceGroup)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
