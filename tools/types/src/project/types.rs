use std::collections::BTreeMap;

use serde::{Deserialize, Serialize};
use serde_with::{apply, skip_serializing_none};
use url::Url;

use crate::util::{Int, OneOrVec, TotalF64, Uint};
use crate::ws::{Inst, InstMap};

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Project {
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    #[serde_with(skip_apply)]
    pub authors: Vec<String>,
    pub license: String,
    #[serde_with(skip_apply)]
    pub languages: Vec<String>,
    #[serde_with(skip_apply)]
    pub tags: Vec<String>,
    pub date: String,
    pub spec_version: SpecVersion,
    #[serde_with(skip_apply)]
    pub source: Vec<Url>,
    #[serde(default, skip_serializing_if = "is_false")]
    pub source_unavailable: bool,
    pub packages: Vec<Package>,
    pub relations: Vec<Relation>,
    pub challenges: Vec<Challenge>,
    pub bounds: Option<Bounds>,
    pub behavior: Option<Behavior>,
    pub whitespace: Option<Whitespace>,
    pub assembly: Option<Assembly>,
    pub mappings: Vec<Mapping>,
    pub programs: Vec<Program>,
    pub build_errors: Option<String>,
    pub run_errors: Option<String>,
    pub scripts: Option<Scripts>,
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
    pub url: Url,
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

#[skip_serializing_none]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Relation {
    pub id: String,
    #[serde(rename = "type")]
    pub typ: String,
    pub commit: Option<String>,
    pub release: Option<String>,
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

#[skip_serializing_none]
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

#[skip_serializing_none]
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

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Whitespace {
    pub unimplemented: Vec<Inst>,
    pub nonstandard: Vec<NonstandardInst>,
    pub extension: Option<String>,
}

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct NonstandardInst {
    pub name: Option<String>,
    pub aliases: Vec<String>,
    pub seq: Option<String>,
    pub args: Vec<String>,
    pub desc: Option<String>,
}

#[skip_serializing_none]
#[apply(
    Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")],
    BTreeMap => #[serde(default, skip_serializing_if = "BTreeMap::is_empty")],
    InstMap => #[serde(default, skip_serializing_if = "InstMap::is_empty")],
)]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Assembly {
    pub mnemonics: InstMap<OneOrVec<String>>,
    pub macros: Vec<Macro>,
    pub patterns: BTreeMap<String, String>,
    pub case_sensitive_mnemonics: Option<bool>,
    pub instruction_delimiter: Option<InstDelim>,
    pub instruction_wrap: Option<bool>,
    pub line_comments: Vec<String>,
    pub block_comments: Vec<BlockComment>,
    pub indentation: Option<String>,
    pub label_indentation: Option<String>,
    pub block_indentation: Option<bool>,
    pub binary_numbers: Option<bool>,
    pub usage: Vec<String>,
    pub extension: Option<String>,
    pub notes: Option<String>,
}

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Macro {
    pub name: String,
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

#[skip_serializing_none]
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

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Program {
    pub path: String,
    pub generated: Option<String>,
    pub inputs: Vec<String>,
    pub outputs: Vec<String>,
    pub aux: Vec<String>,
    pub polyglot: Vec<String>,
    pub mapping_index: Option<usize>,
    pub equivalent: Option<String>,
    pub spec_version: Option<SpecVersion>,
    pub generate: Option<String>,
    pub authors: Vec<String>,
    pub desc: Option<String>,
    pub notes: Option<String>,
}

#[skip_serializing_none]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Scripts {
    pub run: Option<Script>,
    pub compile: Option<Script>,
    pub assemble: Option<Script>,
    pub disassemble: Option<Script>,
}

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Script {
    pub bin: ScriptBin,
    pub args: Vec<String>,
    pub notes: Option<String>,
}

#[skip_serializing_none]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
#[serde(untagged)]
pub enum ScriptBin {
    Path { path: String },
    Jar { jar: String, main: Option<String> },
}

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Command {
    #[serde(rename = "type")]
    pub typ: Option<String>,
    pub bin: Option<String>,
    pub dependencies: Vec<String>,
    pub install_dependencies: Option<String>,
    pub build: Option<String>,
    pub build_errors: Option<String>,
    pub usage: Option<String>,
    pub example_usages: Vec<String>,
    pub run_errors: Option<String>,
    pub input: Option<String>,
    pub output: Option<String>,
    pub options: Vec<CommandOption>,
    pub option_parse: Option<OptionParse>,
    pub subcommands: Vec<Subcommand>,
    pub unimplemented: Vec<Inst>,
    pub notes: Option<String>,
}

#[skip_serializing_none]
#[apply(
    Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")],
    OneOrVec => #[serde(default, skip_serializing_if = "OneOrVec::is_empty")],
)]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct CommandOption {
    pub short: OneOrVec<String>, // -s
    pub long: Option<String>,    // --long
    pub bare: Option<String>,    // bare
    pub required: Option<bool>,
    pub repeat_allowed: Option<bool>,
    pub arg: Option<String>,
    pub arg_required: Option<bool>,
    #[serde(rename = "type")]
    pub typ: Option<String>,
    pub default: Option<OptionArg>,
    pub min: Option<Int>,
    pub desc: Option<String>,
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

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct OptionValue {
    pub values: Vec<String>,
    pub desc: Option<String>,
}

#[skip_serializing_none]
#[apply(Vec => #[serde(default, skip_serializing_if = "Vec::is_empty")])]
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Eq)]
pub struct Subcommand {
    pub name: String,
    pub desc: Option<String>,
    pub usage: Option<String>,
    pub input: Option<String>,
    pub output: Option<String>,
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

fn is_false(b: &bool) -> bool {
    !b
}
