use std::process;
use super::external::find_path;
use std::env;
use std::path::Path;
pub enum BuiltInCommand{
  Exit(i32),
  Echo(String),
  Type(String),
  Pwd,
  Cd(String),
}
pub fn handle_builtins(cmd:BuiltInCommand)->Result<(),String>{
    match cmd{
      BuiltInCommand::Exit(code)=>{
        process::exit(code)
      },
      BuiltInCommand::Echo(msg)=>{
        println!("{}",msg);
        Ok(())
      }
      BuiltInCommand::Type(cmd)=>{
          let builtins=vec!["type","echo","exit"];
          if builtins.contains(&cmd.as_str()){
            println!("{} is a shell builtin",cmd);
            Ok(())
          }else if let Some(path)=find_path(&cmd){
            println!("{}",path);
            Ok(())
          }else{
            Err(format!("{}:not found",cmd))
          }
      }
      BuiltInCommand::Pwd=>{
        let current_path=env::current_dir().unwrap();
        println!("{}",current_path.display());
        Ok(())
      }
      BuiltInCommand::Cd(destination)=>{
          if destination==String::from("~"){
            let home_path=Path::new(r"C:\Users\lenovo");
            env::set_current_dir(home_path);
            return Ok(());
          }
          let dest_path=Path::new(&destination);
          if dest_path.exists(){
            match env::set_current_dir(dest_path){  
              Ok(())=>{},
              Err(e)=>{
                eprintln!("Failed to change the directory:{}",e);
              }
            }
            Ok(())
          }else{
            Err(format!("cd: {:?}: No such file or directory",dest_path))
          }
      }
    }
}
