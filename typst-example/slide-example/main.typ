// Get Polylux from the official package repository
#import "@preview/polylux:0.3.1": *

// Make the paper dimensions fit for a presentation and the text larger
#set page(paper: "presentation-16-9")
#set text(font: "Noto Sans CJK JP", size: 25pt)

// Use #polylux-slide to create a slide and style it using your favourite Typst functions
#polylux-slide[
  #align(horizon + center)[
    = おもしろスライドツール Typst

    yosaka

    2024/09/01
  ]
]

#polylux-slide[
== 一枚目のスライド

Some static text on this slide. Adding `rbx` to `rcx` gives the desired result.

What is ```rust fn main()``` in Rust would be ```cpp int main()``` in C.

```rust
fn main() {
    println!("Hello World!");
}
```

This has ``` `backticks` ``` in it (but the spaces are trimmed). And
``` here``` the leading space is also trimmed.
]

#polylux-slide[
  == This slide changes!

  You can always see this.
  // Make use of features like #uncover, #only, and others to create dynamic content
  #uncover(2)[But this appears later!]
]
