package aws_json

type ParametersByPath struct {
    Parameters []Parameter
    NextToken string
}

func NewParametersByPath() *ParametersByPath {
    return &ParametersByPath{
        Parameters: []Parameter{},
        NextToken: "",
    }
}
