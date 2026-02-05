const std = @import("std");
const Canvas = @import("renderer/canvas.zig").Canvas;
const Color = @import("nmath/color.zig").Color;

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    var canvas = try Canvas.init(allocator, 10, 20);

    const canvas_ppm = try canvas.toPPM(allocator);

    std.debug.print("{s}", .{canvas_ppm});
}
