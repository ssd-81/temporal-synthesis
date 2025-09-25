#[allow(unused_imports)]
use std::io::{self, Write};

fn main() {
    loop {
        let exit = "exit 0".to_string();
        print!("$ ");
        io::stdout().flush().unwrap();
        let mut input = String::new();
        io::stdin().read_line(&mut input).unwrap();
        if input.trim() == exit {
            std::process::exit(0);
        }
        //For echo_command command I am thinking splitting the input, echo_command is just 5 words including space what if I just simply cut that part off and print the other part
        //for that I would have done conversion too cumbersome so sticked with &str
        let echo_command = "echo ";
        let type_command="type ";
        if input.trim().starts_with(echo_command){
            let (first,last)=input.trim().split_at(5);
            println!("{}",last);
        }else if input.trim().starts_with(type_command){
             let (first,last)=input.trim().split_at(5);
             println!("{} is a shell builtin",last);
        }
        
        else{ println!("{}: command not found\r\n", input.trim());}

       
    }
}
