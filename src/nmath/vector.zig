const std = @import("std");
const expect = std.testing.expect;
const builtin = @import("builtin");
const math = std.math;

const constants = @import("constants.zig");
/// A 4D Vector used for points and directions and colors.
pub const Vec4 = @Vector(4, f32);
pub const Point = Vec4;
pub const Vector = Vec4;

//pub const Vec3 = @Vector(3, f32);

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

pub fn vec3(v: Vec4) @Vector(3, f32) {
    return .{ v[0], v[1], v[2] };
}

pub const FmtVec4 = std.fmt.Alt(Vec4, format);
fn format(v: Vec4, wr: *std.Io.Writer) !void {
    try wr.print("{d} {d} {d} {d}", .{ v[0], v[1], v[2], v[3] });
}
pub fn fmt(v: Vec4) FmtVec4 {
    return .{
        .data = v,
    };
}

pub fn approxEq(lhs: Vec4, rhs: Vec4) bool {
    return @reduce(.And, @abs(lhs - rhs) < splat(constants.epsilon));
}

pub fn approxZero(v: Vec4) bool {
    return @reduce(.And, @abs(v) < splat(constants.epsilon));
}

pub fn magnitude(v: Vec4) f32 {
    return @sqrt(magnitudeSquared(v));
}

pub fn magnitudeSquared(v: Vec4) f32 {
    return @reduce(.Add, v * v);
}

pub fn normalize(v: Vec4) Vec4 {
    std.debug.assert(!std.math.approxEqAbs(f32, magnitude(v), 0, constants.epsilon));

    const mag: Vec4 = @splat(magnitude(v));
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

pub inline fn splat(s: f32) Vec4 {
    return @as(Vec4, @splat(s));
}

pub fn reflect(in: Vector, normal: Vector) Vector {
    return in - normal * splat(2) * splat(dot(in, normal));
}

test "vector:reflect works at 45 degrees" {
    const v = vector(1, -1, 0);
    const n = vector(0, 1, 0);
    const r = reflect(v, n);
    const expected = vector(1, 1, 0);
    try expect(approxEq(r, expected));
}

test "vector:reflect works at slanted angles" {
    const v = vector(0, -1, 0);
    const root_2_over_2 = @sqrt(2.0) / 2.0;
    const n = vector(root_2_over_2, root_2_over_2, 0);
    const r = reflect(v, n);
    const expected = vector(1, 0, 0);
    try expect(approxEq(r, expected));
}
