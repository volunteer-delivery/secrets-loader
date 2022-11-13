use crate::parameter::Parameter;

pub struct ExportsOutput {}

impl ExportsOutput {
    pub fn print(parameters: Vec<Parameter>) {
        for parameter in parameters {
            println!("{}={}", parameter.name, parameter.value)
        }
    }
}
