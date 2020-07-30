package find_unique_number

/*
There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55

It’s guaranteed that array contains at least 3 numbers.

The tests contain some very huge arrays, so think about performance.
*/


func FindUniq(arr []float32) float32 {
  listToValidate := arr[0:3]
  counter := make(map[float32]int)
  var numberToFilter float32
  for _, num := range listToValidate{
    counter[num] = counter[num] + 1
    if counter[num] >= 2 {
      numberToFilter = num
    }
  }

  var resultNumber float32
  for _, num := range arr {
    if num != numberToFilter {
      resultNumber = num
      break
    }
  }
  return resultNumber
}
