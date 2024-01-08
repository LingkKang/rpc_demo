//! The checksum submodule provides functions for generating
//! and verifying checksums by using XOR.

use super::message::Byte;

/// Use XOR to generate checksum.
pub fn generate_checksum(data: &[Byte]) -> Byte {
    let mut checksum: Byte = 0;
    for &byte in data {
        checksum ^= byte;
    }
    checksum
}

/// Verify checksum by comparing it with the generated checksum
pub fn verify_checksum(data: &[Byte], checksum: Byte) -> bool {
    generate_checksum(data) == checksum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_generate_checksum() {
        let data: [Byte; 4] = [0x01, 0x02, 0x03, 0x04];
        let checksum: Byte = generate_checksum(&data);
        assert_eq!(checksum, 0x04);

        // Test with all zeros
        let data: [Byte; 4] = [0x00, 0x00, 0x00, 0x00];
        assert_eq!(generate_checksum(&data), 0x00);
    }

    #[test]
    fn test_verify_checksum() {
        let data: [Byte; 4] = [0x01, 0x02, 0x03, 0x04];
        let checksum: Byte = generate_checksum(&data);
        assert!(verify_checksum(&data, checksum));

        // Test with wrong checksum
        assert!(!verify_checksum(&data, checksum + 1));
        assert!(!verify_checksum(&data, checksum >> 1));

        // Test with all zeros
        let data: [Byte; 4] = [0x00, 0x00, 0x00, 0x00];
        let checksum: Byte = generate_checksum(&data);
        assert!(verify_checksum(&data, checksum));
    }
}
