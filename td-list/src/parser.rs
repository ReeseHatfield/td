use crate::command::Command;

pub fn parse_args(args: &Vec<String>) -> Command{
    let action: &String = &args[1];
    println!("Current action: {}", action);

    let action_args: Vec<String> = (&args[2..]).to_vec();
    println!("action args: {:?}", action_args);

    let parsed_cmd = Command {
        function: action.clone(),
        args: action_args
    };

    return parsed_cmd;
}