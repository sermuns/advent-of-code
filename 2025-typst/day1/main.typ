#import calc: ceil, floor, rem

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

  let calculate-zeroes(input) = {
    let lines = input.split("\n")

    let num-zeroes = 0
    let dial = 50

    let out = [[#sym.space #sym.space] 50 (0)\ ]
    for (dir, value) in lines
      .filter(l => l.len() > 0)
      .map(l => (l.at(0), int(l.slice(1)))) {
      let new-zeroes = 0

      if dir == "R" {
        dial += value
        new-zeroes += floor(dial / 100)
        dial = rem(dial, 100)
      } else {
        dial -= value
        new-zeroes += -floor(dial / 100)
        dial = rem(dial, 100)
        if dial < 0 {
          dial += 100
        }
      }

      out += [[#dir;#value] #dial (#new-zeroes) \ ]
      num-zeroes += new-zeroes
    }
    [*#num-zeroes* \ ]
    out
  }

  calculate-zeroes(input)
}

// 7273 incorrect
// 6789 is correct
