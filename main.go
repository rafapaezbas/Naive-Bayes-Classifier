package main

import(
	"strings"
	"os"
	"bufio"
	//"fmt"
)

func main(){

}

func calculateNBC(targets []string, dataset []string, labels []string) map[string]float32{

	nbc := make(map[string]float32)

	for _, label := range labels {
		labelPos := getLabelsPos(labels)[label]
		var targetsGivenClassPos float32 = 1;
		var targetsPos float32 = 1;
		for _,target := range targets{
			targetsGivenClassPos *= getTargetPosGivenClass(target,label,dataset,labels)
			targetsPos *= getTargetsPos(targets,dataset)[target]
		}

		nbc[label] = (labelPos * targetsGivenClassPos) / targetsPos
	}

	return nbc
}

func getLabelsPos(labels []string) map[string]float32 {

	tempLabelsPos := make(map[string]float32)

	for _,label := range labels {
		if tempLabelsPos[label] == 0 {
			var labelFreq float32
			for _, checkedLabel := range labels {
				if checkedLabel == label {
					labelFreq++
				}
			}
			tempLabelsPos[label] = labelFreq / float32(len(labels))
		}
	}
	return tempLabelsPos
}

//Given a set of words, how often do they appear in a given dataset
func getTargetsPos(targets []string, dataset []string)map[string]float32 {

	tempTargetsPos := make(map[string]float32)

	for _,target := range targets{
		var targetFreq float32
		for _,row := range dataset{
			if(strings.Contains(row,target)){
				targetFreq++
			}
		}
		tempTargetsPos[target] = targetFreq /float32(len(dataset))
	}
	return tempTargetsPos;
}

func getTargetPosGivenClass(target string, targetLabel string, dataset []string , labels []string) float32{

	var labelFreq float32
	var targetGivenLabelFreq float32

	for i,label := range labels{
		if targetLabel == label{
			labelFreq++
			if strings.Contains(dataset[i],target){
				targetGivenLabelFreq++
			}
		}
	}
	return targetGivenLabelFreq /labelFreq
}



//Returns a []string where each entry is a line o the sample
func readSample (path string) []string{
	sampleFile, _ := os.Open(path)
	defer sampleFile.Close()
	var sample []string
	scanner := bufio.NewScanner(sampleFile)
	for scanner.Scan(){
		var line = scanner.Text()
		sample = append(sample,line)
	}
	return sample
}
