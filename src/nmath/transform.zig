const vec = @import("vector.zig");
const matrix = @import("matrix.zig");
const quat = @import("quaternion.zig");
const std = @import("std");
const constants = @import("constants.zig");

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

test "Transform:translation invert can be applied to points w=1" {
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

pub fn scaling(x: f32, y: f32, z: f32) matrix.Mat4 {
    return .{
        x, 0, 0, 0,
        0, y, 0, 0,
        0, 0, z, 0,
        0, 0, 0, 1,
    };
}

test "Transform:scaling can be applied to points w=1" {
    const t = scaling(2, 3, 4);
    const p = vec.point(-4, 6, 8);
    const expected = vec.point(-8, 18, 32);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(t, p), expected));
}

test "Transform:scaling can be applied to vectors w=0" {
    const t = scaling(2, 3, 4);
    const v = vec.vector(-4, 6, 8);
    const expected = vec.vector(-8, 18, 32);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(t, v), expected));
}

test "Transform:scaling invert can be applied" {
    const t = scaling(2, 3, 4);
    const inv = matrix.mat4Inverse(t);
    const v = vec.vector(-4, 6, 8);
    const expected = vec.vector(-2, 2, 2);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(inv, v), expected));
}

test "Transform:scaling with negative value can reflect points" {
    const t = scaling(-1, 1, 1);
    const p = vec.point(2, 3, 4);
    const expected = vec.point(-2, 3, 4);

    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(t, p), expected));
}

pub fn rotation(angle: f32, axis: vec.Vec4) matrix.Mat4 {
    const rotation_quat = quat.Quaternion.angleAxisRotation(angle, axis);
    return rotation_quat.toMatrix();
}

pub fn rotation_x(angle: f32) matrix.Mat4 {
    return rotation(angle, vec.Vec4{ 1, 0, 0, 0 });
}

pub fn rotation_y(angle: f32) matrix.Mat4 {
    return rotation(angle, vec.Vec4{ 0, 1, 0, 0 });
}

pub fn rotation_z(angle: f32) matrix.Mat4 {
    return rotation(angle, vec.Vec4{ 0, 0, 1, 0 });
}

test "Transform:rotate_x rotates a point around the x axis" {
    const p = vec.point(0, 1, 0);
    const half_quarter = rotation_x(constants.pi / 4.0);
    const full_quarter = rotation_x(constants.pi / 2.0);
    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(half_quarter, p), vec.point(0, @sqrt(2.0) / 2.0, @sqrt(2.0) / 2.0)));
    try std.testing.expect(vec.approxEq(matrix.mat4MultiplyVec(full_quarter, p), vec.point(0, 0, 1)));
}
