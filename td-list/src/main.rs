use std::env;
use std::process;

mod info;

fn main() {
    let args: Vec<String> = get_args();
    debug_args(&args);

}

fn get_args() -> Vec<String> {
    let command_line_args: Vec<String> = env::args().collect();

    if command_line_args.len() == 1 {
        println!("{}", info::get_help_info());
        // handles errors here
    }


    return command_line_args;
}

fn debug_args(args: &Vec<String> ) {

    for (index, arg) in args.iter().enumerate() {
        println!("{}: {}", index, arg)
    }

}
