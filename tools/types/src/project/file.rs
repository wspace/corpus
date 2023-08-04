use std::fs::{self, File};
use std::io::{self, Write};
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
    Serde(#[from] serde_json::Error),
    #[error("error parsing with underscore: {}", LossyString::new(.0))]
    Underscore(Vec<u8>),
}

#[derive(Error, Debug)]
pub enum WriteProjectError {
    #[error(transparent)]
    Io(#[from] io::Error),
    #[error("error serializing: {0}")]
    Serde(#[from] serde_json::Error),
    #[error("error formatting with underscore: {}", LossyString::new(.0))]
    Underscore(Vec<u8>),
}

impl Project {
    pub fn from_json(input: &[u8]) -> Result<Self, ReadProjectError> {
        Ok(serde_json::from_slice(input)?)
    }

    pub fn from_json_lax(input: &[u8]) -> Result<Self, ReadProjectError> {
        if let Ok(project) = serde_json::from_slice(input) {
            return Ok(project);
        }

        let mut printer = Command::new("underscore")
            .arg("print")
            .stdin(Stdio::piped())
            .stdout(Stdio::piped())
            .spawn()?;
        let mut stdin = printer.stdin.take().unwrap();
        let stdout = printer.stdout.take().unwrap();

        stdin.write_all(input)?;
        drop(stdin);
        let project = serde_json::from_reader(stdout)?;

        let output = printer.wait_with_output()?;
        if !output.stderr.is_empty() {
            return Err(ReadProjectError::Underscore(output.stderr));
        }
        Ok(project)
    }

    pub fn to_json(&self) -> Result<Vec<u8>, WriteProjectError> {
        Ok(serde_json::to_vec(self)?)
    }

    pub fn to_json_pretty(&self) -> Result<Vec<u8>, WriteProjectError> {
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
        Ok(output.stdout)
    }

    pub fn from_file<P: AsRef<Path>>(path: P) -> Result<Self, ReadProjectError> {
        let f = File::open(path)?;
        Ok(serde_json::from_reader(f)?)
    }

    pub fn from_file_lax<P: AsRef<Path>>(path: P) -> Result<(Self, Vec<u8>), ReadProjectError> {
        let input = fs::read(path)?;
        let project = Project::from_json_lax(&input)?;
        Ok((project, input))
    }

    pub fn write_file<P: AsRef<Path>>(&self, path: P) -> Result<(), WriteProjectError> {
        let f = File::create(path)?;
        serde_json::to_writer(f, self)?;
        Ok(())
    }

    pub fn write_file_pretty<P: AsRef<Path>>(&self, path: P) -> Result<(), WriteProjectError> {
        let pretty = self.to_json_pretty()?;
        fs::write(path, &pretty)?;
        Ok(())
    }
}

pub fn all_project_json() -> Paths {
    glob::glob("*/*/project.json").unwrap()
}
