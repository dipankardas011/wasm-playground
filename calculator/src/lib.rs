cargo_component_bindings::generate!();

use bindings::exports::docs::calculator::calculate::Guest;

// Bring the imported add function into scope
use bindings::docs::adder::add::add;

struct Component;

impl Guest for Component {
    fn evalexpression(a: u32, b: u32) -> u32 {
        // Cleverly parse `expr` into values and operations, and evaluate
        // them meticulously.
        add(a,b)
    }
}

