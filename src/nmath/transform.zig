const vec = @import("vector.zig");
const matrix = @import("matrix.zig");
const std = @import("std");

pub fn translation(x: f32, y: f32, z: f32) matrix.Mat4 {
    return .{
        1, 0, 0, x,
        0, 1, 0, y,
        0, 0, 1, z,
        0, 0, 0, 1,
    };
}

test "Transform:translation can be applied to points" {
    const t = translation(5, -3, 2);
    const p = vec.point(-3, 4, 5);
    const expected = vec.point(2, 1, 7);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(t, p), expected));
}

test "Transform:translation invert can be applied to points" {
    const t = translation(5, -3, 2);
    const inv = matrix.mat4Inverse(t);
    const p = vec.point(-3, 4, 5);
    const expected = vec.point(-8, 7, 3);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(inv, p), expected));
}

test "Transform:translation cannot be applied to vectors w=0" {
    const t = translation(5, -3, 2);
    const v = vec.vector(-3, 4, 5);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(t, v), v));
}
