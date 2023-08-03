use std::io::Write;
use std::path::Path;
use std::process::{self, Command, Stdio};
use std::{env, fs};

use anyhow::{bail, Result};
use corpus_types::project::{all_project_json, Project};
use corpus_types::util::MultiError;

fn main() {
    if let Err(err) = format_all() {
        eprintln!("{err}");
        process::exit(1);
    }
}

fn format_all() -> Result<()> {
    let args = env::args_os().skip(1);
    let mut errs = MultiError::new();
    if args.len() != 0 {
        for path in args {
            if let Err(err) = format_file(path.as_ref()) {
                errs.push(err);
            }
        }
    } else {
        for path in all_project_json() {
            if let Err(err) = format_file(path?.as_path()) {
                errs.push(err);
            }
        }
    }
    errs.err()?;
    Ok(())
}

fn format_file(path: &Path) -> Result<bool> {
    let unformatted = fs::read(path)?;

    let project: Project = if let Ok(project) = serde_json::from_slice(&*unformatted) {
        project
    } else {
        let mut pre_printer = Command::new("underscore")
            .arg("print")
            .stdin(Stdio::piped())
            .stdout(Stdio::piped())
            .spawn()?;
        let mut stdin = pre_printer.stdin.take().unwrap();
        let stdout = pre_printer.stdout.take().unwrap();

        stdin.write_all(&unformatted)?;
        drop(stdin);
        let project = serde_json::from_reader(stdout)?;

        let output = pre_printer.wait_with_output()?;
        if !output.stderr.is_empty() {
            let err = String::from_utf8_lossy(&output.stderr);
            bail!("JSON pre-formatting exited with error: {err}");
        }
        project
    };

    // TODO: Processing

    let mut post_printer = Command::new("underscore")
        .arg("print")
        .stdin(Stdio::piped())
        .stdout(Stdio::piped())
        .spawn()?;
    let mut stdin = post_printer.stdin.take().unwrap();

    serde_json::to_writer(&mut stdin, &project)?;
    drop(stdin);

    let output = post_printer.wait_with_output()?;
    if !output.stderr.is_empty() {
        let err = String::from_utf8_lossy(&output.stderr);
        bail!("JSON post-formatting exited with error: {err}");
    }
    let formatted = output.stdout;

    if unformatted == formatted {
        return Ok(false);
    }
    println!("Formatted {}", path.to_string_lossy());
    fs::write(path, formatted)?;
    Ok(true)
}
