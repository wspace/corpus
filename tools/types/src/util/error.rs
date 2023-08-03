use std::error::Error;
use std::fmt::{self, Debug, Display, Formatter};

#[derive(Debug)]
pub struct MultiError<T>(pub Vec<T>);

impl<T> MultiError<T> {
    pub fn new() -> Self {
        MultiError(Vec::new())
    }

    pub fn push(&mut self, err: T) {
        self.0.push(err);
    }

    pub fn err(self) -> Result<(), Self> {
        if self.0.is_empty() {
            Ok(())
        } else {
            Err(self)
        }
    }
}

impl<T: Debug + Display> Error for MultiError<T> {}

impl<T: Display> Display for MultiError<T> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        for err in &self.0 {
            writeln!(f, "- {err}")?;
        }
        Ok(())
    }
}
