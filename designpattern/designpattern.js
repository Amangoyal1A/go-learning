/* -------------------------------------------------------------------------- */
/*                                 single instace                             */
/* -------------------------------------------------------------------------- */

class Singleton {
    constructor() {
        if (Singleton.instance) return Singleton.instance;
        this.timestamp = Date.now();
        Singleton.instance = this;
    }
}

const a = new Singleton();
const b = new Singleton();
console.log(a === b); // true


/* -------------------------------------------------------------------------- */
/*                                 factory pattern                            */
/* -------------------------------------------------------------------------- */

class Car { drive() { console.log("Driving a car"); } pen() { console.log("pen")} }
class Truck { drive() { console.log("Driving a truck"); } }

class VehicleFactory {
  static getVehicle(type) {
    if (type === "car") return new Car();
    if (type === "truck") return new Truck();
  }
}

const vehicle = VehicleFactory.getVehicle("car");
vehicle.pen(); // Driving a car
