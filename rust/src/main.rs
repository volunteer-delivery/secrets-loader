mod parameter;
mod command_arguments;
mod loader;
mod exports_output;

use command_arguments::CommandArguments;
use exports_output::ExportsOutput;
use loader::Loader;


fn main() {
    let args = CommandArguments::new();
    let loader = Loader::new(args);
    loader.load();
    ExportsOutput::print(loader.parameters());
}
