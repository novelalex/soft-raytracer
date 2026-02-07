const v = @import("vector.zig");
const std = @import("std");
const quadratics = @import("quadratics.zig");
const Sphere = @import("sphere.zig").Sphere;

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

    pub fn intersectSphere(ray: Ray, sphere: Sphere) RaySphereIntersection {
        const D = ray.direction;
        const S = ray.origin;
        const C = sphere.center;
        const r = sphere.radius;

        const a = v.dot(D, D);
        const b = 2 * v.dot(S, D) - 2 * v.dot(D, C);
        const c = v.dot(S, S) - 2 * v.dot(S, C) + v.dot(C, C) - r * r;

        return quadratics.solve(a, b, c);
    }

    test "Ray:intersectSphere when ray passes through two points on a sphere" {
        const r = Ray.init(v.point(0, 0, -5), v.vector(0, 0, 1));
        const s = Sphere.unit;

        const expected = RaySphereIntersection{ .Two = .{ 4, 6 } };

        try std.testing.expect(quadratics.approxEq(r.intersectSphere(s), expected));
    }

    test "Ray:intersectSphere when ray passes through one point on a sphere" {
        const r = Ray.init(v.point(0, 1, -5), v.vector(0, 0, 1));
        const s = Sphere.unit;
        const expected = RaySphereIntersection{ .One = 5 };
        try std.testing.expect(quadratics.approxEq(r.intersectSphere(s), expected));
    }

    test "Ray:intersectSphere when ray misses the sphere" {
        const r = Ray.init(v.point(0, 2, -5), v.vector(0, 0, 1));
        const s = Sphere.unit;
        const expected = RaySphereIntersection{ .Zero = {} };
        try std.testing.expect(quadratics.approxEq(r.intersectSphere(s), expected));
    }

    test "Ray:intersectSphere when ray originates in the sphere" {
        const r = Ray.init(v.point(0, 0, 0), v.vector(0, 0, 1));
        const s = Sphere.unit;
        const expected = RaySphereIntersection{ .Two = .{ -1, 1 } };
        try std.testing.expect(quadratics.approxEq(r.intersectSphere(s), expected));
    }

    test "Ray:intersectSphere when sphere is behind the ray" {
        const r = Ray.init(v.point(0, 0, 5), v.vector(0, 0, 1));
        const s = Sphere.unit;

        const expected = RaySphereIntersection{ .Two = .{ -6, -4 } };

        try std.testing.expect(quadratics.approxEq(r.intersectSphere(s), expected));
    }
};

pub const RaySphereIntersection = quadratics.Solution;

test {
    _ = Ray;
}
