//! The checksum submodule provides functions for generating 
//! and verifying checksums by using XOR.

/// Use XOR to generate checksum.
pub fn generate_checksum(data: &[u8]) -> u8 {
    let mut checksum: u8 = 0;
    for &byte in data {
        checksum ^= byte;
    }
    checksum
}

/// Verify checksum by comparing it with the generated checksum
pub fn verify_checksum(data: &[u8], checksum: u8) -> bool {
    generate_checksum(data) == checksum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_generate_checksum() {
        let data = [0x01, 0x02, 0x03, 0x04];
        let checksum = generate_checksum(&data);
        assert_eq!(checksum, 0x04);

        // Test with all zeros
        let data = [0x00, 0x00, 0x00, 0x00];
        assert_eq!(generate_checksum(&data), 0x00);
    }

    #[test]
    fn test_verify_checksum() {
        let data = [0x01, 0x02, 0x03, 0x04];
        let checksum = generate_checksum(&data);
        assert!(verify_checksum(&data, checksum));

        // Test with wrong checksum
        assert!(!verify_checksum(&data, checksum + 1));
        assert!(!verify_checksum(&data, checksum >> 1));

        // Test with all zeros
        let data = [0x00, 0x00, 0x00, 0x00];
        let checksum = generate_checksum(&data);
        assert!(verify_checksum(&data, checksum));
    }
}
