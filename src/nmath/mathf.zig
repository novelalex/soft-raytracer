const constants = @import("constants.zig");
const std = @import("std");

pub fn approxEq(n: f32, m: f32) bool {
    return std.math.approxEqAbs(f32, n, m, constants.epsilon);
}
