// use std::env;
// use std::fs;
#[allow(unused_imports)]
use std::io::{self, Write};
// use std::path::Path;
mod shell;
fn main() {
    loop {
        // let exit = "exit 0".to_string();
        print!("$ ");
        io::stdout().flush().unwrap();
        let mut input = String::new();
        io::stdin().read_line(&mut input).unwrap();
        let mut pure_input=input.trim();
        if input.is_empty(){
            continue;
        }
        if let Some(par)=shell::parse(pure_input){
            if let Err(err)=shell::execute(par){
                    eprintln!("{}", err); 
            }
        }
    }
        // if input.trim() == exit {
        //     std::process::exit(0);
        // }
        //For echo_command command I am thinking splitting the input, echo_command is just 5 words including space what if I just simply cut that part off and print the other part
        //for that I would have done conversion too cumbersome so sticked with &str
    //     let echo_command = "echo ";
    //     let type_command = "type ";
    //     let type_impl = vec!["echo", "type", "exit"];
    //     if input.trim().starts_with(echo_command) {
    //         let (first, last) = input.trim().split_at(5);
    //         println!("{}", last);
    //     } else if input.trim().starts_with(type_command) {
    //         let (first, last) = input.trim().split_at(5);
    //         if type_impl.iter().any(|e| *e == last) {
    //             println!("{}: is a shell builtin", last);
    //         } else {
    //             match env::var("PATH") {
    //                 Ok(path_string) => {
    //                     let paths: Vec<&str> = path_string.split(';').collect();
    //                     let mut found = false;
    //                     for path in paths.iter() {
    //                         let mut full_path_str = format!("{}\\{}.exe", path, last);
    //                         let full_path = Path::new(&full_path_str);
    //                         if full_path.is_file() {
    //                            found = true;
    //                            println!("{} is {:?}",last,full_path);
    //                             break;
    //                         }
    //                     }
    //                     if !found {
    //                         println!("{}: not found", last);
    //                     }
    //                 }
    //                 Err(e) => {
    //                     eprintln!("Error: Could not find the PATH variable. {}", e);
    //                 }
    //             }
    //         }
    //     } else {
    //         println!("{}: command not found\r\n", input.trim());
    //     }
    //     //Trying to find executable files
    // }
}
