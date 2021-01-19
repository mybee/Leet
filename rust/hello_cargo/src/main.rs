use  std::collections::HashMap;
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

const Max: u32 = 21000;

fn main() {

    let mut  v: Vec<i32> = Vec::new();
    v.push(1);

    let mut scores: HashMap<String, i32> = HashMap::new();
    scores.insert(String::from("bule"),10);

    let k = String::from("yello");


    // let name: type
    println!("{}", Max);
    println!("{}", usize::max_value());
    let inte = Point{x: 1, y:3};
    println!("{:#?}", inte)
}

fn return_clo() -> fn(i32) -> i32 {

}

trait OutPrint: fmt::Dispaly{ // 要求实现display traint
    fn out_print(&self) {
    let out = self.to_string();

}
}

fn func(v: i32) -> i32 {
    v+1
}


#[derive(Debug)]
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn get_x(&self) -> &T {
    &self.x
}
   fn get_y(&self) -> &T {
&self.y
}

}

fn largest<T: PartialOrd + Copy> (list: &[T]) {

}

enum Option<T, E> {
    Ok(T),
    Err(E)
}