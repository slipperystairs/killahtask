use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Task {
    pub task_id: String,
    pub description: String,
    pub created: String,
    pub completed: String,
}