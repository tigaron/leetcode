function countOfAtoms(formula: string): string {
  const stack = [new Map()];

  let i = 0;
  const n = formula.length

  for (; i < n ;) {
      if (formula[i] === "(") {
          stack.push(new Map());
          i++
      } else if (formula[i] === ")") {
          i++
          const start = i;
          while (i < n && /^\d+$/.test(formula[i])) {
              i++
          }
          const elementMultiply = formula.substring(start, i);
          
          let multiply;
          if (Number(elementMultiply) === 0) {
              multiply = 1;
          } else {
              multiply = elementMultiply;
          }

          const lastElement = stack.at(-1);
          stack.pop();

          for (let [key, value] of lastElement) {
              let total = stack.at(-1).get(key)
              if (total !== undefined) {
                  total += value*multiply
              } else {
                  total = value*multiply
              }
              stack.at(-1).set(key, total)
          }
      } else {
          let start = i;
          i++

          while (i < n && /^[a-z]$/.test(formula[i])) {
              i++
          }

          const element = formula.substring(start, i)

          start = i
          while (i < n && /^\d$/.test(formula[i])) {
              i++
          }

          const elementCount = formula.substring(start, i);

          let count;
          if (Number(elementCount) === 0) {
              count = 1;
          } else {
              count = elementCount;
          }

          let total = 0;
          let existingElement = stack.at(-1).get(element)
          if (existingElement !== undefined) {
              total = Number(count) + Number(existingElement)
          } else {
              total += Number(count)
          }
          stack.at(-1).set(element, total)
      }
  }

  const sorted = new Map([...stack[0]].sort((a, b) => String(a[0]).localeCompare(b[0])))

  let result = ""
  for (let [key, value] of sorted) {
      result += key
      if (value > 1) {
          result += value
      }
  }

  return result;
};
