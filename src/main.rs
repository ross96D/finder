use std::{env, path::Path, process::exit, rc::Rc, vec};

mod winit_helpers;
mod matcher;
use matcher::preview;
use winit_helpers::center_window;

use slint::{ComponentHandle, LogicalPosition, ModelRc, SharedString, VecModel, Weak};

slint::include_modules!();

fn main() -> Result<(), slint::PlatformError> {
    let app = MyApp::new()?;
    let weak_app: Weak<MyApp> = app.as_weak();
    // app.window().set_position(LogicalPosition::new(500.0, 500.0));

    center_window(app.window());

    let vec_model: Rc<VecModel<SharedString>> = Rc::new(VecModel::from(vec![]));
    let vec_model_rc = ModelRc::from(vec_model.clone());
    app.set_elements(vec_model_rc);

    app.on_search(move |item: SharedString| {
        let app = weak_app.upgrade().unwrap();
        let cwd = match env::current_dir() {
            Ok(cwd) => cwd,
            Err(err) => {
                // ! TODO handle the error..
                panic!("{}", err)
            },
        };
        // ! TODO handle the error
        let resp = matcher::search(&item, cwd.as_os_str()).unwrap();
        for e in resp {
            vec_model.push(SharedString::from(e.to_string()));
            if e.focus {
                let preview = match preview(Path::new(&e.file_path), e.line_number, 10) {
                    Ok(preview) => preview,
                    Err(err) => {
                        // TODO log error
                        eprint!("error getting preview {}", err);
                        exit(1)
                    },
                };
                app.set_preview(SharedString::from(preview));
            }
        }
    });

    app.run()
}
