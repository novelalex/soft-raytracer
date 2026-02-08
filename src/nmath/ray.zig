const v = @import("vector.zig");
const matrix = @import("matrix.zig");
const xf = @import("transform.zig");
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

    pub fn approxEq(self: Ray, other: Ray) bool {
        return v.approxEq(self.origin, other.origin) and
            v.approxEq(self.direction, other.direction);
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

    pub fn transform(ray: Ray, m: matrix.Mat4) Ray {
        return .{
            .origin = matrix.mat4MultiplyVec(m, ray.origin),
            .direction = matrix.mat4MultiplyVec(m, ray.direction),
        };
    }

    test "Ray:transform can translate" {
        const r = Ray.init(v.point(1, 2, 3), v.vector(0, 1, 0));
        const m = xf.translation(3, 4, 5);
        const expected = Ray.init(v.point(4, 6, 8), v.vector(0, 1, 0));
        try std.testing.expect(r.transform(m).approxEq(expected));
    }

    test "Ray:transform can scale" {
        const r = Ray.init(v.point(1, 2, 3), v.vector(0, 1, 0));
        const m = xf.scaling(2, 3, 4);
        const expected = Ray.init(v.point(2, 6, 12), v.vector(0, 3, 0));
        try std.testing.expect(r.transform(m).approxEq(expected));
    }
};

test {
    _ = Ray;
}
