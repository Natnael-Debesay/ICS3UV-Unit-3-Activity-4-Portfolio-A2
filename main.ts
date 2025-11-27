/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-11-26
 * @fileoverview This program decides whether you should keep going on a car ride or not depending on your gas tank.
 */

// variables
let capacity: string
let gasGauge: string
let averageMileage: string

// input
capacity = prompt("Car Gas Capacity (L): ") || "0";
gasGauge = prompt("Gas Gauge (in percentage): ") || "0";
averageMileage = prompt("Average km/L mileage value for the car: ") || "0";

// convert inputs
const tankCapicity = parseFloat(capacity);
const gasPercent = parseFloat(gasGauge);
const mileage = parseFloat(averageMileage);

const fuelLeft = (gasPercent / 100) * tankCapicity;
const distancePossible = fuelLeft * mileage;

// output
if (distancePossible >= 50) {
  console.log(
    `Safe to proceed`
  );
} else {
  console.log(
    `Get Gas`
  );
}

console.log("\nDone.");
