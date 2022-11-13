use clap::Parser;

/// CLI to generate .env file based on aws parameter store
#[derive(Parser)]
#[command(author, version, about, long_about = None)]
pub struct CommandArguments {
    /// Select parameters by path
    #[arg(short, long)]
    pub path: String,

    /// Filter parameters by label
    #[arg(short, long)]
    pub label: Option<String>,

    /// AWS regiosn
    #[arg(short, long)]
    pub region: Option<String>,
}

impl CommandArguments {
    pub fn new() -> CommandArguments {
        CommandArguments::parse()
    }
}
