

type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var sb strings.Builder

	for _, str:= range strs{
		sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteByte('#')
        sb.WriteString(str)
	}

	return sb.String()

}


func (s *Solution) Decode(encoded string) []string {

	eLength := len(encoded)
	results := make([]string, 0)

	for i := 0; i < eLength ;{

		j := i

		for encoded[j] != '#'{
			j++
		}

		length,_ := strconv.Atoi(encoded[i:j])

		start := j+1

		results = append(results, encoded[start: start+length])

		i = start+length

	}
    return results
}
