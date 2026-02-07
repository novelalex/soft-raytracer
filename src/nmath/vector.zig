const std = @import("std");
const builtin = @import("builtin");
const math = std.math;
const testing = std.testing;

const constants = @import("constants.zig");
/// A 4D Vector used for points and directions and colors.
pub const Vec4 = @Vector(4, f32);

pub fn x(v: Vec4) f32 {
    return v[0];
}
pub fn y(v: Vec4) f32 {
    return v[1];
}
pub fn z(v: Vec4) f32 {
    return v[2];
}
pub fn w(v: Vec4) f32 {
    return v[3];
}

pub fn point(_x: f32, _y: f32, _z: f32) Vec4 {
    return .{ _x, _y, _z, 1 };
}

pub fn vector(_x: f32, _y: f32, _z: f32) Vec4 {
    return .{ _x, _y, _z, 0 };
}

pub fn isPoint(v: Vec4) bool {
    return v[3] == 1;
}

pub fn isVector(v: Vec4) bool {
    return v[3] == 0;
}

pub const Fmt = std.fmt.Alt(Vec4, format);
fn format(v: Vec4, wr: *std.Io.Writer) !void {
    try wr.print("{d} {d} {d}", .{ v[0], v[1], v[2] });
}

pub fn approxEq(lhs: Vec4, rhs: Vec4) bool {
    return @reduce(.And, @abs(lhs - rhs) < @as(Vec4, @splat(constants.epsilon)));
}

pub fn approxZero(v: Vec4) bool {
    return @reduce(.And, @abs(v) < @as(Vec4, @splat(constants.epsilon)));
}

pub fn magnitude(v: Vec4) f32 {
    return @sqrt(magnitudeSquared(v));
}

pub fn magnitudeSquared(v: Vec4) f32 {
    return @reduce(.Add, v * v);
}

pub fn normalize(v: Vec4) Vec4 {
    std.debug.assert(std.math.approxEqAbs(f32, magnitude(v), 0, constants.epsilon));

    const mag: Vec4 = @splat(v.magnitude());
    return v / mag;
}

/// We should only be using dot and cross product on vectors and not points.
pub fn dot(lhs: Vec4, rhs: Vec4) f32 {
    return @reduce(.Add, lhs * rhs);
}

/// 3D cross product, w is dropped
pub fn cross(lhs: Vec4, rhs: Vec4) Vec4 {
    return .{
        lhs[1] * rhs[2] - lhs[2] * rhs[1],
        lhs[2] * rhs[0] - lhs[0] * rhs[2],
        lhs[0] * rhs[1] - lhs[1] * rhs[0],
        0,
    };
}
