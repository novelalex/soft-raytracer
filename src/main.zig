const std = @import("std");
const Canvas = @import("renderer/canvas.zig").Canvas;
const Color = @import("nmath/color.zig").Color;
const Vec4 = @import("nmath/vector.zig").Vec4;

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    var canvas = try Canvas.init(allocator, 10, 10);

    //canvas.writePixel(5, 5, Color.init(1, 0, 0));

    const canvas_ppm = try canvas.toPPM(allocator);

    //std.debug.print("{s}", .{canvas_ppm});
    _ = try std.fs.File.stdout().write(canvas_ppm);
}

test {
    _ = @import("nmath/math_tests.zig");
}
