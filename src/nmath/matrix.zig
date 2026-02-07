const builtin = @import("builtin");
const std = @import("std");
const constants = @import("constants.zig");
const vec = @import("vector.zig");
const mathf = @import("mathf.zig");

pub inline fn element(comptime M: type, m: M, x: usize, y: usize) f32 {
    const info = @typeInfo(M);

    const dim = comptime @as(usize, @intFromFloat(std.math.sqrt(info.vector.len)));
    if (dim * dim != info.vector.len)
        @compileError("matrix vector must be square");

    if (builtin.mode == .Debug) {
        std.debug.assert(x < dim);
        std.debug.assert(y < dim);
    }

    return m[y * dim + x];
}

pub fn approxEq(comptime M: type, lhs: M, rhs: M) bool {
    return @reduce(.And, @abs(lhs - rhs) < @as(M, @splat(constants.epsilon_loose)));
}

test "Mat4:approxEq returns true on equal Mat4s" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const m2 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    try std.testing.expect(approxEq(Mat4, m1, m2));
}

test "Mat4:approxEq returns false on diffrent Mat4s" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const m2 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 8,
        5, 4, 3, 2,
    };

    try std.testing.expect(!approxEq(Mat4, m1, m2));
}

/// Mat4 is a 4D row major Matrix indexed like so:
///```
/// 0  1  2  3
/// 4  5  6  7
/// 8  9  10 11
/// 12 13 14 15
/// ```
pub const Mat4 = @Vector(16, f32);

test "Mat4:element index works" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const el = element(Mat4, m1, 1, 2);
    const expected = 8;

    try std.testing.expect(el == expected);
}

pub const mat4_identity =
    Mat4{
        1, 0, 0, 0,
        0, 1, 0, 0,
        0, 0, 1, 0,
        0, 0, 0, 1,
    };



pub fn mat4Multiply(a: Mat4, b: Mat4) Mat4 {
    return .{
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
    };
}

test "Mat4:mat4Multiply works" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const m2 = Mat4{
        -2, 1, 2, 3,
        3,  2, 1, -1,
        4,  3, 6, 5,
        1,  2, 7, 8,
    };

    const expected = Mat4{
        20, 22, 50,  48,
        44, 54, 114, 108,
        40, 58, 110, 102,
        16, 26, 46,  42,
    };

    try std.testing.expect(approxEq(Mat4, mat4Multiply(m1, m2), expected));
}

test "Mat4:mat4Multiply with identity produces startng matrix" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const m2 = mat4_identity;

    try std.testing.expect(approxEq(Mat4, mat4Multiply(m1, m2), m1));
}

pub fn mat4MultiplyVec(m: Mat4, v: vec.Vec4) vec.Vec4 {
    return .{
        m[0] * vec.x(v) + m[1] * vec.y(v) + m[2] * vec.z(v) + m[3] * vec.w(v),
        m[4] * vec.x(v) + m[5] * vec.y(v) + m[6] * vec.z(v) + m[7] * vec.w(v),
        m[8] * vec.x(v) + m[9] * vec.y(v) + m[10] * vec.z(v) + m[11] * vec.w(v),
        m[12] * vec.x(v) + m[13] * vec.y(v) + m[14] * vec.z(v) + m[15] * vec.w(v),
    };
}

test "Mat4:mat4MultiplyVec works" {
    const m = Mat4{
        1, 2, 3, 4,
        2, 4, 4, 2,
        8, 6, 4, 1,
        0, 0, 0, 1,
    };

    const v = vec.Vec4{ 1, 2, 3, 1 };

    const expected = vec.Vec4{ 18, 24, 33, 1 };

    try std.testing.expect(vec.approxEq(mat4MultiplyVec(m, v), expected));
}

pub fn mat4Transpose(m: Mat4) Mat4 {
    return Mat4{
        m[0], m[4], m[8],  m[12],
        m[1], m[5], m[9],  m[13],
        m[2], m[6], m[10], m[14],
        m[3], m[7], m[11], m[15],
    };
}

test "Mat4:mat4Transpose works" {
    const m1 = Mat4{
        1, 2, 3, 4,
        5, 6, 7, 8,
        9, 8, 7, 6,
        5, 4, 3, 2,
    };

    const expected = Mat4{
        1, 5, 9, 5,
        2, 6, 8, 4,
        3, 7, 7, 3,
        4, 8, 6, 2,
    };

    try std.testing.expect(approxEq(Mat4, mat4Transpose(m1), expected));
}

pub fn mat4Inverse(m: Mat4) Mat4 {
    var adj: Mat4 = undefined;

    // Row 0
    adj[0] = m[5] * m[10] * m[15] - m[5] * m[11] * m[14] - m[9] * m[6] * m[15] + m[9] * m[7] * m[14] + m[13] * m[6] * m[11] - m[13] * m[7] * m[10];
    adj[1] = -m[1] * m[10] * m[15] + m[1] * m[11] * m[14] + m[9] * m[2] * m[15] - m[9] * m[3] * m[14] - m[13] * m[2] * m[11] + m[13] * m[3] * m[10];
    adj[2] = m[1] * m[6] * m[15] - m[1] * m[7] * m[14] - m[5] * m[2] * m[15] + m[5] * m[3] * m[14] + m[13] * m[2] * m[7] - m[13] * m[3] * m[6];
    adj[3] = -m[1] * m[6] * m[11] + m[1] * m[7] * m[10] + m[5] * m[2] * m[11] - m[5] * m[3] * m[10] - m[9] * m[2] * m[7] + m[9] * m[3] * m[6];
    // Row 1
    adj[4] = -m[4] * m[10] * m[15] + m[4] * m[11] * m[14] + m[8] * m[6] * m[15] - m[8] * m[7] * m[14] - m[12] * m[6] * m[11] + m[12] * m[7] * m[10];
    adj[5] = m[0] * m[10] * m[15] - m[0] * m[11] * m[14] - m[8] * m[2] * m[15] + m[8] * m[3] * m[14] + m[12] * m[2] * m[11] - m[12] * m[3] * m[10];
    adj[6] = -m[0] * m[6] * m[15] + m[0] * m[7] * m[14] + m[4] * m[2] * m[15] - m[4] * m[3] * m[14] - m[12] * m[2] * m[7] + m[12] * m[3] * m[6];
    adj[7] = m[0] * m[6] * m[11] - m[0] * m[7] * m[10] - m[4] * m[2] * m[11] + m[4] * m[3] * m[10] + m[8] * m[2] * m[7] - m[8] * m[3] * m[6];
    // Row 2
    adj[8] = m[4] * m[9] * m[15] - m[4] * m[11] * m[13] - m[8] * m[5] * m[15] + m[8] * m[7] * m[13] + m[12] * m[5] * m[11] - m[12] * m[7] * m[9];
    adj[9] = -m[0] * m[9] * m[15] + m[0] * m[11] * m[13] + m[8] * m[1] * m[15] - m[8] * m[3] * m[13] - m[12] * m[1] * m[11] + m[12] * m[3] * m[9];
    adj[10] = m[0] * m[5] * m[15] - m[0] * m[7] * m[13] - m[4] * m[1] * m[15] + m[4] * m[3] * m[13] + m[12] * m[1] * m[7] - m[12] * m[3] * m[5];
    adj[11] = -m[0] * m[5] * m[11] + m[0] * m[7] * m[9] + m[4] * m[1] * m[11] - m[4] * m[3] * m[9] - m[8] * m[1] * m[7] + m[8] * m[3] * m[5];
    // Row 3
    adj[12] = -m[4] * m[9] * m[14] + m[4] * m[10] * m[13] + m[8] * m[5] * m[14] - m[8] * m[6] * m[13] - m[12] * m[5] * m[10] + m[12] * m[6] * m[9];
    adj[13] = m[0] * m[9] * m[14] - m[0] * m[10] * m[13] - m[8] * m[1] * m[14] + m[8] * m[2] * m[13] + m[12] * m[1] * m[10] - m[12] * m[2] * m[9];
    adj[14] = -m[0] * m[5] * m[14] + m[0] * m[6] * m[13] + m[4] * m[1] * m[14] - m[4] * m[2] * m[13] - m[12] * m[1] * m[6] + m[12] * m[2] * m[5];
    adj[15] = m[0] * m[5] * m[10] - m[0] * m[6] * m[9] - m[4] * m[1] * m[10] + m[4] * m[2] * m[9] + m[8] * m[1] * m[6] - m[8] * m[2] * m[5];

    const det =
        m[0] * adj[0] +
        m[1] * adj[4] +
        m[2] * adj[8] +
        m[3] * adj[12];

    std.debug.assert(!std.math.approxEqAbs(f32, det, 0.0, constants.epsilon));

    return adj / @as(Mat4, @splat(det));
}

test "Mat4:mat4Inverse works" {
    const m = Mat4{
        -5, 2,  6,  -8,
        1,  -5, 1,  8,
        7,  7,  -6, -7,
        1,  -3, 7,  4,
    };
    const expected = Mat4{
        0.21805,  0.45113,  0.24060,  -0.04511,
        -0.80827, -1.45677, -0.44361, 0.52068,
        -0.07895, -0.22368, -0.05263, 0.19737,
        -0.52256, -0.81391, -0.30075, 0.30639,
    };

    try std.testing.expect(approxEq(Mat4, mat4Inverse(m), expected));
}

// TODO: (Novel) use an error union or option for non invertable matrix
// test "Mat4:mat4inverse fails for non invertable matrix" {
//     const m = Mat4{
//         -4, 2,  -2, -3,
//         9,  6,  2,  6,
//         0,  -5, 1,  -5,
//         0,  0,  0,  0,
//     };
// }

pub const Mat3 = @Vector(9, f32);

pub const mat3_identity = Mat3{
    1, 0, 0,
    0, 1, 0,
    0, 0, 1,
};

pub const Mat2 = @Vector(4, f32);

pub const mat2_identity = Mat2{
    1, 0,
    0, 1,
};

pub fn mat2Determinant(m: Mat2) f32 {
    return m[0] * m[3] - m[1] * m[2];
}

test "2D matrix determinant works" {
    const m = Mat2{
        1,  5,
        -3, 2,
    };
    const expected = 17;

    try std.testing.expect(mathf.approxEq(mat2Determinant(m), expected));
}
