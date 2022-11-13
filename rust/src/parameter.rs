use serde_json::Value;

pub struct Parameter {
    pub name: String,
    pub value: String
}

impl Parameter {
    pub fn from_aws_param(path: String, param: &Value) -> Parameter {
        let name = Self::prepare_aws_path_name(path, param["Name"].to_string());
        let value = Self::strip_quotes(param["Value"].to_string());
        Parameter { name,value }
    }

    fn prepare_aws_path_name(path: String, name: String) -> String {
        let unquoted = Self::strip_quotes(name.to_string());
        unquoted.strip_prefix(&format!("{}/", path)).unwrap().to_string()
    }

    pub fn strip_quotes(str: String) -> String {
        str.strip_prefix('"').map(|str| str.strip_suffix('"')).flatten().unwrap().to_string()
    }
}
