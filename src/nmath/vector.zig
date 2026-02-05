const std = @import("std");
const math = std.math;
const testing = std.testing;
/// A 4D Vector used for points and directions
/// ... and colors.
pub const Vec4 = struct {
    x: f32,
    y: f32,
    z: f32,
    w: f32,

    pub fn init(x: f32, y: f32, z: f32, w: f32) Vec4 {
        return .{
            .x = x,
            .y = y,
            .z = z,
            .w = w,
        };
    }

    pub fn point(x: f32, y: f32, z: f32) Vec4 {
        return .{
            .x = x,
            .y = y,
            .z = z,
            .w = 1,
        };
    }

    pub fn vector(x: f32, y: f32, z: f32) Vec4 {
        return .{
            .x = x,
            .y = y,
            .z = z,
            .w = 0,
        };
    }

    pub fn isPoint(self: Vec4) bool {
        return self.w == 1;
    }

    pub fn isVector(self: Vec4) bool {
        return self.w == 0;
    }

    const epsilon = 1e-6;
    pub fn approxEq(self: Vec4, other: Vec4) bool {
        return math.approxEqAbs(f32, self.x, other.x, epsilon) and
            math.approxEqAbs(f32, self.y, other.y, epsilon) and
            math.approxEqAbs(f32, self.z, other.z, epsilon) and
            math.approxEqAbs(f32, self.w, other.w, epsilon);
    }

    // comptime could be used to do component wise operators but
    // "Simple is better than complex." - The Zen of Python

    pub fn add(self: Vec4, other: Vec4) Vec4 {
        return .{
            .x = self.x + other.x,
            .y = self.y + other.y,
            .z = self.z + other.z,
            .w = self.w + other.w,
        };
    }

    pub fn sub(self: Vec4, other: Vec4) Vec4 {
        return .{
            .x = self.x - other.x,
            .y = self.y - other.y,
            .z = self.z - other.z,
            .w = self.w - other.w,
        };
    }

    pub fn negate(self: Vec4) Vec4 {
        return .{
            .x = -self.x,
            .y = -self.y,
            .z = -self.z,
            .w = -self.w,
        };
    }

    pub fn mult(self: Vec4, scalar: f32) Vec4 {
        return .{
            .x = self.x * scalar,
            .y = self.y * scalar,
            .z = self.z * scalar,
            .w = self.w * scalar,
        };
    }

    pub fn div(self: Vec4, scalar: f32) Vec4 {
        return .{
            .x = self.x / scalar,
            .y = self.y / scalar,
            .z = self.z / scalar,
            .w = self.w / scalar,
        };
    }

    pub fn magnitude(self: Vec4) f32 {
        return math.sqrt(self.x * self.x +
            self.y * self.y +
            self.z * self.z +
            self.w * self.w);
    }

    pub fn magnitudeSquared(self: Vec4) f32 {
        return self.x * self.x +
            self.y * self.y +
            self.z * self.z +
            self.w * self.w;
    }

    pub fn normalize(self: Vec4) Vec4 {
        const mag = self.magnitude();
        return .{
            .x = self.x / mag,
            .y = self.y / mag,
            .z = self.z / mag,
            .w = self.w / mag,
        };
    }

    /// We should only be using dot and cross product on vectors and not points.
    pub fn dot(self: Vec4, other: Vec4) f32 {
        return self.x * other.x +
            self.y * other.y +
            self.z * other.z +
            self.w * other.w;
    }

    /// 3D cross product, w is dropped
    pub fn cross(self: Vec4, other: Vec4) Vec4 {
        return .{
            .x = self.y * other.z - self.z * other.y,
            .y = self.z * other.x - self.x * other.z,
            .z = self.x * other.y - self.y * other.x,
            .w = 0,
        };
    }
};
