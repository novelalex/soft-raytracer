const std = @import("std");
const Color = @import("../nmath/color.zig").Color;

pub const Canvas = struct {
    width: usize,
    height: usize,
    buffer: []Color,

    pub fn init(allocator: std.mem.Allocator, width: usize, height: usize) !Canvas {
        const buffer_size = width * height;
        const buffer = try allocator.alloc(Color, buffer_size);

        for (buffer) |*p| {
            p.* = Color.init(0, 0, 0);
        }

        return .{
            .width = width,
            .height = height,
            .buffer = buffer,
        };
    }

    pub fn deinit(self: *Canvas, allocator: std.mem.Allocator) void {
        allocator.free(self.buffer);
    }

    pub fn pixelAt(self: Canvas, x: usize, y: usize) Color {
        return self.buffer[y * self.width + x];
    }

    pub fn writePixel(self: *Canvas, x: usize, y: usize, color: Color) void {
        self.buffer[y * self.width + x] = color;
    }
};
