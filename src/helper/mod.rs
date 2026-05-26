use std::error::Error;
use whoami;

pub fn get_current_user_home_dir() -> Result<std::path::PathBuf, Box<dyn Error>> {
    let home_dir = match home::home_dir() {
        Some(path) => path,
        None => return Err("Could not find home directory".into()),
    };
    Ok(home_dir)
}

pub fn get_current_username() -> Result<String, Box<dyn Error>> {
    let user = whoami::username()?;
    Ok(user)
}
