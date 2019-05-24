mod host;

#[no_mangle]
pub extern fn exported() {
    host::println("test string");
}
