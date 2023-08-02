use corpus_types::project::Project;
use glob::glob;

fn main() {
    for json_path in glob("../../*/*/project.json").unwrap() {
        let json_path = json_path.unwrap();
        let project = Project::from_file(&json_path).unwrap();
        println!("{project:?}");
    }
}
