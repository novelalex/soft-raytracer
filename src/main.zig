const std = @import("std");
const Canvas = @import("renderer/canvas.zig").Canvas;
const Color = @import("nmath/color.zig").Color;

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    var canvas = try Canvas.init(allocator, 10, 10);
    std.debug.print("canvas[4,5]={any}\n", .{canvas.pixelAt(4, 5)});

    const c1 = Color.init(1, 4, 2);
    const c2 = Color.init(3, 4, 2);
    _ = c1.hadamard_product(c2);
}
