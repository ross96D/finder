use std::{rc::Rc, vec};

mod winit_helpers;
use winit_helpers::center_window;

use slint::{ComponentHandle, LogicalPosition, ModelRc, SharedString, VecModel};

slint::include_modules!();

fn main() -> Result<(), slint::PlatformError> {
    let ui = MyApp::new()?;
    ui.window().set_position(LogicalPosition::new(500.0, 500.0));

    center_window(ui.window());

    let vec_model: Rc<VecModel<SharedString>> = Rc::new(VecModel::from(vec![]));
    let vec_model_rc = ModelRc::from(vec_model.clone());
    ui.set_elements(vec_model_rc);

    ui.on_add_item(move |item: SharedString| vec_model.push(item));

    ui.run()
}
