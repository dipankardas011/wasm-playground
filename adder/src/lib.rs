cargo_component_bindings::generate!();

// Separating out the interface puts it in a sub-module
use bindings::exports::docs::adder::add::Guest;

struct Component;

impl Guest for Component {
    fn add(a: u32, b: u32) -> u32 {
        a + b
    }
}
