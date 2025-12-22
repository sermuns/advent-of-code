#import calc: ceil, floor, rem

#{
  let input1 = read("input1")

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

  let get-num-zeroes(input) = {
    let out = []
    let num-zeroes = 0
    let dial-position = 50
    out += [#dial-position (0) \ ]

    for line in input.split("\n").filter(l => l.len() > 0) {
      let dir = line.at(0)
      let value = int(line.slice(1))

      if dir == "L" {
        dial-position -= value
      } else if dir == "R" {
        dial-position += value
      }


      let new-zeroes = 0

      new-zeroes += calc.abs(calc.div-euclid(dial-position, 100))
      while dial-position < 0 {
        dial-position += 100
      }
      dial-position = calc.rem-euclid(dial-position, 100)
      if dial-position == 0 {
        // new-zeroes += 1
      }

      num-zeroes += new-zeroes
      out += [#dial-position (#new-zeroes) \ ]
    }
    [
      *#num-zeroes* \
      #out
    ]
  }

  get-num-zeroes(input1)
}

// 7273 incorrect
// 6789 is correct
