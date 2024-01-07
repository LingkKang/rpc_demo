//! Custom logger module.
//!
//! A custom logger implementation over [`log`] crate
//! that uses stdout to print logs.
//!
//! Note that `log` crate should be set with `"sdt"` feature.
//!
//! ``` toml
//! # Cargo.toml
//! [dependencies]
//! log = {version = "0.4.20", features = ["std"]}
//! ```
extern crate log;

use log::{LevelFilter, Log, Metadata, Record};

/// A custom logger struct that uses stdout to print logs.
///
/// # Example
///
/// ``` rust
/// extern crate log;
/// use crate::logger::logger::Logger;
/// fn main() {
///     Logger::init(Some(log::LevelFilter::Debug));
///     log::debug!("This is a debug message");
/// }
/// ```
pub struct Logger {
    /// The default level of the logger.
    default_level: LevelFilter,
}

impl Logger {
    /// Static method to initialize the logger
    /// with an optional logging level.
    pub fn init(level: Option<LevelFilter>) {
        // If the level is not specified, use `Trace` as default level.
        let record_level = level.unwrap_or(LevelFilter::Trace);
        log::set_max_level(record_level);
        let logger = Logger {
            default_level: record_level,
        };
        // Set the logger as the global logger.
        // Note that `log` crate should be set with "sdt" feature.
        // See Cargo.toml for more details.
        log::set_boxed_logger(Box::new(logger)).unwrap();
    }
}

impl Log for Logger {
    /// Check if current message level is enabled for logging.
    fn enabled(&self, metadata: &Metadata) -> bool {
        metadata.level() <= self.default_level
    }

    /// Log the message.
    fn log(&self, record: &Record) {
        if self.enabled(record.metadata()) {
            let level_string = {
                {
                    record.level().to_string().to_uppercase()
                }
            };
            let message = format!("[{}]\t {}", level_string, record.args());
            println!("{}", message);
        }
    }

    /// Flush the logger.
    /// As stdout is used, no need to flush, so this method is empty.
    fn flush(&self) {}
}
