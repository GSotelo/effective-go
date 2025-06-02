# Numeric types
- [ ] TODO: Add more details about numeric types, including examples and usage.

## Integers

| Type     | Size   | Range                                       |
|----------|--------|---------------------------------------------|
| `uint8`  | 8-bit  | 0 to 255                                    |
| `uint16` | 16-bit | 0 to 65535                                  |
| `uint32` | 32-bit | 0 to 4294967295                             |
| `uint64` | 64-bit | 0 to 18,446,744,073,709,551,615             |
| `int8`   | 8-bit  | -128 to 127                                 |
| `int16`  | 16-bit | -32768 to 32767                             |
| `int32`  | 32-bit | -2147483648 to 2147483647                   |
| `int64`  | 64-bit | -9223372036854775808 to 9223372036854775807 |

### Aliases

| Alias  | Actual Type | Description                                             |
|--------|-------------|---------------------------------------------------------|
| `byte` | `uint8`     | Represents raw data, binary values, or ASCII characters |
| `rune` | `int32`     | Represents Unicode code points                          |
