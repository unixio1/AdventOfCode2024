package main(
	"adventOfCode2024/customScanner"
	"strings"
	"strconv"
	"adventOfCode2024/myMath"
)

func main(){
	

}

func partOne(fileName string) int{
	scanner := CustomScanner.NewScanner(fileName)
	defer scanner.Close()
	rules := getRulesMap(scanner)

}

func getRulesMap(scanner CustomScanner) map[int][]int{
	rules := make(map[int][]int)
	for scanner.Scan(){
		line := scanner.Text()
		if line == ""{
			break
		}
		splitted, _ :=  strings.Split(line, "|")
		firstPage := strconv.Atoi(splitted[0])
		secondPage := strconv.Atoi(splitted[1])
		_, ok := rules[firstPage]
		if ok{
			rules[firstPage] = append(rules[firstPage], secondPage)
		}else{
			rules[firstPage] = []int{secondPage}
		}
	}
	return rules
}

func processData(scanner CustomScanner, rules map[int][]int) int{
	for scanner.Scan(){
		line := scanner.Text()
		splitted := strings.Split(line, ",")

	}
}

func processLine(line []string, rules map[int][]int) int, err {
	valuesSaw := make([]int, 0, len(string))
	for _, elem := range line{
		intValue, _ := strconv.Atoi(elem)
		values, ok := rules[intValue]
		if !ok{
			valuesSaw = append(valuesSaw, intValue)
			continue
		}
		if valuesSaw.Contains(intValue){
			continue
		}
		for _, value := range values{
			if valuesSaw.Contains(value){
				return -1, newError("Invalid line")
			}
			valuesSaw = append(valuesSaw, value)
		}
	}
	middleIndex := math.Ceil(float32(len(valuesSaw))/2)
	middleElement := strconv.Atoi(valuesSaw[middleIndex])
	return middleElement, nil
}