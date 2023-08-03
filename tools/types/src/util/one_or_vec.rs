use std::cmp::Ordering;
use std::fmt::{self, Debug, Formatter};
use std::hash::{Hash, Hasher};

use serde::{Deserialize, Deserializer, Serialize, Serializer};

#[repr(transparent)]
pub struct OneOrVec<T>(pub Vec<T>);

impl<T> OneOrVec<T> {
    pub fn is_empty(&self) -> bool {
        self.0.is_empty()
    }

    pub fn len(&self) -> usize {
        self.0.len()
    }
}

impl<T> AsRef<Vec<T>> for OneOrVec<T> {
    fn as_ref(&self) -> &Vec<T> {
        &self.0
    }
}

impl<T> AsMut<Vec<T>> for OneOrVec<T> {
    fn as_mut(&mut self) -> &mut Vec<T> {
        &mut self.0
    }
}

impl<T> AsRef<[T]> for OneOrVec<T> {
    fn as_ref(&self) -> &[T] {
        &self.0
    }
}

impl<T> AsMut<[T]> for OneOrVec<T> {
    fn as_mut(&mut self) -> &mut [T] {
        &mut self.0
    }
}

impl<T> Default for OneOrVec<T> {
    fn default() -> Self {
        OneOrVec(Vec::new())
    }
}

impl<T: Clone> Clone for OneOrVec<T> {
    fn clone(&self) -> Self {
        OneOrVec(self.0.clone())
    }
}

impl<T: PartialEq> PartialEq for OneOrVec<T> {
    fn eq(&self, other: &Self) -> bool {
        self.0 == other.0
    }
}

impl<T: Eq> Eq for OneOrVec<T> {}

impl<T: PartialOrd> PartialOrd for OneOrVec<T> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.0.partial_cmp(&other.0)
    }
}

impl<T: Ord> Ord for OneOrVec<T> {
    fn cmp(&self, other: &Self) -> Ordering {
        self.0.cmp(&other.0)
    }
}

impl<T: Hash> Hash for OneOrVec<T> {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.0.hash(state);
    }
}

impl<T: Debug> Debug for OneOrVec<T> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        if let [v] = &*self.0 {
            v.fmt(f)
        } else {
            self.0.fmt(f)
        }
    }
}

impl<T: Serialize> Serialize for OneOrVec<T> {
    fn serialize<S: Serializer>(&self, s: S) -> Result<S::Ok, S::Error> {
        if let [v] = &*self.0 {
            v.serialize(s)
        } else {
            self.0.serialize(s)
        }
    }
}

impl<'de, T: Deserialize<'de>> Deserialize<'de> for OneOrVec<T> {
    fn deserialize<D: Deserializer<'de>>(d: D) -> Result<Self, D::Error> {
        #[derive(Deserialize)]
        #[serde(untagged)]
        enum OneOrVecEnum<T> {
            One(T),
            Vec(Vec<T>),
        }
        match OneOrVecEnum::deserialize(d)? {
            OneOrVecEnum::One(val) => Ok(OneOrVec(vec![val])),
            OneOrVecEnum::Vec(vals) => Ok(OneOrVec(vals)),
        }
    }
}
