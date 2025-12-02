#{
  let input = read("input")

  let sample = ```
  L68
  L30
  R48
  L5
  R60
  L55
  L1
  L99
  R14
  L82
  ```.text

  let final-dial-position(input) = {
    let lines = input.split("\n")

    let num-zeroes = 0
    let dial-position = 50

    for line in lines {
      if line.len() == 0 {
        continue
      }

      let dir = line.at(0)
      let value = int(line.slice(1))

      if dir == "L" {
        dial-position -= value
      } else {
        dial-position += value
      }

      if dial-position == 0 {
        num-zeroes += 1
      }

      if dial-position >= 100 {
        dial-position -= 100
      } else if dial-position < 0 {
        dial-position += 100
      }
      [#dial-position \ ]
    }
    // num-zeroes
  }

  final-dial-position(sample)
}
