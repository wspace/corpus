use std::path::Path;
use std::process;
use std::{env, fs};

use anyhow::Result;
use corpus_types::project::{all_project_json, Project};
use corpus_types::util::MultiError;

fn main() {
    let args = env::args_os().skip(1);
    let mut errs = MultiError::new();
    if args.len() != 0 {
        for path in args {
            if let Err(err) = format_file(path.as_ref()) {
                errs.push(err);
            }
        }
    } else {
        for path_res in all_project_json() {
            let res = match path_res {
                Ok(path) => format_file(path.as_ref()),
                Err(err) => Err(err.into()),
            };
            if let Err(err) = res {
                errs.push(err);
            }
        }
    }
    if let Err(err) = errs.err() {
        eprintln!("{err}");
        process::exit(1);
    }
}

fn format_file(path: &Path) -> Result<bool> {
    let (project, unformatted) = Project::from_file_lax(path)?;
    // TODO: Processing
    let formatted = project.to_json_pretty()?;
    if unformatted == formatted {
        return Ok(false);
    }
    println!("Formatted {}", path.to_string_lossy());
    fs::write(path, formatted)?;
    Ok(true)
}
