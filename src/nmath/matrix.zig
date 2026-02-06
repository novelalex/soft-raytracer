const builtin = @import("builtin");
const std = @import("std");
const constants = @import("constants.zig");
const Vec4 = @import("vector.zig").Vec4;

pub const Mat4 = extern struct {
    data: [16]f32,
    
    // ```
    // 0  1  2  3
    // 4  5  6  7
    // 8  9  10 11
    // 12 13 14 15
    // ```
    pub fn init(values: [16]f32) Mat4 {
        return .{ .data = values };
    }
    
    pub fn identity() Mat4 {
        return .{
            .data = [_]f32{
                1, 0, 0, 0,
                0, 1, 0, 0,
                0, 0, 1, 0,
                0, 0, 0, 1,
            },
        };
    }

    pub inline fn at(self: Mat4, x: usize, y: usize) f32 {
        if (builtin.mode == .Debug) {
            std.debug.assert(x < 4);
            std.debug.assert(y < 4);
        }
        return self.data[y * 4 + x];
    }

    pub fn approxEq(self: Mat4, other: Mat4) bool {
        for (self.data, other.data) |n, m| {
            const cmp = std.math.approxEqAbs(f32, n, m, constants.epsilon);
            if (!cmp) {
                return false;
            }
        }
        return true;
    }

    pub fn multiply(self: Mat4, other: Mat4) Mat4 {
        const a = self.data;
        const b = other.data;

        return .{
            .data = [_]f32{
                // row 0
                a[0] * b[0] + a[1] * b[4] + a[2] * b[8] + a[3] * b[12],
                a[0] * b[1] + a[1] * b[5] + a[2] * b[9] + a[3] * b[13],
                a[0] * b[2] + a[1] * b[6] + a[2] * b[10] + a[3] * b[14],
                a[0] * b[3] + a[1] * b[7] + a[2] * b[11] + a[3] * b[15],

                // row 1
                a[4] * b[0] + a[5] * b[4] + a[6] * b[8] + a[7] * b[12],
                a[4] * b[1] + a[5] * b[5] + a[6] * b[9] + a[7] * b[13],
                a[4] * b[2] + a[5] * b[6] + a[6] * b[10] + a[7] * b[14],
                a[4] * b[3] + a[5] * b[7] + a[6] * b[11] + a[7] * b[15],

                // row 2
                a[8] * b[0] + a[9] * b[4] + a[10] * b[8] + a[11] * b[12],
                a[8] * b[1] + a[9] * b[5] + a[10] * b[9] + a[11] * b[13],
                a[8] * b[2] + a[9] * b[6] + a[10] * b[10] + a[11] * b[14],
                a[8] * b[3] + a[9] * b[7] + a[10] * b[11] + a[11] * b[15],

                // row 3
                a[12] * b[0] + a[13] * b[4] + a[14] * b[8] + a[15] * b[12],
                a[12] * b[1] + a[13] * b[5] + a[14] * b[9] + a[15] * b[13],
                a[12] * b[2] + a[13] * b[6] + a[14] * b[10] + a[15] * b[14],
                a[12] * b[3] + a[13] * b[7] + a[14] * b[11] + a[15] * b[15],
            },
        };
    }

    pub fn multiplyVec(self: Mat4, v: Vec4) Vec4 {
        const m = self.data;

        return .{
            .x = m[0] * v.x + m[1] * v.y + m[2] * v.z + m[3] * v.w,
            .y = m[4] * v.x + m[5] * v.y + m[6] * v.z + m[7] * v.w,
            .z = m[8] * v.x + m[9] * v.y + m[10] * v.z + m[11] * v.w,
            .w = m[12] * v.x + m[13] * v.y + m[14] * v.z + m[15] * v.w,
        };
    }
    
    pub fn transpose(self: Mat4) Mat4 {
        return .{
            .data = [_]f32{
                self.data[1]
            }
        }
    }
};

pub const Mat3 = extern struct {
    data: [9]f32,

    pub fn init(values: [9]f32) Mat3 {
        return .{ .data = values };
    }

    pub fn identity() Mat3 {
        return .{
            .data = [_]f32{
                1, 0, 0,
                0, 1, 0,
                0, 0, 1,
            },
        };
    }

    pub inline fn at(self: Mat3, x: usize, y: usize) f32 {
        if (builtin.mode == .Debug) {
            std.debug.assert(x < 3);
            std.debug.assert(y < 3);
        }
        return self.data[y * 3 + x];
    }

    pub fn approxEq(self: Mat3, other: Mat3) bool {
        for (self.data, other.data) |n, m| {
            const cmp = std.math.approxEqAbs(f32, n, m, constants.epsilon);
            if (!cmp) {
                return false;
            }
        }
        return true;
    }
};

pub const Mat2 = extern struct {
    data: [4]f32,

    pub fn init(values: [4]f32) Mat2 {
        return .{ .data = values };
    }

    pub fn identity() Mat2 {
        return .{
            .data = [_]f32{
                1, 0,
                0, 1,
            },
        };
    }

    pub inline fn at(self: Mat2, x: usize, y: usize) f32 {
        if (builtin.mode == .Debug) {
            std.debug.assert(x < 2);
            std.debug.assert(y < 2);
        }
        return self.data[y * 2 + x];
    }

    pub fn approxEq(self: Mat2, other: Mat2) bool {
        for (self.data, other.data) |n, m| {
            const cmp = std.math.approxEqAbs(f32, n, m, constants.epsilon);
            if (!cmp) {
                return false;
            }
        }
        return true;
    }
};

test "4x4 matrix multiplication works" {
    const m1 = Mat4.init(.{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    });

    const m2 = Mat4.init(.{
        -2, 1, 2, 3,
        3,  2, 1, -1,
        4,  3, 6, 5,
        1,  2, 7, 8,
    });

    const expected = Mat4.init(.{
        20, 22, 50,  48,
        44, 54, 114, 108,
        40, 58, 110, 102,
        16, 26, 46,  42,
    });

    try std.testing.expect(m1.multiply(m2).approxEq(expected));
}

test "4x4 matrix and vec4 multiplication works" {
    const m = Mat4.init(.{
        1, 2, 3, 4,
        2, 4, 4, 2,
        8, 6, 4, 1,
        0, 0, 0, 1,
    });

    const v = Vec4.init(1, 2, 3, 1);

    const expected = Vec4.init(18, 24, 33, 1);

    try std.testing.expect(m.multiplyVec(v).approxEq(expected));
}

test "4x4 matrix multiplication with identity produces startng matrix" {
    const m1 = Mat4.init(.{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    });

    const m2 = Mat4.identity();

    try std.testing.expect(m1.multiply(m2).approxEq(m1));
}
