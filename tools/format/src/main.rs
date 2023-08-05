use std::env;
use std::fs;
use std::path::Path;
use std::process;
use std::process::Command;

use anyhow::{bail, Result};
use corpus_types::project::Project;
use corpus_types::util::MultiError;

fn main() {
    let args = env::args_os().skip(1);
    if args.len() == 0 {
        eprintln!("Usage: corpus-format <file>...");
        process::exit(2);
    }
    let mut errs = MultiError::new();
    for path in args {
        if let Err(err) = format_file(path.as_ref()) {
            errs.push(err);
        }
    }
    if let Err(err) = errs.err() {
        eprintln!("{err}");
        process::exit(1);
    }
}

fn format_file(path: &Path) -> Result<bool> {
    let (project, unformatted) = Project::from_json5_file(path)?;

    add_submodule(&project)?;

    let formatted = project.to_json_pretty()?;
    if unformatted == formatted {
        return Ok(false);
    }
    println!("Formatted {}", path.to_string_lossy());
    fs::write(path, &formatted)?;
    Ok(true)
}

fn add_submodule(project: &Project) -> Result<()> {
    if let Some(submodule) = project.submodules.first() {
        let path = format!("{}/{}", project.id, submodule.path);
        if Path::new(&path).exists() {
            return Ok(());
        }

        let mut command = Command::new("git");
        command.args(&["submodule", "add"]);
        if let Some(branch) = &submodule.branch {
            println!("git submodule add -b {branch} {} {path}", submodule.url);
            command.args(&["-b", &branch]);
        } else {
            println!("git submodule add {} {path}", submodule.url);
        }
        command.arg(submodule.url.as_str()).arg(&path);

        let mut child = command.spawn().unwrap();
        let status = child.wait().unwrap();
        if !status.success() {
            bail!("git submodule exited with code {status}");
        }
    }
    Ok(())
}
