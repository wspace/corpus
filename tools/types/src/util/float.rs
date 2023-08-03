pub use std::fmt::{self, Debug, Formatter};
use std::{cmp::Ordering, fmt::Display};

use serde::{Deserialize, Serialize};

macro_rules! TotalFloat(
    ($Type:ident of $inner:ty) => {
        #[derive(Serialize, Deserialize, Clone, Copy)]
        #[serde(transparent)]
        #[repr(transparent)]
        pub struct $Type(pub $inner);

        impl Debug for $Type {
            fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
                Debug::fmt(&self.0, f)
            }
        }

        impl Display for $Type {
            fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
                Display::fmt(&self.0, f)
            }
        }

        impl PartialEq for $Type {
            fn eq(&self, other: &Self) -> bool {
                self.0.total_cmp(&other.0) == Ordering::Equal
            }
        }

        impl Eq for $Type {}

        impl PartialOrd for $Type {
            fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
                Some(self.0.total_cmp(&other.0))
            }
        }

        impl Ord for $Type {
            fn cmp(&self, other: &Self) -> Ordering {
                self.0.total_cmp(&other.0)
            }
        }
    }
);

TotalFloat!(TotalF64 of f64);
TotalFloat!(TotalF32 of f32);
