const std = @import("std");
const Canvas = @import("renderer/canvas.zig").Canvas;
const clr = @import("nmath/color.zig");
const v = @import("nmath/vector.zig");
const sphere = @import("nmath/sphere.zig");
const ray = @import("nmath/ray.zig");
const intersect = @import("nmath/intersect.zig");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    var canvas = try Canvas.init(allocator, 100, 100);

    const eye = v.point(0, 0, -5);

    const canvas_pixels: f32 = @floatFromInt(canvas.width);
    const wall_z: f32 = 10;
    const wall_size: f32 = 7;
    const pixel_size: f32 = wall_size / canvas_pixels;
    const half_wall_size: f32 = wall_size / 2;
    const color = clr.init(1, 1, 1);
    const shape = sphere.Sphere.init();

    for (0..canvas.height - 1) |y| {
        const world_y: f32 = half_wall_size - pixel_size * @as(f32, @floatFromInt(y));
        for (0..canvas.width - 1) |x| {
            const world_x: f32 = -half_wall_size + pixel_size * @as(f32, @floatFromInt(x));
            const position = v.point(world_x, world_y, wall_z);
            const r = ray.Ray.init(eye, v.normalize(position - eye));
            const xs = intersect.raySphere(r, shape);

            if (!xs.hit().approxEq(intersect.Intersection.empty)) {
                canvas.writePixel(x, y, color);
            }
        }
    }

    const canvas_ppm = try canvas.toPPM(allocator);

    //std.debug.print("{s}", .{canvas_ppm});
    _ = try std.fs.File.stdout().write(canvas_ppm);
}

test {
    _ = @import("nmath/math_tests.zig");
}
