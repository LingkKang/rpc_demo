//! Time utility for the logger.

use std::fmt::Write;
use std::time::{SystemTime, UNIX_EPOCH};

/// Returns the current time in a formatted string.
/// The format is `HH:MM:SS.NNNN`.
pub fn get_formatted_time() -> String {
    let now = SystemTime::now();
    let since_the_epoch = now.duration_since(UNIX_EPOCH).unwrap();

    let seconds = since_the_epoch.as_secs();
    let nanos = since_the_epoch.subsec_nanos();

    let seconds_of_the_day = seconds % (60 * 60 * 24);
    let hour = seconds_of_the_day / (60 * 60);
    let minute = (seconds_of_the_day % (60 * 60)) / 60;
    let second = seconds_of_the_day % 60;

    let mut time_string = String::new();
    write!(
        &mut time_string,
        "{:02}:{:02}:{:02}.{:04}",
        hour,
        minute,
        second,
        nanos / 1_000_000
    )
    .unwrap();
    time_string
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_system_time_to_date() {
        println!("{}", get_formatted_time());
    }
}
