use std::cell::RefCell;
use std::process::Command;
use serde_json::Value;

use crate::command_arguments::CommandArguments;
use crate::parameter::Parameter;

pub struct Loader {
    command_arguments: CommandArguments,
    parameters_ptr: RefCell<Vec<Parameter>>,
}

impl Loader {
    pub fn new(command_arguments: CommandArguments) -> Loader {
        Loader {
            command_arguments,
            parameters_ptr: RefCell::new(Vec::new())
        }
    }

    pub fn load(&self) {
        self.next_parameters(None);
    }

    pub fn parameters(&self) -> Vec<Parameter> {
        self.parameters_ptr.take()
    }

    fn next_parameters(&self, token: Option<String>) {
        let json = self.fetch_parameters(token);
        let parameters = json["Parameters"].as_array().unwrap();

        for value in parameters {
            let parameter = Parameter::from_aws_param(self.command_arguments.path.to_string(), value);
            self.parameters_ptr.borrow_mut().push(parameter);
        }

        if json["NextToken"].is_string() {
            self.next_parameters(Some(json["NextToken"].to_string()));
        }
    }

    fn fetch_parameters(&self, token: Option<String>) -> Value {
        let mut binding = Command::new("aws");
        let mut command = binding.arg("ssm").arg("get-parameters-by-path")
            .arg("--path").arg(self.command_arguments.path.to_string())
            .arg("--with-decryption");

        if self.command_arguments.label.is_some() {
            let filter = format!("Key=Label,Values={}", self.command_arguments.label.as_ref().unwrap());
            command = command.arg("--parameter-filters").arg(filter);
        }

        if self.command_arguments.region.is_some() {
            command = command.arg("--region").arg(self.command_arguments.region.as_ref().unwrap());
        }

        if token.is_some() {
            let raw_token = Parameter::strip_quotes(token.unwrap());
            command = command.arg("--starting-token").arg(raw_token);
        }

        let output = command.output().expect("failed to execute process");
        serde_json::from_slice(output.stdout.as_slice()).map_err(|err| println!("PARSE ERROR: {}", err)).unwrap()
    }
}