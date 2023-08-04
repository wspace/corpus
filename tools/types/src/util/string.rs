use std::fmt::{self, Display, Formatter};

#[derive(Clone, Debug, Default, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub struct LossyString<'a>(&'a [u8]);

impl<'a> LossyString<'a> {
    pub fn new<B: AsRef<[u8]>>(bytes: &'a B) -> Self {
        LossyString(bytes.as_ref())
    }
}

impl Display for LossyString<'_> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        f.write_str(&String::from_utf8_lossy(&self.0))
    }
}
