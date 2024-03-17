use grep_regex::RegexMatcher;
use grep_searcher::{sinks::UTF8, BinaryDetection, SearcherBuilder};
use ignore::WalkBuilder;
use std::{
    error::Error,
    ffi::OsStr,
    fmt,
    fs::File,
    io::{self, BufRead},
    path::Path,
};

#[derive(Debug)]
pub struct SearchResult {
    pub file_path: String,
    pub line_number: u64,
    pub line: String,
    pub focus: bool,
}
impl fmt::Display for SearchResult {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // Customize so only `x` and `y` are denoted.
        write!(
            f,
            "file_path: {}, line_number: {}, line: {}, focus: {}",
            self.file_path, self.line_number, self.line, self.focus
        )
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
    for result in WalkBuilder::new(path)
        .hidden(false)
        .ignore_case_insensitive(true)
        .build()
    {
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
            }
            None => {
                continue;
            }
        };

        let result = searcher.search_path(
            &matcher,
            entry.path(),
            UTF8(|lnum, line| {
                response.push(SearchResult {
                    file_path: String::from(entry.path().to_str().unwrap()),
                    line_number: lnum,
                    line: String::from(line),
                    focus: first,
                });
                first = false;
                Ok(true)
            }),
        );
        if let Err(err) = result {
            eprintln!("{}: {}", entry.path().display(), err);
        }
    }
    Ok(response)
}

pub fn preview(path: &Path, line_number: u64, limit: u64) -> Result<String, Box<dyn Error>> {
    let mut result: String = String::new();
    if let Ok(lines) = read_lines(path) {
        let mut count = 0;
        for line in lines {
            match line {
                Ok(line) => {
                    if count <= line_number - limit {
                        // do nothing
                    } else if count < line_number + limit {
                        result.push_str(&line);
                        result.push('\n');
                    } else {
                        break;
                    }
                    count += 1;
                }
                Err(err) => {
                    // TODO improve error logging
                    eprintln!("error reading line {}", err);
                    continue;
                }
            }
        }
    }
    Ok(result)
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
