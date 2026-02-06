const std = @import("std");
const Canvas = @import("renderer/canvas.zig").Canvas;
const Color = @import("nmath/color.zig").Color;
const Vec4 = @import("nmath/vector.zig").Vec4;

pub const Body = struct {
    position: Vec4,
    velocity: Vec4,
    acceleration: Vec4,
    color: Color,
};

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    var canvas = try Canvas.init(allocator, 300, 200);

    //canvas.writePixel(5, 5, Color.init(1, 0, 0));

    var body = Body{
        .position = Vec4.point(0, 200, 0),
        .velocity = Vec4.vector(5, -5, 0),
        .acceleration = Vec4.point(0, 0.1, 0),
        .color = Color.init(1, 0, 0),
    };

    for (0..300) |_| {
        body.position = body.position.add(body.velocity);
        body.velocity = body.velocity.add(body.acceleration);

        canvas.writePixel(
            @intFromFloat(std.math.clamp(body.position.x, 0, 299)),
            @intFromFloat(std.math.clamp(body.position.y, 0, 199)),
            body.color,
        );
    }

    const canvas_ppm = try canvas.toPPM(allocator);

    //std.debug.print("{s}", .{canvas_ppm});
    _ = try std.fs.File.stdout().write(canvas_ppm);
}
