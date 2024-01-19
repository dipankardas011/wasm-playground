cargo_component_bindings::generate!();

use std::time::Duration;

use bindings::docs::calculator::calculate::evalexpression;

fn main() {
    let result = evalexpression(1 as u32,1 as u32);
    println!("1 + 1 = {result}");
    std::thread::sleep(Duration::from_millis(10));
}
