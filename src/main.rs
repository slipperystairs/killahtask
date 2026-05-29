mod helper;
mod models;
mod file_handler;
mod commands;

use clap::{Parser, CommandFactory};
use clap_complete::generate;
use commands::Commands;
use std::error::Error;
use std::io;


#[derive(Parser)]
#[command(name = "rkt", about = "Rusty Killah Task (RKT) is a todo CLI tool.", long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

fn main() -> Result<(), Box<dyn Error>> {
    let cli = Cli::parse();

    if let Commands::Completion { shell } = &cli.command {
        let mut cmd = Cli::command();
        let name = cmd.get_name().to_string();
        generate(*shell, &mut cmd, name, &mut io::stdout());
        return Ok(());
    }
    
    // Set up file path: ~/killahtask_<user>.csv, maintaining file name for backwards compatibility with KAT
    let username = helper::get_current_username()?;
    let home_dir = helper::get_current_user_home_dir()?;
    
    let mut path = home_dir;
    path.push(format!("killahtask_{}.csv", username));

    let file = file_handler::load_file(&path)?;

    match &cli.command {
        Commands::Add { description } => commands::handle_add(&file, description.clone())?,
        Commands::Complete { task_id } => commands::handle_complete(&file, task_id.clone())?,
        Commands::Completion { .. } => {}, // Handled above --- IGNORE ---
        Commands::Delete { task_id } => commands::handle_delete(&file, task_id.clone())?,
        Commands::List { all } => commands::handle_list(&file, *all)?,
    }

    Ok(())
}