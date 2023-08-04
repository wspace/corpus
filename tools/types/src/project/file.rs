use std::fs::{self, File};
use std::io;
use std::path::Path;
use std::process::{Command, Stdio};

use glob::Paths;
use thiserror::Error;

use crate::project::Project;
use crate::util::LossyString;

#[derive(Error, Debug)]
pub enum ReadProjectError {
    #[error(transparent)]
    Io(#[from] io::Error),
    #[error("error deserializing: {0}")]
    DeserializeJson(#[from] serde_json::Error),
    #[error("error deserializing: {0}")]
    DeserializeJson5(#[from] json5::Error),
}

#[derive(Error, Debug)]
pub enum WriteProjectError {
    #[error(transparent)]
    Io(#[from] io::Error),
    #[error("error serializing: {0}")]
    Serialize(#[from] serde_json::Error),
    #[error("error formatting with underscore: {}", LossyString::new(.0))]
    Underscore(Vec<u8>),
}

impl Project {
    pub fn from_json(input: &str) -> serde_json::Result<Self> {
        serde_json::from_str(input)
    }

    pub fn from_json5(input: &str) -> json5::Result<Self> {
        json5::from_str(input)
    }

    pub fn to_json(&self) -> serde_json::Result<Vec<u8>> {
        serde_json::to_vec(self)
    }

    pub fn to_json_pretty(&self) -> Result<String, WriteProjectError> {
        let mut post_printer = Command::new("underscore")
            .arg("print")
            .stdin(Stdio::piped())
            .stdout(Stdio::piped())
            .spawn()?;
        let mut stdin = post_printer.stdin.take().unwrap();

        serde_json::to_writer(&mut stdin, self)?;
        drop(stdin);

        let output = post_printer.wait_with_output()?;
        if !output.stderr.is_empty() {
            return Err(WriteProjectError::Underscore(output.stderr));
        }
        let pretty = String::from_utf8(output.stdout).unwrap();
        Ok(pretty)
    }

    pub fn from_json_file<P: AsRef<Path>>(path: P) -> Result<Self, ReadProjectError> {
        let f = File::open(path)?;
        Ok(serde_json::from_reader(f)?)
    }

    pub fn from_json5_file<P: AsRef<Path>>(path: P) -> Result<(Self, String), ReadProjectError> {
        let input = fs::read_to_string(path)?;
        let project = Project::from_json5(&input)?;
        Ok((project, input))
    }

    pub fn write_json_file<P: AsRef<Path>>(&self, path: P) -> Result<(), WriteProjectError> {
        let f = File::create(path)?;
        serde_json::to_writer(f, self)?;
        Ok(())
    }

    pub fn write_json_file_pretty<P: AsRef<Path>>(&self, path: P) -> Result<(), WriteProjectError> {
        let pretty = self.to_json_pretty()?;
        fs::write(path, &pretty)?;
        Ok(())
    }
}

pub fn all_project_json() -> Paths {
    glob::glob("*/*/project.json").unwrap()
}
