extern "C" {
    #[link_name = "println"]
    fn host_println(ptr: *const u8, len: usize);
}

pub fn println(s: &str) {
    unsafe {
        host_println(s.as_ptr(), s.len());
    }
}
