#![allow(non_snake_case)]
use dioxus::prelude::*;
use dioxus_slides::prelude::*;

fn main() {
    dioxus_tui::launch(App);
}

#[derive(Slidable, Clone, Default)]
enum Slides {
    #[default]
    Intro,
    WhatIsRust,
    WhyLearnRust,
    WhyYouShouldntLearnRust,
    JumpingShip,
    ErrorHandling,
    Toolchain,
    Ecosystem,
    Final,
}

fn App(cx: Scope) -> Element {
    cx.render(rsx! {
        div {
            style: "position: relative; min-width: 10000px;",
            SlideContainer::<Slides> { width: "10000px", height: "100%", enable_keyboard_navigation: true }
        }
    })
}

#[derive(Props)]
pub struct HeaderProps<'a> {
    title: &'a str,
}

fn Header<'a>(cx: Scope<'a, HeaderProps>) -> Element<'a> {
    cx.render(rsx! {
        div {
            flex_direction : "column",
            border_width : "1px",

            h1 {
                height : "2px",
                color : "#b7410e",
                justify_content: "center",
                cx.props.title,
            }
        }
    })
}

#[derive(Props)]
pub struct ListProps<'a> {
    children: Vec<ListItem<'a>>,
}

#[derive(Props)]
pub struct ListItem<'a> {
    content: &'a str,
}

fn List<'a>(cx: Scope<'a, ListProps>) -> Element<'a> {
    cx.render(rsx! {
        div {
            flex_direction : "column",

            ul {
                flex_direction: "column",
                cx.props.children.iter().map(|item| {
                    rsx! {
                        li {
                            div{ color : "#b7410e", "- " } div {"{item.content}"}
                        }
                    }
                })
            }
        }
    })
}

fn Intro(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Hello, Gophers!" }
                p {
                    "I'm exploring Rust, and you should give it a try too."
                }
            }
        }
    }))
}

fn WhatIsRust(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "What is Rust?" }

                List {
                    children: vec![
                        ListItem { content: "A systems programming language" },
                        ListItem { content: "Designed for safety" },
                        ListItem { content: "Designed for concurrency" },
                        ListItem { content: "Designed for speed" },
                    ]
                }
            }
        }
    }))
}

fn WhyLearnRust(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Why learn Rust?" }

                List {
                    children: vec![
                        ListItem { content: "Better control over low level details" },
                        ListItem { content: "Memory safety without garbage collection" },
                        ListItem { content: "Macros/Generics/other fun features" },
                        ListItem { content: "Fearless concurrency" },
                        ListItem { content: "Thriving ecosystem" },
                        ListItem { content: "Frequent language updates" },
                    ]
                }
            }
        }
    }))
}

fn WhyYouShouldntLearnRust(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Why you shouldn't learn Rust" }

                List {
                    children: vec![
                        ListItem { content: "You're happy with your current language" },
                        ListItem { content: "You're unfamiliar and have tight deadlines" },
                        ListItem { content: "You're not interested in low level details" },
                    ]
                }
            }
        }
    }))
}

fn JumpingShip(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Jumping ship" }

                p {
                    "Rust needs a different mindset."
                }
            }
        }
    }))
}

fn ErrorHandling(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Error handling" }

                List {
                    children: vec![
                        ListItem { content: "Go functions return 2+ parameters: `($THING, error)`" },
                        ListItem { content: "Typically the error is then checked, if ok $THING is used." },
                    ]
                }
                div {
                    border_width : "0.5px",
                }
                List {
                    children: vec![
                        ListItem { content: "Rust uses Result<T, E> for handling results" },
                        ListItem { content: "Result<T, E> is an enum with Ok(T) and Err(E) variants" },
                    ]
                }
            }
        }
    }))
}

fn Toolchain(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Toolchain Considerations" }

                List {
                    children: vec![
                        ListItem { content: "Compile times are slow (initially)" },
                        ListItem { content: "Formatting code is roughly the same as Go" },
                        ListItem { content: "Easy to update version within the tools provided" },
                        ListItem { content: "Easy to target different architectures - kinda" },
                    ]
                }
            }
        }
    }))
}

fn Ecosystem(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "What can I use rust for today?" }

                List {
                    children: vec![
                        ListItem { content: "Critical systems" },
                        ListItem { content: "Embedded systems" },
                        ListItem { content: "Video Games" },
                        ListItem { content: "WebAssembly" },
                        ListItem { content: "Microservices" },
                        ListItem { content: "Apparently making a slide deck" },
                    ]
                }

                div {
                    border_width : "0.5px",
                    "Really, not much you can't do with it."
                }
            }
        }
    }))
}

fn Final(cx: Scope) -> Element {
    cx.render(rsx!(Slide::<Slides> {
        content: render! {
            div {
                flex_direction : "column",
                border_width : "1px",

                Header { title: "Thank you!" }
                p {
                    justify_content: "center",
                    "Any questions so far?"
                }

                p {
                    border_width : "0.5px",
                    "I'm @alexanderjophus on Twitter and GitHub"
                }
            }
        }
    }))
}
