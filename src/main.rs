use std::fs;
use std::io;
use std::io::prelude::*;
use std::io::{BufReader, Write};
#[allow(unused_imports)]
use std::net::{TcpListener, TcpStream};
use std::thread;
//This is what i got in http_request
// Request: [
//     "GET /test/path HTTP/1.1",
//     "Host: 127.0.0.1:4221",
//     "User-Agent: curl/8.13.0",
//     "Accept: */*",
// ]
fn handle_client(mut stream: TcpStream) -> io::Result<()> {
    let mut buf_reader = BufReader::new(&mut stream);
    let http_request: Vec<_> = (&mut buf_reader)
        .lines()
        .map(|result| result.unwrap())
        .take_while(|line| !line.is_empty())
        .collect();
    // println!("{:?}",http_request);
    // let mut response = "HTTP/1.1 200 OK\r\n\r\n";
    // stream.write_all(response.as_bytes());
    let specific_path: Vec<_> = http_request[0].split_whitespace().collect();
    //this produces this output ---- ["POST", "/sarvil", "HTTP/1.1"]
    // println!("{:?}", specific_path);
    if specific_path[1].contains("/user-agent") {
        let got_header = &http_request[3];
        let returned_values: Vec<&str> = got_header.split(":").collect();
        let returned_value = returned_values[1].trim();
        let passing_return = format!(
            "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: {}\r\n\r\n{}",
            returned_value.len(),
            returned_value
        );
        stream.write_all(passing_return.as_bytes());
        return Ok(());
    }
    if specific_path[1].contains("/files") {
        match specific_path[0] {
            "GET" => {
                let impure_destination = specific_path[1].split_once("/files/");
                let pure_destination = match impure_destination {
                    Some((_, inneed)) => inneed,
                    None => "",
                };
                let message_send: String = match fs::read_to_string(pure_destination) {
                    Ok(message) => {
                        format!(
                            "HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: {}\r\n\r\n{}",
                            message.len(),
                            message
                        )
                    }
                    Err(_) => {
                        format!("HTTP/1.1 404 Not Found\r\n\r\n")
                    }
                };

                stream.write_all(message_send.as_bytes());
            }
            "POST" => {
                let impure_destination = specific_path[1].split_once("/files/");
                let pure_destination = match impure_destination {
                    Some((_, inneed)) => inneed,
                    None => "",
                };
                println!("{}", pure_destination);
                let mut f = fs::File::create(pure_destination).unwrap();
                // ["POST /files/testingcreation HTTP/1.1", "Host: localhost:4221", "User-Agent: curl/8.13.0", "Accept: */*", "Content-Length: 27", "Content-Type: application/x-www-form-urlencoded"]
                //taking content length from here
                let content_len: u32 = http_request[4]
                    .split(":")
                    .nth(1)
                    .unwrap()
                    .trim()
                    .parse()
                    .unwrap();
                println!("{}", content_len);
                let mut body_buf = vec![0u8; content_len as usize];
                buf_reader.read_exact(&mut body_buf)?;
                // let body=String::from_utf8(body_buf).unwrap();
                f.write_all(&body_buf);
                let response = format!("HTTP/1.1 201 Created\r\nContent-Length: 0\r\n\r\n");
                stream.write_all(response.as_bytes())
            }?,
            _ => {
                println!("error");
            }
        };
        // println!("connection is asking for files");
        return Ok(());
    }
    // println!("{:?}",specific_path[1]);
    if specific_path[1].contains("/echo") {
        //["GET /echo/foo HTTP/1.1", "Host: localhost:4221", "User-Agent: curl/8.13.0", "Accept: */*", "Accept-Encoding: gzip"]
        let parts_path: Vec<_> = specific_path[1].split("/").collect();
        let lenght_returned = parts_path[2].len();
        let encoding_came = &http_request[4];
        let encoding_split: Vec<&str> = encoding_came.split(":").collect();
        let encodings:Vec<&str>=encoding_split[1].split(",").collect();
        /* println!("{:?}",encodings);
        [" gzip", " invalid-encoding-1", " invalid-encoding-2"] */
        // for encoding in encodings {}
        // Or if encodings.iter().any(...) {}
        if encodings.into_iter().any(|x| x.trim().starts_with("gzip")){
            let returned_response= format!(
                     "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: {}\r\n\r\n{}",
                     lenght_returned, parts_path[2]
                 );
                 stream.write_all(returned_response.as_bytes());
                 
        }else{
            let returned_response= format!(
                     "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: {}\r\n\r\n{}",
                     lenght_returned, parts_path[2]
                 );
                 stream.write_all(returned_response.as_bytes());
        }
        
        // let returned_value=match encoding_split[1].trim() {
        //     "gzip" => {
        //          format!(
        //             "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: {}\r\n\r\n{}",
        //             lenght_returned, parts_path[2]
        //         )            }
        //     _ =>{
        //           format!(
        //             "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: {}\r\n\r\n{}",
        //             lenght_returned, parts_path[2]
        //         )
        //     } 
        // };
        return Ok(());
    }
    // match http_request[0].trim() {
    //     // "GET /abcdefg HTTP/1.1" => stream.write_all(b"HTTP/1.1 404 Not Found\r\n\r\n"),
    //     "GET / HTTP/1.1" => stream.write_all(b"HTTP/1.1 200 OK\r\n\r\n"),
    //     _ => stream.write_all(b"HTTP/1.1 400 BAD REQUEST\r\n\r\n"),
    // };
    // println!("{:?}",specific_path);
    Ok(())
    // buf_reader.read_line(&mut url_path);
}
fn main() {
    let listener = TcpListener::bind("127.0.0.1:4221").unwrap();
    // let mut total_connection=0;
    for stream in listener.incoming() {
        let stream = stream.unwrap();
        // println!("recieved a connection");
        // total_connection+=1;
        // println!("{}",total_connection);
        thread::spawn(move || {
            handle_client(stream);
        });
        // stream.write_all("HTTP/1.1 200 OK\r\n\r\n".as_bytes()).unwrap();
    }
}
