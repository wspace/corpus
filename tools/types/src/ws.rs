use std::fmt::{self, Debug, Formatter};
use std::marker::PhantomData;
use std::mem;
use std::ops::{Index, IndexMut};

use serde::de::MapAccess;
use serde::{de::Visitor, Deserialize, Deserializer, Serialize, Serializer};
use strum::{Display, EnumCount, EnumString, FromRepr};

#[derive(
    Display,
    EnumCount,
    EnumString,
    FromRepr,
    Serialize,
    Deserialize,
    Clone,
    Copy,
    Debug,
    PartialEq,
    Eq,
    Hash,
)]
#[strum(serialize_all = "lowercase")]
#[serde(rename_all = "lowercase")]
pub enum Inst {
    Push,
    Dup,
    Copy,
    Swap,
    Drop,
    Slide,
    Add,
    Sub,
    Mul,
    Div,
    Mod,
    Store,
    Retrieve,
    Label,
    Call,
    Jmp,
    Jz,
    Jn,
    Ret,
    End,
    Printc,
    Printi,
    Readc,
    Readi,

    Shuffle,
    DumpStack,
    DumpHeap,
    DumpTrace,
}

#[derive(Clone, PartialEq, Eq)]
pub struct InstMap<T>(Box<[Option<T>; Inst::COUNT]>);

impl<T> InstMap<T> {
    const NONE: Option<T> = None;

    pub fn new() -> Self {
        InstMap(Box::new([Self::NONE; Inst::COUNT]))
    }

    pub fn get(&self, key: Inst) -> &Option<T> {
        &self.0[key as u8 as usize]
    }

    pub fn get_mut(&mut self, key: Inst) -> &mut Option<T> {
        &mut self.0[key as u8 as usize]
    }

    pub fn insert(&mut self, key: Inst, value: T) -> Option<T> {
        mem::replace(self.get_mut(key), Some(value))
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn len(&self) -> usize {
        self.0.iter().filter(|v| v.is_some()).count()
    }
}

impl<T> Default for InstMap<T> {
    fn default() -> Self {
        InstMap::new()
    }
}

impl<T> Index<Inst> for InstMap<T> {
    type Output = T;

    fn index(&self, key: Inst) -> &Self::Output {
        self.get(key).as_ref().unwrap()
    }
}

impl<T> IndexMut<Inst> for InstMap<T> {
    fn index_mut(&mut self, key: Inst) -> &mut Self::Output {
        self.get_mut(key).as_mut().unwrap()
    }
}

impl<T: Debug> Debug for InstMap<T> {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        f.debug_map()
            .entries(
                self.0
                    .iter()
                    .enumerate()
                    .filter_map(|(k, v)| v.as_ref().map(|v| (Inst::from_repr(k).unwrap(), v))),
            )
            .finish()
    }
}

impl<T: Serialize> Serialize for InstMap<T> {
    fn serialize<S: Serializer>(&self, s: S) -> Result<S::Ok, S::Error> {
        s.collect_map(
            self.0
                .iter()
                .enumerate()
                .filter_map(|(k, v)| v.as_ref().map(|v| (Inst::from_repr(k).unwrap(), v))),
        )
    }
}

impl<'de, T: Deserialize<'de>> Deserialize<'de> for InstMap<T> {
    fn deserialize<D: Deserializer<'de>>(d: D) -> Result<Self, D::Error> {
        struct MapVisitor<T>(PhantomData<fn() -> InstMap<T>>);

        impl<'de, T: Deserialize<'de>> Visitor<'de> for MapVisitor<T> {
            type Value = InstMap<T>;

            fn expecting(&self, f: &mut Formatter) -> fmt::Result {
                f.write_str("an instruction map")
            }

            fn visit_map<A: MapAccess<'de>>(self, mut access: A) -> Result<Self::Value, A::Error> {
                let mut map = InstMap::new();
                while let Some((k, v)) = access.next_entry()? {
                    map.insert(k, v);
                }
                Ok(map)
            }
        }

        d.deserialize_map(MapVisitor(PhantomData))
    }
}

#[macro_export]
macro_rules! inst_map {
    ($($key:ident => $val:expr),* $(,)?) => {{
        // Check uniqueness
        match Inst::Push {
            $(Inst::$key => {},)*
            _ => {},
        }
        let mut m = InstMap::new();
        $(m[Inst::$key] = Some($val);)*
        m
    }};
}
