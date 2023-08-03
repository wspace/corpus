use std::cmp::Ordering;
use std::fmt::{self, Formatter};
use std::marker::PhantomData;

use rug::Integer;
use serde::{
    de::{self, Visitor},
    Deserialize, Deserializer, Serialize, Serializer,
};
use thiserror::Error;

/// An arbitrary-precision integer type, that can be represented in JSON as an
/// integer literal, when it is within the precision of float64, or as a string.
#[derive(Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
#[repr(transparent)]
pub struct Int(Integer);

/// An arbitrary-precision unsigned integer type, that can be represented in
/// JSON as an integer literal, when it is within the precision of float64, or
/// as a string.
#[derive(Deserialize, Clone, Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
#[serde(try_from = "Int")]
#[repr(transparent)]
pub struct Uint(Integer);

impl AsRef<Integer> for Int {
    fn as_ref(&self) -> &Integer {
        &self.0
    }
}

impl AsMut<Integer> for Int {
    fn as_mut(&mut self) -> &mut Integer {
        &mut self.0
    }
}

impl AsRef<Integer> for Uint {
    fn as_ref(&self) -> &Integer {
        &self.0
    }
}

impl From<Integer> for Int {
    fn from(n: Integer) -> Self {
        Int(n)
    }
}

impl From<Uint> for Int {
    fn from(n: Uint) -> Self {
        Int(n.0)
    }
}

#[derive(Error, Clone, Debug, PartialEq, Eq)]
#[error("negative unsigned integer")]
pub struct NegativeUintError;

impl TryFrom<Integer> for Uint {
    type Error = NegativeUintError;

    fn try_from(n: Integer) -> Result<Self, Self::Error> {
        if n.cmp0() == Ordering::Less {
            Err(NegativeUintError)
        } else {
            Ok(Uint(n))
        }
    }
}

impl TryFrom<Int> for Uint {
    type Error = NegativeUintError;

    fn try_from(n: Int) -> Result<Self, Self::Error> {
        Uint::try_from(n.0)
    }
}

impl Serialize for Int {
    fn serialize<S: Serializer>(&self, s: S) -> Result<S::Ok, S::Error> {
        serialize_integer(&self.0, s)
    }
}

impl Serialize for Uint {
    fn serialize<S: Serializer>(&self, s: S) -> Result<S::Ok, S::Error> {
        serialize_integer(&self.0, s)
    }
}

fn serialize_integer<S: Serializer>(n: &Integer, s: S) -> Result<S::Ok, S::Error> {
    match n.to_i64() {
        Some(n) if n.unsigned_abs() < 1 << 53 => n.serialize(s),
        _ => n.to_string().serialize(s),
    }
}

#[derive(Error, Clone, Debug, PartialEq, Eq)]
#[error("integer exceeds precision of f64")]
pub struct IntPrecisionError;

impl<'de> Deserialize<'de> for Int {
    fn deserialize<D: Deserializer<'de>>(d: D) -> Result<Self, D::Error> {
        struct IntVisitor(PhantomData<fn() -> Int>);

        impl<'de> Visitor<'de> for IntVisitor {
            type Value = Int;

            fn expecting(&self, f: &mut Formatter) -> fmt::Result {
                f.write_str("integer or integer as string")
            }

            fn visit_i64<E: de::Error>(self, n: i64) -> Result<Self::Value, E> {
                if n.unsigned_abs() < 1 << 53 {
                    Ok(Int(Integer::from(n)))
                } else {
                    Err(de::Error::custom(IntPrecisionError))
                }
            }

            fn visit_u64<E: de::Error>(self, n: u64) -> Result<Self::Value, E> {
                if n < 1 << 53 {
                    Ok(Int(Integer::from(n)))
                } else {
                    Err(de::Error::custom(IntPrecisionError))
                }
            }

            fn visit_str<E: de::Error>(self, s: &str) -> Result<Self::Value, E> {
                match parse_json_integer(s) {
                    Ok(n) => Ok(Int(n)),
                    Err(err) => Err(de::Error::custom(err)),
                }
            }
        }

        d.deserialize_any(IntVisitor(PhantomData))
    }
}

#[derive(Error, Clone, Debug, PartialEq, Eq)]
pub enum ParseJsonIntegerError {
    #[error("empty integer")]
    Empty,
    #[error("leading zero in integer")]
    LeadingZero,
    #[error("invalid digit in integer")]
    InvalidDigit,
    #[error("fraction in integer")]
    Fraction,
    #[error("exponent in integer")]
    Exponent,
}

/// Parses a JSON integer literal.
///
/// The grammar for JSON number literals is
/// `-?([1-9][0-9]*|0)(\.[0-9]*)?([eE][+-]?[0-9]+)?`, but excluding the fraction
/// and exponent components simplifies it to `-?([1-9][0-9]*|0)`.
fn parse_json_integer(s: &str) -> Result<Integer, ParseJsonIntegerError> {
    use ParseJsonIntegerError as E;
    let b = s.as_bytes();
    let (neg, b) = if !b.is_empty() && b[0] == b'-' {
        (true, &b[1..])
    } else {
        (false, b)
    };
    if b.is_empty() {
        return Err(E::Empty);
    }
    let mut digits = Vec::with_capacity(b.len());
    for &ch in b {
        let digit = ch.wrapping_sub(b'0');
        if digit > 9 {
            let err = match ch {
                b'.' => E::Fraction,
                b'e' => E::Exponent,
                _ => E::InvalidDigit,
            };
            return Err(err);
        }
        digits.push(digit);
    }
    if b.len() > 1 && b[0] == b'0' {
        return Err(E::LeadingZero);
    }
    let mut n = Integer::new();
    // SAFETY: Digits are in bounds for radix.
    unsafe {
        Integer::assign_bytes_radix_unchecked(&mut n, &digits, 10, neg);
    }
    Ok(n)
}

#[cfg(test)]
mod tests {
    use serde_json::from_str as from_json;
    use serde_json::to_string as to_json;

    use super::*;

    const MAX: u64 = (1u64 << 53) - 1;

    #[test]
    fn serialize() {
        assert_eq!("42", to_json(&Int(42.into())).unwrap());
        assert_eq!("9007199254740991", to_json(&Int(MAX.into())).unwrap(),);
        assert_eq!(
            "\"9007199254740992\"",
            to_json(&Int((MAX + 1).into())).unwrap(),
        );
    }

    #[test]
    fn deserialize() {
        assert_eq!(Int(42.into()), from_json("42").unwrap());
        assert_eq!(Int(42.into()), from_json("\"42\"").unwrap());
        assert_eq!(Int(MAX.into()), from_json("9007199254740991").unwrap());
        assert_eq!(
            Int((MAX + 1).into()),
            from_json("\"9007199254740992\"").unwrap()
        );
        assert!(from_json::<Int>("9007199254740992").is_err());
    }
}
