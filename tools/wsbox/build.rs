use std::process;

use corpus_types::project::Project;
use glob::glob;

fn main() {
    let mut errs = Vec::new();
    for json_path in glob("../../*/*/project.json").unwrap() {
        let json_path = json_path.unwrap();
        match Project::from_json5_file(&json_path) {
            Ok(project) => {
                println!("{project:?}");
            }
            Err(err) => {
                errs.push((json_path, err));
            }
        }
    }
    if !errs.is_empty() {
        for (json_path, err) in errs {
            eprintln!("{json_path:?}: {err}");
        }
        process::exit(1);
    }
}
