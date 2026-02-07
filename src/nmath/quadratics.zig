const constants = @import("constants.zig");
const mathf = @import("mathf.zig");
const std = @import("std");

pub const NumSolutions = enum {
    Zero,
    One,
    Two,
};

pub const Solution = union(NumSolutions) {
    Zero: void,
    One: f32,
    Two: struct { f32, f32 },
};

pub fn solve(a: f32, b: f32, c: f32) Solution {
    const discriminant = b * b - 4 * a * c;
    return if (discriminant < -constants.epsilon)
        Solution{ .Zero = {} }
    else if (mathf.approxEq(discriminant, 0.0))
        Solution{ .One = (-b) / (2.0 * a) }
    else
        Solution{ .Two = .{
            ((-b) + @sqrt(discriminant)) / (2.0 * a),
            ((-b) - @sqrt(discriminant)) / (2.0 * a),
        } };
}

pub fn approxEq(a: Solution, b: Solution) bool {
    return switch (a) {
        .Zero => b == .Zero,

        .One => |x| switch (b) {
            .One => |y| mathf.approxEq(x, y),
            else => false,
        },

        .Two => |x| switch (b) {
            .Two => |y| (mathf.approxEq(x[0], y[0]) and mathf.approxEq(x[1], y[1])) or
                (mathf.approxEq(x[0], y[1]) and mathf.approxEq(x[1], y[0])),
            else => false,
        },
    };
}

test "quadratics:solve works" {
    const s = solve(1, -4, -8);
    const expected = Solution{ .Two = .{
        2.0 + 2.0 * @sqrt(3.0),
        2.0 - 2.0 * @sqrt(3.0),
    } };

    try std.testing.expect(approxEq(s, expected));
}
