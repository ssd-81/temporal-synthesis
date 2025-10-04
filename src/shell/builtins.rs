use std::process;
use super::external::find_path;
pub enum BuiltInCommand{
  Exit(i32),
  Echo(String),
  Type(String),
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
    }
}
