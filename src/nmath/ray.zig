const v = @import("vector.zig");
const std = @import("std");

pub const Ray = struct {
    origin: v.Point,
    direction: v.Vector,

    pub fn init(origin: v.Point, direction: v.Vector) Ray {
        return .{
            .origin = origin,
            .direction = direction,
        };
    }

    pub fn position(ray: Ray, t: f32) v.Point {
        return ray.origin + ray.direction * @as(v.Vec4, @splat(t));
    }

    test "Ray:position works" {
        const cases = .{
            .{ 0, v.point(2, 3, 4) },
            .{ 1, v.point(3, 3, 4) },
            .{ -1, v.point(1, 3, 4) },
            .{ 2.5, v.point(4.5, 3, 4) },
        };

        const ray = Ray.init(v.point(2, 3, 4), v.vector(1, 0, 0));

        inline for (cases) |c| {
            try std.testing.expect(v.approxEq(ray.position(c[0]), c[1]));
        }
    }
};


test {
    _ = Ray;
}
