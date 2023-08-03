use std::collections::HashMap;
use std::fs::File;
use std::io;
use std::path::Path;

use serde::{Deserialize, Serialize};
use thiserror::Error;

use crate::util::{Int, OneOrVec, TotalF64, Uint};
use crate::ws::InstMap;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Project {
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    pub authors: Vec<String>,
    pub license: String,
    pub languages: Vec<String>,
    pub tags: Vec<String>,
    pub date: String,
    pub spec_version: SpecVersion,
    pub source: Vec<String>,
    #[serde(default)]
    pub source_unavailable: bool,
    #[serde(default)]
    pub packages: Vec<Package>,
    #[serde(default)]
    pub relations: Vec<Relation>,
    #[serde(default)]
    pub challenges: Vec<Challenge>,
    pub bounds: Option<Bounds>,
    pub behavior: Option<Behavior>,
    pub whitespace: Option<Whitespace>,
    pub assembly: Option<Assembly>,
    #[serde(default)]
    pub mappings: Vec<Mapping>,
    #[serde(default)]
    pub programs: Vec<Program>,
    pub build_errors: Option<String>,
    pub run_errors: Option<String>,
    pub scripts: Option<Scripts>,
    #[serde(default)]
    pub commands: Vec<Command>,
    pub notes: Option<String>,
    pub todo: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
pub enum SpecVersion {
    #[serde(rename = "0.1")]
    Ws0_1,
    #[serde(rename = "0.2")]
    Ws0_2,
    #[serde(rename = "0.3")]
    Ws0_3,
    #[serde(rename = "unknown")]
    Unknown,
    #[serde(rename = "-")]
    NA,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Package {
    pub name: String,
    pub manager: PackageManager,
    pub url: String,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
pub enum PackageManager {
    #[serde(rename = "crates.io")]
    CratesIo,
    #[serde(rename = "Docker Hub")]
    DockerHub,
    #[serde(rename = "Hackage")]
    Hackage,
    #[serde(rename = "npm")]
    Npm,
    #[serde(rename = "NuGet")]
    NuGet,
    #[serde(rename = "PyPI")]
    PyPi,
    #[serde(rename = "RubyGems")]
    RubyGems,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Relation {
    pub id: String,
    #[serde(rename = "type")]
    pub typ: String,
    #[serde(default)]
    pub commit: String,
    #[serde(default)]
    pub release: String,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
#[serde(rename_all = "snake_case")]
pub enum RelationType {
    Fork,
    Submodule,
    Embedded,
    Library,
    Binary,
    Assembly,
    Design,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
pub enum Challenge {
    #[serde(rename = "Advent of Code 2015")]
    AdventOfCode2015,
    #[serde(rename = "Advent of Code 2016")]
    AdventOfCode2016,
    #[serde(rename = "Advent of Code 2017")]
    AdventOfCode2017,
    #[serde(rename = "Advent of Code 2018")]
    AdventOfCode2018,
    #[serde(rename = "Advent of Code 2019")]
    AdventOfCode2019,
    #[serde(rename = "Advent of Code 2020")]
    AdventOfCode2020,
    #[serde(rename = "Advent of Code 2021")]
    AdventOfCode2021,
    #[serde(rename = "Advent of Code 2022")]
    AdventOfCode2022,
    #[serde(rename = "Cerner 2^5")]
    Cerner2Pow5,
    #[serde(rename = "Code Golf")]
    CodeGolf,
    #[serde(rename = "Codewars")]
    Codewars,
    #[serde(rename = "Hackergame")]
    Hackergame,
    #[serde(rename = "Hacktoberfest 2016")]
    Hacktoberfest2016,
    #[serde(rename = "Hacktoberfest 2017")]
    Hacktoberfest2017,
    #[serde(rename = "Hacktoberfest 2018")]
    Hacktoberfest2018,
    #[serde(rename = "Kattis")]
    Kattis,
    #[serde(rename = "noxCTF")]
    NoxCtf,
    #[serde(rename = "Project Euler")]
    ProjectEuler,
    #[serde(rename = "Rosetta Code")]
    RosettaCode,
    #[serde(rename = "Sphere Online Judge")]
    SphereOnlineJudge,
    #[serde(rename = "yukicoder")]
    Yukicoder,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Bounds {
    pub precision: Option<Precision>,
    pub arg_precision: Option<Precision>,
    pub label_precision: Option<Precision>,
    pub stack_cap: Option<Bound<Uint>>,
    pub call_stack_cap: Option<Bound<Uint>>,
    pub heap_min: Option<Bound<Int>>,
    pub heap_max: Option<Bound<Int>>,
    pub heap_cap: Option<Bound<Uint>>,
    pub label_cap: Option<Bound<Uint>>,
    pub instruction_cap: Option<Bound<Uint>>,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
#[serde(rename_all = "lowercase")]
pub enum IntegerType {
    Int64,
    Int32,
    Uint64,
    Uint32,
    Float64,
    Float32,
    #[serde(rename = "C int")]
    CInt,
    #[serde(rename = "Go int")]
    GoInt,
    #[serde(rename = "Go uint")]
    GoUint,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
#[serde(rename_all = "snake_case")]
pub enum Precision {
    Arbitrary,
    Fixed,
    #[serde(untagged)]
    Type(IntegerType),
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
#[serde(rename_all = "snake_case")]
pub enum Bound<T> {
    Unbounded,
    #[serde(untagged)]
    Type(IntegerType),
    #[serde(untagged)]
    Exact(T),
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Behavior {
    pub buffered_output: Option<bool>,
    pub eof: Option<EofBehavior>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub enum EofBehavior {
    #[serde(rename = "error")]
    Error,
    #[serde(untagged)]
    Value(Int),
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Whitespace {
    #[serde(default)]
    pub unimplemented: Vec<String>,
    #[serde(default)]
    pub nonstandard: Vec<NonstandardInst>,
    pub extension: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct NonstandardInst {
    pub name: Option<String>,
    #[serde(default)]
    pub aliases: Vec<String>,
    pub seq: Option<String>,
    #[serde(default)]
    pub args: Vec<String>,
    pub desc: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Assembly {
    pub mnemonics: Option<InstMap<OneOrVec<String>>>,
    #[serde(default)]
    pub macros: Vec<Macro>,
    #[serde(default)]
    pub patterns: HashMap<String, String>,
    pub case_sensitive_mnemonics: Option<bool>,
    #[serde(default)]
    pub instruction_delimiter: Option<InstDelim>,
    pub instruction_wrap: Option<bool>,
    #[serde(default)]
    pub line_comments: Vec<String>,
    #[serde(default)]
    pub block_comments: Vec<BlockComment>,
    pub indentation: Option<String>,
    pub label_indentation: Option<String>,
    pub block_indentation: Option<bool>,
    pub binary_numbers: Option<bool>,
    #[serde(default)]
    pub usage: Vec<String>,
    pub extension: Option<String>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Macro {
    pub name: String,
    #[serde(default)]
    pub args: Vec<String>,
    pub replace: Option<Vec<String>>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
pub enum InstDelim {
    #[serde(rename = "none")]
    None,
    #[serde(rename = "space")]
    Space,
    #[serde(rename = ";")]
    Semi,
    #[serde(rename = "/")]
    Slash,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct BlockComment {
    pub start: String,
    pub end: String,
    pub nestable: bool,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Mapping {
    pub space: String,
    pub tab: String,
    pub lf: String,
    pub spaces_between: Option<bool>,
    pub space_before_arg: Option<bool>,
    pub line_comment: Option<String>,
    pub before_stl: Option<bool>,
    pub ignore_case: Option<bool>,
    pub extension: Option<String>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Program {
    pub path: String,
    pub generated: Option<String>,
    #[serde(default)]
    pub inputs: Vec<String>,
    #[serde(default)]
    pub outputs: Vec<String>,
    #[serde(default)]
    pub aux: Vec<String>,
    #[serde(default)]
    pub polyglot: Vec<String>,
    pub mapping_index: Option<usize>,
    pub equivalent: Option<String>,
    pub spec_version: Option<SpecVersion>,
    pub generate: Option<String>,
    #[serde(default)]
    pub authors: Vec<String>,
    pub desc: Option<String>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Scripts {
    pub run: Option<Script>,
    pub compile: Option<Script>,
    pub assemble: Option<Script>,
    pub disassemble: Option<Script>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Script {
    pub bin: ScriptBin,
    #[serde(default)]
    pub args: Vec<String>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
#[serde(untagged)]
pub enum ScriptBin {
    Path { path: String },
    Jar { jar: String, main: Option<String> },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Command {
    #[serde(rename = "type")]
    pub typ: Option<String>,
    pub bin: Option<String>,
    #[serde(default)]
    pub dependencies: Vec<String>,
    pub install_dependencies: Option<String>,
    pub build: Option<String>,
    pub build_errors: Option<String>,
    pub usage: Option<String>,
    #[serde(default)]
    pub example_usages: Vec<String>,
    pub run_errors: Option<String>,
    pub input: Option<String>,
    pub output: Option<String>,
    #[serde(default)]
    pub options: Vec<CommandOption>,
    pub option_parse: Option<OptionParse>,
    #[serde(default)]
    pub subcommands: Vec<Subcommand>,
    #[serde(default)]
    pub unimplemented: Vec<String>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct CommandOption {
    #[serde(default)]
    pub short: OneOrVec<String>, // -s
    pub long: Option<String>, // --long
    pub bare: Option<String>, // bare
    pub required: Option<bool>,
    pub repeat_allowed: Option<bool>,
    pub arg: Option<String>,
    pub arg_required: Option<bool>,
    #[serde(rename = "type")]
    pub typ: Option<String>,
    pub default: Option<OptionArg>,
    pub min: Option<Int>,
    pub desc: Option<String>,
    #[serde(default)]
    pub values: Vec<OptionValue>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
#[serde(untagged)]
pub enum OptionArg {
    Int(Int),
    F64(TotalF64),
    Bool(bool),
    String(String),
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct OptionValue {
    #[serde(default)]
    pub values: Vec<String>,
    pub desc: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Subcommand {
    pub name: String,
    pub desc: Option<String>,
    pub usage: Option<String>,
    pub input: Option<String>,
    pub output: Option<String>,
    #[serde(default)]
    pub options: Vec<CommandOption>,
    pub notes: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Copy, Debug, PartialEq, Eq)]
pub enum OptionParse {
    #[serde(rename = "manual")]
    Manual,
    #[serde(rename = "bash getopt")]
    BashGetOpt,
    #[serde(rename = "bash getopts")]
    BashGetOpts,
    #[serde(rename = "C getopt")]
    CGetOpt,
    #[serde(rename = "Clojure clojure.tools.cli")]
    ClojureClojureToolsCli,
    #[serde(rename = "Crystal option_parser")]
    CrystalOptionParser,
    #[serde(rename = "Go flag")]
    GoFlag,
    #[serde(rename = "Go github.com/alexflint/go-arg")]
    GoGitHubComAlexFlintGoArg,
    #[serde(rename = "Haskell cmdargs")]
    HaskellCmdArgs,
    #[serde(rename = "Haskell optparse-applicative")]
    HaskellOptParseApplicative,
    #[serde(rename = "Java org.apache.commons.cli")]
    JavaOrgApacheCommonsCli,
    #[serde(rename = "JavaScript Commander.js")]
    JavaScriptCommanderJs,
    #[serde(rename = "Objective-C getopt")]
    ObjectiveCGetOpt,
    #[serde(rename = "OCaml Arg")]
    OCamlArg,
    #[serde(rename = "Perl Getopt::Std")]
    PerlGetOptStd,
    #[serde(rename = "Python argparse")]
    PythonArgParse,
    #[serde(rename = "Python click")]
    PythonClick,
    #[serde(rename = "Python optparse")]
    PythonOptParse,
    #[serde(rename = "Ruby optparse")]
    RubyOptParse,
    #[serde(rename = "Ruby Thor")]
    RubyThor,
    #[serde(rename = "Rust clap")]
    RustClap,
    #[serde(rename = "Rust structopt")]
    RustStructOpt,
}

#[derive(Error, Debug)]
pub enum ReadProjectError {
    #[error(transparent)]
    Io(#[from] io::Error),
    #[error(transparent)]
    Json(#[from] serde_json::Error),
}

impl Project {
    pub fn from_file(path: &Path) -> Result<Self, ReadProjectError> {
        let f = File::open(path)?;
        Ok(serde_json::from_reader(f)?)
    }
}
