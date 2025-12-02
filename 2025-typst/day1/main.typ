
#{
  set page(margin: 1em, fill: luma(5%))
  set text(fill: luma(90%), font: "MonaspiceKr NFM")

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

    let out = [50 (0)\ ]
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

      let new-zeroes = 0


      while dial-position >= 100 {
        dial-position -= 100
        new-zeroes += 1
      }
      while dial-position < 0 {
        dial-position += 100
        new-zeroes += 1
      }

      if new-zeroes == 0 and dial-position == 0 {
        new-zeroes += 1
      }

      out += [#dial-position (#new-zeroes) \ ]

      num-zeroes += new-zeroes
    }
    [*#num-zeroes* \ ]
    out
  }

  final-dial-position(input)
}

// 7273 incorrect
