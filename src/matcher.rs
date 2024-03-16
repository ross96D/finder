use std::{error::Error, ffi::OsStr, fmt};
use grep_searcher::{SearcherBuilder, BinaryDetection, sinks::UTF8};
use grep_regex::RegexMatcher;
use ignore::WalkBuilder;

#[derive(Debug)]
pub struct SearchResult {
    pub file_path: String,
    pub line_number: u64,
    pub line: String,
    pub focus: bool
}
impl fmt::Display for SearchResult {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // Customize so only `x` and `y` are denoted.
        write!(f, "file_path: {}, line_number: {}, line: {}, focus: {}", self.file_path, self.line_number, self.line, self.focus)
    }
}

impl SearchResult {
    pub fn to_string(&self) -> String {
        String::from(format!("{}: {}", self.line_number, self.line))
    }
}

pub fn search(pattern: &str, path: &OsStr) -> Result<Vec<SearchResult>, Box<dyn Error>> {
    let matcher = RegexMatcher::new_line_matcher(pattern)?;
    let mut searcher = SearcherBuilder::new()
        .binary_detection(BinaryDetection::quit(b'\x00'))
        .build();

    let mut response: Vec<SearchResult> = vec![];
    let mut first = true;
    for result in WalkBuilder::new(path).hidden(false).ignore_case_insensitive(true).build() {
        let entry = match result {
            Ok(entry) => entry,
            Err(err) => {
                eprintln!("{}", err);
                continue;
            }
        };
        match entry.file_type() {
            Some(t) => {
                if !t.is_file() {
                    continue;
                }
            },
            None => {
                continue;
            },
        };

        let result = searcher.search_path(
            &matcher, 
            entry.path(), 
            UTF8(|lnum, line| {
                response.push(SearchResult{
                    file_path: String::from(entry.path().to_str().unwrap()), 
                    line_number: lnum, 
                    line: String::from(line),
                    focus: first,
                });
                first = false;
                Ok(true)
            })
        );
        if let Err(err) = result {
            eprintln!("{}: {}", entry.path().display(), err);
        }
    }
    Ok(response)   
}