use std::{fs::File, io};

fn main() {
    let mut file = File::open("examples/README.md").unwrap();
    let mut stdout = io::stdout();
    io::copy(&mut file, &mut stdout).unwrap();
    // let r = io::copy(&mut file, &mut stdout);
    // match r {
    //     Ok(_) => println!("ok"),
    //     Err(e) => println!("error: {}", e),
    // }
}
