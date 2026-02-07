const std = @import("std");
const Writer = std.Io.Writer;
const vec = @import("vector.zig");

pub const Color = vec.Vec4;

pub fn init(_r: f32, _g: f32, _b: f32) Color {
    return .{ _r, _g, _b, 1 };
}

pub fn r(c: Color) f32 {
    c[0];
}
pub fn g(c: Color) f32 {
    c[1];
}
pub fn b(c: Color) f32 {
    c[2];
}

pub fn hadamardProduct(lhs: Color, rhs: Color) Color {
    return lhs * rhs;
}
