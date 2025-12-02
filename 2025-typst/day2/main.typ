#set page(margin: 1em, fill: luma(5%))
#set text(fill: luma(90%), font: "MonaspiceKr NFM")

#let input = read("input")
#let sample = (
  ```
    11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
  1698522-1698528,446443-446449,38593856-38593862,565653-565659,
  824824821-824824827,2121212118-2121212124
  ```
    .text
    .replace("\n", "")
    .trim()
)

#let is-repeated-twice(id) = {
  let id-str = str(id)
  let len = id-str.len()
  if calc.odd(len) {
    return false // impossible for id of odd length
  }
  let middle-idx = calc.floor(len / 2)
  id-str.slice(0, middle-idx) == id-str.slice(middle-idx)
}

#let get-invalid-ids(input) = {
  let invalid-ids = ()
  let id-ranges = input.split(",")
  for id-range in id-ranges {
    let (lower, upper) = id-range.split("-")
    for id in range(int(lower.trim()), int(upper.trim()) + 1) {
      if is-repeated-twice(id) {
        invalid-ids.push(id)
      }
    }
  }
  invalid-ids
}

#let invalid-ids = get-invalid-ids(input)
#invalid-ids.sum()
