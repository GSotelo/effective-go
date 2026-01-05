# Strings

In Go, a string is a sequence of characters stored as UTF-8 bytes, and once it’s created, it cannot be changed. The default value of a string is an empty string (""), and strings can be compared with each other using operators like == or <.

Runes are an alias for the int32 type and represent Unicode code points. When referring to individual characters, use the rune type rather than int32, as it's more idiomatic and clearly conveys your intent.