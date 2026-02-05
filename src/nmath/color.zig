const std = @import("std");
const Writer = std.Io.Writer;
const vector = @import("vector.zig");

pub const Color = union {
    rgba: struct {
        r: f32,
        g: f32,
        b: f32,
        a: f32,
    },
    vec: vector.Vec4,

    pub fn init(r: f32, g: f32, b: f32) Color {
        return .{ .rgba = .{ .r = r, .g = g, .b = b, .a = 1 } };
    }

    pub fn format(self: Color, writer: *Writer) !void {
        try writer.print("Color({d}, {d}, {d})", .{ self.rgba.r, self.rgba.g, self.rgba.b });
    }

    pub fn hadamard_product(self: Color, other: Color) Color {
        return .{
            .rgba = .{
                .r = self.rgba.r * other.rgba.r,
                .g = self.rgba.g * other.rgba.g,
                .b = self.rgba.b * other.rgba.b,
                .a = self.rgba.a * other.rgba.a,
            },
        };
    }
};
