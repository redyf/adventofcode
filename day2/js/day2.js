/* Day 2: I Was Told There Would Be No Math-- -

  The elves are running low on wrapping paper, and so they need to submit an order for more.They have a list of the dimensions(length l, width w, and height h) of each present, and only want to order exactly as much as they need.

    Fortunately, every present is a box(a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2 * l * w + 2 * w * h + 2 * h * l.The elves also need a little extra paper for each present: the area of the smallest side.

For example:

    A present with dimensions 2x3x4 requires 2 * 6 + 2 * 12 + 2 * 8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
    A present with dimensions 1x1x10 requires 2 * 1 + 2 * 10 + 2 * 10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.

All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

/* Observações: l (length), w (width), h (height)
 * Cada presente é uma caixa (paralelepípedo)
 * Para achar a área da superfície da caixa, precisamos usar a fórmula w*l*w + 2*w*h + 2*h*l*/

// Parte 1
function calculaArea(length, width, height) {
  let area1 = 2 * length * width
  let area2 = 2 * width * height
  let area3 = 2 * height * length
  let areaTotal = area1 + area2 + area3
  let arrayAreas = [area1, area2, area3]
  let areaMenor = Math.min(...arrayAreas) / 2
  console.log(`Área 1 é:${area1}`)
  console.log(`Área 2 é:${area2}`)
  console.log(`Área 3 é:${area3}`)
  console.log(`A área total da caixa é: ${areaTotal}`)
  console.log(`Você vai precisar de ${areaMenor} papel extra`)
  // return areaTotal
  return areaMenor
}

console.log(calculaArea(2, 3, 4))
console.log(calculaArea(1, 1, 10))

// Parte 2 (Ainda não acabei) 
const fs = require('fs');
const DIMENSOES = fs.readFileSync('./dimensoes.txt', 'utf-8').split('\n');
