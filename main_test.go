package main

import (
	"testing"
	"fmt"
)

func TestGetLabelsPos(t *testing.T){
	labels := readSample("./datasets/labels")
	labelsPos := getLabelsPos(labels)

	if(labelsPos["flu"] != 0.5){
		t.Errorf("Labels posibilities was incorrect")
	}

	if(labelsPos["concussion"] < 0.3 || labelsPos["concussion"] > 0.4){ //Real value is 0.33
		t.Errorf("Labels posibilities was incorrect")
	}

}

func TestGetTargetPos(t *testing.T){
	targets := []string{"builder","sneezing"}
	dataset := readSample("./datasets/samples")
	targetsPos := getTargetsPos(targets,dataset)
	if(targetsPos["sneezing"] != 0.5){
		t.Errorf("Targets posibilities was incorrect")
	}
}

func TestGetTargetPosGivenClass( t *testing.T){
	labels := readSample("./datasets/labels")
	dataset := readSample("/home/rafa/Escritorio/dataset/samples")

	if(getTargetPosGivenClass("sneezing","flu",dataset,labels) < 0.6 || getTargetPosGivenClass("builder","flu",dataset,labels) > 0.7){
		t.Errorf("Target posibilitie given class was incorrect")
	}

	if(getTargetPosGivenClass("builder","concussion",dataset,labels) != 0.5){
		t.Errorf("Target posibilitie given class was incorrect")
	}
}

func TestCalculateNBC(t *testing.T){

	labels := readSample("./datasets/labels")
	dataset := readSample("./datasets/samples")

	nbc := calculateNBC([]string{"sneezing","builder"},dataset,labels);
	fmt.Println(nbc)
	if(nbc["flu"] < 0.6 || nbc["flu"] > 0.7){
		t.Errorf("nbc  was incorrect")
	}
}
