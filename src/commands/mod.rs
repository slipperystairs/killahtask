use clap::Subcommand;
use clap_complete::Shell;
use chrono::{DateTime, Utc};
use crate::models::Task;
use crate::file_handler;
use prettytable::{Table, row, format};
use std::error::Error;
use std::fs::File;

#[derive(Subcommand)]
pub enum Commands {
    Add { description: String },
    Complete { task_id: String },
    Completion { shell: Shell },
    Delete { task_id: String },
    List {
        #[arg(short, long)]
        all: bool,
    },
}

pub fn handle_add(file: &File, description: String) -> Result<(), Box<dyn Error>> {
    let mut tasks = match file_handler::read_tasks(&file) {
        Ok(t) => t,
        Err(_) => Vec::new(),
    };
    
    if tasks.iter().any(|t| &t.description == &description) {
        return Err(format!("Task description isn't unique! \"{}\" already exists.", description).into());
    }

    let last_id_val = match tasks.last() {
        Some(task) => match task.task_id.parse::<u32>() {
            Ok(id) => id + 1,
            Err(_) => 0,
        },
        None => 0, 
    };

    tasks.push(Task {
        task_id: last_id_val.to_string(),
        description: description.clone(),
        created: Utc::now().to_rfc3339(),
        completed: "false".to_string(),
    });

    file_handler::write_tasks(&file, tasks)?;
    println!("Task \"{}\" added successfully!", description);
    Ok(())
}

pub fn handle_list(file: &File, all: bool) -> Result<(), Box<dyn Error>> {
    let tasks = match file_handler::read_tasks(&file) {
        Ok(t) => t,
        Err(_) => {
            println!("No task to show! Try adding a task by running killahtask add \"my task\"");
            return Ok(());
        }
    };

    if tasks.is_empty() {
        println!("No task to show! Try adding a task by running killahtask add \"my task\"");
        return Ok(());
    }

    let mut table = Table::new();
    table.set_format(*format::consts::FORMAT_NO_BORDER_LINE_SEPARATOR);
    
    if all {
        table.add_row(row!["ID", "Description", "Created", "Completed"]);
    } else {
        table.add_row(row!["ID", "Description", "Created"]);
    }

    for t in tasks {
        if !all && t.completed == "true" { continue; }
        
        let created_at = match DateTime::parse_from_rfc3339(&t.created) {
            Ok(dt) => dt,
            Err(_) => return Err(format!("Invalid timestamp for task ID {}", t.task_id).into()),
        };
        
        let now = Utc::now();
        let duration = now.signed_duration_since(created_at);
        let time_str = format!("{}m ago", duration.num_minutes()); 

        if all {
            table.add_row(row![t.task_id, t.description, time_str, t.completed]);
        } else {
            table.add_row(row![t.task_id, t.description, time_str]);
        }
    }
    table.printstd();
    Ok(())
}

pub fn handle_complete(file: &File, task_id: String) -> Result<(), Box<dyn Error>> {
    let mut tasks = file_handler::read_tasks(&file)?;
    let mut found = false;
    for t in &mut tasks {
        if &t.task_id == &task_id {
            t.completed = "true".to_string();
            found = true;
            break;
        }
    }
    if found {
        file_handler::write_tasks(&file, tasks)?;
        println!("ID {} was marked as complete!", task_id);
    } else {
        println!("Task ID could not be found.");
    }
    Ok(())
}

pub fn handle_delete(file: &File, task_id: String) -> Result<(), Box<dyn Error>> {
    let tasks = file_handler::read_tasks(&file)?;
    let original_len = tasks.len();
    let new_tasks: Vec<Task> = tasks.into_iter().filter(|t| &t.task_id != &task_id).collect();
    
    if new_tasks.len() == original_len {
        println!("Task ID could not be found.");
    } else {
        file_handler::write_tasks(&file, new_tasks)?;
        println!("Task removed successfully!");
    }
    Ok(())
}