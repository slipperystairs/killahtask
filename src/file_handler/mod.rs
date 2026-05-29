use std::fs::{File, OpenOptions};
use std::io::{Seek, SeekFrom};
use fs2::FileExt;
use crate::models::Task;
use std::error::Error;

pub fn load_file(path: &std::path::PathBuf) -> Result<File, Box<dyn Error>> {
    let file = OpenOptions::new()
        .read(true)
        .write(true)
        .create(true)
        .open(path)?;

    file.lock_exclusive()?;
    Ok(file)
}

pub fn write_tasks(mut file: &File, tasks: Vec<Task>) -> Result<(), Box<dyn Error>> {
    file.set_len(0)?; // Truncate(0)
    file.seek(SeekFrom::Start(0))?; // Seek(0, 0)

    let mut wtr = csv::Writer::from_writer(file);
    for task in tasks {
        wtr.serialize(task)?;
    }
    wtr.flush()?;
    Ok(())
}

pub fn read_tasks(file: &File) -> Result<Vec<Task>, Box<dyn Error>> {
    let mut rdr = csv::Reader::from_reader(file);
    let mut tasks = Vec::new();
    for result in rdr.deserialize() {
        let task: Task = result?;
        tasks.push(task);
    }
    Ok(tasks)
}