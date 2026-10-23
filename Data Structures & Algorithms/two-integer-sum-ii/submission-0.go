func twoSum(numbers []int, target int) []int {


      i,j := 0, len(numbers)-1
	  for i < j {

		value := numbers[i]+ numbers[j]
         
		 if  value > target{
			j--
		 }else if  value < target{
			i++
		 }else{
			return []int{i+1,j+1}
		 }
	}

	return []int{}
}
